package script

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

/*
* VariableDataType 定义脚本变量的数据类型
 */
type VariableDataType string

const (
	/*
	* VariableBoolean 布尔类型
	 */
	VariableBoolean VariableDataType = "boolean"
	/*
	* VariableString 字符串类型
	 */
	VariableString VariableDataType = "string"
	/*
	* VariableFilePath 文件路径类型
	 */
	VariableFilePath VariableDataType = "filepath"
	/*
	* VariableInteger 整数类型
	 */
	VariableInteger VariableDataType = "integer"
	/*
	* VariableNumber 浮点数类型
	 */
	VariableNumber VariableDataType = "number"
	/*
	* VariableJson JSON 对象类型
	 */
	VariableJson VariableDataType = "json"
)

/*
* ModuleInfo 定义脚本模组元信息
 */
type ModuleInfo struct {
	/*
	* Specifier 模块标识符，包含域、所有者、名称和版本，从 go.mod 文件中解析得到
	 */
	Specifier string `json:"specifier"`
	/*
	* Package 脚本包名，有且仅有该包名的脚本可以被引入和调用
	 */
	Package string `toml:"package" json:"package"`
	/*
	* Description 脚本描述，介绍脚本的功能和用途
	 */
	Description string `toml:"description" json:"description"`
	/*
	* Tags 脚本主题标签列表，如 ["web automation", "data extraction"]
	 */
	Tags []string `toml:"tags" json:"tags"`
	/*
	* Website 脚本官方网站
	 */
	Website string `toml:"website" json:"website"`
	/*
	* UpdateTime 更新时间，最后一次提交时间，格式为 RFC3339
	 */
	UpdateTime string `json:"update_time"`
	/*
	* Author 作者
	 */
	Author string `toml:"author" json:"author"`
	/*
	* Email 联系方式
	 */
	Email string `toml:"email" json:"email"`
	/*
	* License 许可证，根据 LICENSE 文件或声明
	 */
	License string `toml:"license" json:"license"`
	/*
	* Devices 适用的设备列表，如 ["windows", "mac", "ubuntu"]
	 */
	Devices []string `toml:"devices" json:"devices"`
	/*
	* EngineVersion 引擎版本要求
	 */
	EngineVersion string `toml:"engine_version" json:"engine_version"`
}

/*
* GuiApplication 定义 GUI 应用配置
 */
type GuiApplication struct {
	/*
	* Id GUI 应用的唯一标识符
	 */
	Id string `toml:"id" json:"id"`
	/*
	* Name GUI 应用名称
	 */
	Name string `toml:"name" json:"name"`
	/*
	* Launcher 启动器命令或路径
	 */
	Launcher string `toml:"launcher" json:"launcher"`
	/*
	* Process GUI 窗口所属的进程名称
	 */
	Process string `toml:"process" json:"process"`
	/*
	* Url 打开 GUI 窗口的 URL，例如 weixin://launchapplet?appid=xxxx
	 */
	Url string `toml:"url" json:"url"`
}

/*
* WebApplication 定义浏览器应用配置
 */
type WebApplication struct {
	/*
	* Id 浏览器应用的唯一标识符
	 */
	Id string `toml:"id" json:"id"`
	/*
	* Name 浏览器应用名称
	 */
	Name string `toml:"name" json:"name"`
	/*
	* Url 浏览器应用 URL，例如 https://www.google.com
	 */
	Url string `toml:"url" json:"url"`
}

/*
* ScriptVariable 定义脚本变量配置
 */
type ScriptVariable struct {
	/*
	* Name 变量名称
	 */
	Name string `toml:"name" json:"name"`
	/*
	* Abstract 变量描述
	 */
	Abstract string `toml:"abstract" json:"abstract"`
	/*
	* Type 变量类型，如 string、integer、boolean、json 等
	 */
	Type VariableDataType `toml:"type" json:"type"`
	/*
	* Default 变量默认值
	 */
	Default string `toml:"default" json:"default"`
	/*
	* Required 是否必填
	 */
	Required bool `toml:"required" json:"required"`
	/*
	* Save 是否保存到项目存储
	 */
	Save bool `toml:"save" json:"save"`
}

