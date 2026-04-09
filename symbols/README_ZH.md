# RPA 脚本 API 参考手册

## 📋 概述

本文档列出了 RPA 脚本可以使用的所有标准库和业务 API 接口。所有接口均经过安全审查，确保脚本无法访问本地文件系统（受控访问除外）、无法直接访问网络（通过受控 HTTP 客户端提供）。

**版本**: 1.0  
**最后更新**: 2025-11-07

---

## 🔒 安全说明

### ✅ 脚本可以做的事情
- 字符串、数字、日期等数据处理
- JSON/XML/CSV 等数据格式转换
- 正则表达式匹配
- 图像处理（内存中）
- 通过受控接口访问 HTTP、文件系统、本地存储
- 浏览器自动化操作
- GUI 窗口自动化操作

### ❌ 脚本不能做的事情
- 直接访问本地文件系统（必须通过 `LocalFilesystem` 接口）
- 直接发起网络连接（必须通过 `HttpClient` 接口）
- 执行系统命令
- 访问环境变量或系统信息（除通过 `DeviceInfo` 接口提供的信息）

---

## 📦 标准库包

### 1. 字符串和文本处理

#### `strings` - 字符串操作
常用函数：`Contains`, `Split`, `Join`, `Replace`, `Trim`, `ToUpper`, `ToLower` 等

#### `strconv` - 字符串转换
常用函数：`Atoi`, `Itoa`, `ParseInt`, `ParseFloat`, `FormatInt`, `FormatFloat` 等

#### `regexp` - 正则表达式
常用函数：`Compile`, `Match`, `MatchString`, `FindString`, `ReplaceAllString` 等

#### `unicode` - Unicode 字符属性
常用函数：`IsLetter`, `IsDigit`, `IsSpace`, `ToUpper`, `ToLower` 等

#### `unicode/utf8` - UTF-8 编码
常用函数：`RuneCount`, `Valid`, `DecodeRune`, `EncodeRune` 等

---

### 2. 数据编解码

#### `encoding/json` - JSON 编解码
常用函数：`Marshal`, `Unmarshal`, `MarshalIndent` 等

#### `encoding/xml` - XML 编解码
常用函数：`Marshal`, `Unmarshal` 等

#### `encoding/csv` - CSV 处理
用于读写 CSV 格式数据

#### `encoding/base64` - Base64 编解码
常用函数：`StdEncoding.EncodeToString`, `StdEncoding.DecodeString` 等

#### `encoding/hex` - 十六进制编解码
常用函数：`EncodeToString`, `DecodeString` 等

---

### 3. 加密和哈希

#### `crypto/md5` - MD5 哈希
常用函数：`New`, `Sum` 等

#### `crypto/sha256` - SHA256 哈希
常用函数：`New`, `Sum256` 等

#### `hash` - 哈希接口
提供通用哈希接口

---

### 4. 数学计算

#### `math` - 数学函数
常用函数：`Abs`, `Ceil`, `Floor`, `Max`, `Min`, `Pow`, `Sqrt`, `Sin`, `Cos` 等

#### `math/big` - 大整数和高精度计算
常用类型：`Int`, `Float`, `Rat`  
用于需要高精度或超大数值计算的场景

#### `math/rand` - 随机数生成
常用函数：`Intn`, `Float64`, `Shuffle`, `Seed` 等  
**注意**：这是伪随机数生成器，适用于一般场景

---

### 5. 数据结构和排序

#### `sort` - 排序
常用函数：`Ints`, `Strings`, `Float64s`, `Slice`, `SliceStable` 等

#### `container/heap` - 堆
用于实现优先队列等数据结构

#### `container/list` - 双向链表
提供链表数据结构

#### `container/ring` - 环形链表
提供环形链表数据结构

---

### 6. 字节和缓冲

#### `bytes` - 字节操作
常用函数：`Contains`, `Split`, `Join`, `Replace`, `Trim` 等  
类似 `strings` 但操作字节切片

#### `bufio` - 缓冲 I/O
提供带缓冲的读写器

