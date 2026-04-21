# Module Overview

- Module: `yanlingrpa.com/yanling/protocol`
- Generated At: `2026-04-21T04:00:18Z`
- Packages: `6`
- Public Interfaces: `15`
- Public Functions: `6`
- Exported Structs: `14`

## Packages

### yanlingrpa.com/yanling/protocol/basic

- Package Name: `basic`
- Directory: `/basic`
- Interfaces: `0`
- Functions: `5`
- Structs: `6`

Public Functions:
- `MaxAreaRect(rectangles []Rect) Rect`: MaxAreaRect 返回面积最大的矩形
- `MergeAllRectangles(rectangles []Rect) Rect`: MergeAllRectangles 将所有矩形合并为一个包含所有矩形的最小矩形
- `MergeGroupRectangles(groupRects ...[]Rect) []Rect`: MergeGroupRectangles 将多个矩形组中的矩形进行各种组合方式合并
- `MergeOverlappingRectangles(rectangles []Rect) []Rect`: MergeOverlappingRectangles 将重叠的矩形合并为一个
- `MinAreaRect(rectangles []Rect) Rect`: MinAreaRect 返回面积最小的矩形

Exported Structs:
- `FPoint`: FPoint 表示一个二维浮点坐标点
- `OcrResult`: OcrResult 表示整张图像的 OCR 识别结果
- `OcrText`: OcrText 表示单条 OCR 文本识别结果
- `Point`: Point 表示一个二维整数坐标点
- `Rect`: Rect 表示一个矩形区域，定义了左上角坐标和宽高
- `Size`: Size 表示宽度和高度

### yanlingrpa.com/yanling/protocol/browser

- Package Name: `browser`
- Directory: `/browser`
- Interfaces: `4`
- Functions: `0`
- Structs: `0`

Public Interfaces:
- `BrowserElement`: BrowserElement 定义浏览器元素的通用操作接口
- `BrowserFramePage`: BrowserFramePage 定义浏览器页面或框架页的通用接口
- `BrowserTabPage`: BrowserTabPage 定义浏览器标签页接口
- `BrowserWindow`: BrowserWindow 定义浏览器窗口接口

### yanlingrpa.com/yanling/protocol/extention

- Package Name: `extention`
- Directory: `/extention`
- Interfaces: `1`
- Functions: `0`
- Structs: `0`

Public Interfaces:
- `VisionExtension`

### yanlingrpa.com/yanling/protocol/osgui

- Package Name: `osgui`
- Directory: `/osgui`
- Interfaces: `2`
- Functions: `0`
- Structs: `0`

Public Interfaces:
- `GuiWindow`: GuiWindow 定义了一个图形界面窗口的接口，提供了获取窗口信息、操作窗口、模拟输入等功能
- `Locator`: Locator 定义了一个定位器接口，提供了获取定位器信息、操作定位器、模拟输入等功能

### yanlingrpa.com/yanling/protocol/ossys

- Package Name: `ossys`
- Directory: `/ossys`
- Interfaces: `6`
- Functions: `0`
- Structs: `0`

Public Interfaces:
- `DeviceInfo`: DeviceInfo 定义设备信息查询能力
- `HttpClient`: HttpClient 定义 HTTP 请求与域名级配置能力
- `LocalFilesystem`: LocalFilesystem 定义本地文件系统操作能力
- `LocalStorage`: LocalStorage 定义本地键值存储能力
- `MonitorInfo`: MonitorInfo 定义显示器信息查询能力
- `ScriptLogger`: ScriptLogger 定义脚本运行日志能力

### yanlingrpa.com/yanling/protocol/script

- Package Name: `script`
- Directory: `/script`
- Interfaces: `2`
- Functions: `1`
- Structs: `8`

Public Interfaces:
- `ModuleRuntime`: ModuleRuntime 定义脚本运行时上下文能力
- `Subscriber`: Subscriber 定义事件订阅者信息

Public Functions:
- `ParseSpecifier(spec string) (Specifier, error)`: ParseSpecifier 解析模块标识符字符串

Exported Structs:
- `Event`: Event 定义本地事件总线的事件数据
- `GuiApplication`: GuiApplication 定义 GUI 应用配置
- `ModuleInfo`: ModuleInfo 定义脚本模组元信息
- `ScriptVariable`: ScriptVariable 定义脚本变量配置
- `Specifier`: Specifier 定义模块规范说明
- `UrlPermission`: UrlPermission 定义网络访问权限
- `WebApplication`: WebApplication 定义浏览器应用配置
- `YScript`: YScript 定义脚本配置总结构

