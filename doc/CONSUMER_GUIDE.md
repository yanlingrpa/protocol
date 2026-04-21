# .yanling 使用指南（面向 AI / 工具）

本指南用于说明：AI 代理与工具应如何消费 `.yanling` 目录下的 machine-first 元数据。

## 读取优先级

1. module.json（入口清单）
- 目的：提供全局清单与包级索引。
- 先读取它，用于发现 schema 版本、包列表、包文件路径与聚合统计信息。
- 关键字段：
  - schema_version
  - files.symbol_index_lite
  - files.symbol_index
  - files.package_dir
  - packages[].package_file

2. symbols.lite.json（轻量全局索引）
- 目的：以最小字段实现上下文友好的快速发现。
- 默认第二步读取，优先用于 AI/工具检索，避免上下文超限。
- 适用场景：
  - 符号路由（name/kind/package/package_file）
  - 源码跳转定位

3. symbols.json（完整全局索引）
- 目的：提供更完整的全局符号检索能力。
- 当你明确需要索引级别的 doc/signature/type_refs/method_count/field_count 时再读取。
- 适用场景：
  - name/qualified_name 检索
  - kind 过滤（interface/struct/function/const/...）
  - 源码位置与包路由
- 典型流程：查询 symbols.json -> 选出候选符号 -> 打开对应 package 文件。

4. packages/*.json（权威明细）
- 目的：提供每个包的完整符号明细。
- 最后按需读取，用于精确生成与校验。
- 包含：method/field 签名、params/results、type_refs、embeds、const_type/const_value、imports/dependencies。

## 推荐检索流程

1. 先加载 module.json，并缓存 package_file 映射。
2. 查询 symbols.lite.json，筛选候选符号。
3. 仅打开必要的 packages/*.json 获取最终上下文。
4. 如需更丰富索引信息，再查询 symbols.json。
5. 优先使用结构化字段（type_refs、params、results、methods、fields），避免依赖自由文本 doc。

## 稳定性说明

- 解析前必须先检查 schema_version。
- $schema 应指向本地规范 schema 文件（schema/yanling.machine-first.v1），以避免编辑器将远程地址判定为不受信任。
- source（file/line/column）是回溯源码的锚点。
- signature 是紧凑摘要，不是完整 AST 转储。

## 标准 Schema 位置

- schema/README.md
- schema/yanling.machine-first.v1/module.schema.json
- schema/yanling.machine-first.v1/symbols.schema.json
- schema/yanling.machine-first.v1/package.schema.json

这些 schema 文件可复用于任意 Go module 的 machine-first 元数据导出。

## 生成范围排除规则

生成器在扫描时会有意忽略以下顶层目录：
- .yanling
- doc
- cmd
- tests
- symbols

这些目录不属于可被 import 的模块，因此生成器不会生成它们下的符号。
