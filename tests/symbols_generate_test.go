/*
Package tests 包含了自动生成和检查 symbols.go 文件的测试脚本。

主要测试函数：
1. TestGenerateSymbols - 自动生成或更新 symbols.go 文件
  - 扫描项目中除 tests 目录和 symbols.go 以外的所有 Go 文件
  - 提取所有导出的符号（类型、函数、方法、变量、常量、接口）
  - 生成完整的 symbols.go 文件，如果存在则覆盖
  - 提供详细的生成统计信息

2. TestGenerateSymbolsQuick - 快速生成 symbols.go 文件
  - 与 TestGenerateSymbols 功能相同，但输出较少的信息
  - 适合日常使用

使用方法：

	# 生成或更新 symbols.go 文件（带详细信息）
	go test -v -run TestGenerateSymbols

	# 快速生成 symbols.go 文件
	go test -v -run TestGenerateSymbolsQuick

特性：
- 自动处理泛型类型（如 ApiResponse[any]）
- 正确处理常量和变量
- 支持结构体、接口、函数、方法的导出
- 自动排序输出，保持文件整洁
- 跳过 main 包和隐藏文件
- 提供详细的错误报告和统计信息
*/
package tests

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestGenerateSymbols 生成或更新 symbols/protocol.go 文件
func TestGenerateSymbols(t *testing.T) {
	// 获取项目根目录
	rootDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	rootDir = filepath.Dir(rootDir) // 从 tests 目录回到项目根目录

	// 收集所有需要检查的 Go 文件
	var goFiles []string
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过 tests 目录和 symbols.go 文件
		if strings.Contains(path, "tests") || strings.Contains(path, "symbols") || strings.HasSuffix(path, "symbols.go") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// 跳过隐藏目录和文件
		if strings.HasPrefix(filepath.Base(path), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// 收集 Go 文件
		if strings.HasSuffix(path, ".go") {
			goFiles = append(goFiles, path)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// 解析所有 Go 文件，收集导出的符号
	packageSymbols := make(map[string]*PackageInfo)
	fileSet := token.NewFileSet()

	for _, goFile := range goFiles {
		// 解析文件
		node, err := parser.ParseFile(fileSet, goFile, nil, parser.ParseComments)
		if err != nil {
			t.Logf("Warning: failed to parse %s: %v", goFile, err)
			continue
		}

		packageName := node.Name.Name
		if packageName == "main" {
			continue // 跳过 main 包
		}

		// 获取包的完整路径
		packagePath := getPackageKey(goFile, packageName)

		// 跳过根目录的文件（避免创建错误的包路径）
		if packagePath == "github.com/yanlingrpa/protocol/protocol" {
			continue
		}

		if packageSymbols[packagePath] == nil {
			packageSymbols[packagePath] = &PackageInfo{
				Name:       packageName,
				Path:       packagePath,
				Types:      make(map[string]*TypeInfo),
				Functions:  make(map[string]*FunctionInfo),
				Variables:  make(map[string]*VariableInfo),
				Interfaces: make(map[string]*InterfaceInfo),
			}
		}

		pkg := packageSymbols[packagePath]

		// 遍历文件中的声明
		for _, decl := range node.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				// 处理类型、常量、变量声明
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						// 类型声明
						if isExported(s.Name.Name) {
							typeInfo := &TypeInfo{
								Name:    s.Name.Name,
								Methods: make(map[string]*MethodInfo),
							}

							// 检查类型种类
							switch t := s.Type.(type) {
							case *ast.StructType:
								typeInfo.Kind = "struct"
								// 收集结构体字段
								if t.Fields != nil {
									for _, field := range t.Fields.List {
										for _, name := range field.Names {
											if isExported(name.Name) {
												typeInfo.Fields = append(typeInfo.Fields, name.Name)
											}
										}
									}
								}
							case *ast.InterfaceType:
								typeInfo.Kind = "interface"
								// 将接口单独处理
								interfaceInfo := &InterfaceInfo{
									Name:    s.Name.Name,
									Methods: make(map[string]*MethodInfo),
								}
								// 收集接口方法
								if t.Methods != nil {
									for _, method := range t.Methods.List {
										for _, name := range method.Names {
											if isExported(name.Name) {
												interfaceInfo.Methods[name.Name] = &MethodInfo{
													Name: name.Name,
												}
											}
										}
									}
								}
								pkg.Interfaces[s.Name.Name] = interfaceInfo
								continue
							default:
								typeInfo.Kind = "type"
							}

							pkg.Types[s.Name.Name] = typeInfo
						}
					case *ast.ValueSpec:
						// 变量/常量声明
						for _, name := range s.Names {
							if isExported(name.Name) {
								varInfo := &VariableInfo{
									Name: name.Name,
								}
								if d.Tok == token.CONST {
									varInfo.Kind = "const"
								} else {
									varInfo.Kind = "var"
								}
								pkg.Variables[name.Name] = varInfo
							}
						}
					}
				}
			case *ast.FuncDecl:
				// 函数声明
				if d.Name != nil && isExported(d.Name.Name) {
					if d.Recv != nil && len(d.Recv.List) > 0 {
						// 这是一个方法
						recvType, isPtr := getReceiverType(d.Recv.List[0].Type)
						if recvType != "" {
							// 添加到对应类型的方法中
							if typeInfo, exists := pkg.Types[recvType]; exists {
								typeInfo.Methods[d.Name.Name] = &MethodInfo{
									Name:              d.Name.Name,
									Receiver:          recvType,
									IsPointerReceiver: isPtr,
								}
							} else if interfaceInfo, exists := pkg.Interfaces[recvType]; exists {
								interfaceInfo.Methods[d.Name.Name] = &MethodInfo{
									Name:              d.Name.Name,
									Receiver:          recvType,
									IsPointerReceiver: isPtr,
								}
							}
						}
					} else {
						// 这是一个函数
						pkg.Functions[d.Name.Name] = &FunctionInfo{
							Name: d.Name.Name,
						}
					}
				}
			}
		}
	}

	// 生成 symbols.go 文件
	content := generateSymbolsFile(packageSymbols)

	// 写入文件
	symbolsFile := filepath.Join(rootDir, "symbols", "protocol.go")
	err = os.WriteFile(symbolsFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write symbols/protocol.go: %v", err)
	}

	// 统计信息
	totalSymbols := 0
	for _, pkg := range packageSymbols {
		totalSymbols += len(pkg.Types) + len(pkg.Functions) + len(pkg.Variables) + len(pkg.Interfaces)
		for _, typeInfo := range pkg.Types {
			totalSymbols += len(typeInfo.Methods)
		}
		for _, interfaceInfo := range pkg.Interfaces {
			totalSymbols += len(interfaceInfo.Methods)
		}
	}

	t.Logf("✅ Successfully generated symbols.go")
	t.Logf("📊 Statistics:")
	t.Logf("  • Processed %d Go files", len(goFiles))
	t.Logf("  • Found %d packages", len(packageSymbols))
	t.Logf("  • Generated %d symbol entries", totalSymbols)

	for packagePath, pkg := range packageSymbols {
		pkgSymbols := len(pkg.Types) + len(pkg.Functions) + len(pkg.Variables) + len(pkg.Interfaces)
		for _, typeInfo := range pkg.Types {
			pkgSymbols += len(typeInfo.Methods)
		}
		for _, interfaceInfo := range pkg.Interfaces {
			pkgSymbols += len(interfaceInfo.Methods)
		}
		t.Logf("  • Package %s: %d symbols", packagePath, pkgSymbols)
	}
}

