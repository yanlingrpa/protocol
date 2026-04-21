package script

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

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
	* VariableNumber floating-point number type.
	 */
	VariableNumber VariableDataType = "number"
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
	* Launcher is the launcher command or path.
	 */
	Launcher string `json:"launcher"`
	/*
	* Process is the process name associated with the GUI window.
	 */
	Process string `json:"process"`
	/*
	* Url is the URL used to open the GUI window, for example weixin://launchapplet?appid=xxxx.
	 */
	Url string `json:"url"`
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
	* Url is the browser application URL, for example https://www.google.com.
	 */
	Url string `json:"url"`
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
	* Variables is the list of global script variable definitions.
	 */
	Variables []ScriptVariable `json:"variables"`
	/*
	* FilePerms is the list of filesystem permissions.
	 */
	FilePerms []UrlPermission `json:"file_perms"`
	/*
	* ApiPerms is the list of network API permissions.
	 */
	ApiPerms []UrlPermission `json:"api_perms"`
}
