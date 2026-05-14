package script

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

/*
* PublicMethod defines the only valid public script method type with a fixed function signature.
* Any function that matches this signature will be exported as a public interface for external calls.
* T is the input parameter type, and R is the return value type.
 */
type PublicMethod[T, R any] func(rt ModuleRuntime, dto T) (R, error)

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
	* Specifier is the module identifier, including domain, owner, name, and version,
	* parsed from the go.mod file.
	 */
	Specifier string `json:"specifier"`
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
	* Process is the process name associated with the GUI window, used for window detection and management.
	 */
	Process string `json:"process"`
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

/*
* ScriptVariable defines script variable configuration.
 */
type ScriptVariable struct {
	/*
	* Name is the variable name.
	 */
	Name string `json:"name"`
	/*
	* Abstract is the variable description.
	 */
	Abstract string `json:"abstract"`
	/*
	* Type is the variable type, such as string, integer, boolean, or json.
	 */
	Type VariableDataType `json:"type"`
	/*
	* Default is the default value of the variable.
	 */
	Default string `json:"default"`
	/*
	* Required indicates whether the variable is required.
	 */
	Required bool `json:"required"`
	/*
	* Save indicates whether to save the variable to project storage.
	 */
	Save bool `json:"save"`
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

/*
* UrlPermission defines network access permissions.
 */
type UrlPermission struct {
	/*
	* Url is the URL pattern, supporting wildcard (*) and script variable ${var_name}.
	 */
	Url string `json:"url"`
	/*
	* Permission is the access permission, such as r, w, d, rw, or rwd.
	 */
	Permission string `json:"permission"`
	/*
	* Abstract is the permission description.
	 */
	Abstract string `json:"abstract"`
}

/*
* Specifier defines module specification information.
 */
type Specifier struct {
	/*
	* Domain is the source domain of the module, such as github.com.
	 */
	Domain string `json:"domain"`
	/*
	* Owner is the module owner, such as a GitHub username.
	 */
	Owner string `json:"owner"`
	/*
	* Name is the module name, such as utils.
	 */
	Name string `json:"name"`
	/*
	* Version is the required module version, such as v1.2.3.
	 */
	Version string `json:"version"`
}

/*
* String returns the full module specifier string in the format
* "domain/owner/name@version" or "owner/name@version".
* If the name part contains "/", it is also handled correctly,
* for example "domain/owner/name/subname@version".
 */
func (s Specifier) String() string {
	var result string

	/*
	* Build the name part.
	 */
	if s.Domain != "" {
		result = fmt.Sprintf("%s/%s/%s", s.Domain, s.Owner, s.Name)
	} else {
		result = fmt.Sprintf("%s/%s", s.Owner, s.Name)
	}

	/*
	* Append version if present.
	 */
	if s.Version != "" {
		result = fmt.Sprintf("%s@%s", result, s.Version)
	}

	return result
}

/*
* IsPseudoVersion checks whether the version of the current Specifier
* uses the Go pseudo-version format.
* Go pseudo-version format: v0.0.0-yyyymmddhhmmss-commithash.
 */
func (s Specifier) IsPseudoVersion() bool {
	if s.Version == "" {
		return false
	}

	/*
	* Go pseudo-version format: v0.0.0-yyyymmddhhmmss-commithash.
	 */
	if !strings.HasPrefix(s.Version, "v0.0.0-") {
		return false
	}

	/*
	* Remove the v0.0.0- prefix.
	 */
	suffix := strings.TrimPrefix(s.Version, "v0.0.0-")

	/*
	* Split by '-', should get [timestamp, hash].
	 */
	parts := strings.Split(suffix, "-")
	if len(parts) != 2 {
		return false
	}

	timestamp, hash := parts[0], parts[1]

	/*
	* Validate timestamp: 14 digits (yyyymmddhhmmss).
	 */
	if len(timestamp) != 14 {
		return false
	}
	if _, err := strconv.ParseInt(timestamp, 10, 64); err != nil {
		return false
	}

	/*
	* Validate hash: hexadecimal string with length 10-12.
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
* ModulePath generates a module path based on the given root directory.
* If the name part contains "/", it is joined as nested directories.
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
* Identifier returns the module identifier without version information.
* Format: "domain/owner/name" or "owner/name".
 */
func (s Specifier) Identifier() string {
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
* ParseSpecifier parses a module specifier string.
* It supports owner/name or domain/owner/name, optionally with @version.
 */
func ParseSpecifier(spec string) (Specifier, error) {
	var domain, owner, name, version string

	/*
	* First split by '@' to separate version.
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
	* Split the name part by '/'.
	 */
	segments := strings.Split(nameParts, "/")

	/*
	* Determine whether domain exists based on segment count.
	 */
	switch len(segments) {
	case 0, 1:
		return Specifier{}, fmt.Errorf("invalid specifier format: %s, expected 'owner/name' or 'domain/owner/name'", spec)
	case 2:
		/*
		* owner/name (domain is empty).
		 */
		owner = segments[0]
		name = segments[1]
	case 3:
		/*
		* domain/owner/name.
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
	* Validate required fields.
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
* YScript defines the top-level script configuration structure.
 */
type YScript struct {
	/*
	* Module is the basic script module information.
	 */
	Module ModuleInfo `json:"module"`
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
	 */
	Variables []ScriptVariable `json:"variables"`
	/*
	* FilePermissions is the list of filesystem permissions.
	 */
	FilePermissions []UrlPermission `json:"file_permissions"`
	/*
	* ApiPermissions is the list of network API permissions.
	 */
	ApiPermissions []UrlPermission `json:"api_permissions"`
	/*
	* ScriptDependencies is the list of script module specifiers imported by this script project.
	 */
	ScriptDependencies []string `json:"script_dependencies"`
	/*
	* WorkerDependencies is the list of IPC service dependencies.
	 */
	WorkerDependencies []string `json:"worker_dependencies"`
}
