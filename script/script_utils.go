package script

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

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
* JsonStruct converts an input value into a target struct via JSON serialization/deserialization.
* The `jsonStructInterface` must be a non-nil pointer (for example: `&MyStruct{}`).
 */
func JsonStruct(obj any, jsonStructInterface any) error {
	if jsonStructInterface == nil {
		return fmt.Errorf("jsonStructInterface must be a non-nil pointer")
	}
	v := reflect.ValueOf(jsonStructInterface)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return fmt.Errorf("jsonStructInterface must be a non-nil pointer")
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %w", err)
	}
	if err := json.Unmarshal(bytes, jsonStructInterface); err != nil {
		return fmt.Errorf("failed to unmarshal into struct: %w", err)
	}
	return nil
}
