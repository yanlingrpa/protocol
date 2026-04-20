package ossys

/*
* ScriptLogger 定义脚本运行日志能力
* 该接口提供按级别输出格式化日志的方法
 */
type ScriptLogger interface {
	/*
	* Debug 输出调试级别日志
	* format 为格式字符串，args 为格式化参数
	 */
	Debug(format string, args ...any)

	/*
	* Info 输出信息级别日志
	* format 为格式字符串，args 为格式化参数
	 */
	Info(format string, args ...any)

	/*
	* Warn 输出警告级别日志
	* format 为格式字符串，args 为格式化参数
	 */
	Warn(format string, args ...any)

	/*
	* Error 输出错误级别日志
	* format 为格式字符串，args 为格式化参数
	 */
	Error(format string, args ...any)
}
