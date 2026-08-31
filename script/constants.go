package script

import (
	"encoding/json"
	"fmt"
)

const (
	SCRIPT_ROUTE_INITIALIZE = "Initialize" // 脚本初始化路由
	SCRIPT_ROUTE_FINALIZE   = "Finalize"   // 脚本终结路由
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
	* VariableNumber floating-point float type.
	 */
	VariableNumber VariableDataType = "float"
	/*
	* VariableJson JSON object type.
	 */
	VariableJson VariableDataType = "json"
)

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
