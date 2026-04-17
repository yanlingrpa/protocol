# 协议文档生成和 AI Prompt 完整指南

## 总体方案

这个方案将你的 RPA 协议转换为 AI Agent 可理解的格式，通过以下三个步骤实现：

```
代码注释 ──go generate──> JSON/Markdown 文档 ──> AI Prompt 
```

---

## 第一步：规范代码注释

### 1.1 在协议接口和结构体上添加标准注释

遵循 [PROTOCOL_COMMENT_GUIDE.md](./PROTOCOL_COMMENT_GUIDE.md) 中的格式。

#### 示例改造前后

**改造前** (browser_window.go):
```go
type BrowserWindow interface {
	osgui.GuiWindow
	DefaultPage() BrowserTabPage
	CurrentPage() BrowserTabPage
	IDTabPage() BrowserTabPage
	NewTabPage(id string, url string) (BrowserTabPage, error)
}
```

**改造后**:
```go
// BrowserWindow 代表浏览器窗口对象
//
// 用途：提供浏览器窗口的管理操作，包括标签页切换、导航等
//
// 主要能力：
// - 获取默认页面和当前活跃页面
// - 创建新标签页
// - 继承GuiWindow的所有窗口操作能力
//
// 继承的接口：
// - osgui.GuiWindow（窗口操作）
type BrowserWindow interface {
	osgui.GuiWindow
	
	// DefaultPage 获取浏览器的首页（第一个标签页）
	//
	// 返回:
	//   BrowserTabPage: 首页标签页对象，无法获取时返回nil
	//
	// 示例:
	//   page := window.DefaultPage()
	DefaultPage() BrowserTabPage
	
	// CurrentPage 获取当前激活的标签页
	//
	// 返回:
	//   BrowserTabPage: 当前活跃标签页对象
	CurrentPage() BrowserTabPage
	
	// IDTabPage 通过ID获取指定的标签页
	//
	// 返回:
	//   BrowserTabPage: 对应ID的标签页对象
	IDTabPage() BrowserTabPage
	
	// NewTabPage 创建新的浏览器标签页
	//
	// 参数:
	//   id  - 标签页的唯一标识符，字符串类型
	//   url - 要打开的URL，字符串类型，如 "https://www.google.com"
	//
	// 返回:
	//   BrowserTabPage: 新创建的标签页对象
	//   error: 创建失败时返回错误信息
	//
	// 示例:
	//   page, err := window.NewTabPage("google", "https://www.google.com")
	//   if err != nil {
	//       log.Fatal("Failed to create tab:", err)
	//   }
	NewTabPage(id string, url string) (BrowserTabPage, error)
}
```

### 1.2 优先级补充注释的文件

根据重要性，按以下顺序补充注释：

1. **高优先级** (核心接口)
   - `osgui/osgui_window.go` - GuiWindow 接口
   - `osgui/osgui_locator.go` - Locator 接口
   - `browser/browser_window.go` - BrowserWindow 接口
   - `browser/browser_tabpage.go` - BrowserTabPage 接口
   - `script/yanling_script.go` - 模块和脚本定义

2. **中优先级** (数据结构)
   - `basic/basic_data_struct.go` - Point, Rect, Size 等
   - `script/yanling_script.go` - ModuleInfo, ScriptVariable 等

3. **低优先级** (辅助接口)
   - `ossys/` 模块中的其他接口
   - `extention/` 模块

---

## 第二步：生成文档

### 2.1 运行生成命令

```bash
# 进入项目目录
cd e:\cylab\yanlingrpa\protocol

# 运行 go generate 
go generate ./symbols
```

### 2.2 生成的输出文件

执行后会生成：

1. **protocol_docs.json** - 机器可读格式
   ```
   e:\cylab\yanlingrpa\protocol\protocol_docs.json
   ```
   
2. **PROTOCOL_API.md** - 人类可读格式
   ```
   e:\cylab\yanlingrpa\protocol\PROTOCOL_API.md
   ```

### 2.3 输出格式示例

**protocol_docs.json** 结构：
```json
{
  "name": "YanLing RPA Protocol",
  "version": "1.0.0",
  "description": "RPA执行协议 - 定义了RPA系统的核心接口和数据结构",
  "modules": [
    {
      "name": "osgui",
      "description": "GUI操作协议 - 提供GUI窗口、元素定位、交互操作等能力",
      "apis": [
        {
          "type": "interface",
          "name": "GuiWindow",
          "package": "osgui",
          "description": "窗口操作接口...",
          "methods": [
            {
              "name": "GetWindowTitle",
              "description": "获取窗口标题...",
              "parameters": [...],
              "returns": [...]
            }
          ]
        }
      ]
    }
  ]
}
```

---

## 第三步：给 AI Agent 提供 Prompt

### 3.1 组合完整的 AI Prompt

结合以下三个文件内容给 AI：

1. **AI_PROMPT_TEMPLATE.md** - 包含基本指导和最佳实践
2. **protocol_docs.json** - 自动生成的协议定义
3. **PROTOCOL_COMMENT_GUIDE.md** - 注释规范（可选）

### 3.2 Prompt 传递方式

#### 方式1：直接复制粘贴（简单）
```
User: [粘贴 AI_PROMPT_TEMPLATE.md 全部内容]
     [粘贴 protocol_docs.json 全部内容]
     现在请帮我编写一个脚本来...
```

