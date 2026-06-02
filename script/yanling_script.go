package script

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var domainPattern = regexp.MustCompile(`(?i)^[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)+$`)

var ownerPattern = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

/*
* VariableDataType defines data types for script variables.
 */
type VariableDataType string

const (
	/*
	* VariableBoolean boolean type.
	 */
	VariableBoolean VariableDataType = "boolean"
	/*
	* VariableString string type.
	 */
	VariableString VariableDataType = "string"
	/*
	* VariableFilePath file path type.
	 */
	VariableFilePath VariableDataType = "filepath"
	/*
	* VariableInteger integer type.
	 */
	VariableInteger VariableDataType = "integer"
	/*
	* VariableNumber floating-point float type.
	 */
	VariableNumber VariableDataType = "float"
	/*
	* VariableJson JSON object type.
	 */
	VariableJson VariableDataType = "json"
)

/*
* ModuleInfo defines script module metadata.
 */
type ModuleInfo struct {
	/*
	* Name is the module name.
	 */
	Name string `json:"name"`
	/*
	* Version is the module version.
	 */
	Version string `json:"version"`
	/*
	* Package is the script package name. Only scripts under this package can be imported and called.
	 */
	Package string `json:"package"`
	/*
	* Description describes script functionality and usage.
	 */
	Description string `json:"description"`
	/*
	* Tags is a list of script topic tags, for example ["web automation", "data extraction"].
	 */
	Tags []string `json:"tags"`
	/*
	* Website is the official website of the script.
	 */
	Website string `json:"website"`
	/*
	* UpdateTime is the last update time (last commit time), formatted as RFC3339.
	 */
	UpdateTime string `json:"update_time"`
	/*
	* Author is the script author.
	 */
	Author string `json:"author"`
	/*
	* Email is the contact address.
	 */
	Email string `json:"email"`
	/*
	* License is the license, based on the LICENSE file or declaration.
	 */
	License string `json:"license"`
	/*
	* Devices is the list of applicable devices, for example ["windows", "mac", "ubuntu"].
	 */
	Devices []string `json:"devices"`
	/*
	* EngineVersion is the required engine version.
	 */
	EngineVersion string `json:"engine_version"`
}

func (info *ModuleInfo) ToMap() map[string]any {
	return map[string]any{
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

func getMapString(data map[string]any, key string) string {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case string:
			return v
		case VariableDataType:
			return string(v)
		}
	}
	return ""
}

func getMapStringSlice(data map[string]any, key string) []string {
	value, ok := data[key]
	if !ok || value == nil {
		return []string{}
	}

	switch v := value.(type) {
	case []string:
		if v == nil {
			return []string{}
		}
		return append([]string{}, v...)
	case []any:
		items := make([]string, 0, len(v))
		for _, item := range v {
			if str, ok := item.(string); ok {
				items = append(items, str)
			}
		}
		return items
	}

	return []string{}
}

func getMapInt(data map[string]any, key string) int {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case int:
			return v
		case int8:
			return int(v)
		case int16:
			return int(v)
		case int32:
			return int(v)
		case int64:
			return int(v)
		case uint:
			return int(v)
		case uint8:
			return int(v)
		case uint16:
			return int(v)
		case uint32:
			return int(v)
		case uint64:
			return int(v)
		case float32:
			return int(v)
		case float64:
			return int(v)
		case json.Number:
			if i, err := v.Int64(); err == nil {
				return int(i)
			}
			if f, err := v.Float64(); err == nil {
				return int(f)
			}
		case string:
			if i, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
				return i
			}
		}
	}
	return 0
}

func getMapBool(data map[string]any, key string) bool {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case bool:
			return v
		case string:
			normalized := strings.TrimSpace(strings.ToLower(v))
			return normalized == "true" || normalized == "1" || normalized == "yes" || normalized == "on" || normalized == "y"
		case int:
			return v != 0
		case int8:
			return v != 0
		case int16:
			return v != 0
		case int32:
			return v != 0
		case int64:
			return v != 0
		case uint:
			return v != 0
		case uint8:
			return v != 0
		case uint16:
			return v != 0
		case uint32:
			return v != 0
		case uint64:
			return v != 0
		case float32:
			return v != 0
		case float64:
			return v != 0
		}
	}
	return false
}

func getMapVariableDataType(data map[string]any, key string) VariableDataType {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case VariableDataType:
			return v
		case string:
			return VariableDataType(v)
		}
	}
	return ""
}

func getMapObject(data map[string]any, key string) map[string]any {
	value, ok := data[key]
	if !ok || value == nil {
		return map[string]any{}
	}

	if item, ok := value.(map[string]any); ok {
		return item
	}

	return map[string]any{}
}

