package script

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/*
* ModuleInfo 定义脚本模组元数据。
 */
type ModuleInfo struct {
	/*
	* 模组唯一标识，一般用repo路径
	 */
	Id string `json:"id"`
	/*
	* 模组名称，可读性名称
	 */
	Name string `json:"name"`
	/*
	* 模型代码版本号
	 */
	Version string `json:"version"`
	/*
	* 模组包名。只有该包下的脚本才能被导入和调用(非golang项目请忽略)
	 */
	Package string `json:"package"`
	/*
	* 模组描述，说明脚本功能和使用方法。
	 */
	Description string `json:"description"`
	/*
	* 模组标签，说明脚本的主题，例如 ["web automation", "data extraction"]。
	 */
	Tags []string `json:"tags"`
	/*
	* 模组官网。
	 */
	Website string `json:"website"`
	/*
	* 模组最后更新时间（最后一次提交时间），格式为RFC3339。
	 */
	UpdateTime string `json:"update_time"`
	/*
	* 模组作者。
	 */
	Author string `json:"author"`
	/*
	* 联系方式。
	 */
	Email string `json:"email"`
	/*
	* 模组许可协议，根据LICENSE文件或声明。
	 */
	License string `json:"license"`
	/*
	* 适用设备列表，例如 ["windows", "mac", "ubuntu"]。
	 */
	Devices []string `json:"devices"`
	/*
	* 所需引擎版本。
	 */
	EngineVersion string `json:"engine_version"`
}

func (info *ModuleInfo) GetSpecifier() *Specifier {
	spec, err := ParseSpecifier(fmt.Sprintf("%s@%s", info.Id, info.Version))
	if err != nil {
		return nil
	}
	return spec
}

func (info *ModuleInfo) ToMap() map[string]any {
	return map[string]any{
		"id":             info.Id,
		"name":           info.Name,
		"version":        info.Version,
		"package":        info.Package,
		"description":    info.Description,
		"tags":           info.Tags,
		"website":        info.Website,
		"update_time":    info.UpdateTime,
		"author":         info.Author,
		"email":          info.Email,
		"license":        info.License,
		"devices":        info.Devices,
		"engine_version": info.EngineVersion,
	}
}

func (info *ModuleInfo) FromMap(data map[string]any) {
	info.Name = getMapString(data, "name")
	info.Version = getMapString(data, "version")
	info.Package = getMapString(data, "package")
	info.Description = getMapString(data, "description")
	info.Tags = getMapStringSlice(data, "tags")
	info.Website = getMapString(data, "website")
	info.UpdateTime = getMapString(data, "update_time")
	info.Author = getMapString(data, "author")
	info.Email = getMapString(data, "email")
	info.License = getMapString(data, "license")
	info.Devices = getMapStringSlice(data, "devices")
	info.EngineVersion = getMapString(data, "engine_version")
}

/*
* GuiApplication 定义图形界面应用配置。
 */
type GuiApplication struct {
	/*
	* Id 是图形界面应用的唯一标识。
	 */
	Id string `json:"id"`
	/*
	* Name 是图形界面应用的名称。
	 */
	Name string `json:"name"`
	/*
	* Requirements 是图形界面应用启动后需要满足的条件，例如已登录。
	 */
	Requirements string `json:"requirements,omitempty"`
	/*
	* Launcher 是启动命令或可执行文件路径，例如 C:\Program Files\App\app.exe 或 /usr/bin/app。
	 */
	Launcher string `json:"launcher"`
	/*
	* Args 是传递给启动器的命令行参数，例如 ["--headless", "--no-sandbox"]。
	 */
	Args []string `json:"args,omitempty"`
	/*
	* WorkDir 是应用启动后的工作目录。
	* 为空时，默认为启动器可执行文件所在目录。
	 */
	WorkDir string `json:"work_dir,omitempty"`
	/*
	* Env 是环境变量列表，格式为 KEY=VALUE，例如 ["PATH=/usr/bin", "DEBUG=1"]。
	 */
	Env []string `json:"env,omitempty"`
	/*
	* ProcessName 是与图形界面窗口关联的进程名称，用于检测和管理窗口。
	 */
	ProcessName string `json:"process_name"`
	/*
	* LaunchUri 是用于打开图形界面窗口的 URI，例如 weixin://launchapplet?appid=xxxx。
	* 如果提供此值，将启动应用并打开该 URI。
	 */
	LaunchUri string `json:"launch_uri,omitempty"`
	/*
	* Timeout 是启动应用的超时时间，单位为毫秒。
	* 为 0 时不限制超时时间。
	 */
	Timeout int `json:"timeout,omitempty"`
	/*
	* WaitTime 是应用启动后到返回前的等待时间，单位为毫秒。
	* 可用于等待图形界面完成初始化。
	 */
	WaitTime int `json:"wait_time,omitempty"`
	/*
	* WindowWidth 是首选窗口宽度，单位为像素。为 0 时使用默认值。
	 */
	WindowWidth int `json:"window_width,omitempty"`
	/*
	* WindowHeight 是首选窗口高度，单位为像素。为 0 时使用默认值。
	 */
	WindowHeight int `json:"window_height,omitempty"`
}

