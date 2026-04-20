package symbols

import "reflect"

//go:generate go run ../cmd/generate-symbols

// var Symbols = make(map[string]map[string]reflect.Value)

// 手工操作
// 保持 package-path/package-name 作为key
// 1. 删除 text-template.go 文件里的以下方法
// ParseFiles()   // 可以读取文件系统！
// ParseGlob()    // 可以读取文件系统！
// ParseFS()      // 可以读取文件系统！
//
// 2. 删除 path-filepath.go 文件里的以下方法/类型
// Abs()           // 可能泄露工作目录
// EvalSymlinks()  // 可能探测文件系统
// Glob()          // 高危，可能泄露文件系统信息
// Walk()          // 高危，遍历文件系统
// WalkDir()       // 高危，遍历文件系统
// WalkFunc        // Walk 的回调类型，关联 Walk 方法

var Symbols = make(map[string]map[string]reflect.Value)
