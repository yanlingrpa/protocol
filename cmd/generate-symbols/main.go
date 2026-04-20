//go:generate go run .

package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

const modulePrefix = "yanlingrpa.com/yanling/protocol"

// excludedDirs are skipped during Go file collection.
var excludedDirs = map[string]struct{}{
	"tests":   {},
	"symbols": {},
	"doc":     {},
	"cmd":     {},
}

func main() {
	rootDir, err := findProjectRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to locate project root: %v\n", err)
		os.Exit(1)
	}

	goFiles, err := collectGoFiles(rootDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to collect Go files: %v\n", err)
		os.Exit(1)
	}

	packageSymbols := parsePackages(goFiles)
	content := generateSymbolsFile(packageSymbols)

	outFile := filepath.Join(rootDir, "symbols", "protocol.go")
	if err := os.WriteFile(outFile, []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write %s: %v\n", outFile, err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", outFile)
	fmt.Printf("Processed %d Go files, %d packages, %d symbol entries\n",
		len(goFiles), len(packageSymbols), countSymbols(packageSymbols))
}

// findProjectRoot walks up from cwd until go.mod is found.
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

// collectGoFiles returns all .go files under rootDir, skipping excluded dirs.
func collectGoFiles(rootDir string) ([]string, error) {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		base := filepath.Base(path)
		if strings.HasPrefix(base, ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if info.IsDir() {
			if _, skip := excludedDirs[base]; skip {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// --- symbol data types ---

type packageInfo struct {
	name       string
	path       string
	types      map[string]*typeInfo
	functions  map[string]*functionInfo
	variables  map[string]*variableInfo
	interfaces map[string]*interfaceInfo
}

type typeInfo struct {
	name    string
	kind    string // struct | type
	fields  []string
	methods map[string]*methodInfo
}

type functionInfo struct{ name string }

type variableInfo struct {
	name string
	kind string // var | const
}

type interfaceInfo struct {
	name    string
	methods map[string]*methodInfo
}

type methodInfo struct {
	name              string
	receiver          string
	isPointerReceiver bool
}

// --- parsing ---

func parsePackages(goFiles []string) map[string]*packageInfo {
	packageSymbols := make(map[string]*packageInfo)
	fset := token.NewFileSet()

	for _, goFile := range goFiles {
		node, err := parser.ParseFile(fset, goFile, nil, parser.ParseComments)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to parse %s: %v\n", goFile, err)
			continue
		}

		pkgName := node.Name.Name
		if pkgName == "main" {
			continue
		}

		pkgPath := packageKey(goFile, pkgName)
		if pkgPath == modulePrefix+"/protocol" {
			continue
		}

		if packageSymbols[pkgPath] == nil {
			packageSymbols[pkgPath] = &packageInfo{
				name:       pkgName,
				path:       pkgPath,
				types:      make(map[string]*typeInfo),
				functions:  make(map[string]*functionInfo),
				variables:  make(map[string]*variableInfo),
				interfaces: make(map[string]*interfaceInfo),
			}
		}
		pkg := packageSymbols[pkgPath]

		for _, decl := range node.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						if !isExported(s.Name.Name) {
							continue
						}
						switch t := s.Type.(type) {
						case *ast.InterfaceType:
							iface := &interfaceInfo{
								name:    s.Name.Name,
								methods: make(map[string]*methodInfo),
							}
							if t.Methods != nil {
								for _, m := range t.Methods.List {
									for _, n := range m.Names {
										if isExported(n.Name) {
											iface.methods[n.Name] = &methodInfo{name: n.Name}
										}
									}
								}
							}
							pkg.interfaces[s.Name.Name] = iface
						case *ast.StructType:
							ti := &typeInfo{
								name:    s.Name.Name,
								kind:    "struct",
								methods: make(map[string]*methodInfo),
							}
							if t.Fields != nil {
								for _, field := range t.Fields.List {
									for _, n := range field.Names {
										if isExported(n.Name) {
											ti.fields = append(ti.fields, n.Name)
										}
									}
								}
							}
							pkg.types[s.Name.Name] = ti
						default:
							pkg.types[s.Name.Name] = &typeInfo{
								name:    s.Name.Name,
								kind:    "type",
								methods: make(map[string]*methodInfo),
							}
						}
					case *ast.ValueSpec:
						for _, n := range s.Names {
							if isExported(n.Name) {
								kind := "var"
								if d.Tok == token.CONST {
									kind = "const"
								}
								pkg.variables[n.Name] = &variableInfo{name: n.Name, kind: kind}
							}
						}
					}
				}
			case *ast.FuncDecl:
				if d.Name == nil || !isExported(d.Name.Name) {
					continue
				}
				if d.Recv != nil && len(d.Recv.List) > 0 {
					recvType, isPtr := receiverType(d.Recv.List[0].Type)
					if recvType != "" {
						mi := &methodInfo{
							name:              d.Name.Name,
							receiver:          recvType,
							isPointerReceiver: isPtr,
						}
						if ti, ok := pkg.types[recvType]; ok {
							ti.methods[d.Name.Name] = mi
						} else if iface, ok := pkg.interfaces[recvType]; ok {
							iface.methods[d.Name.Name] = mi
						}
					}
				} else {
					pkg.functions[d.Name.Name] = &functionInfo{name: d.Name.Name}
				}
			}
		}
	}
	return packageSymbols
}