func (app *GuiApplication) ToMap() map[string]any {
	return map[string]any{
		"id":            app.Id,
		"name":          app.Name,
		"launcher":      app.Launcher,
		"args":          app.Args,
		"requirements":  app.Requirements,
		"work_dir":      app.WorkDir,
		"env":           app.Env,
		"process_name":  app.ProcessName,
		"launch_uri":    app.LaunchUri,
		"timeout":       app.Timeout,
		"wait_time":     app.WaitTime,
		"window_width":  app.WindowWidth,
		"window_height": app.WindowHeight,
	}
}

func (app *GuiApplication) FromMap(data map[string]any) {
	app.Id = getMapString(data, "id")
	app.Name = getMapString(data, "name")
	app.Requirements = getMapString(data, "requirements")
	app.Launcher = getMapString(data, "launcher")
	app.Args = getMapStringSlice(data, "args")
	app.WorkDir = getMapString(data, "work_dir")
	app.Env = getMapStringSlice(data, "env")
	app.ProcessName = getMapString(data, "process_name")
	app.LaunchUri = getMapString(data, "launch_uri")
	app.Timeout = getMapInt(data, "timeout")
	app.WaitTime = getMapInt(data, "wait_time")
	app.WindowWidth = getMapInt(data, "window_width")
	app.WindowHeight = getMapInt(data, "window_height")
}

/*
* WebApplication 定义浏览器应用配置。
 */
type WebApplication struct {
	/*
	* Id 是浏览器应用的唯一标识。
	 */
	Id string `json:"id"`
	/*
	* Name 是浏览器应用的名称。
	 */
	Name string `json:"name"`
	/*
	* Requirements 指定浏览器应用打开后需要满足的条件，例如已登录。
	 */
	Requirements string `json:"requirements,omitempty"`
	/*
	* BrowserType 是要使用的浏览器类型，例如 "chrome"、"firefox"、"edge"、"safari"。
	* 为空时，默认使用系统默认浏览器。
	 */
	BrowserType string `json:"browser_type,omitempty"`
	/*
	* Url 是要打开的浏览器应用 URL，例如 https://www.google.com。
	 */
	Url string `json:"url"`
	/*
	* Incognito 表示是否以隐私/无痕模式打开浏览器。
	 */
	Incognito bool `json:"incognito,omitempty"`
	/*
	* Args 是浏览器的附加命令行参数列表，例如 ["--proxy-server=localhost:8080", "--disable-plugins"]。
	 */
	Args []string `json:"args,omitempty"`
	/*
	* UserDataDir 是浏览器配置文件的自定义用户数据目录。
	* 为空时使用默认配置文件目录。
	 */
	UserDataDir string `json:"user_data_dir,omitempty"`
	/*
	* LoadTimeout 是页面加载的超时时间，单位为毫秒。
	* 为 0 时不限制超时时间。
	 */
	LoadTimeout int `json:"load_timeout,omitempty"`
	/*
	* WaitTime 是打开 URL 后到返回前的等待时间，单位为毫秒。
	* 可用于等待页面完全加载。
	 */
	WaitTime int `json:"wait_time,omitempty"`
	/*
	* WindowWidth 是首选窗口宽度，单位为像素。为 0 时使用默认值。
	 */
	WindowWidth int `json:"window_width,omitempty"`
	/*
	* WindowHeight 是首选窗口高度，单位为像素。为 0 时使用默认值。
	 */
	WindowHeight int `json:"window_height,omitempty"`
}

func (app *WebApplication) ToMap() map[string]any {
	return map[string]any{
		"id":            app.Id,
		"name":          app.Name,
		"requirements":  app.Requirements,
		"browser_type":  app.BrowserType,
		"url":           app.Url,
		"incognito":     app.Incognito,
		"args":          app.Args,
		"user_data_dir": app.UserDataDir,
		"load_timeout":  app.LoadTimeout,
		"wait_time":     app.WaitTime,
		"window_width":  app.WindowWidth,
		"window_height": app.WindowHeight,
	}
}

