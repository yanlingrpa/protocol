package script

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

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