// TestGenerateSymbolsQuick 快速生成或更新 symbols/protocol.go 文件
func TestGenerateSymbolsQuick(t *testing.T) {
	// 获取项目根目录
	rootDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	rootDir = filepath.Dir(rootDir) // 从 tests 目录回到项目根目录

	// 收集所有需要检查的 Go 文件
	var goFiles []string
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过 tests 目录和 symbols.go 文件
		if strings.Contains(path, "tests") || strings.Contains(path, "symbols") || strings.HasSuffix(path, "symbols.go") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// 跳过隐藏目录和文件
		if strings.HasPrefix(filepath.Base(path), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// 收集 Go 文件
		if strings.HasSuffix(path, ".go") {
			goFiles = append(goFiles, path)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// 解析所有 Go 文件，收集导出的符号
	packageSymbols := make(map[string]*PackageInfo)
	fileSet := token.NewFileSet()

	for _, goFile := range goFiles {
		// 解析文件
		node, err := parser.ParseFile(fileSet, goFile, nil, parser.ParseComments)
		if err != nil {
			continue
		}

		packageName := node.Name.Name
		if packageName == "main" {
			continue
		}

		packagePath := getPackageKey(goFile, packageName)

		if packagePath == "github.com/yanlingrpa/protocol/protocol" {
			continue
		}

		if packageSymbols[packagePath] == nil {
			packageSymbols[packagePath] = &PackageInfo{
				Name:       packageName,
				Path:       packagePath,
				Types:      make(map[string]*TypeInfo),
				Functions:  make(map[string]*FunctionInfo),
				Variables:  make(map[string]*VariableInfo),
				Interfaces: make(map[string]*InterfaceInfo),
			}
		}

		pkg := packageSymbols[packagePath]

		for _, decl := range node.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						if isExported(s.Name.Name) {
							typeInfo := &TypeInfo{
								Name:    s.Name.Name,
								Methods: make(map[string]*MethodInfo),
							}

							switch t := s.Type.(type) {
							case *ast.StructType:
								typeInfo.Kind = "struct"
								if t.Fields != nil {
									for _, field := range t.Fields.List {
										for _, name := range field.Names {
											if isExported(name.Name) {
												typeInfo.Fields = append(typeInfo.Fields, name.Name)
											}
										}
									}
								}
							case *ast.InterfaceType:
								typeInfo.Kind = "interface"
								interfaceInfo := &InterfaceInfo{
									Name:    s.Name.Name,
									Methods: make(map[string]*MethodInfo),
								}
								if t.Methods != nil {
									for _, method := range t.Methods.List {
										for _, name := range method.Names {
											if isExported(name.Name) {
												interfaceInfo.Methods[name.Name] = &MethodInfo{
													Name: name.Name,
												}
											}
										}
									}
								}
								pkg.Interfaces[s.Name.Name] = interfaceInfo
								continue
							default:
								typeInfo.Kind = "type"
							}

							pkg.Types[s.Name.Name] = typeInfo
						}
					case *ast.ValueSpec:
						for _, name := range s.Names {
							if isExported(name.Name) {
								varInfo := &VariableInfo{
									Name: name.Name,
								}
								if d.Tok == token.CONST {
									varInfo.Kind = "const"
								} else {
									varInfo.Kind = "var"
								}
								pkg.Variables[name.Name] = varInfo
							}
						}
					}
				}
			case *ast.FuncDecl:
				if d.Name != nil && isExported(d.Name.Name) {
					if d.Recv != nil && len(d.Recv.List) > 0 {
						recvType, isPtr := getReceiverType(d.Recv.List[0].Type)
						if recvType != "" {
							if typeInfo, exists := pkg.Types[recvType]; exists {
								typeInfo.Methods[d.Name.Name] = &MethodInfo{
									Name:              d.Name.Name,
									Receiver:          recvType,
									IsPointerReceiver: isPtr,
								}
							} else if interfaceInfo, exists := pkg.Interfaces[recvType]; exists {
								interfaceInfo.Methods[d.Name.Name] = &MethodInfo{
									Name:              d.Name.Name,
									Receiver:          recvType,
									IsPointerReceiver: isPtr,
								}
							}
						}
					} else {
						pkg.Functions[d.Name.Name] = &FunctionInfo{
							Name: d.Name.Name,
						}
					}
				}
			}
		}
	}

	// 生成 symbols.go 文件
	content := generateSymbolsFile(packageSymbols)

	// 写入文件
	symbolsFile := filepath.Join(rootDir, "symbols", "protocol.go")
	err = os.WriteFile(symbolsFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write symbols/protocol.go: %v", err)
	}

	t.Logf("✅ Successfully generated symbols.go")
}