/*
* Parse 将字符串解析为对应的数据类型值
* 解析失败时返回原始字符串
 */
func (vdt VariableDataType) Parse(value string) any {
	switch vdt {
	case VariableBoolean:
		return value == "true" || value == "1" || strings.ToLower(value) == "yes" || strings.ToLower(value) == "on" || strings.ToLower(value) == "y"
	case VariableString:
		return value
	case VariableInteger:
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	case VariableNumber:
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		}
	case VariableJson:
		var result any
		if err := json.Unmarshal([]byte(value), &result); err == nil {
			return result
		}
	}

	return value
}

/*
* ToString 将任意值转换为对应类型的字符串表示
* 转换失败时返回 fmt.Sprintf("%v", value)
 */
func (vdt VariableDataType) ToString(value any) string {
	switch vdt {
	case VariableBoolean:
		if b, ok := value.(bool); ok {
			if b {
				return "true"
			}
			return "false"
		}
	case VariableString:
		if s, ok := value.(string); ok {
			return s
		}
	case VariableInteger:
		if i, ok := value.(int); ok {
			return fmt.Sprintf("%d", i)
		}
	case VariableNumber:
		if f, ok := value.(float64); ok {
			return fmt.Sprintf("%f", f)
		}
	case VariableJson:
		if jsonBytes, err := json.Marshal(value); err == nil {
			return string(jsonBytes)
		}
	}
	return fmt.Sprintf("%v", value)
}

/*
* UrlPermission 定义网络访问权限
 */
type UrlPermission struct {
	/*
	* Url URL 模式，支持通配符(*)和脚本变量 ${var_name}
	 */
	Url string `toml:"url" json:"url"`
	/*
	* Permission 访问权限，r、w、d、rw、rwd 等
	 */
	Permission string `toml:"permission" json:"permission"`
	/*
	* Abstract 权限说明
	 */
	Abstract string `toml:"abstract" json:"abstract"`
}

/*
* Specifier 定义模块规范说明
 */
type Specifier struct {
	/*
	* Domain 模块来源域，如 github.com
	 */
	Domain string `json:"domain"`
	/*
	* Owner 模块所有者，如 github 用户名
	 */
	Owner string `json:"owner"`
	/*
	* Name 模块名称，如 utils
	 */
	Name string `json:"name"`
	/*
	* Version 模块版本要求，如 v1.2.3
	 */
	Version string `json:"version"`
}

/*
* String 返回完整的模块标识符字符串，格式为 "domain/owner/name@version" 或 "owner/name@version"
* 其中 name 部分如果包含 /，也会被正确处理，例如 "domain/owner/name/subname@version"
 */
func (s Specifier) String() string {
	var result string

	/*
	* 构建名称部分
	 */
	if s.Domain != "" {
		result = fmt.Sprintf("%s/%s/%s", s.Domain, s.Owner, s.Name)
	} else {
		result = fmt.Sprintf("%s/%s", s.Owner, s.Name)
	}

	/*
	* 如果有版本，添加版本号
	 */
	if s.Version != "" {
		result = fmt.Sprintf("%s@%s", result, s.Version)
	}

	return result
}

/*
* IsPseudoVersion 判断当前 Specifier 的版本是否为 Go 伪版本格式
* Go 伪版本格式: v0.0.0-yyyymmddhhmmss-commithash
 */
func (s Specifier) IsPseudoVersion() bool {
	if s.Version == "" {
		return false
	}

	/*
	* Go 伪版本格式: v0.0.0-yyyymmddhhmmss-commithash
	 */
	if !strings.HasPrefix(s.Version, "v0.0.0-") {
		return false
	}

	/*
	* 移除 v0.0.0- 前缀
	 */
	suffix := strings.TrimPrefix(s.Version, "v0.0.0-")

	/*
	* 按 - 分割，应该得到 [时间戳, hash]
	 */
	parts := strings.Split(suffix, "-")
	if len(parts) != 2 {
		return false
	}

	timestamp, hash := parts[0], parts[1]

	/*
	* 检查时间戳: 14 位数字 (yyyymmddhhmmss)
	 */
	if len(timestamp) != 14 {
		return false
	}
	if _, err := strconv.ParseInt(timestamp, 10, 64); err != nil {
		return false
	}

	/*
	* 检查 hash: 十六进制字符串，长度 10-12 位
	 */
	if len(hash) < 10 || len(hash) > 12 {
		return false
	}
	if _, err := strconv.ParseInt(hash, 16, 64); err != nil {
		return false
	}

	return true
}