func (app *WebApplication) FromMap(data map[string]any) {
	app.Id = getMapString(data, "id")
	app.Name = getMapString(data, "name")
	app.Requirements = getMapString(data, "requirements")
	app.BrowserType = getMapString(data, "browser_type")
	app.Url = getMapString(data, "url")
	app.Incognito = getMapBool(data, "incognito")
	app.Args = getMapStringSlice(data, "args")
	app.UserDataDir = getMapString(data, "user_data_dir")
	app.LoadTimeout = getMapInt(data, "load_timeout")
	app.WaitTime = getMapInt(data, "wait_time")
	app.WindowWidth = getMapInt(data, "window_width")
	app.WindowHeight = getMapInt(data, "window_height")
}

/*
* MobileApplication 定义 Android/iOS 设备上的移动应用配置。
 */
type MobileApplication struct {
	/*
	* Id 是移动应用的唯一标识。
	 */
	Id string `json:"id"`
	/*
	* Name 是移动应用的名称。
	 */
	Name string `json:"name"`
	/*
	* Requirements 是移动应用启动后需要满足的条件，例如已登录。
	 */
	Requirements string `json:"requirements,omitempty"`
	/*
	* Package 是应用包名，例如 com.example.app。
	 */
	Package string `json:"package"`
	/*
	* Activity 是要启动的 Activity 类名，例如 .MainActivity。
	* 与 Package 一起组成 adb 命令的组件参数：adb shell am start -n package/activity。
	 */
	Activity string `json:"activity"`
	/*
	* Action 是 Intent 动作，例如 android.intent.action.MAIN 或 android.intent.action.VIEW。
	* 为空时，默认为 android.intent.action.MAIN。
	 */
	Action string `json:"action,omitempty"`
	/*
	* Flags 是以逗号分隔的 Intent 标志列表，例如 "FLAG_ACTIVITY_NEW_TASK,FLAG_ACTIVITY_CLEAR_TOP"。
	* 这些标志通过 --flags 参数传递给 adb。
	 */
	Flags string `json:"flags,omitempty"`
	/*
	* Extras 是包含 Intent 附加数据（键值对）的 JSON 对象。
	* 例如：{"key1": "value1", "key2": "value2"}。
	 */
	Extras string `json:"extras,omitempty"`
	/*
	* Timeout 是启动应用的超时时间，单位为毫秒。
	* 为 0 时不限制超时时间。
	 */
	Timeout int `json:"timeout,omitempty"`
	/*
	* WaitTime 是应用启动后到返回前的等待时间，单位为毫秒。
	* 可用于等待应用完成初始化。
	 */
	WaitTime int `json:"wait_time,omitempty"`
}

func (app *MobileApplication) ToMap() map[string]any {
	return map[string]any{
		"id":           app.Id,
		"name":         app.Name,
		"requirements": app.Requirements,
		"package":      app.Package,
		"activity":     app.Activity,
		"action":       app.Action,
		"flags":        app.Flags,
		"extras":       app.Extras,
		"timeout":      app.Timeout,
		"wait_time":    app.WaitTime,
	}
}

func (app *MobileApplication) FromMap(data map[string]any) {
	app.Id = getMapString(data, "id")
	app.Name = getMapString(data, "name")
	app.Requirements = getMapString(data, "requirements")
	app.Package = getMapString(data, "package")
	app.Activity = getMapString(data, "activity")
	app.Action = getMapString(data, "action")
	app.Flags = getMapString(data, "flags")
	app.Extras = getMapString(data, "extras")
	app.Timeout = getMapInt(data, "timeout")
	app.WaitTime = getMapInt(data, "wait_time")
}

/*
* ScriptVariable 定义脚本变量配置。
 */
type ScriptVariable struct {
	/*
	* Name 是变量名称。
	 */
	Name string `json:"name"`
	/*
	* Description 是变量说明。
	 */
	Description string `json:"description"`
	/*
	* Type 是变量类型，例如 string、integer、boolean 或 json。
	 */
	Type VariableDataType `json:"type"`
	/*
	* DefaultValue 是变量的默认值。
	 */
	DefaultValue string `json:"default_value"`
	/*
	* Required 表示该变量是否必填。
	 */
	Required bool `json:"required"`
	/*
	* Save 表示是否将变量保存到项目存储中。
	 */
	Save bool `json:"save"`
}

