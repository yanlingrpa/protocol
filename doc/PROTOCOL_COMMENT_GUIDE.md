# RPA 协议文档注释规范

本指南定义了在协议代码中添加注释的标准方式，用于自动生成协议文档和 AI 使用的提示。

## 1. 接口文档注释规范

### 基本格式

```go
// InterfaceName 简短一句话描述
//
// 用途：详细说明此接口的用途
//
// 主要能力：
// - 能力1
// - 能力2
// - 能力3
//
// 继承的接口（如有）：
// - ParentInterface
type InterfaceName interface {
	// 方法简要说明
	//
	// 参数:
	//   param1 - 参数1说明，类型为 string
	//   param2 - 参数2说明，类型为 int，可选
	//
	// 返回:
	//   返回值类型1: 返回值说明
	//   error: 错误说明
	//
	// 示例:
	//   result, err := obj.Method(param1, param2)
	Method(param1 string, param2 int) (string, error)
}
```

### 实例 - 好的注释

```go
// BrowserWindow 代表浏览器窗口对象
//
// 用途：提供浏览器窗口的管理操作，包括标签页切换、导航、脚本执行等
//
// 主要能力：
// - 标签页管理（默认页、当前页、按ID获取、新建标签页）
// - 继承GuiWindow的窗口操作（移动、缩放、截图、键盘鼠标操作）
// - Cookie和本地存储管理
//
// 继承的接口：
// - osgui.GuiWindow
type BrowserWindow interface {
	osgui.GuiWindow
	
	// DefaultPage 获取浏览器的首页（第一个标签页）
	//
	// 返回:
	//   BrowserTabPage: 首页标签页对象
	//
	// 说明：
	//   这是浏览器启动时自动创建的第一个标签页
	DefaultPage() BrowserTabPage
	
	// NewTabPage 创建新的浏览器标签页
	//
	// 参数:
	//   id  - 标签页的唯一标识符，字符串类型
	//   url - 要打开的URL，字符串类型
	//
	// 返回:
	//   BrowserTabPage: 新创建的标签页对象
	//   error: 创建失败时返回错误信息
	//
	// 示例:
	//   page, err := window.NewTabPage("tab1", "https://www.google.com")
	//   if err != nil {
	//       log.Fatal(err)
	//   }
	//   page.Activate()
	NewTabPage(id string, url string) (BrowserTabPage, error)
}
```

## 2. 结构体/数据类型注释规范

```go
// StructName 简短说明
//
// 用途：详细说明此结构体的用途
//
// 典型使用场景：
// - 场景1
// - 场景2
type StructName struct {
	// FieldName 字段简要说明
	//
	// 类型: 字段类型
	// 可选: 是否可选（字段是否可以为空）
	// 示例: "示例值"
	FieldName FieldType `json:"field_name"`
}
```

### 实例 - 数据结构

```go
// ModuleInfo 脚本模块信息
//
// 用途：描述RPA脚本的元数据，包括版本、作者、支持的设备等
//
// 典型使用场景：
// - 注册脚本到系统
// - 显示脚本信息到UI
// - 检查兼容性
type ModuleInfo struct {
	// Specifier 模块标识符
	//
	// 格式: domain/owner/name@version (如 yanlingrpa.com/company/my-script@1.0.0)
	// 来源: 从 go.mod 文件中解析
	Specifier string `json:"specifier"`
	
	// Description 脚本功能描述
	//
	// 用途: 向用户介绍脚本的功能和用途
	// 长度: 建议 100-500 字符
	// 示例: "自动登录并爬取用户信息的脚本"
	Description string `toml:"description" json:"description"`
	
	// Tags 脚本主题标签列表
	//
	// 格式: 字符串数组
	// 用途: 分类和搜索
	// 示例: ["web automation", "data extraction", "login"]
	Tags []string `toml:"tags" json:"tags"`
	
	// Devices 适用的设备列表
	//
	// 可选值: "windows", "mac", "ubuntu", "android", "ios"
	// 示例: ["windows", "mac"]
	Devices []string `toml:"devices" json:"devices"`
}
```

## 3. 函数/方法注释的最佳实践

### 参数说明

- **参数名**: 使用参数的实际名称
- **类型**: 清晰说明参数类型
- **说明**: 说明参数的含义和约束
- **可选标记**: 如果是可选参数，标记为 "可选"

### 返回值说明

- **类型**: 清晰说明返回值类型
- **含义**: 说明返回值代表什么
- **错误情况**: 说明什么情况下会返回 error

### 示例说明

- 提供真实可用的代码示例
- 示例应该展示最常见的用法
- 包括错误处理示例

## 4. 代码注释示例集合

### 示例1：GUI窗口操作

```go
// PressKeys 模拟键盘按键输入
//
// 参数:
//   keys - 按键序列，可变长参数，类型为 Keyboard
//
// 返回:
//   error: 输入失败时返回错误
//
// 支持的按键：
//   ctrl, alt, shift, win (修饰键)
//   enter, esc, tab, space, home, end, pageup, pagedown (功能键)
//   0-9, a-z (字符键)
//   组合: ctrl+a (同时按下多个键)
//
// 示例:
//   // 按下 Ctrl+C 复制
//   window.PressKeys(ctrl, c)
//   
//   // 按下 Ctrl+Alt+Delete
//   window.PressKeys(ctrl+alt+del)
PressKeys(keys ...Keyboard) error
```

### 示例2：图像定位

```go
// ImageLocator 在当前区域内查找包含相似图片的子定位器
//
// 参数:
//   image - 图片来源，支持以下格式：
//           1. 相对路径: "../images/button.png" (相对于项目根目录)
//           2. URL: "https://example.com/image.png"
//           3. Base64: "data:image/png;base64,iVBORw0KGgoAAAANS..."
//           4. 绝对路径: "/absolute/path/to/image.png"
//   sim   - 相似度阈值，范围 0.1-1.0
//           0.1: 极度相似（容错极少）
//           0.5: 中等相似（推荐值）
//           0.9: 宽松相似（容错很大）
//
// 返回:
//   Locator: 找到的子定位器对象
//   error: 找不到或参数无效时返回错误
//
// 示例:
//   // 查找相似度>=80%的图片
//   loc, err := bodyLoc.ImageLocator("./images/submit_btn.png", 0.8)
//   if err == nil {
//       loc.LeftClick()  // 点击找到的按钮
//   }
ImageLocator(image string, sim float32) (Locator, error)
```

## 5. 生成文档

### 执行生成

```bash
cd protocol/symbols
go generate
```

### 输出文件

生成两个文件：

1. **protocol_docs.json** - 机器可读的JSON格式（用于AI）
2. **PROTOCOL_API.md** - 人类可读的Markdown格式

## 6. 注释的DO和DON'T

### ✅ DO

- 使用清晰的英文或中文描述
- 对复杂参数给出多个示例
- 说明参数的约束和有效范围
- 说明方法的副作用（如修改状态）
- 说明错误情况和异常

### ❌ DON'T

- 不要重复代码中显而易见的内容
- 不要使用过于简洁的缩写（除非在代码中已定义）
- 不要留下空注释或不完整注释
- 不要在注释中包含个人信息或草稿内容

## 7. 文档更新流程

1. **修改代码和注释**
2. **运行生成器**: `go generate ./symbols`
3. **提交生成的文件**: `protocol_docs.json` 和 `PROTOCOL_API.md`
4. **在CI/CD中运行生成**: 确保文档始终是最新的
