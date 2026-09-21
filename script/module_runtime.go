package script

import (
	"yanlingrpa.com/yanling/protocol/appgui"
	"yanlingrpa.com/yanling/protocol/browser"
	"yanlingrpa.com/yanling/protocol/component"
	"yanlingrpa.com/yanling/protocol/osgui"
	"yanlingrpa.com/yanling/protocol/ossys"
)

/*
* Subscriber 定义事件订阅者信息。
* 用于标识订阅来源、主题和状态。
 */
type Subscriber interface {
	/*
	* GetModuleId 获取拥有该主题的模块 ID。
	 */
	GetModuleId() string
	/*
	* GetTopic 获取已订阅的事件主题。
	 */
	GetTopic() string
	/*
	* IsActive 表示订阅是否仍然处于激活状态。
	 */
	IsActive() bool
}

/*
* Event 定义本地事件总线中的事件数据。
 */
type Event struct {
	/*
	* ModuleId 是事件发布者模块 ID。
	 */
	ModuleId string
	/*
	* Topic 是事件主题。
	 */
	Topic string
	/*
	* Data 是经过 JSON 反序列化后的事件载荷。
	* 因为它表示为 `any`，调用方应根据 JSON 结构做类型断言，
	* 例如：对象使用 `map[string]any`，数组使用 `[]any`，原生类型使用 Go 基础类型。
	* 若要映射为强类型结构体，可使用 `JsonStruct(event.Data, &YourStruct{})`。
	 */
	Data any
	/*
	* Ts 是事件发生时间，使用 Unix 时间戳格式。
	 */
	Ts int64
}

/*
* EventHandler 定义事件处理函数签名。
* `event.Data` 字段是经过 JSON 反序列化后的 `any` 值。
* 若要将 `event.Data` 转换为强类型结构体，可使用 `JsonStruct(event.Data, &YourStruct{})`。
 */
type EventHandler func(event Event)

/*
* ModuleRuntime 定义脚本运行时上下文能力。
* 该接口提供对窗口、系统服务、缓存、变量以及本地事件总线的访问。
 */
type ModuleRuntime interface {
	/*
	* MainModuleId 获取入口脚本的模块 ID。
	 */
	MainModuleId() string
	/*
	* CurrentModuleId 获取当前正在执行脚本的模块 ID。
	 */
	CurrentModuleId() string
	/*
	* OsGuiWindow 根据窗口 ID 获取 OS GUI 窗口。
	 */
	OsGuiWindow(id string) (osgui.OSGuiWindow, bool)
	/*
	* BrowserWindow 根据窗口 ID 获取浏览器窗口。
	 */
	BrowserWindow(id string) (browser.BrowserWindow, bool)
	/*
	* MobileWindow 根据窗口 ID 获取移动端窗口。
	 */
	MobileWindow(id string) (appgui.AppGuiWindow, bool)
	/*
	* BrokerInfo 获取代理信息。
	 */
	BrokerInfo() ossys.BrokerInfo
	/*
	* Logger 获取脚本日志记录器。
	 */
	Logger() ossys.ScriptLogger
	/*
	* Storage 获取项目本地存储接口。
	 */
	Storage() ossys.LocalStorage
	/*
	* HttpClient 获取 HTTP 客户端接口。
	 */
	HttpClient() ossys.HttpClient
	/*
	* FileSystem 获取本地文件系统接口。
	 */
	FileSystem() ossys.LocalFilesystem
	/*
	* SetCacheData 在运行时临时存储数据。
	 */
	SetCacheData(key string, value string)
	/*
	* GetCacheData 获取运行时临时存储的数据。
	 */
	GetCacheData(key string) (string, bool)
	/*
	* GetWriteBackCache 获取当前脚本执行的回写缓存映射。
	* 回写缓存用于保存执行结束后需要写回脚本上下文的数据。
	 */
	GetWriteBackCache() map[string]string
	/*
	* StringVariable 获取字符串脚本变量值。
	 */
	StringVariable(name string) (string, bool)
	/*
	* IntegerVariable 获取整数脚本变量值。
	 */
	IntegerVariable(name string) (int, bool)
	/*
	* FloatVariable 获取浮点脚本变量值。
	 */
	FloatVariable(name string) (float64, bool)
	/*
	* BooleanVariable 获取布尔脚本变量值。
	 */
	BooleanVariable(name string) (bool, bool)
	/*
	* JsonVariable 获取 JSON 脚本变量值。
	 */
	JsonVariable(name string) (map[string]any, bool)
	/*
	* FilePathVariable 获取文件路径脚本变量值。
	 */
	FilePathVariable(name string) (string, bool)

	/*
	* Invoke 调用另一个本地 IPC worker 暴露的方法。
	* `moduleId` 标识目标 worker，对应该 worker 的模块名。
	* 注意：`moduleId` 不包含版本号，因为系统会根据 `go.mod` 文件确定版本。
	* `route` 是要调用的暴露方法名。
	* `dto`（数据传输对象）可以是原始类型，也可以是带 JSON 标签的结构体。
	* 返回值是方法调用后的 JSON 反序列化结果。
	* 因为方法签名使用 `any`，调用方应根据 JSON 结构做类型断言，
	* 例如：对象使用 `map[string]any`，数组使用 `[]any`，原生类型使用 Go 基础类型。
	* 若要将结果映射为强类型结构体，可使用 `JsonStruct(result, &YourStruct{})`。
	* 如果调用失败，将返回错误。
	 */
	InvokeWorker(workerId string, route string, dto any) (any, error)

	/*
	* Subscribe 订阅来自另一个本地 IPC worker/yscript 或当前 worker/yscript 暴露的事件。
	* `moduleId` 标识目标 worker/yscript，`topic` 是要订阅的事件主题名。
	* `handler` 是事件发布时调用的函数。
	* `handler` 接收到的 `event.Data` 是 JSON 反序列化后的 `any`，
	* 因此调用方应根据预期载荷结构进行类型断言。
	* 若要将 `event.Data` 映射为强类型结构体，可使用 `JsonStruct(event.Data, &YourStruct{})`。
	* 它返回一个表示订阅的 Subscriber 对象，若订阅失败则返回错误。
	 */
	Subscribe(moduleId string, topic string, handler EventHandler) (Subscriber, error)
	/*
	* Unsubscribe 取消对另一个本地 IPC worker/yscript 或当前 worker/yscript 暴露事件的订阅。
	* `subscriber` 是要取消订阅的 Subscriber 对象。
	 */
	Unsubscribe(subscriber Subscriber) error
	/*
	* Publish 向本地事件总线发布事件。
	* `topic` 是要发布的事件主题名，`data` 是事件载荷。
	* `data` 必须是可 JSON 序列化的值或对象。
	* 如果发布失败，则返回错误。
	 */
	Publish(topic string, data any) error
	/*
	* VisionWorker 获取视觉能力 worker 接口。
	 */
	VisionWorker() component.VisionWorker
}