#### 方式2：嵌入到系统提示（推荐）
如果使用 OpenAI API 或类似服务：
```python
system_prompt = open("AI_PROMPT_TEMPLATE.md").read() + "\n" + open("protocol_docs.json").read()

response = client.chat.completions.create(
    model="gpt-4",
    system_prompt=system_prompt,
    messages=[
        {"role": "user", "content": "请帮我编写一个脚本来..."}
    ]
)
```

#### 方式3：引用文件（最灵活）
```
当前项目使用 RPA 协议定义在:
- 核心规范: /path/to/protocol_docs.json
- 完整指南: /path/to/AI_PROMPT_TEMPLATE.md
- 注释规范: /path/to/PROTOCOL_COMMENT_GUIDE.md

请基于这些文件中的协议定义编写脚本...
```

---

## 工作流完整示例

### 场景：实现新功能并更新文档

```
┌─────────────────────────────────────────────────────────┐
│ 1. 新增功能到 osgui/osgui_locator.go                     │
│    添加: GetElementText() 方法                           │
└─────────────────────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────┐
│ 2. 为新方法添加标准注释                                  │
│    // GetElementText 获取定位器内的文本内容             │
│    // ...                                               │
└─────────────────────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────┐
│ 3. 运行: go generate ./symbols                          │
│    生成:                                                 │
│    - protocol_docs.json (更新)                          │
│    - PROTOCOL_API.md (更新)                             │
└─────────────────────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────┐
│ 4. 将更新的 protocol_docs.json 提供给 AI Agent          │
│    AI 现在了解新的 GetElementText() 方法                │
└─────────────────────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────┐
│ 5. AI 可以在生成的脚本中使用新方法                       │
└─────────────────────────────────────────────────────────┘
```

---

## 维护建议

### 长期维护流程

1. **定期同步** 
   - 每次 API 变更后运行 `go generate`
   - 提交更新的 JSON 文档

2. **版本控制**
   ```
   git add protocol_docs.json PROTOCOL_API.md
   git commit -m "Update protocol docs for v1.1.0"
   ```

3. **CI/CD 集成**
   在 GitHub Actions 或类似工具中添加：
   ```yaml
   - name: Generate Protocol Docs
     run: cd symbols && go generate
   - name: Check if docs changed
     run: git diff --exit-code protocol_docs.json || (echo "Docs out of sync"; exit 1)
   ```

4. **AI Agent 更新频率**
   - 每周或每次重大 API 更新时更新 AI Prompt
   - 维护一个版本化的 Prompt 库

### 最佳实践清单

- [ ] 所有公共接口都有标准注释
- [ ] 所有复杂参数都有示例说明
- [ ] 定期运行 `go generate` 保持文档同步
- [ ] 在 Git 中追踪 protocol_docs.json 变更
- [ ] AI Agent 使用最新生成的文档
- [ ] 文档中的示例代码定期验证可运行性
- [ ] 为新功能添加到 PROTOCOL_COMMENT_GUIDE 中
- [ ] 在变更日志中记录 API 变更

---

## 常见问题

### Q1: 注释改了但 go generate 出错？

**A**: 检查：
1. generateDocs.go 在 `symbols/` 目录中
2. 运行 `go mod tidy` 确保依赖完整
3. 查看具体错误信息中的行号

### Q2: JSON 文件太大，如何分割？

**A**: 可以修改 generateDocs.go 生成多个文件：
```go
// 按模块生成单独的 JSON
for _, module := range modules {
    filename := fmt.Sprintf("protocol_%s.json", module.Name)
    writeJSON(filename, module)
}
```

### Q3: AI 总是生成超出协议的代码？

**A**: 
1. 在 Prompt 中明确说明"仅使用协议中定义的接口"
2. 提供反例说明"不允许的操作"
3. 在示例中强调错误处理

### Q4: 如何处理协议版本兼容性？

**A**: 在 ProtocolDoc 结构中添加版本信息：
```json
{
  "version": "1.0.0",
  "breaking_changes": [
    "v1.0: NewTabPage 参数从 (url) 改为 (id, url)"
  ]
}
```

### Q5: 可以生成其他格式的文档吗？

**A**: 完全可以，修改 generateDocs.go：
```go
// 生成 YAML
writeYAML("protocol_docs.yaml", doc)

// 生成 CSV (用于表格)
writeCSV("protocol_docs.csv", doc)

// 生成 HTML (用于网页展示)
writeHTML("protocol_docs.html", doc)
```

---

## 下一步

1. **立即开始**
   ```bash
   # 从高优先级文件开始补充注释
   # 然后运行:
   go generate ./symbols
   ```

2. **测试 Prompt**
   - 用生成的 protocol_docs.json + AI_PROMPT_TEMPLATE.md
   - 要求 AI 生成一个简单脚本
   - 验证生成的脚本是否遵从协议

3. **迭代改进**
   - 收集 AI 的反馈
   - 根据反馈改进注释内容
   - 循环运行 `go generate` 直到满意

---

## 示例使用

创建文件 `example_prompt.txt`：

```
[从 AI_PROMPT_TEMPLATE.md 复制内容]

[从 protocol_docs.json 复制内容]

## 具体任务

请帮我编写一个 RPA 脚本，完成以下操作：

1. 打开 Chrome 浏览器，导航到 https://example.com
2. 填充登录表单（邮箱和密码）
3. 点击登录按钮
4. 等待页面加载完成
5. 提取页面上所有用户名
6. 保存到本地存储

要求：
- 使用 BrowserWindow 和 BrowserTabPage 接口
- 添加错误处理和超时控制
- 添加日志记录关键步骤
- 相似度设置为 0.8
```

然后复制到 ChatGPT 或其他 AI 服务，即可获得完整的 RPA 脚本！