// PackageInfo 包信息
type PackageInfo struct {
	Name       string
	Path       string
	Types      map[string]*TypeInfo
	Functions  map[string]*FunctionInfo
	Variables  map[string]*VariableInfo
	Interfaces map[string]*InterfaceInfo
}

// TypeInfo 类型信息
type TypeInfo struct {
	Name    string
	Kind    string // struct, interface, type
	Fields  []string
	Methods map[string]*MethodInfo
}

// FunctionInfo 函数信息
type FunctionInfo struct {
	Name string
}

// VariableInfo 变量信息
type VariableInfo struct {
	Name string
	Kind string // var, const
}

// InterfaceInfo 接口信息
type InterfaceInfo struct {
	Name    string
	Methods map[string]*MethodInfo
}

// MethodInfo 方法信息
type MethodInfo struct {
	Name              string
	Receiver          string
	IsPointerReceiver bool
}

// generateSymbolsFile 生成 symbols.go 文件内容
func generateSymbolsFile(packageSymbols map[string]*PackageInfo) string {
	var content strings.Builder

	// 文件头
	content.WriteString("package symbols\n\n")
	content.WriteString("import (\n")
	content.WriteString("\t\"reflect\"\n\n")

	// 添加包导入
	var packages []string
	for packagePath := range packageSymbols {
		packages = append(packages, packagePath)
	}
	sort.Strings(packages)
	for _, packagePath := range packages {
		content.WriteString(fmt.Sprintf("\t\"%s\"\n", packagePath))
	}
	content.WriteString(")\n\n")

	// 声明 Symbols 变量
	content.WriteString("var Symbols = make(map[string]map[string]reflect.Value)\n\n")

	// 生成 init 方法
	content.WriteString("func init() {\n")
	for _, packagePath := range packages {
		pkg := packageSymbols[packagePath]
		content.WriteString(fmt.Sprintf("\tSymbols[\"%s\"] = map[string]reflect.Value{\n", packagePath))

		// 获取所有符号并排序
		var allSymbols []string
		symbolMap := make(map[string]string)

		// 类型
		for _, typeInfo := range pkg.Types {
			symbolName := typeInfo.Name
			allSymbols = append(allSymbols, symbolName)
			if typeInfo.Name == "ApiResponse" {
				symbolMap[symbolName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf((*%s.%s[any])(nil)).Elem(), // Export %s %s with any type",
					symbolName, pkg.Name, typeInfo.Name, typeInfo.Name, typeInfo.Kind)
			} else {
				symbolMap[symbolName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf((*%s.%s)(nil)).Elem(), // Export %s %s",
					symbolName, pkg.Name, typeInfo.Name, typeInfo.Name, typeInfo.Kind)
			}
			for _, method := range typeInfo.Methods {
				methodName := fmt.Sprintf("%s.%s", typeInfo.Name, method.Name)
				allSymbols = append(allSymbols, methodName)
				// 根据接收器类型生成不同的代码
				if method.IsPointerReceiver {
					symbolMap[methodName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf((*%s.%s).%s),",
						methodName, pkg.Name, typeInfo.Name, method.Name)
				} else {
					symbolMap[methodName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf(%s.%s.%s),",
						methodName, pkg.Name, typeInfo.Name, method.Name)
				}
			}
		}
		// 接口
		for _, interfaceInfo := range pkg.Interfaces {
			symbolName := interfaceInfo.Name
			allSymbols = append(allSymbols, symbolName)
			symbolMap[symbolName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf((*%s.%s)(nil)), // Export %s interface pointer type",
				symbolName, pkg.Name, interfaceInfo.Name, interfaceInfo.Name)
			for _, method := range interfaceInfo.Methods {
				methodName := fmt.Sprintf("%s.%s", interfaceInfo.Name, method.Name)
				allSymbols = append(allSymbols, methodName)
				symbolMap[methodName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf((*%s.%s)(nil)).MethodByName(\"%s\"),",
					methodName, pkg.Name, interfaceInfo.Name, method.Name)
			}
		}
		// 函数
		for _, funcInfo := range pkg.Functions {
			symbolName := funcInfo.Name
			allSymbols = append(allSymbols, symbolName)
			symbolMap[symbolName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf(%s.%s), // Export %s function",
				symbolName, pkg.Name, funcInfo.Name, funcInfo.Name)
		}
		// 变量
		for _, varInfo := range pkg.Variables {
			symbolName := varInfo.Name
			allSymbols = append(allSymbols, symbolName)
			if varInfo.Kind == "const" {
				symbolMap[symbolName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf(%s.%s), // Export %s constant",
					symbolName, pkg.Name, varInfo.Name, varInfo.Name)
			} else {
				symbolMap[symbolName] = fmt.Sprintf("\t\t\"%s\": reflect.ValueOf(&%s.%s).Elem(), // Export %s variable",
					symbolName, pkg.Name, varInfo.Name, varInfo.Name)
			}
		}
		sort.Strings(allSymbols)
		for _, symbol := range allSymbols {
			content.WriteString(symbolMap[symbol])
			content.WriteString("\n")
		}
		content.WriteString("\t}\n")
	}
	content.WriteString("}\n")
	return content.String()
}
