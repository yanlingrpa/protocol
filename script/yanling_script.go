package script

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

type VariableDataType string

const (
	VariableBoolean  VariableDataType = "boolean"
	VariableString   VariableDataType = "string"
	VariableFilePath VariableDataType = "filepath" // 文件路径
	VariableInteger  VariableDataType = "integer"
	VariableNumber   VariableDataType = "number" // 浮点数
	VariableJson     VariableDataType = "json"   // JSON对象
)

type ModuleInfo struct {
	Specifier     string   `json:"specifier"`                            // 模块标识符，包含域、所有者、名称和版本, 从 go.mod 文件中解析得到
	Package       string   `toml:"package" json:"package"`               // 脚本包名，有且仅有该包名的脚本可以被引入和调用
	Description   string   `toml:"description" json:"description"`       // 脚本描述，介绍脚本的功能和用途
	Tags          []string `toml:"tags" json:"tags"`                     // 脚本主题标签列表，如["web automation", "data extraction"]
	Website       string   `toml:"website" json:"website"`               // 脚本官方网站
	UpdateTime    string   `json:"update_time"`                          // 更新时间，通过git命令获取，最后一次提交时间，格式为RFC3339，如2024-01-02T15:04:05Z07:00
	Author        string   `toml:"author" json:"author"`                 // 作者
	Email         string   `toml:"email" json:"email"`                   // 联系方式
	License       string   `toml:"license" json:"license"`               // 许可证, 根据license文件或声明
	Devices       []string `toml:"devices" json:"devices"`               // 适用的设备列表，如["windows", "mac", "ubuntu"]
	EngineVersion string   `toml:"engine_version" json:"engine_version"` // 引擎版本要求
}

type GuiApplication struct {
	Id       string `toml:"id" json:"id"`             // GUI应用的唯一标识符
	Name     string `toml:"name" json:"name"`         // GUI应用名称
	Launcher string `toml:"launcher" json:"launcher"` // 启动器命令或路径
	Process  string `toml:"process" json:"process"`   // GUI窗口所属的进程名称
	Url      string `toml:"url" json:"url"`           // 打开GUI窗口的Url, 例如: weixin://launchapplet?appid=xxxx
}

type WebApplication struct {
	Id   string `toml:"id" json:"id"`     // 浏览器应用的唯一标识符
	Name string `toml:"name" json:"name"` // 浏览器应用名称
	Url  string `toml:"url" json:"url"`   // 浏览器应用的URL, 例如: https://www.google.com
}

type ScriptVariable struct {
	Name     string           `toml:"name" json:"name"`         // 变量名称
	Abstract string           `toml:"abstract" json:"abstract"` // 变量描述
	Type     VariableDataType `toml:"type" json:"type"`         // 变量类型（如string、int、bool, json等）
	Default  string           `toml:"default" json:"default"`   // 变量默认值
	Required bool             `toml:"required" json:"required"` // 是否必填
	Save     bool             `toml:"save" json:"save"`         // 是否保存到项目存储
}

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

// 网络访问权限
type UrlPermission struct {
	Url        string `toml:"url" json:"url"`               // URL 模式，支持通配符(*)和脚本变量${var_name}
	Permission string `toml:"permission" json:"permission"` // 访问权限，r,w,d,rw,rwd 等
	Abstract   string `toml:"abstract" json:"abstract"`     // 权限说明
}

type Argument struct {
	Name     string `json:"name"`     // 参数名称
	Type     string `json:"type"`     // 参数类型
	Abstract string `json:"abstract"` // 参数描述
	Example  string `json:"example"`  // 参数的JSON示例值
}

type ExportApi struct {
	Name       string     `json:"name"`       // 接口名称
	Abstract   string     `json:"abstract"`   // 接口描述
	Parameters []Argument `json:"parameters"` // 接口参数列表
	Result     Argument   `json:"result"`     // 接口返回值
}

type ExportTopic struct {
	Name     string   `json:"name"`     // 事件主题名称
	Abstract string   `json:"abstract"` // 事件主题描述
	Data     Argument `json:"data"`     // 事件数据结构定义
}

type Specifier struct {
	Domain  string `json:"domain"`  // 模块来源域，如github.com
	Owner   string `json:"owner"`   // 模块所有者，如github用户名
	Name    string `json:"name"`    // 模块名称，如utils
	Version string `json:"version"` // 模块版本要求，如v1.2.3
}

/**
 * String 返回完整的模块标识符字符串，格式为 "domain/owner/name@version" 或 "owner/name@version"，包含版本信息
 * 其中name部分如果包含/，也会被正确处理，例如 "domain/owner/name/subname@version"
 */
