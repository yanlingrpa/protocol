package script

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/*
* ModuleInfo defines script module metadata.
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
* GuiApplication defines GUI application configuration.
 */
type GuiApplication struct {
	/*
	* Id is the unique identifier of the GUI application.
	 */
	Id string `json:"id"`
	/*
	* Name is the GUI application name.
	 */
	Name string `json:"name"`
	/*
	* Launcher is the launcher command or executable path, for example C:\Program Files\App\app.exe or /usr/bin/app.
	 */
	Launcher string `json:"launcher"`
	/*
	* Args is the command-line arguments passed to the launcher, for example ["--headless", "--no-sandbox"].
	 */
	Args []string `json:"args,omitempty"`
	/*
	* WorkDir is the working directory for the launched application.
	* If empty, defaults to the directory of the launcher executable.
	 */
	WorkDir string `json:"work_dir,omitempty"`
	/*
	* Env is a list of environment variables in the format KEY=VALUE, for example ["PATH=/usr/bin", "DEBUG=1"].
	 */
	Env []string `json:"env,omitempty"`
	/*
	* ProcessName is the process name associated with the GUI window, used for window detection and management.
	 */
	ProcessName string `json:"process_name"`
	/*
	* LaunchUri is the URI used to open the GUI window, for example weixin://launchapplet?appid=xxxx.
	* If provided, the application will be launched and this URI will be opened.
	 */
	LaunchUri string `json:"launch_uri,omitempty"`
	/*
	* Timeout is the timeout in milliseconds for launching the application.
	* If 0, no timeout limit is applied.
	 */
	Timeout int `json:"timeout,omitempty"`
	/*
	* WaitTime is the wait time in milliseconds after launching the application before returning.
	* Useful for allowing the GUI to fully initialize.
	 */
	WaitTime int `json:"wait_time,omitempty"`
	/*
	* WindowWidth is the preferred window width in pixels. If 0, uses default.
	 */
	WindowWidth int `json:"window_width,omitempty"`
	/*
	* WindowHeight is the preferred window height in pixels. If 0, uses default.
	 */
	WindowHeight int `json:"window_height,omitempty"`
}

