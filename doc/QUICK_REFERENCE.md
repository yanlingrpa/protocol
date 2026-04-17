# 快速参考 - 协议 AI Prompt 方案

## 🎯 方案概述

```
协议代码 + 标准注释 
         ↓
    go generate
         ↓
 JSON + Markdown 文档
         ↓
    AI Agent Prompt
```

---

## 📋 三个关键文件说明

| 文件 | 目的 | 用途 |
|------|------|------|
| **PROTOCOL_COMMENT_GUIDE.md** | 编写标准 | 教你如何给协议添加注释 |
| **generateDocs.go** | 代码生成器 | 从注释自动生成文档 |
| **AI_PROMPT_TEMPLATE.md** | AI 指导 | 包含完整的 AI Agent 使用指南 |

---

## 🚀 快速开始（3分钟）

### 第1步：理解格式

查看 [PROTOCOL_COMMENT_GUIDE.md](./PROTOCOL_COMMENT_GUIDE.md) 中的示例，了解注释格式。

### 第2步：生成文档

```bash
cd e:\cylab\yanlingrpa\protocol
go generate ./symbols
```

### 第3步：给 AI 提示

将以下内容发送给 AI Agent：

```
[AI_PROMPT_TEMPLATE.md 全部内容]

[protocol_docs.json 全部内容]

现在请帮我编写一个脚本来：
[你的具体需求]
```

---

## 📝 协议快速查询

### 核心模块

```
basic/       - 数据结构 (Point, Rect, Size)
osgui/       - GUI操作 (窗口、定位、交互)
browser/     - 浏览器 (标签页、脚本)
ossys/       - 系统 (设备、网络、文件)
script/      - 脚本运行时和配置
extention/   - 扩展功能
```

### 最常用接口

```go
// GUI 操作
GuiWindow      // 窗口操作
Locator        // 元素定位

// 浏览器
BrowserWindow  // 浏览器窗口
BrowserTabPage // 浏览器标签页

// 系统
DeviceInfo     // 设备信息
HttpClient     // HTTP 请求
LocalStorage   // 本地存储
```

---

## 💡 常见操作

### 查找界面元素

```go
// 方式1：使用文本查找
locators, _ := window.BodyLocator().TextLocators("保存", 0.8)
button := locators[0]

// 方式2：使用图像查找
loc, _ := window.BodyLocator().ImageLocator("./button.png", 0.75)

// 方式3：浏览器 CSS 选择器
elem, _ := page.Query("button.submit")
```

### 点击和输入

```go
// 点击
locator.LeftClick()          // 左键
locator.RightClick()         // 右键  
locator.DoubleClick()        // 双击

// 输入
locator.TypeText("hello")    // 输入文本
window.PressKeys(ctrl, c)    // 按键
```

### 等待和验证

```go
// 等待元素出现
page.WaitForElement("input[name='email']", 10*time.Second)

// 等待条件
time.Sleep(2 * time.Second)

// 获取内容
text, _ := locator.GetText()
screenshot, _ := window.Snapshot(false)
```

---

## ⚙️ 配置建议

### 相似度 (0.1-1.0)

```
0.9+  ❌ 完全匹配（找不到）
0.75  ✅ 推荐（平衡）
0.6   🆗 宽松（变化较大）
0.5   ⚠️  极度宽松（特殊情况）
```

### 等待时间

```
0s     - UI 稳定
1-3s   - 快速操作
5-10s  - 页面加载
15-30s - 复杂操作
60s+   - 不推荐
```

---

## 🔍 协议生成时间表

| 步骤 | 时间 | 说明 |
|------|------|------|
| 添加注释 | 2-4小时 | 给核心接口补充标准注释 |
| 运行 generate | <1分钟 | `go generate ./symbols` |
| 获得文档 | 即时 | JSON + Markdown 同时生成 |
| 给 AI Prompt | 即时 | 复制到 ChatGPT/Claude |
| 验证脚本 | 5-10分钟 | 测试生成的 RPA 脚本 |

---

## 📦 输出文件

运行 `go generate ./symbols` 后获得：

```
e:\cylab\yanlingrpa\protocol\
├── protocol_docs.json      ← 用于 AI Agent
├── PROTOCOL_API.md         ← 用于人类阅读
└── symbols/
    └── generateDocs.go     ← 生成器（已配置）
```

---

## 🤖 给 AI 的 Prompt 模板

```
请基于以下 RPA 协议，生成自动化脚本：

[粘贴 AI_PROMPT_TEMPLATE.md]

[粘贴 protocol_docs.json]

## 具体任务：

1. [第一步操作]
2. [第二步操作]
3. [第三步操作]

要求：
- 使用协议中定义的接口
- 添加错误处理
- 添加相关日志
- 相似度设置为 0.75-0.8
```

---

## ✅ 成功标志

✓ 协议有完整的标准注释  
✓ 运行 `go generate` 无错误  
✓ 生成了 protocol_docs.json  
✓ AI 理解协议接口  
✓ AI 生成的脚本能正常运行  

---

## 📚 详细文档

- [完整实现指南](./IMPLEMENTATION_GUIDE.md) - 详细的实现步骤
- [注释规范](./PROTOCOL_COMMENT_GUIDE.md) - 如何写注释
- [AI 提示模板](./AI_PROMPT_TEMPLATE.md) - AI Agent 完整指南

---

## 🆘 快速故障排除

| 问题 | 解决方案 |
|------|--------|
| go generate 出错 | 检查 generateDocs.go 是否在 symbols/ 目录 |
| JSON 文件是空的 | 确保接口/结构体有注释，运行 go generate 前检查代码 |
| AI 生成的脚本不工作 | 1) 验证 protocol_docs.json 内容 2) 更新 Prompt 中的示例 |
| 找不到定位器 | 降低相似度 (0.75→0.7)，检查截图 |
| 脚本超时 | 增加等待时间或检查网络连接 |

---

## 🎓 推荐学习路径

1. **理解阶段** (10分钟)
   - 读一遍 AI_PROMPT_TEMPLATE.md 中的"关键接口说明"
   - 看一遍示例脚本

2. **实践阶段** (30分钟)
   - 在 2-3 个关键接口上添加标准注释
   - 运行 `go generate ./symbols`
   - 验证 protocol_docs.json 内容

3. **使用阶段** (5分钟)
   - 复制 AI_PROMPT_TEMPLATE.md + protocol_docs.json
   - 发送给 ChatGPT/Claude
   - 获取第一个自动化脚本

---

**版本**: 1.0.0  
**最后更新**: 2024-04-17  
**维护者**: 协议文档团队