func (sv *ScriptVariable) ToMap() map[string]any {
	return map[string]any{
		"name":          sv.Name,
		"description":   sv.Description,
		"type":          sv.Type,
		"default_value": sv.DefaultValue,
		"required":      sv.Required,
		"save":          sv.Save,
	}
}

func (sv *ScriptVariable) FromMap(data map[string]any) {
	sv.Name = getMapString(data, "name")
	sv.Description = getMapString(data, "description")
	sv.Type = getMapVariableDataType(data, "type")
	sv.DefaultValue = getMapString(data, "default_value")
	sv.Required = getMapBool(data, "required")
	sv.Save = getMapBool(data, "save")
}

/*
* Parse 将字符串解析为相应数据类型的值。
* 如果解析失败，则返回原始字符串。
 */
func (vdt VariableDataType) Parse(value string) any {
	switch vdt {
	case VariableBoolean:
		return value == "true" || value == "1" || strings.ToLower(value) == "yes" || strings.ToLower(value) == "on" || strings.ToLower(value) == "y"
	case VariableString:
	case VariableFilePath:
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

type PathPermission struct {
	/*
	* Path 是文件或目录的路径。
	 */
	Path string `json:"path"`
	/*
	* Permission 是访问权限，例如 r(读取)/w(写入)/d(删除)/x(执行)/a(全部)。
	 */
	Permission string `json:"permission"`
	/*
	* Description 是权限说明。
	 */
	Description string `json:"description"`
}

func (pp *PathPermission) ToMap() map[string]any {
	return map[string]any{
		"path":        pp.Path,
		"permission":  pp.Permission,
		"description": pp.Description,
	}
}

func (pp *PathPermission) FromMap(data map[string]any) {
	pp.Path = getMapString(data, "path")
	pp.Permission = getMapString(data, "permission")
	pp.Description = getMapString(data, "description")
}

/*
* UrlPermission 定义网络访问权限。
 */
type UrlPermission struct {
	/*
	* Url 是 URL 模式，支持通配符 (*) 和脚本变量 ${var_name}。
	 */
	Url string `json:"url"`
	/*
	* HTTP Methods 是 HTTP 请求方法，例如 GET、POST、PUT、DELETE、HEAD、OPTIONS、PATCH 等。
	 */
	Methods []string `json:"permission"`
	/*
	* Description 是权限说明。
	 */
	Description string `json:"description"`
}

func (up *UrlPermission) ToMap() map[string]any {
	return map[string]any{
		"url":         up.Url,
		"methods":     up.Methods,
		"description": up.Description,
	}
}

func (up *UrlPermission) FromMap(data map[string]any) {
	up.Url = getMapString(data, "url")
	up.Methods = getMapStringSlice(data, "methods")
	up.Description = getMapString(data, "description")
}

/*
* YScript 定义顶层脚本配置结构。
 */
type YScript struct {
	/*
	* 脚本模组的基本信息。
	 */
	Module ModuleInfo `json:"module"`

	/*
	* 脚本运行需要的图形界面应用列表。
	 */
	GuiApps []GuiApplication `json:"gui_apps"`
	/*
	* 脚本运行需要的网页应用列表。
	 */
	WebApps []WebApplication `json:"web_apps"`
	/*
	* 脚本运行需要的移动应用列表。
	 */
	MobileApps []MobileApplication `json:"mobile_apps"`
	/*
	* 脚本全局变量定义列表。
	* ${script_root} && ${data_root} 是保留变量，分别表示脚本根目录和数据目录，可在变量默认值和权限定义中使用。
	 */
	Variables []ScriptVariable `json:"variables"`
	/*
	* 文件系统权限列表。
	 */
	PathPermissions []PathPermission `json:"path_permissions"`
	/*
	* 网络URL权限列表。
	 */
	UrlPermissions []UrlPermission `json:"url_permissions"`
	/*
	* 脚本依赖的脚本模组列表，格式为 moduleId@version。
	 */
	ScriptDependencies []string `json:"script_dependencies"`
	/*
	* 脚本依赖的IPC服务模组列表，格式为 moduleId@version。
	 */
	WorkerDependencies []string `json:"worker_dependencies"`
}

/*
* ApplyVariables 将变量值应用到脚本配置中。
 */
func (s *YScript) ApplyVariables(variableValues map[string]any) {
	for i, app := range s.GuiApps {
		s.GuiApps[i].Launcher = apply_velocity_variables(app.Launcher, variableValues)
		s.GuiApps[i].WorkDir = apply_velocity_variables(app.WorkDir, variableValues)
		s.GuiApps[i].ProcessName = apply_velocity_variables(app.ProcessName, variableValues)
		s.GuiApps[i].LaunchUri = apply_velocity_variables(app.LaunchUri, variableValues)
	}
	for i, app := range s.WebApps {
		s.WebApps[i].Url = apply_velocity_variables(app.Url, variableValues)
		s.WebApps[i].UserDataDir = apply_velocity_variables(app.UserDataDir, variableValues)
	}
	for i, app := range s.MobileApps {
		s.MobileApps[i].Activity = apply_velocity_variables(app.Activity, variableValues)
		s.MobileApps[i].Package = apply_velocity_variables(app.Package, variableValues)
		s.MobileApps[i].Extras = apply_velocity_variables(app.Extras, variableValues)
	}
	for i, permission := range s.PathPermissions {
		s.PathPermissions[i].Path = apply_velocity_variables(permission.Path, variableValues)
	}
	for i, permission := range s.UrlPermissions {
		s.UrlPermissions[i].Url = apply_velocity_variables(permission.Url, variableValues)
	}
}

func (s *YScript) ToMap() map[string]any {
	data := map[string]any{
		"module":              s.Module.ToMap(),
		"script_dependencies": s.ScriptDependencies,
		"worker_dependencies": s.WorkerDependencies,
	}

	data["gui_apps"] = make([]map[string]any, 0, len(s.GuiApps))
	for _, app := range s.GuiApps {
		data["gui_apps"] = append(data["gui_apps"].([]map[string]any), app.ToMap())
	}

	data["web_apps"] = make([]map[string]any, 0, len(s.WebApps))
	for _, app := range s.WebApps {
		data["web_apps"] = append(data["web_apps"].([]map[string]any), app.ToMap())
	}

	data["mobile_apps"] = make([]map[string]any, 0, len(s.MobileApps))
	for _, app := range s.MobileApps {
		data["mobile_apps"] = append(data["mobile_apps"].([]map[string]any), app.ToMap())
	}

	data["variables"] = make([]map[string]any, 0, len(s.Variables))
	for _, variable := range s.Variables {
		data["variables"] = append(data["variables"].([]map[string]any), variable.ToMap())
	}

	data["path_permissions"] = make([]map[string]any, 0, len(s.PathPermissions))
	for _, permission := range s.PathPermissions {
		data["path_permissions"] = append(data["path_permissions"].([]map[string]any), permission.ToMap())
	}

	data["url_permissions"] = make([]map[string]any, 0, len(s.UrlPermissions))
	for _, permission := range s.UrlPermissions {
		data["url_permissions"] = append(data["url_permissions"].([]map[string]any), permission.ToMap())
	}

	return data
}

func (s *YScript) FromMap(data map[string]any) {
	s.Module = ModuleInfo{}
	s.Module.FromMap(getMapObject(data, "module"))

	s.GuiApps = make([]GuiApplication, 0)
	for _, item := range getMapObjectSlice(data, "gui_apps") {
		app := GuiApplication{}
		app.FromMap(item)
		s.GuiApps = append(s.GuiApps, app)
	}

	s.WebApps = make([]WebApplication, 0)
	for _, item := range getMapObjectSlice(data, "web_apps") {
		app := WebApplication{}
		app.FromMap(item)
		s.WebApps = append(s.WebApps, app)
	}

	s.MobileApps = make([]MobileApplication, 0)
	for _, item := range getMapObjectSlice(data, "mobile_apps") {
		app := MobileApplication{}
		app.FromMap(item)
		s.MobileApps = append(s.MobileApps, app)
	}

	s.Variables = make([]ScriptVariable, 0)
	for _, item := range getMapObjectSlice(data, "variables") {
		variable := ScriptVariable{}
		variable.FromMap(item)
		s.Variables = append(s.Variables, variable)
	}

	s.PathPermissions = make([]PathPermission, 0)
	for _, item := range getMapObjectSlice(data, "path_permissions") {
		permission := PathPermission{}
		permission.FromMap(item)
		s.PathPermissions = append(s.PathPermissions, permission)
	}

	s.UrlPermissions = make([]UrlPermission, 0)
	for _, item := range getMapObjectSlice(data, "url_permissions") {
		permission := UrlPermission{}
		permission.FromMap(item)
		s.UrlPermissions = append(s.UrlPermissions, permission)
	}

	s.ScriptDependencies = getMapStringSlice(data, "script_dependencies")
	s.WorkerDependencies = getMapStringSlice(data, "worker_dependencies")
}