// --- code generation ---

func generateSymbolsFile(packageSymbols map[string]*packageInfo) string {
	var b strings.Builder

	var pkgPaths []string
	for p := range packageSymbols {
		pkgPaths = append(pkgPaths, p)
	}
	sort.Strings(pkgPaths)

	b.WriteString("package symbols\n\n")
	b.WriteString("import (\n\t\"reflect\"\n\n")
	for _, p := range pkgPaths {
		fmt.Fprintf(&b, "\t\"%s\"\n", p)
	}
	b.WriteString(")\n\n")
	b.WriteString("var Symbols = make(map[string]map[string]reflect.Value)\n\n")
	b.WriteString("func init() {\n")

	for _, pkgPath := range pkgPaths {
		pkg := packageSymbols[pkgPath]
		fmt.Fprintf(&b, "\tSymbols[\"%s\"] = map[string]reflect.Value{\n", pkgPath)

		var allSymbols []string
		symbolMap := make(map[string]string)

		for _, ti := range pkg.types {
			allSymbols = append(allSymbols, ti.name)
			if ti.name == "ApiResponse" {
				symbolMap[ti.name] = fmt.Sprintf(
					"\t\t\"%s\": reflect.ValueOf((*%s.%s[any])(nil)).Elem(), // Export %s %s with any type",
					ti.name, pkg.name, ti.name, ti.name, ti.kind)
			} else {
				symbolMap[ti.name] = fmt.Sprintf(
					"\t\t\"%s\": reflect.ValueOf((*%s.%s)(nil)).Elem(), // Export %s %s",
					ti.name, pkg.name, ti.name, ti.name, ti.kind)
			}
			for _, m := range ti.methods {
				key := ti.name + "." + m.name
				allSymbols = append(allSymbols, key)
				if m.isPointerReceiver {
					symbolMap[key] = fmt.Sprintf(
						"\t\t\"%s\": reflect.ValueOf((*%s.%s).%s),",
						key, pkg.name, ti.name, m.name)
				} else {
					symbolMap[key] = fmt.Sprintf(
						"\t\t\"%s\": reflect.ValueOf(%s.%s.%s),",
						key, pkg.name, ti.name, m.name)
				}
			}
		}

		for _, iface := range pkg.interfaces {
			allSymbols = append(allSymbols, iface.name)
			symbolMap[iface.name] = fmt.Sprintf(
				"\t\t\"%s\": reflect.ValueOf((*%s.%s)(nil)), // Export %s interface pointer type",
				iface.name, pkg.name, iface.name, iface.name)
			for _, m := range iface.methods {
				key := iface.name + "." + m.name
				allSymbols = append(allSymbols, key)
				symbolMap[key] = fmt.Sprintf(
					"\t\t\"%s\": reflect.ValueOf((*%s.%s)(nil)).MethodByName(\"%s\"),",
					key, pkg.name, iface.name, m.name)
			}
		}

		for _, fi := range pkg.functions {
			allSymbols = append(allSymbols, fi.name)
			symbolMap[fi.name] = fmt.Sprintf(
				"\t\t\"%s\": reflect.ValueOf(%s.%s), // Export %s function",
				fi.name, pkg.name, fi.name, fi.name)
		}

		for _, vi := range pkg.variables {
			allSymbols = append(allSymbols, vi.name)
			if vi.kind == "const" {
				symbolMap[vi.name] = fmt.Sprintf(
					"\t\t\"%s\": reflect.ValueOf(%s.%s), // Export %s constant",
					vi.name, pkg.name, vi.name, vi.name)
			} else {
				symbolMap[vi.name] = fmt.Sprintf(
					"\t\t\"%s\": reflect.ValueOf(&%s.%s).Elem(), // Export %s variable",
					vi.name, pkg.name, vi.name, vi.name)
			}
		}

		sort.Strings(allSymbols)
		for _, sym := range allSymbols {
			b.WriteString(symbolMap[sym])
			b.WriteByte('\n')
		}
		b.WriteString("\t}\n")
	}
	b.WriteString("}\n")
	return b.String()
}

// --- helpers ---

func isExported(name string) bool {
	return name != "" && unicode.IsUpper(rune(name[0]))
}

// packageKey derives the full import path from a Go source file path.
func packageKey(goFile, pkgName string) string {
	parts := strings.Split(filepath.ToSlash(goFile), "/")
	for i, p := range parts {
		if p == "protocol" && i > 0 {
			var sub []string
			for _, d := range parts[i+1:] {
				if strings.HasSuffix(d, ".go") {
					break
				}
				sub = append(sub, d)
			}
			if len(sub) > 0 {
				return modulePrefix + "/" + strings.Join(sub, "/")
			}
		}
	}
	return modulePrefix + "/" + pkgName
}

func receiverType(expr ast.Expr) (string, bool) {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name, false
	case *ast.StarExpr:
		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name, true
		}
	}
	return "", false
}

func countSymbols(packageSymbols map[string]*packageInfo) int {
	total := 0
	for _, pkg := range packageSymbols {
		total += len(pkg.types) + len(pkg.functions) + len(pkg.variables) + len(pkg.interfaces)
		for _, ti := range pkg.types {
			total += len(ti.methods)
		}
		for _, iface := range pkg.interfaces {
			total += len(iface.methods)
		}
	}
	return total
}