#### `io` - I/O 接口
提供 `Reader`, `Writer` 等基础接口  
**注意**：仅提供接口，无法直接访问文件

---

### 7. 格式化和错误处理

#### `fmt` - 格式化输出
常用函数：`Sprintf`, `Printf`, `Errorf` 等

#### `errors` - 错误处理
常用函数：`New`, `Is`, `As`, `Unwrap` 等

---

### 8. 路径和 URL

#### `path` - 路径操作（POSIX 风格）
常用函数：`Base`, `Dir`, `Ext`, `Join`, `Clean` 等  
**注意**：仅字符串操作，不访问文件系统

#### `path/filepath` - 文件路径操作
常用函数：`Base`, `Dir`, `Ext`, `Join`, `Clean`, `Split` 等  
**注意**：已移除 `Walk`, `Glob`, `Abs` 等可能访问文件系统的函数

#### `net/url` - URL 解析
常用函数：`Parse`, `ParseQuery`, `QueryEscape`, `JoinPath` 等  
**注意**：仅解析 URL 字符串，不访问网络

---

### 9. 时间和上下文

#### `time` - 时间处理
常用函数：`Now`, `Parse`, `Format`, `Sleep`, `Since`, `Until` 等  
常用类型：`Time`, `Duration`, `Timer`, `Ticker`

#### `context` - 上下文管理
常用函数：`Background`, `TODO`, `WithCancel`, `WithTimeout`, `WithValue` 等  
用于控制超时和取消操作

---

### 10. 图像处理

#### `image` - 图像结构
常用函数：`Decode`, `NewRGBA`, `Rect`, `Pt` 等  
常用类型：`Image`, `RGBA`, `Gray`, `Point`, `Rectangle`

#### `image/color` - 颜色模型
常用类型：`RGBA`, `Gray`, `Alpha` 等  
常用函数：`RGBToYCbCr`, `YCbCrToRGB` 等

#### `image/png` - PNG 图像
常用函数：`Decode`, `Encode`

#### `image/jpeg` - JPEG 图像
常用函数：`Decode`, `Encode`

#### `image/gif` - GIF 图像
常用函数：`Decode`, `DecodeAll`, `Encode`, `EncodeAll`

---

### 11. HTML 和模板

#### `html` - HTML 转义
常用函数：`EscapeString`, `UnescapeString`

#### `text/template` - 文本模板
常用函数：`New`, `Parse` (仅字符串解析)  
**注意**：已移除 `ParseFiles` 等文件访问函数

---

### 12. 并发控制

#### `sync` - 同步原语
常用类型：`Mutex`, `RWMutex`, `WaitGroup`, `Once`, `Pool`, `Map`  
用于并发安全控制

#### `sync/atomic` - 原子操作
提供原子级别的数值操作

---

## 🏢 业务 API 包

### 1. `basic` - 基础数据结构

#### 数据类型
- `Point` - 坐标点 (X, Y)
- `FPoint` - 浮点坐标点
- `Size` - 尺寸 (Width, Height)
- `Rect` - 矩形区域
- `OcrText` - OCR 识别的文本
- `OcrResult` - OCR 识别结果
- `Task` - 任务接口
- `DispatchTaskData` - 任务分发数据

#### 实用函数
- `MaxAreaRect` - 获取最大面积矩形
- `MinAreaRect` - 获取最小面积矩形
- `MergeOverlappingRectangles` - 合并重叠矩形
- `MergeGroupRectangles` - 分组合并矩形
- `MergeAllRectangles` - 合并所有矩形

---

### 2. `browser` - 浏览器自动化

#### `BrowserWindow` - 浏览器窗口
- `CurrentPage()` - 获取当前页面
- `DefaultPage()` - 获取默认页面
- `NewTabPage(url)` - 打开新标签页
- `IDTabPage(id)` - 根据 ID 获取标签页

