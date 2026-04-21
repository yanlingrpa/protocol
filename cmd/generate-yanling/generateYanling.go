package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var excludedDirs = map[string]struct{}{
	".yanling": {},
	"cmd":      {},
	"doc":      {},
	"tests":    {},
	"symbols":  {},
}

type packageAggregate struct {
	Name       string
	ImportPath string
	RelDir     string
	Files      []string
	Interfaces []InterfaceDoc
	Functions  []FunctionDoc
	Structs    map[string]*StructDoc
}

type apiOutput struct {
	Module   string            `json:"module"`
	Packages []apiIndexPackage `json:"packages"`
}

type apiIndexPackage struct {
	Name       string `json:"name"`
	ImportPath string `json:"import_path"`
	APIFile    string `json:"api_file"`
}

type apiPackage struct {
	Name       string         `json:"name"`
	ImportPath string         `json:"import_path"`
	Interfaces []InterfaceDoc `json:"interfaces,omitempty"`
	Functions  []FunctionDoc  `json:"functions,omitempty"`
}

type apiPackageOutput struct {
	Module  string     `json:"module"`
	Package apiPackage `json:"package"`
}

type splitAPIDoc struct {
	FileName string
	Doc      apiPackageOutput
}

type structOutput struct {
	Module   string          `json:"module"`
	Packages []structPackage `json:"packages"`
}

type structPackage struct {
	Name       string      `json:"name"`
	ImportPath string      `json:"import_path"`
	Structs    []StructDoc `json:"structs"`
}

type InterfaceDoc struct {
	Name    string      `json:"name"`
	Doc     string      `json:"doc,omitempty"`
	Methods []MethodDoc `json:"methods,omitempty"`
}

type FunctionDoc struct {
	Name      string     `json:"name"`
	Doc       string     `json:"doc,omitempty"`
	Params    []ParamDoc `json:"params"`
	Results   []ParamDoc `json:"results"`
	Signature string     `json:"signature"`
}

type MethodDoc struct {
	Name      string     `json:"name"`
	Doc       string     `json:"doc,omitempty"`
	Params    []ParamDoc `json:"params"`
	Results   []ParamDoc `json:"results"`
	Signature string     `json:"signature"`
}

type ParamDoc struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type StructDoc struct {
	Name    string      `json:"name"`
	Doc     string      `json:"doc,omitempty"`
	Fields  []FieldDoc  `json:"fields"`
	Methods []MethodDoc `json:"methods,omitempty"`
}

type FieldDoc struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type"`
	Tag      string `json:"tag,omitempty"`
	Doc      string `json:"doc,omitempty"`
	Embedded bool   `json:"embedded,omitempty"`
}

func main() {
	rootDir, err := findProjectRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to locate project root: %v\n", err)
		os.Exit(1)
	}

	moduleName, err := parseModuleName(filepath.Join(rootDir, "go.mod"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse module name: %v\n", err)
		os.Exit(1)
	}

	aggMap, err := scanPackages(rootDir, moduleName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to scan packages: %v\n", err)
		os.Exit(1)
	}

	packages := buildSortedPackages(aggMap)
	apiIndexDoc, splitAPIDocs, structDoc := buildOutputs(moduleName, packages)

	outputDir := filepath.Join(rootDir, ".yanling")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	if err := writeJSON(filepath.Join(outputDir, "api.json"), apiIndexDoc); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write api.json: %v\n", err)
		os.Exit(1)
	}

	apiDir := filepath.Join(outputDir, "api")
	if err := os.RemoveAll(apiDir); err != nil {
		fmt.Fprintf(os.Stderr, "failed to clean api output directory: %v\n", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create api output directory: %v\n", err)
		os.Exit(1)
	}
	for _, split := range splitAPIDocs {
		if err := writeJSON(filepath.Join(apiDir, split.FileName), split.Doc); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write split api file %s: %v\n", split.FileName, err)
			os.Exit(1)
		}
	}

	if err := writeJSON(filepath.Join(outputDir, "struct.json"), structDoc); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write struct.json: %v\n", err)
		os.Exit(1)
	}

	if err := writeInfoMD(filepath.Join(outputDir, "info.md"), moduleName, packages); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write info.md: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("generated %s\n", filepath.Join(outputDir, "info.md"))
	fmt.Printf("generated %s\n", filepath.Join(outputDir, "api.json"))
	fmt.Printf("generated %s\n", filepath.Join(outputDir, "api"))
	fmt.Printf("generated %s\n", filepath.Join(outputDir, "struct.json"))
}

func findProjectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, nil
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			return "", errors.New("go.mod not found")
		}
		wd = parent
	}
}

func parseModuleName(goModPath string) (string, error) {
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(content), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "module ") {
			parts := strings.Fields(trimmed)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	return "", errors.New("module line not found in go.mod")
}

func scanPackages(rootDir, moduleName string) (map[string]*packageAggregate, error) {
	fset := token.NewFileSet()
	packages := make(map[string]*packageAggregate)

	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if d.IsDir() {
			name := d.Name()
			if name == ".git" || strings.HasPrefix(name, ".") {
				if path != rootDir {
					return filepath.SkipDir
				}
			}
			if _, ok := excludedDirs[name]; ok {
				if path != rootDir {
					return filepath.SkipDir
				}
			}
			return nil
		}

		if !strings.HasSuffix(d.Name(), ".go") || strings.HasSuffix(d.Name(), "_test.go") {
			return nil
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath)

		fileAst, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("parse %s: %w", relPath, err)
		}

		relDir := filepath.ToSlash(filepath.Dir(relPath))
		if relDir == "." {
			relDir = ""
		}
		importPath := moduleName
		if relDir != "" {
			importPath = moduleName + "/" + relDir
		}

		pkg := packages[importPath]
		if pkg == nil {
			pkg = &packageAggregate{
				Name:       fileAst.Name.Name,
				ImportPath: importPath,
				RelDir:     relDir,
				Structs:    make(map[string]*StructDoc),
			}
			packages[importPath] = pkg
		}
		pkg.Files = append(pkg.Files, relPath)

		extractDeclarations(pkg, fileAst, relPath)
		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, pkg := range packages {
		sort.Strings(pkg.Files)
	}

	return packages, nil
}

func extractDeclarations(pkg *packageAggregate, fileAst *ast.File, relPath string) {
	for _, decl := range fileAst.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}
			for _, spec := range d.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok || !typeSpec.Name.IsExported() {
					continue
				}

				switch t := typeSpec.Type.(type) {
				case *ast.InterfaceType:
					iface := InterfaceDoc{
						Name:    typeSpec.Name.Name,
						Doc:     cleanComment(pickDoc(d.Doc, typeSpec.Doc)),
						Methods: extractInterfaceMethods(t),
					}
					sortMethods(iface.Methods)
					pkg.Interfaces = append(pkg.Interfaces, iface)
				case *ast.StructType:
					fields := extractStructFields(t)
					structDoc := &StructDoc{
						Name:    typeSpec.Name.Name,
						Doc:     cleanComment(pickDoc(d.Doc, typeSpec.Doc)),
						Fields:  fields,
						Methods: []MethodDoc{},
					}
					sortFields(structDoc.Fields)
					pkg.Structs[typeSpec.Name.Name] = structDoc
				}
			}
		case *ast.FuncDecl:
			if d.Name == nil || !d.Name.IsExported() {
				continue
			}
			if d.Recv == nil {
				fn := buildFunctionDoc(d)
				pkg.Functions = append(pkg.Functions, fn)
				continue
			}

			receiverType := receiverTypeName(d.Recv)
			if receiverType == "" {
				continue
			}
			if structDoc, ok := pkg.Structs[receiverType]; ok {
				method := buildMethodDoc(d)
				structDoc.Methods = append(structDoc.Methods, method)
			}
		}
	}
}

func buildFunctionDoc(fn *ast.FuncDecl) FunctionDoc {
	params := extractFieldList(fn.Type.Params)
	results := extractFieldList(fn.Type.Results)
	return FunctionDoc{
		Name:      fn.Name.Name,
		Doc:       cleanComment(pickDoc(fn.Doc, nil)),
		Params:    params,
		Results:   results,
		Signature: buildSignature(fn.Name.Name, params, results),
	}
}