func (s Specifier) String() string {
	var result string

	// 构建名称部分
	if s.Domain != "" {
		result = fmt.Sprintf("%s/%s/%s", s.Domain, s.Owner, s.Name)
	} else {
		result = fmt.Sprintf("%s/%s", s.Owner, s.Name)
	}

	// 如果有版本，添加版本号
	if s.Version != "" {
		result = fmt.Sprintf("%s@%s", result, s.Version)
	}

	return result
}

/**
 * IsPseudoVersion 判断当前Specifier的版本是否为Go伪版本格式
 * Go伪版本格式: v0.0.0-yyyymmddhhmmss-commithash
 * 示例: v0.0.0-20231101134539-556fd59b42f6
 */
func (s Specifier) IsPseudoVersion() bool {
	if s.Version == "" {
		return false
	}

	// Go伪版本格式: v0.0.0-yyyymmddhhmmss-commithash
	// 示例: v0.0.0-20231101134539-556fd59b42f6
	if !strings.HasPrefix(s.Version, "v0.0.0-") {
		return false
	}

	// 移除 v0.0.0- 前缀
	suffix := strings.TrimPrefix(s.Version, "v0.0.0-")

	// 按 - 分割，应该得到 [时间戳, hash]
	parts := strings.Split(suffix, "-")
	if len(parts) != 2 {
		return false
	}

	timestamp, hash := parts[0], parts[1]

	// 检查时间戳: 14位数字 (yyyymmddhhmmss)
	if len(timestamp) != 14 {
		return false
	}
	if _, err := strconv.ParseInt(timestamp, 10, 64); err != nil {
		return false
	}

	// 检查hash: 十六进制字符串，长度10-12位
	if len(hash) < 10 || len(hash) > 12 {
		return false
	}
	if _, err := strconv.ParseInt(hash, 16, 64); err != nil {
		return false
	}

	return true
}

/**
 * ModulePath 返回模块路径字符串，格式为 "domain/owner/name@version" 或 "owner/name@version"，包含版本信息
 * 其中name部分如果包含/，也会被正确处理，例如 "domain/owner/name/subname@version"
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

/**
 * Identifier 返回模块标识符字符串，格式为 "domain/owner/name" 或 "owner/name"，不包含版本信息
 */
func (s Specifier) Identifier() string {
	var result string

	// 构建名称部分
	if s.Domain != "" {
		result = fmt.Sprintf("%s/%s/%s", s.Domain, s.Owner, s.Name)
	} else {
		result = fmt.Sprintf("%s/%s", s.Owner, s.Name)
	}

	return result
}

func ParseSpecifier(spec string) (Specifier, error) {
	var domain, owner, name, version string

	// 先按 @ 分割，分离 version
	atParts := strings.Split(spec, "@")
	if len(atParts) > 2 {
		return Specifier{}, fmt.Errorf("invalid specifier format: %s, multiple @ found", spec)
	}

	nameParts := atParts[0]
	if len(atParts) == 2 {
		version = atParts[1]
	}

	// 按 / 分割名称部分
	segments := strings.Split(nameParts, "/")

	// 根据段数判断 domain 是否存在
	switch len(segments) {
	case 0, 1:
		return Specifier{}, fmt.Errorf("invalid specifier format: %s, expected 'owner/name' or 'domain/owner/name'", spec)
	case 2:
		// owner/name（domain为空）
		owner = segments[0]
		name = segments[1]
	case 3:
		// domain/owner/name
		domain = segments[0]
		owner = segments[1]
		name = segments[2]
	default:
		domain = segments[0]
		owner = segments[1]
		name = strings.Join(segments[2:], "/") // 允许name中包含/，例如：domain/owner/name/subname
	}

	// 验证必须字段
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

type YScript struct {
	Module    ModuleInfo       `toml:"module" json:"module"`   // 脚本模组基本信息
	GuiApps   []GuiApplication `toml:"gui" json:"gui_apps"`    // 需要操作的GUI应用列表
	WebApps   []WebApplication `toml:"web" json:"web_apps"`    // 需要操作的Web应用列表
	Variables []ScriptVariable `toml:"var" json:"variables"`   // 脚本全局变量定义列表
	FilePerms []UrlPermission  `toml:"file" json:"file_perms"` // 文件系统授权列表
	ApiPerms  []UrlPermission  `toml:"api" json:"api_perms"`   // 网络接口授权列表
	Exports   []ExportApi      `json:"exports"`                // 导出接口列表
	Topics    []ExportTopic    `json:"topics"`                 // 可订阅事件主题列表
	Requires  []string         `json:"requires"`               // 模块依赖列表
}