#### `BrowserTabPage` - 浏览器标签页
- `Activate()` - 激活标签页
- `Destroy()` - 关闭标签页
- `SaveCookies()` / `LoadCookies()` - Cookie 管理
- `SaveLocalStorage()` / `LoadLocalStorage()` - LocalStorage 管理
- `ClearCookies()` / `ClearLocalStorage()` - 清除数据
- `WaitForNewTab()` - 等待新标签页

#### `BrowserFramePage` - 页面/框架
- `GetURL()` / `GetTitle()` / `GetDomain()` - 获取页面信息
- `QuerySelector(css)` - CSS 选择器查询
- `QuerySelectorAll(css)` - 查询所有匹配元素
- `QueryXPath(xpath)` / `QueryXPathAll(xpath)` - XPath 查询
- `WaitSelector(css, timeout)` - 等待元素出现
- `Evaluate(script)` - 执行 JavaScript
- `Reload()` - 刷新页面

#### `BrowserElement` - 页面元素
- `Click()` / `DoubleClick()` / `RightClick()` - 点击操作
- `Hover()` / `MouseMove()` - 鼠标移动
- `Input(text)` / `WriteText(text)` - 输入文本
- `Focus()` / `Blur()` - 焦点控制
- `Text()` / `Html()` - 获取文本/HTML
- `GetAttribute(name)` / `SetAttribute(name, value)` - 属性操作
- `GetProperty(name)` / `SetProperty(name, value)` - 属性操作
- `Visible()` / `Disabled()` / `Interactable()` - 状态查询
- `WaitVisible()` / `WaitEnabled()` - 等待状态
- `ScrollIntoView()` - 滚动到可见
- `SelectByText()` / `SelectByCss()` - 下拉选择
- `SetFiles(paths)` - 文件上传

---

### 3. `osgui` - GUI 窗口自动化

#### `GuiWindow` - GUI 窗口
- `GetWindowTitle()` - 获取窗口标题
- `GetWindowRect()` / `GetClientRect()` - 获取窗口区域
- `GetMonitor()` - 获取所在显示器
- `Activate()` / `DeActivate()` - 激活/失活窗口
- `Close()` - 关闭窗口
- `MoveTo(x, y)` / `ResizeTo(w, h)` - 移动/调整大小
- `Snapshot(rect)` - 截图
- `BodyLocator()` / `RectLocator(rect)` - 创建定位器
- `PressKeys(keys)` - 按键操作
- `ReadClipboard()` / `WriteClipboard(text)` - 剪贴板操作
- `GetWindowCursorPos()` / `GetWindowCaretPos()` - 获取光标位置
- `TransToScreen(point)` / `TransFromScreen(point)` - 坐标转换

#### `Locator` - 元素定位器
- `Click(point)` / `DoubleClick(point)` / `RightClick(point)` - 点击
- `MouseMove(point)` - 鼠标移动
- `Focus()` - 聚焦
- `Snapshot(rect)` - 截图
- `Ocr(rect)` - OCR 识别
- `ReadText()` / `WriteText(text)` / `ClearText()` - 文本操作
- `IsEditing()` / `WaitForEditing()` - 编辑状态
- `ScrollVertical(amount)` / `ScrollHorizontal(amount)` - 滚动
- `ImageLocator(template, similarity)` - 图像匹配定位
- `ImageLocators(template, similarity)` - 多个图像匹配
- `TextLocator(text)` / `TextLocators(text)` - 文本定位
- `ShapeLocator(shape)` / `ShapeLocators(shape)` - 图形定位
- `SubLocator(rect)` - 创建子定位器
- `TransToScreen()` / `TransFromScreen()` - 坐标转换

#### 键盘常量
提供所有常用按键常量：`Enter`, `Esc`, `Ctrl`, `Alt`, `Shift`, `F1`-`F24`, `KeyA`-`KeyZ`, `Key0`-`Key9` 等

#### 图形形状常量
- `GraphicShape_Circle` - 圆形
- `GraphicShape_Rectangle` - 矩形
- `GraphicShape_Triangle` - 三角形
- `GraphicShape_Star` - 星形
- 等等