func (app *GuiApplication) ToMap() map[string]any {
	return map[string]any{
		"id":            app.Id,
		"name":          app.Name,
		"launcher":      app.Launcher,
		"args":          app.Args,
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
* WebApplication defines browser application configuration.
 */
type WebApplication struct {
	/*
	* Id is the unique identifier of the browser application.
	 */
	Id string `json:"id"`
	/*
	* Name is the browser application name.
	 */
	Name string `json:"name"`
	/*
	* BrowserType is the browser type to use, for example "chrome", "firefox", "edge", "safari".
	* If empty, defaults to system default browser.
	 */
	BrowserType string `json:"browser_type,omitempty"`
	/*
	* Url is the browser application URL to open, for example https://www.google.com.
	 */
	Url string `json:"url"`
	/*
	* Incognito indicates whether to open the browser in private/incognito mode.
	 */
	Incognito bool `json:"incognito,omitempty"`
	/*
	* Args is a list of additional browser command-line arguments, for example ["--proxy-server=localhost:8080", "--disable-plugins"].
	 */
	Args []string `json:"args,omitempty"`
	/*
	* UserDataDir is the custom user data directory for the browser profile.
	* If empty, uses the default profile directory.
	 */
	UserDataDir string `json:"user_data_dir,omitempty"`
	/*
	* LoadTimeout is the timeout in milliseconds for page loading.
	* If 0, no timeout limit is applied.
	 */
	LoadTimeout int `json:"load_timeout,omitempty"`
	/*
	* WaitTime is the wait time in milliseconds after opening the URL before returning.
	* Useful for allowing the page to fully load.
	 */
	WaitTime int `json:"wait_time,omitempty"`
	/*
	* WindowWidth is the preferred window width in pixels. If 0, uses default.
	 */
	WindowWidth int `json:"window_width,omitempty"`
	/*
	* WindowHeight is the preferred window height in pixels. If 0, uses default.
	 */
	WindowHeight int `json:"window_height,omitempty"`
}

func (app *WebApplication) ToMap() map[string]any {
	return map[string]any{
		"id":            app.Id,
		"name":          app.Name,
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
* MobileApplication defines mobile application configuration for Android/iOS devices.
 */
type MobileApplication struct {
	/*
	* Id is the unique identifier of the mobile application.
	 */
	Id string `json:"id"`
	/*
	* Name is the mobile application name.
	 */
	Name string `json:"name"`
	/*
	* Package is the application package name, for example com.example.app.
	 */
	Package string `json:"package"`
	/*
	* Activity is the Activity class name to launch, for example .MainActivity.
	* Used with Package to form the component for adb command: adb shell am start -n package/activity.
	 */
	Activity string `json:"activity"`
	/*
	* Action is the Intent action, for example android.intent.action.MAIN or android.intent.action.VIEW.
	* If empty, defaults to android.intent.action.MAIN.
	 */
	Action string `json:"action,omitempty"`
	/*
	* Flags is the Intent flags as a comma-separated list, for example "FLAG_ACTIVITY_NEW_TASK,FLAG_ACTIVITY_CLEAR_TOP".
	* These are passed to adb with --flags parameter.
	 */
	Flags string `json:"flags,omitempty"`
	/*
	* Extras is a JSON object containing Intent extras (key-value pairs).
	* For example: {"key1": "value1", "key2": "value2"}.
	 */
	Extras string `json:"extras,omitempty"`
	/*
	* Timeout is the timeout in milliseconds for starting the application.
	* If 0, no timeout limit is applied.
	 */
	Timeout int `json:"timeout,omitempty"`
	/*
	* WaitTime is the wait time in milliseconds after launching the application before returning.
	* Useful for allowing the app to fully initialize.
	 */
	WaitTime int `json:"wait_time,omitempty"`
}

func (app *MobileApplication) ToMap() map[string]any {
	return map[string]any{
		"id":        app.Id,
		"name":      app.Name,
		"package":   app.Package,
		"activity":  app.Activity,
		"action":    app.Action,
		"flags":     app.Flags,
		"extras":    app.Extras,
		"timeout":   app.Timeout,
		"wait_time": app.WaitTime,
	}
}

func (app *MobileApplication) FromMap(data map[string]any) {
	app.Id = getMapString(data, "id")
	app.Name = getMapString(data, "name")
	app.Package = getMapString(data, "package")
	app.Activity = getMapString(data, "activity")
	app.Action = getMapString(data, "action")
	app.Flags = getMapString(data, "flags")
	app.Extras = getMapString(data, "extras")
	app.Timeout = getMapInt(data, "timeout")
	app.WaitTime = getMapInt(data, "wait_time")
}

/*
* ScriptVariable defines script variable configuration.
 */
type ScriptVariable struct {
	/*
	* Name is the variable name.
	 */
	Name string `json:"name"`
	/*
	* Description is the variable description.
	 */
	Description string `json:"description"`
	/*
	* Type is the variable type, such as string, integer, boolean, or json.
	 */
	Type VariableDataType `json:"type"`
	/*
	* DefaultValue is the default value of the variable.
	 */
	DefaultValue string `json:"default_value"`
	/*
	* Required indicates whether the variable is required.
	 */
	Required bool `json:"required"`
	/*
	* Save indicates whether to save the variable to project storage.
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
* Parse parses a string into a value of the corresponding data type.
* If parsing fails, it returns the original string.
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
	* Path is the path to the file or directory.
	 */
	Path string `json:"path"`
	/*
	* Permission is the access permission, such as r(read)/w(write)/d(delete)/x(execute)/a(all).
	 */
	Permission string `json:"permission"`
	/*
	* Description is the permission description.
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
* UrlPermission defines network access permissions.
 */
type UrlPermission struct {
	/*
	* Url is the URL pattern, supporting wildcard (*) and script variable ${var_name}.
	 */
	Url string `json:"url"`
	/*
	* HTTP Methods is the HTTP methods, such as GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, etc.
	 */
	Methods []string `json:"permission"`
	/*
	* Description is the permission description.
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
* YScript defines the top-level script configuration structure.
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
* ApplyVariables applies variable values to the script configuration.
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
