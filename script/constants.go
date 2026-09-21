package script

import (
	"encoding/json"
	"fmt"
)

const (
	YSCRIPT_ROUTE_INITIALIZE = "Initialize" // 脚本初始化路由
	YSCRIPT_ROUTE_FINALIZE   = "Finalize"   // 脚本终结路由
)

/*
* VariableDataType 定义脚本变量的数据类型。
 */
type VariableDataType string

const (
	/*
	* VariableBoolean 布尔类型。
	 */
	VariableBoolean VariableDataType = "boolean"
	/*
	* VariableString 字符串类型。
	 */
	VariableString VariableDataType = "string"
	/*
	* VariableFilePath 文件路径类型。
	 */
	VariableFilePath VariableDataType = "filepath"
	/*
	* VariableInteger 整数类型。
	 */
	VariableInteger VariableDataType = "integer"
	/*
	* VariableNumber 浮点数类型。
	 */
	VariableNumber VariableDataType = "float"
	/*
	* VariableJson JSON 对象类型。
	 */
	VariableJson VariableDataType = "json"
)

/*
* ToString 将任意值转换为对应类型的字符串表示。
* 如果转换失败，则返回 fmt.Sprintf("%v", value)。
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