func buildMethodDoc(fn *ast.FuncDecl) MethodDoc {
	params := extractFieldList(fn.Type.Params)
	results := extractFieldList(fn.Type.Results)
	return MethodDoc{
		Name:      fn.Name.Name,
		Doc:       cleanComment(pickDoc(fn.Doc, nil)),
		Params:    params,
		Results:   results,
		Signature: buildSignature(fn.Name.Name, params, results),
	}
}

func extractInterfaceMethods(iface *ast.InterfaceType) []MethodDoc {
	if iface.Methods == nil {
		return []MethodDoc{}
	}

	methods := make([]MethodDoc, 0, len(iface.Methods.List))
	for _, m := range iface.Methods.List {
		if len(m.Names) == 0 {
			continue
		}
		name := m.Names[0].Name
		funcType, ok := m.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		params := extractFieldList(funcType.Params)
		results := extractFieldList(funcType.Results)
		methods = append(methods, MethodDoc{
			Name:      name,
			Doc:       cleanComment(pickDoc(m.Doc, nil)),
			Params:    params,
			Results:   results,
			Signature: buildSignature(name, params, results),
		})
	}
	return methods
}

func extractStructFields(st *ast.StructType) []FieldDoc {
	if st.Fields == nil {
		return []FieldDoc{}
	}

	fields := make([]FieldDoc, 0, len(st.Fields.List))
	for _, field := range st.Fields.List {
		typeText := exprToString(field.Type)
		doc := cleanComment(pickDoc(field.Doc, field.Comment))
		tag := ""
		if field.Tag != nil {
			tag = strings.Trim(field.Tag.Value, "`")
		}

		if len(field.Names) == 0 {
			if !isExportedEmbedded(field.Type) {
				continue
			}
			embeddedName := embeddedFieldName(field.Type)
			fields = append(fields, FieldDoc{
				Name:     embeddedName,
				Type:     typeText,
				Tag:      tag,
				Doc:      doc,
				Embedded: true,
			})
			continue
		}

		for _, name := range field.Names {
			if !name.IsExported() {
				continue
			}
			fields = append(fields, FieldDoc{
				Name: name.Name,
				Type: typeText,
				Tag:  tag,
				Doc:  doc,
			})
		}
	}
	return fields
}

func extractFieldList(list *ast.FieldList) []ParamDoc {
	if list == nil || len(list.List) == 0 {
		return []ParamDoc{}
	}

	items := make([]ParamDoc, 0)
	for _, field := range list.List {
		typeText := exprToString(field.Type)
		if len(field.Names) == 0 {
			items = append(items, ParamDoc{Type: typeText})
			continue
		}
		for _, name := range field.Names {
			items = append(items, ParamDoc{Name: name.Name, Type: typeText})
		}
	}
	return items
}

func buildSignature(name string, params, results []ParamDoc) string {
	paramParts := make([]string, 0, len(params))
	for _, p := range params {
		if p.Name == "" {
			paramParts = append(paramParts, p.Type)
		} else {
			paramParts = append(paramParts, p.Name+" "+p.Type)
		}
	}

	resultParts := make([]string, 0, len(results))
	for _, r := range results {
		if r.Name == "" {
			resultParts = append(resultParts, r.Type)
		} else {
			resultParts = append(resultParts, r.Name+" "+r.Type)
		}
	}

	sig := fmt.Sprintf("%s(%s)", name, strings.Join(paramParts, ", "))
	switch len(resultParts) {
	case 0:
		return sig
	case 1:
		return sig + " " + resultParts[0]
	default:
		return sig + " (" + strings.Join(resultParts, ", ") + ")"
	}
}

func receiverTypeName(recv *ast.FieldList) string {
	if recv == nil || len(recv.List) == 0 {
		return ""
	}

	switch t := recv.List[0].Type.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name
		}
	}
	return ""
}

func exprToString(expr ast.Expr) string {
	return nodeToString(token.NewFileSet(), expr)
}

