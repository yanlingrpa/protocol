package ossys

/*
* ScriptLogger defines script runtime logging capabilities.
* This interface provides methods for outputting formatted logs by level.
 */
type ScriptLogger interface {
	/*
	* Debug outputs debug-level logs.
	* format is the format string, and args are formatting arguments.
	 */
	Debug(format string, args ...any)

	/*
	* Info outputs info-level logs.
	* format is the format string, and args are formatting arguments.
	 */
	Info(format string, args ...any)

	/*
	* Warn outputs warning-level logs.
	* format is the format string, and args are formatting arguments.
	 */
	Warn(format string, args ...any)

	/*
	* Error outputs error-level logs.
	* format is the format string, and args are formatting arguments.
	 */
	Error(format string, args ...any)
}
