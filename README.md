# yanling protocol

`yanlingrpa.com/yanling/protocol` 是 Yanling RPA 的协议层模块，定义了脚本运行时与各类能力边界（浏览器、桌面 GUI、系统能力、扩展能力等）的 Go 接口与数据结构。

该模块的目标是：

- 提供稳定、可被实现方和调用方共同依赖的协议定义
- 让脚本引擎与能力提供方通过统一接口解耦
- 为 AI/工具提供 machine-first 的结构化自述产物（`.yanling`）

## 模块信息

- 模块名：`yanlingrpa.com/yanling/protocol`
- Go 版本：`1.22`

## 目录说明

- `basic/`：基础数据结构（点、矩形、OCR 结果等）
- `browser/`：浏览器自动化协议（窗口、页面、元素）
- `osgui/`：桌面 GUI 自动化协议（窗口、定位器、键鼠能力）
- `ossys/`：系统能力协议（文件、存储、HTTP、设备、日志等）
- `script/`：脚本运行时协议与脚本元数据结构
- `extention/`：扩展能力协议（如视觉能力）
- `symbols/`：Yaegi 符号文件目录，用于脚本在 Yaegi 环境执行时的默认导入（不参与 machine-first 对外 API 扫描）
- `cmd/generate-symbols/`：Yaegi 符号文件生成器
- `cmd/generate-yanling/`：machine-first 自述文件生成器
- `schema/`：machine-first 标准 schema 定义
- `doc/`：补充文档

## 生成 Yaegi symbols（脚本默认导入）

在仓库根目录执行：

```bash
go run ./cmd/generate-symbols
```

该命令会调用 `yaegi extract`，按模块内目标包生成 `symbols/` 下的符号文件。

用途说明：

- `symbols/` 下文件用于 Yaegi 运行脚本时的默认导入能力。
- 这套文件服务于脚本执行期，不等同于 `.yanling/symbols.lite.json` 或 `.yanling/symbols.json`。

## 生成 machine-first 自述

在仓库根目录执行：

```bash
go run ./cmd/generate-yanling
```

会生成 `.yanling/` 目录，主要产物包括：

- `module.json`：全局入口清单
- `symbols.lite.json`：轻量全局索引（优先用于 AI 检索）
- `symbols.json`：完整全局索引
- `packages/*.json`：按包拆分的权威明细

## AI / 工具推荐读取顺序

建议按以下顺序加载，避免上下文超限：

1. `module.json`
2. `symbols.lite.json`
3. 按需加载对应 `packages/*.json`
4. 仅在需要时补充 `symbols.json`

完整说明见：`doc/CONSUMER_GUIDE.md`

## 标准 Schema

本仓库将 machine-first 规范沉淀在 `schema/` 目录，可复用于任意 Go 模块：

- `schema/yanling.machine-first.v1/module.schema.json`
- `schema/yanling.machine-first.v1/symbols.schema.json`
- `schema/yanling.machine-first.v1/package.schema.json`

说明文档见：`schema/README.md`

## 生成范围排除

为确保导出内容聚焦于“可被 import 的模块 API”，生成器会忽略以下顶层目录：

- `.yanling`
- `doc`
- `cmd`
- `tests`
- `symbols`
- `schema`


