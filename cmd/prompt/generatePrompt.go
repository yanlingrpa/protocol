package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// APIDoc 代表单个API的文档
type APIDoc struct {
	Type         string      `json:"type"`        // interface / struct / function
	Name         string      `json:"name"`        // 名称
	Package      string      `json:"package"`     // 所属包
	Description  string      `json:"description"` // 详细描述
	Methods      []MethodDoc `json:"methods,omitempty"`
	Fields       []FieldDoc  `json:"fields,omitempty"`
	Examples     []string    `json:"examples,omitempty"`
	RelatedTypes []string    `json:"related_types,omitempty"`
}

// MethodDoc 代表接口方法的文档
type MethodDoc struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []ParamDoc  `json:"parameters,omitempty"`
	Returns     []ReturnDoc `json:"returns,omitempty"`
}

// FieldDoc 代表结构体字段的文档
type FieldDoc struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Tag         string `json:"tag,omitempty"` // JSON标签等
}

// ParamDoc 代表参数文档
type ParamDoc struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Optional    bool   `json:"optional,omitempty"`
}

// ReturnDoc 代表返回值文档
type ReturnDoc struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ProtocolDoc 代表整个协议的文档
type ProtocolDoc struct {
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	Modules     []ModuleDoc `json:"modules"`
	GeneratedAt string      `json:"generated_at"`
}

// ModuleDoc 代表模块文档
type ModuleDoc struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	APIs        []APIDoc `json:"apis"`
}

var (
	// 用于提取参数说明的正则
	paramRegex   = regexp.MustCompile(`(?m)//\s*参数:\s*\n((?://.*\n)*)`)
	returnRegex  = regexp.MustCompile(`(?m)//\s*返回:\s*(.+)`)
	exampleRegex = regexp.MustCompile(`(?m)//\s*示例:\s*\n((?://.*\n)*)`)
)

func main() {
	// 获取协议根目录
	protocolDir := getProtocolDir()
	fmt.Printf("处理协议目录: %s\n", protocolDir)

	modules := []ModuleDoc{
		processModule(protocolDir, "basic"),
		processModule(protocolDir, "osgui"),
		processModule(protocolDir, "browser"),
		processModule(protocolDir, "ossys"),
		processModule(protocolDir, "extention"),
		processModule(protocolDir, "script"),
	}

	doc := ProtocolDoc{
		Name:        "YanLing RPA Protocol",
		Version:     "1.0.0",
		Description: "RPA执行协议 - 定义了RPA系统的核心接口和数据结构",
		Modules:     modules,
		GeneratedAt: getTimeString(),
	}

	// 输出为JSON
	outputPath := filepath.Join(protocolDir, "protocol_docs.json")
	writeJSON(outputPath, doc)
	fmt.Printf("✓ 生成文档: %s\n", outputPath)

	// 也生成为Markdown供人类阅读
	mdPath := filepath.Join(protocolDir, "PROTOCOL_API.md")
	writeMD(mdPath, doc)
	fmt.Printf("✓ 生成文档: %s\n", mdPath)
}

func processModule(rootDir, moduleName string) ModuleDoc {
	moduleDir := filepath.Join(rootDir, moduleName)

	moduleDoc := ModuleDoc{
		Name: moduleName,
		Path: moduleDir,
		APIs: []APIDoc{},
	}

	// 获取模块描述
	if moduleName == "osgui" {
		moduleDoc.Description = "GUI操作协议 - 提供GUI窗口、元素定位、交互操作等能力"
	} else if moduleName == "browser" {
		moduleDoc.Description = "浏览器自动化协议 - 提供浏览器窗口、页面、框架等操作"
	} else if moduleName == "basic" {
		moduleDoc.Description = "基本数据结构 - 定义坐标、矩形、大小等基础类型"
	} else if moduleName == "ossys" {
		moduleDoc.Description = "系统操作协议 - 提供设备信息、HTTP客户端、文件系统等操作"
	} else if moduleName == "script" {
		moduleDoc.Description = "脚本运行时 - 提供脚本执行环境和配置管理"
	} else if moduleName == "extention" {
		moduleDoc.Description = "扩展协议 - 提供视觉识别等扩展能力"
	}

	// 处理模块内的所有Go文件
	files, err := os.ReadDir(moduleDir)
	if err != nil {
		fmt.Printf("警告: 无法读取模块目录 %s: %v\n", moduleDir, err)
		return moduleDoc
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") && !strings.HasSuffix(file.Name(), "_test.go") {
			filePath := filepath.Join(moduleDir, file.Name())
			apis := extractAPIs(filePath, moduleName)
			moduleDoc.APIs = append(moduleDoc.APIs, apis...)
		}
	}

	return moduleDoc
}

func extractAPIs(filePath, moduleName string) []APIDoc {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		fmt.Printf("警告: 无法解析文件 %s: %v\n", filePath, err)
		return nil
	}

	var apis []APIDoc

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok == token.TYPE {
				for _, spec := range d.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						doc := &APIDoc{
							Name:    typeSpec.Name.Name,
							Package: moduleName,
						}

						if d.Doc != nil {
							doc.Description = cleanComment(d.Doc.Text())
						}

						// 处理接口
						if iface, ok := typeSpec.Type.(*ast.InterfaceType); ok {
							doc.Type = "interface"
							doc.Methods = extractMethods(iface, d)
							apis = append(apis, *doc)
						}

						// 处理结构体
						if s, ok := typeSpec.Type.(*ast.StructType); ok {
							doc.Type = "struct"
							doc.Fields = extractFields(s)
							apis = append(apis, *doc)
						}
					}
				}
			}
		}
	}

	return apis
}