---

### 4. `ossys` - 系统服务

#### `DeviceInfo` - 设备信息
- `DeviceId()` / `DeviceName()` - 设备标识
- `GetComputerName()` / `GetUserName()` - 计算机名/用户名
- `OS()` / `OSVersion()` - 操作系统信息
- `NumLogicCPU()` - CPU 核心数
- `GetGpuMemoryMB()` / `HasNvidiaGPU()` - GPU 信息
- `GetMonitors()` / `GetPrimaryMonitor()` - 显示器信息

#### `MonitorInfo` - 显示器信息
- `GetBounds()` / `GetWorkArea()` - 屏幕区域
- `GetDPI()` - DPI 信息
- `IsPrimary()` - 是否主显示器

#### `HttpClient` - HTTP 客户端（受控）
- `Get(url, headers)` - GET 请求
- `Post(url, data, headers)` - POST 请求
- `PostJson(url, jsonData, headers)` - POST JSON
- `PostForm(url, formData, headers)` - POST 表单
- `DownloadFile(url, savePath)` - 下载文件
- `UploadFile(url, filePath, headers)` - 上传文件
- `UploadData(url, data, filename, headers)` - 上传数据
- `SetDomainCookies(domain, cookies)` - 设置 Cookie
- `GetDomainCookies(domain)` - 获取 Cookie
- `SetDomainHeaders(domain, headers)` - 设置请求头
- `GetDomainHeaders(domain)` - 获取请求头

#### `LocalFilesystem` - 本地文件系统（受控）
- `JoinDataPath(paths...)` - 拼接数据目录路径
- `PathExists(path)` - 检查路径是否存在
- `IsFile(path)` / `IsDir(path)` - 判断文件/目录
- `ReadFile(path)` - 读取文件
- `WriteFile(path, data)` - 写入文件
- `CopyFile(src, dst)` - 复制文件
- `Rename(old, new)` - 重命名
- `Remove(path)` / `RemoveAll(path)` - 删除
- `MkdirAll(path)` - 创建目录
- `CreateTmpFile()` - 创建临时文件

#### `LocalStorage` - 本地存储（Key-Value）
- `Get(key)` / `Set(key, value)` - 读写
- `SetEx(key, value, ttl)` - 带过期时间写入
- `MGet(keys...)` / `MSet(kvPairs...)` - 批量操作
- `Del(keys...)` - 删除
- `Keys(pattern)` - 查询键

#### `ScriptLogger` - 日志记录
- `Debug(message)` - 调试日志
- `Info(message)` - 信息日志
- `Warn(message)` - 警告日志
- `Error(message)` - 错误日志

---

### 5. `robot` - 脚本执行框架

#### `ProjectExecutor` - 项目执行器
- `GetID()` - 获取项目 ID
- `GetVariable(key)` - 获取项目变量
- `IsInterrupts()` - 检查是否中断
- `SleepRandom(min, max)` - 随机睡眠
- `LockScreen()` / `UnlockScreen()` - 锁定/解锁屏幕
- `DeviceInfo()` - 获取设备信息
- `HttpClient()` - 获取 HTTP 客户端
- `FileSystem()` - 获取文件系统
- `Storage()` - 获取本地存储
- `Logger()` - 获取日志记录器
- `DispatchTask(taskData)` - 分发任务

#### 脚本类型
- `BrowserScript` - 浏览器脚本接口
- `GuiScript` - GUI 脚本接口
- `GoScript` - Go 脚本接口
- `WxAppScript` - 微信小程序脚本接口

#### 操作类型常量
- `BrowserOperator` - 浏览器操作
- `GuiOperator` - GUI 操作
- `WxAppOperator` - 微信小程序操作

---

### 6. `wxapp` - 微信小程序自动化

#### `WxAppWindow` - 微信小程序窗口
提供微信小程序自动化能力（具体方法根据实现暴露）

---

## 📖 使用示例

