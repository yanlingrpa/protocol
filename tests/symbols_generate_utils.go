package tests

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// isExported 检查标识符是否是导出的（首字母大写）
func isExported(name string) bool {
	if name == "" {
		return false
	}
	return unicode.IsUpper(rune(name[0]))
}

// getPackageKey 根据文件路径和包名生成包键
func getPackageKey(goFile, packageName string) string {
	// 从文件路径中提取相对路径
	parts := strings.Split(goFile, string(filepath.Separator))

	// 找到项目根目录后的路径
	var relevantParts []string
	foundRpa := false
	for _, part := range parts {
		if part == "rpa-execution-protocol" {
			foundRpa = true
			continue
		}
		if foundRpa && part != "" && !strings.HasSuffix(part, ".go") {
			relevantParts = append(relevantParts, part)
		}
	}

	if len(relevantParts) > 0 {
		return fmt.Sprintf("github.com/yanlingrpa/protocol/%s", strings.Join(relevantParts, "/"))
	}

	return fmt.Sprintf("github.com/yanlingrpa/protocol/%s", packageName)
}

// getReceiverType 从接收者类型中提取类型名称和是否是指针接收器
func getReceiverType(expr ast.Expr) (string, bool) {
	switch t := expr.(type) {
	case *ast.Ident:
		// 非指针接收器
		return t.Name, false
	case *ast.StarExpr:
		// 指针接收器
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name, true
		}
	}
	return "", false
}

// parseSymbolsFile 解析 symbols.go 文件并提取符号定义
func parseSymbolsFile(rootDir string) map[string]map[string]bool {
	symbolsFile := filepath.Join(rootDir, "symbols.go")

	// 读取 symbols.go 文件内容
	content, err := os.ReadFile(symbolsFile)
	if err != nil {
		return nil
	}

	// 解析文件
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, symbolsFile, content, parser.ParseComments)
	if err != nil {
		return nil
	}

	symbolsMap := make(map[string]map[string]bool)

	// 遍历文件中的声明
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range valueSpec.Names {
						if name.Name == "Symbols" {
							// 找到 Symbols 变量，解析其内容
							if compositeLit, ok := valueSpec.Values[0].(*ast.CompositeLit); ok {
								parseSymbolsComposite(compositeLit, symbolsMap)
							}
						}
					}
				}
			}
		}
	}

	return symbolsMap
}

// parseSymbolsComposite 解析 Symbols 变量的复合字面量
func parseSymbolsComposite(composite *ast.CompositeLit, symbolsMap map[string]map[string]bool) {
	for _, elt := range composite.Elts {
		if keyValue, ok := elt.(*ast.KeyValueExpr); ok {
			// 获取包名
			var packageName string
			if basicLit, ok := keyValue.Key.(*ast.BasicLit); ok {
				packageName = strings.Trim(basicLit.Value, `"`)
			}

			if packageName == "" {
				continue
			}

			// 初始化包的符号 map
			if symbolsMap[packageName] == nil {
				symbolsMap[packageName] = make(map[string]bool)
			}

			// 解析包内的符号
			if packageComposite, ok := keyValue.Value.(*ast.CompositeLit); ok {
				for _, pkgElt := range packageComposite.Elts {
					if pkgKeyValue, ok := pkgElt.(*ast.KeyValueExpr); ok {
						// 获取符号名
						var symbolName string
						if basicLit, ok := pkgKeyValue.Key.(*ast.BasicLit); ok {
							symbolName = strings.Trim(basicLit.Value, `"`)
						}

						if symbolName != "" {
							symbolsMap[packageName][symbolName] = true
						}
					}
				}
			}
		}
	}
}