func extractMethods(iface *ast.InterfaceType, genDecl *ast.GenDecl) []MethodDoc {
	var methods []MethodDoc

	if iface.Methods == nil {
		return methods
	}

	for _, method := range iface.Methods.List {
		if len(method.Names) == 0 {
			continue
		}

		methodName := method.Names[0].Name

		methodDoc := MethodDoc{
			Name: methodName,
		}

		// 获取方法描述
		if method.Doc != nil {
			methodDoc.Description = cleanComment(method.Doc.Text())
		}

		// 提取参数
		if funcType, ok := method.Type.(*ast.FuncType); ok {
			if funcType.Params != nil && funcType.Params.List != nil {
				for _, param := range funcType.Params.List {
					paramName := ""
					if len(param.Names) > 0 {
						paramName = param.Names[0].Name
					}
					paramType := getTypeString(param.Type)

					methodDoc.Parameters = append(methodDoc.Parameters, ParamDoc{
						Name: paramName,
						Type: paramType,
					})
				}
			}

			// 提取返回值
			if funcType.Results != nil && funcType.Results.List != nil {
				for _, result := range funcType.Results.List {
					resultType := getTypeString(result.Type)
					methodDoc.Returns = append(methodDoc.Returns, ReturnDoc{
						Type: resultType,
					})
				}
			}
		}

		methods = append(methods, methodDoc)
	}

	return methods
}

func extractFields(s *ast.StructType) []FieldDoc {
	var fields []FieldDoc

	if s.Fields == nil {
		return fields
	}

	for _, field := range s.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		fieldName := field.Names[0].Name
		fieldType := getTypeString(field.Type)

		fieldDoc := FieldDoc{
			Name: fieldName,
			Type: fieldType,
		}

		// 获取字段描述
		if field.Doc != nil {
			fieldDoc.Description = cleanComment(field.Doc.Text())
		} else if field.Comment != nil {
			fieldDoc.Description = cleanComment(field.Comment.Text())
		}

		// 获取标签
		if field.Tag != nil {
			fieldDoc.Tag = field.Tag.Value
		}

		fields = append(fields, fieldDoc)
	}

	return fields
}

func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + getTypeString(t.X)
	case *ast.SelectorExpr:
		return getTypeString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + getTypeString(t.Elt)
	case *ast.MapType:
		return "map[" + getTypeString(t.Key) + "]" + getTypeString(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return "unknown"
	}
}

func cleanComment(text string) string {
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}

func getTimeString() string {
	return "2024-04-17"
}

func getProtocolDir() string {
	// 获取当前文件所在目录的父目录
	dir, _ := os.Getwd()
	return dir
}

func writeJSON(path string, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}

func writeMD(path string, doc ProtocolDoc) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s\n\n", doc.Name))
	sb.WriteString(fmt.Sprintf("**版本**: %s\n\n", doc.Version))
	sb.WriteString(fmt.Sprintf("**描述**: %s\n\n", doc.Description))
	sb.WriteString(fmt.Sprintf("**生成时间**: %s\n\n", doc.GeneratedAt))

	sb.WriteString("## 模块概览\n\n")

	for _, module := range doc.Modules {
		sb.WriteString(fmt.Sprintf("### %s\n\n", module.Name))
		sb.WriteString(fmt.Sprintf("%s\n\n", module.Description))
		sb.WriteString(fmt.Sprintf("**API数量**: %d\n\n", len(module.APIs)))

		for _, api := range module.APIs {
			sb.WriteString(fmt.Sprintf("#### `%s` (%s)\n\n", api.Name, api.Type))

			if api.Description != "" {
				sb.WriteString(fmt.Sprintf("%s\n\n", api.Description))
			}

			if len(api.Methods) > 0 {
				sb.WriteString("**方法**:\n\n")
				for _, method := range api.Methods {
					params := ""
					for _, p := range method.Parameters {
						if params != "" {
							params += ", "
						}
						params += fmt.Sprintf("%s %s", p.Name, p.Type)
					}

					returns := ""
					for _, r := range method.Returns {
						if returns != "" {
							returns += ", "
						}
						returns += r.Type
					}

					sb.WriteString(fmt.Sprintf("- `%s(%s) (%s)`\n", method.Name, params, returns))
					if method.Description != "" {
						sb.WriteString(fmt.Sprintf("  - %s\n", method.Description))
					}
				}
				sb.WriteString("\n")
			}

			if len(api.Fields) > 0 {
				sb.WriteString("**字段**:\n\n")
				for _, field := range api.Fields {
					sb.WriteString(fmt.Sprintf("- `%s: %s`\n", field.Name, field.Type))
					if field.Description != "" {
						sb.WriteString(fmt.Sprintf("  - %s\n", field.Description))
					}
				}
				sb.WriteString("\n")
			}
		}
	}

	return os.WriteFile(path, []byte(sb.String()), 0644)
}