### 示例 1: 字符串处理
```go
import "strings"

text := "Hello, RPA World!"
upper := strings.ToUpper(text)        // "HELLO, RPA WORLD!"
parts := strings.Split(text, ", ")    // ["Hello", "RPA World!"]
joined := strings.Join(parts, " - ")  // "Hello - RPA World!"
```

### 示例 2: JSON 处理
```go
import "encoding/json"

// 序列化
data := map[string]interface{}{
    "name": "RPA Task",
    "count": 42,
}
jsonBytes, _ := json.Marshal(data)
jsonStr := string(jsonBytes)

// 反序列化
var result map[string]interface{}
json.Unmarshal(jsonBytes, &result)
```

### 示例 3: HTTP 请求
```go
// 通过 ProjectExecutor 获取 HttpClient
httpClient := executor.HttpClient()

// GET 请求
response, _ := httpClient.Get("https://api.example.com/data", nil)

// POST JSON
jsonData := `{"name": "test"}`
response, _ := httpClient.PostJson("https://api.example.com/create", jsonData, nil)

// 下载文件
httpClient.DownloadFile("https://example.com/file.pdf", "downloaded.pdf")
```

### 示例 4: 浏览器自动化
```go
// 获取浏览器窗口（由框架注入）
page := browserWindow.DefaultPage()

// 导航到页面
page.Evaluate(`window.location.href = "https://example.com"`)

// 查询元素并点击
button := page.QuerySelector("#submit-button")
button.Click()

// 输入文本
input := page.QuerySelector("input[name='username']")
input.Input("myusername")

// 等待元素出现
page.WaitSelector(".success-message", 5000)
```

### 示例 5: GUI 自动化
```go
// 获取 GUI 窗口（由框架注入）
window.Activate()

// 创建定位器
locator := window.BodyLocator()

// 图像识别点击
imageLocator := locator.ImageLocator("button_template.png", 0.9)
imageLocator.Click(basic.Point{X: 10, Y: 10})

// OCR 文本识别
ocrResult := locator.Ocr(basic.Rect{X: 100, Y: 100, Width: 200, Height: 50})
for _, text := range ocrResult.Texts {
    logger.Info("识别到文本: " + text.Text)
}

// 按键操作
window.PressKeys([]string{osgui.Ctrl, osgui.KeyC})  // Ctrl+C
```

### 示例 6: 文件操作
```go
// 通过 ProjectExecutor 获取 FileSystem
fs := executor.FileSystem()

// 读取文件
content, _ := fs.ReadFile(fs.JoinDataPath("config.json"))

// 写入文件
fs.WriteFile(fs.JoinDataPath("output.txt"), []byte("Hello, RPA!"))

// 检查文件是否存在
exists := fs.PathExists(fs.JoinDataPath("data.csv"))
```

### 示例 7: 本地存储
```go
// 通过 ProjectExecutor 获取 Storage
storage := executor.Storage()

// 保存数据
storage.Set("last_run_time", "2025-11-07 10:30:00")
storage.SetEx("temp_token", "abc123", 3600)  // 1小时后过期

// 读取数据
lastRun := storage.Get("last_run_time")

// 批量操作
values := storage.MGet("key1", "key2", "key3")
```

---

## ⚠️ 重要限制

1. **文件系统访问受限**
   - 只能通过 `LocalFilesystem` 接口访问文件
   - 访问路径受沙箱限制
   - 无法使用 `os` 包直接操作文件

2. **网络访问受限**
   - 只能通过 `HttpClient` 接口访问网络
   - 可能有白名单或黑名单限制
   - 无法使用 `net` 包直接建立连接

3. **系统命令禁止**
   - 无法执行外部程序
   - 无法访问环境变量（除非通过受控接口）

4. **反射和不安全操作禁止**
   - 未暴露 `reflect` 包
   - 未暴露 `unsafe` 包
   - 防止绕过安全限制

---

## 📚 更多资源

- 完整 API 文档：[待补充]
- 示例脚本库：[待补充]
- 常见问题：[待补充]

---

**文档版本**: 1.0  
**生成时间**: 2025-11-07