func nodeToString(fset *token.FileSet, node any) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, node); err != nil {
		return ""
	}
	return buf.String()
}

func embeddedFieldName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return t.Sel.Name
	case *ast.StarExpr:
		return embeddedFieldName(t.X)
	default:
		return ""
	}
}

func isExportedEmbedded(expr ast.Expr) bool {
	name := embeddedFieldName(expr)
	if name == "" {
		return false
	}
	return ast.IsExported(name)
}

func pickDoc(primary, fallback *ast.CommentGroup) string {
	if primary != nil {
		return primary.Text()
	}
	if fallback != nil {
		return fallback.Text()
	}
	return ""
}

func cleanComment(text string) string {
	lines := strings.Split(text, "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(strings.TrimPrefix(line, "//"))
		line = strings.TrimSpace(strings.TrimPrefix(line, "*"))
		if line != "" {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}

func sortMethods(methods []MethodDoc) {
	sort.Slice(methods, func(i, j int) bool {
		return methods[i].Name < methods[j].Name
	})
}

func sortFields(fields []FieldDoc) {
	sort.Slice(fields, func(i, j int) bool {
		if fields[i].Name == fields[j].Name {
			return fields[i].Type < fields[j].Type
		}
		return fields[i].Name < fields[j].Name
	})
}

func buildSortedPackages(aggMap map[string]*packageAggregate) []*packageAggregate {
	packages := make([]*packageAggregate, 0, len(aggMap))
	for _, pkg := range aggMap {
		sort.Slice(pkg.Interfaces, func(i, j int) bool {
			return pkg.Interfaces[i].Name < pkg.Interfaces[j].Name
		})
		sort.Slice(pkg.Functions, func(i, j int) bool {
			return pkg.Functions[i].Name < pkg.Functions[j].Name
		})
		for _, st := range pkg.Structs {
			sortMethods(st.Methods)
			sortFields(st.Fields)
		}
		packages = append(packages, pkg)
	}

	sort.Slice(packages, func(i, j int) bool {
		return packages[i].ImportPath < packages[j].ImportPath
	})
	return packages
}

func buildOutputs(moduleName string, packages []*packageAggregate) (apiOutput, []splitAPIDoc, structOutput) {
	apiDoc := apiOutput{
		Module:   moduleName,
		Packages: make([]apiIndexPackage, 0, len(packages)),
	}
	splitDocs := make([]splitAPIDoc, 0, len(packages))
	structDoc := structOutput{
		Module:   moduleName,
		Packages: make([]structPackage, 0, len(packages)),
	}

	for _, pkg := range packages {
		interfaces := append([]InterfaceDoc{}, pkg.Interfaces...)
		functions := append([]FunctionDoc{}, pkg.Functions...)

		apiPkg := apiPackage{
			Name:       pkg.Name,
			ImportPath: pkg.ImportPath,
			Interfaces: interfaces,
			Functions:  functions,
		}
		apiFile := apiPackageFileName(pkg)
		apiDoc.Packages = append(apiDoc.Packages, apiIndexPackage{
			Name:       pkg.Name,
			ImportPath: pkg.ImportPath,
			APIFile:    filepath.ToSlash(filepath.Join("api", apiFile)),
		})
		splitDocs = append(splitDocs, splitAPIDoc{
			FileName: apiFile,
			Doc: apiPackageOutput{
				Module:  moduleName,
				Package: apiPkg,
			},
		})

		structNames := make([]string, 0, len(pkg.Structs))
		for name := range pkg.Structs {
			structNames = append(structNames, name)
		}
		sort.Strings(structNames)
		structs := make([]StructDoc, 0, len(structNames))
		for _, name := range structNames {
			structs = append(structs, *pkg.Structs[name])
		}
		if len(structs) == 0 {
			continue
		}

		structDoc.Packages = append(structDoc.Packages, structPackage{
			Name:       pkg.Name,
			ImportPath: pkg.ImportPath,
			Structs:    structs,
		})
	}

	return apiDoc, splitDocs, structDoc
}

func apiPackageFileName(pkg *packageAggregate) string {
	if pkg.RelDir == "" {
		return "root.json"
	}
	name := strings.ReplaceAll(pkg.RelDir, "/", "__")
	name = strings.ReplaceAll(name, "\\", "__")
	return name + ".json"
}

func writeJSON(path string, data any) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	content = append(content, '\n')
	return os.WriteFile(path, content, 0644)
}

func writeInfoMD(path, moduleName string, packages []*packageAggregate) error {
	var sb strings.Builder
	now := time.Now().UTC().Format(time.RFC3339)

	totalInterfaces := 0
	totalFunctions := 0
	totalStructs := 0
	for _, pkg := range packages {
		totalInterfaces += len(pkg.Interfaces)
		totalFunctions += len(pkg.Functions)
		totalStructs += len(pkg.Structs)
	}

	sb.WriteString("# Module Overview\n\n")
	sb.WriteString(fmt.Sprintf("- Module: `%s`\n", moduleName))
	sb.WriteString(fmt.Sprintf("- Generated At: `%s`\n", now))
	sb.WriteString(fmt.Sprintf("- Packages: `%d`\n", len(packages)))
	sb.WriteString(fmt.Sprintf("- Public Interfaces: `%d`\n", totalInterfaces))
	sb.WriteString(fmt.Sprintf("- Public Functions: `%d`\n", totalFunctions))
	sb.WriteString(fmt.Sprintf("- Exported Structs: `%d`\n\n", totalStructs))

	sb.WriteString("## Packages\n\n")
	for _, pkg := range packages {
		sb.WriteString(fmt.Sprintf("### %s\n\n", pkg.ImportPath))
		sb.WriteString(fmt.Sprintf("- Package Name: `%s`\n", pkg.Name))
		if pkg.RelDir == "" {
			sb.WriteString("- Directory: `/`\n")
		} else {
			sb.WriteString(fmt.Sprintf("- Directory: `/%s`\n", pkg.RelDir))
		}
		sb.WriteString(fmt.Sprintf("- Interfaces: `%d`\n", len(pkg.Interfaces)))
		sb.WriteString(fmt.Sprintf("- Functions: `%d`\n", len(pkg.Functions)))
		sb.WriteString(fmt.Sprintf("- Structs: `%d`\n\n", len(pkg.Structs)))

		if len(pkg.Interfaces) > 0 {
			sb.WriteString("Public Interfaces:\n")
			for _, iface := range pkg.Interfaces {
				desc := oneLineDoc(iface.Doc)
				if desc == "" {
					sb.WriteString(fmt.Sprintf("- `%s`\n", iface.Name))
				} else {
					sb.WriteString(fmt.Sprintf("- `%s`: %s\n", iface.Name, desc))
				}
			}
			sb.WriteString("\n")
		}

		if len(pkg.Functions) > 0 {
			sb.WriteString("Public Functions:\n")
			for _, fn := range pkg.Functions {
				desc := oneLineDoc(fn.Doc)
				if desc == "" {
					sb.WriteString(fmt.Sprintf("- `%s`\n", fn.Signature))
				} else {
					sb.WriteString(fmt.Sprintf("- `%s`: %s\n", fn.Signature, desc))
				}
			}
			sb.WriteString("\n")
		}

		if len(pkg.Structs) > 0 {
			structNames := make([]string, 0, len(pkg.Structs))
			for name := range pkg.Structs {
				structNames = append(structNames, name)
			}
			sort.Strings(structNames)
			sb.WriteString("Exported Structs:\n")
			for _, name := range structNames {
				st := pkg.Structs[name]
				desc := oneLineDoc(st.Doc)
				if desc == "" {
					sb.WriteString(fmt.Sprintf("- `%s`\n", st.Name))
				} else {
					sb.WriteString(fmt.Sprintf("- `%s`: %s\n", st.Name, desc))
				}
			}
			sb.WriteString("\n")
		}
	}

	return os.WriteFile(path, []byte(sb.String()), 0644)
}

func oneLineDoc(text string) string {
	if text == "" {
		return ""
	}
	parts := strings.Split(text, "\n")
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			return trimmed
		}
	}
	return ""
}
