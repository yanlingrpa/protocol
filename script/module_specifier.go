package script

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var domainPattern = regexp.MustCompile(`(?i)^[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)+$`)

var ownerPattern = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

/*
* ModuleName defines module name.
 */
type Specifier struct {
	/*
	* 域名，例如 github.com。
	 */
	Domain string `json:"domain"`
	/*
	* 模组所有者，例如 GitHub 用户名。
	 */
	Owner string `json:"owner"`
	/*
	* 项目名称，例如 utils。
	 */
	Name string `json:"name"`
	/*
	* 版本，例如 v1.0.0。
	 */
	Version string `json:"version"`
}

func (s *Specifier) String() string {
	if s.Domain != "" {
		return fmt.Sprintf("%s/%s/%s@%s", s.Domain, s.Owner, s.Name, s.Version)
	}
	return fmt.Sprintf("%s/%s@%s", s.Owner, s.Name, s.Version)
}

func (s *Specifier) GetId() string {
	if s.Domain != "" {
		return fmt.Sprintf("%s/%s/%s", s.Domain, s.Owner, s.Name)
	}
	return fmt.Sprintf("%s/%s", s.Owner, s.Name)
}

func (s *Specifier) GetPath(modPath string) string {
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

// ParseSpecifier 解析模块标识符。
//
// 模块路径支持 "所有者/项目名[/子路径...]" 和
// "域名/所有者/项目名[/子路径...]" 两种格式。版本号可省略，也可以使用
// "模块路径@版本号" 或 "模块路径 版本号" 的形式指定。
//
// 域名是可选项，但提供时必须包含至少一个点；域名标签只能由字母、数字和
// 中划线组成，且不能以中划线开头或结尾。所有者和项目名为必填项，
// 所有者只能包含 ASCII 字母、数字、下划线和中划线，项目名可以包含多级子路径。
// 格式无效时返回错误。
func ParseSpecifier(spec string) (*Specifier, error) {
	original := spec
	spec = strings.TrimSpace(spec)
	if spec == "" {
		return nil, fmt.Errorf("invalid module spec format: %s, empty spec", original)
	}

	var modulePart string
	var version string

	// 优先按 '@' 分隔模块路径和版本号，并允许分隔符两侧存在空白。
	if strings.Count(spec, "@") > 1 {
		return nil, fmt.Errorf("invalid module spec format: %s, multiple @ found", original)
	}
	if strings.Contains(spec, "@") {
		parts := strings.SplitN(spec, "@", 2)
		modulePart = strings.TrimSpace(parts[0])
		version = strings.TrimSpace(parts[1])
		if modulePart == "" || version == "" {
			return nil, fmt.Errorf("invalid module spec format: %s, expected 'moduleName@version'", original)
		}
	} else {
		// 未使用 '@' 时，允许通过空白分隔模块路径和版本号。
		parts := strings.Fields(spec)
		switch len(parts) {
		case 1:
			modulePart = parts[0]
		case 2:
			modulePart = parts[0]
			version = parts[1]
		default:
			return nil, fmt.Errorf("invalid module spec format: %s, expected 'moduleName' or 'moduleName version'", original)
		}
	}

	segments := strings.Split(modulePart, "/")
	if len(segments) < 2 {
		return nil, fmt.Errorf("invalid module spec format: %s, expected 'owner/name' or 'domain/owner/name'", original)
	}

	for _, segment := range segments {
		if segment == "" {
			return nil, fmt.Errorf("invalid module spec format: %s, contains empty path segment", original)
		}
	}

	domain := ""
	owner := ""
	name := ""

	// 当路径至少有三段且首段是合法域名时，将首段解析为域名；
	// 否则按“所有者/项目名”解析，并将后续路径段保留在项目名中。
	if len(segments) >= 3 && isDomainName(segments[0]) {
		domain = segments[0]
		owner = segments[1]
		name = strings.Join(segments[2:], "/")
	} else {
		owner = segments[0]
		name = strings.Join(segments[1:], "/")
	}

	if owner == "" || name == "" {
		return nil, fmt.Errorf("invalid module spec format: %s, owner and project name are required", original)
	}
	if !isValidOwner(owner) {
		return nil, fmt.Errorf("invalid module spec format: %s, owner must be a single word without '/'", original)
	}

	return &Specifier{Domain: domain, Owner: owner, Name: name, Version: version}, nil
}

// isDomainName 检查字符串是否为包含至少两级标签的合法域名。
func isDomainName(domain string) bool {
	return domainPattern.MatchString(domain)
}

// isValidOwner 检查所有者是否仅包含 ASCII 字母、数字、下划线和中划线。
func isValidOwner(owner string) bool {
	return ownerPattern.MatchString(owner)
}