func getMapObjectSlice(data map[string]any, key string) []map[string]any {
	value, ok := data[key]
	if !ok || value == nil {
		return []map[string]any{}
	}

	switch v := value.(type) {
	case []map[string]any:
		if v == nil {
			return []map[string]any{}
		}
		return append([]map[string]any{}, v...)
	case []any:
		items := make([]map[string]any, 0, len(v))
		for _, item := range v {
			if mapped, ok := item.(map[string]any); ok {
				items = append(items, mapped)
			}
		}
		return items
	}

	return []map[string]any{}
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

/*
* ToString converts any value into the string representation of the corresponding type.
* If conversion fails, it returns fmt.Sprintf("%v", value).
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
	case VariableFilePath:
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
* ModuleName defines module name.
 */
type ModuleName struct {
	/*
	* Domain is the source domain of the module, such as github.com.
	 */
	Domain string `json:"domain"`
	/*
	* Owner is the module owner, such as a GitHub username.
	 */
	Owner string `json:"owner"`
	/*
	* Name is the project name, such as utils.
	 */
	Name string `json:"name"`
}

/*
* String returns the full module name string in the format
* "domain/owner/name" or "owner/name/v2".
* If the name part contains "/", it is also handled correctly,
* for example "domain/owner/name/subname@version".
 */
func (s ModuleName) String() string {
	var result string

	/*
	* Build the name part.
	 */
	if s.Domain != "" {
		result = fmt.Sprintf("%s/%s/%s", s.Domain, s.Owner, s.Name)
	} else {
		result = fmt.Sprintf("%s/%s", s.Owner, s.Name)
	}

	return result
}

/*
* ModulePath generates a module path based on the given root directory.
* If the name part contains "/", it is joined as nested directories.
 */
func (s ModuleName) ModulePath(modPath string, version string) string {
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
	if version != "" {
		path = path + "@" + version
	}
	return path
}

/*
* ParseModuleSpec parses a module spec string.
* Accepted formats are "owner/name[/subname...]" or "domain/owner/name[/subname...]",
* and an optional version suffix can be appended as "moduleName@version".
* The domain part is optional and must be a valid domain name when present.
* The owner part is required and must be a single word (no '/').
* The project name part is required and may contain '/'.
* It returns a ModuleName struct, version string, and an error if parsing fails.
 */
func ParseModuleSpec(spec string) (ModuleName, string, error) {
	original := spec
	spec = strings.TrimSpace(spec)
	if spec == "" {
		return ModuleName{}, "", fmt.Errorf("invalid module spec format: %s, empty spec", original)
	}

	var modulePart string
	var version string

	/*
	* Separate module name and version by '@', allowing spaces around both sides.
	 */
	if strings.Count(spec, "@") > 1 {
		return ModuleName{}, "", fmt.Errorf("invalid module spec format: %s, multiple @ found", original)
	}
	if strings.Contains(spec, "@") {
		parts := strings.SplitN(spec, "@", 2)
		modulePart = strings.TrimSpace(parts[0])
		version = strings.TrimSpace(parts[1])
		if modulePart == "" || version == "" {
			return ModuleName{}, "", fmt.Errorf("invalid module spec format: %s, expected 'moduleName@version'", original)
		}
	} else {
		modulePart = spec
	}

	segments := strings.Split(modulePart, "/")
	if len(segments) < 2 {
		return ModuleName{}, version, fmt.Errorf("invalid module spec format: %s, expected 'owner/name' or 'domain/owner/name'", original)
	}

	for _, segment := range segments {
		if segment == "" {
			return ModuleName{}, version, fmt.Errorf("invalid module spec format: %s, contains empty path segment", original)
		}
	}

	domain := ""
	owner := ""
	name := ""

	/*
	* If the first segment is a domain and there are at least 3 segments,
	* parse as domain/owner/name...; otherwise parse as owner/name....
	 */
	if len(segments) >= 3 && isDomainName(segments[0]) {
		domain = segments[0]
		owner = segments[1]
		name = strings.Join(segments[2:], "/")
	} else {
		owner = segments[0]
		name = strings.Join(segments[1:], "/")
	}

	if owner == "" || name == "" {
		return ModuleName{}, version, fmt.Errorf("invalid module spec format: %s, owner and project name are required", original)
	}
	if !isValidOwner(owner) {
		return ModuleName{}, version, fmt.Errorf("invalid module spec format: %s, owner must be a single word without '/'", original)
	}

	return ModuleName{Domain: domain, Owner: owner, Name: name}, version, nil
}

func isDomainName(domain string) bool {
	return domainPattern.MatchString(domain)
}

func isValidOwner(owner string) bool {
	return ownerPattern.MatchString(owner)
}

/*
* YScript defines the top-level script configuration structure.
 */
type YScript struct {
	/*
	* Module is the basic script module information.
	 */
	Module ModuleInfo `json:"module"`
	/*
	* Version is the module version.
	 */
	Version string `json:"version"`
	/*
	* GuiApps is the list of GUI applications to operate.
	 */
	GuiApps []GuiApplication `json:"gui_apps"`
	/*
	* WebApps is the list of web applications to operate.
	 */
	WebApps []WebApplication `json:"web_apps"`
	/*
	* MobileApps is the list of mobile applications to operate.
	 */
	MobileApps []MobileApplication `json:"mobile_apps"`
	/*
	* Variables is the list of global script variable definitions.
	* ${script_root} && ${data_root} are reserved variables representing the script root directory and data directory, which can be used in variable default values and permission definitions.
	 */
	Variables []ScriptVariable `json:"variables"`
	/*
	* PathPermissions is the list of filesystem permissions.
	 */
	PathPermissions []PathPermission `json:"path_permissions"`
	/*
	* UrlPermissions is the list of network URL permissions.
	 */
	UrlPermissions []UrlPermission `json:"url_permissions"`
	/*
	* ScriptDependencies is the list of script module names (moduleName@version) imported by this script project.
	 */
	ScriptDependencies []string `json:"script_dependencies"`
	/*
	* WorkerDependencies is the list of IPC service module names (moduleName@version) depended by this script project.
	 */
	WorkerDependencies []string `json:"worker_dependencies"`
}

func toVariableString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func apply_velocity_variables(text string, variableValues map[string]any) string {
	if text == "" || len(variableValues) == 0 {
		return text
	}
	if !strings.Contains(text, "${") {
		return text
	}

	result := text
	for key, value := range variableValues {
		placeholder := "${" + key + "}"
		if !strings.Contains(result, placeholder) {
			continue
		}

		result = strings.ReplaceAll(result, placeholder, toVariableString(value))

		if !strings.Contains(result, "${") {
			break
		}
	}
	return result
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
		"version":             s.Version,
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
	s.Version = getMapString(data, "version")

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