/*
* ModulePath 基于给定根目录生成模块路径
* 其中 name 部分如果包含 /，会被按多层目录拼接
 */
func (s Specifier) ModulePath(modPath string) string {
	path := modPath
	if s.Domain != "" {
		path = filepath.Join(path, s.Domain)
	}
	if s.Owner != "" {
		path = filepath.Join(path, s.Owner)
	}
	if s.Name != "" {
		splitName := strings.Split(s.Name, "/")
		for _, segment := range splitName {
			path = filepath.Join(path, segment)
		}
	}
	if s.Version != "" {
		path = path + "@" + s.Version
	}
	return path
}

/*
* Identifier 返回不包含版本信息的模块标识符
* 格式为 "domain/owner/name" 或 "owner/name"
 */
func (s Specifier) Identifier() string {
	var result string

	/*
	* 构建名称部分
	 */
	if s.Domain != "" {
		result = fmt.Sprintf("%s/%s/%s", s.Domain, s.Owner, s.Name)
	} else {
		result = fmt.Sprintf("%s/%s", s.Owner, s.Name)
	}

	return result
}

/*
* ParseSpecifier 解析模块标识符字符串
* 支持 owner/name 或 domain/owner/name 两种形式，可带 @version
 */
func ParseSpecifier(spec string) (Specifier, error) {
	var domain, owner, name, version string

	/*
	* 先按 @ 分割，分离 version
	 */
	atParts := strings.Split(spec, "@")
	if len(atParts) > 2 {
		return Specifier{}, fmt.Errorf("invalid specifier format: %s, multiple @ found", spec)
	}

	nameParts := atParts[0]
	if len(atParts) == 2 {
		version = atParts[1]
	}

	/*
	* 按 / 分割名称部分
	 */
	segments := strings.Split(nameParts, "/")

	/*
	* 根据段数判断 domain 是否存在
	 */
	switch len(segments) {
	case 0, 1:
		return Specifier{}, fmt.Errorf("invalid specifier format: %s, expected 'owner/name' or 'domain/owner/name'", spec)
	case 2:
		/*
		* owner/name（domain 为空）
		 */
		owner = segments[0]
		name = segments[1]
	case 3:
		/*
		* domain/owner/name
		 */
		domain = segments[0]
		owner = segments[1]
		name = segments[2]
	default:
		domain = segments[0]
		owner = segments[1]
		name = strings.Join(segments[2:], "/")
	}

	/*
	* 验证必须字段
	 */
	if owner == "" || name == "" {
		return Specifier{}, fmt.Errorf("invalid specifier format: %s, owner and name are required", spec)
	}

	return Specifier{
		Domain:  domain,
		Owner:   owner,
		Name:    name,
		Version: version,
	}, nil
}

/*
* YScript 定义脚本配置总结构
 */
type YScript struct {
	/*
	* Module 脚本模组基本信息
	 */
	Module ModuleInfo `toml:"module" json:"module"`
	/*
	* GuiApps 需要操作的 GUI 应用列表
	 */
	GuiApps []GuiApplication `toml:"gui" json:"gui_apps"`
	/*
	* WebApps 需要操作的 Web 应用列表
	 */
	WebApps []WebApplication `toml:"web" json:"web_apps"`
	/*
	* Variables 脚本全局变量定义列表
	 */
	Variables []ScriptVariable `toml:"var" json:"variables"`
	/*
	* FilePerms 文件系统授权列表
	 */
	FilePerms []UrlPermission `toml:"file" json:"file_perms"`
	/*
	* ApiPerms 网络接口授权列表
	 */
	ApiPerms []UrlPermission `toml:"api" json:"api_perms"`
	/*
	* Requires 模块依赖列表
	 */
	Requires []string `json:"requires"`
}
