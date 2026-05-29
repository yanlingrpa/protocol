package script

import (
	"time"

	"yanlingrpa.com/yanling/protocol/appgui"
	"yanlingrpa.com/yanling/protocol/browser"
	"yanlingrpa.com/yanling/protocol/component"
	"yanlingrpa.com/yanling/protocol/osgui"
	"yanlingrpa.com/yanling/protocol/ossys"
)

/*
* Subscriber defines event subscriber information.
* It is used to identify subscription source, topic, and status.
 */
type Subscriber interface {
	/*
	* GetModule gets the module name that owns the topic.
	 */
	GetModule() string
	/*
	* GetTopic gets the subscribed event topic.
	 */
	GetTopic() string
	/*
	* IsActive indicates whether the subscription is still active.
	 */
	IsActive() bool
}

/*
* Event defines event data for the local event bus.
 */
type Event struct {
	/*
	* Module is the topic publisher module name.
	 */
	Module string
	/*
	* Topic is the event topic.
	 */
	Topic string
	/*
	* Data is the event payload.
	 */
	Data any
	/*
	* OccuredAt is the event occurrence time.
	 */
	OccuredAt time.Time
}

/*
* EventHandler defines the event handler function signature.
 */
type EventHandler func(event Event)

/*
* ModuleRuntime defines script runtime context capabilities.
* This interface provides access to windows, system services, cache, variables,
* and the local event bus.
 */
type ModuleRuntime interface {
	/*
	* MainModule gets the module name of the entry point.
	 */
	MainModule() string
	/*
	* CurrentModule gets the module name of the currently executing script.
	 */
	CurrentModule() string
	/*
	* OsGuiWindow gets an OS GUI window by window ID.
	 */
	OsGuiWindow(id string) (osgui.OSGuiWindow, bool)
	/*
	* BrowserWindow gets a browser window by window ID.
	 */
	BrowserWindow(id string) (browser.BrowserWindow, bool)
	/*
	* MobileWindow gets a mobile window by window ID.
	 */
	MobileWindow(id string) (appgui.AppGuiWindow, bool)
	/*
	* BrokerInfo gets the broker information.
	 */
	BrokerInfo() ossys.BrokerInfo
	/*
	* Logger gets the script logger.
	 */
	Logger() ossys.ScriptLogger
	/*
	* Storage gets the project local storage interface.
	 */
	Storage() ossys.LocalStorage
	/*
	* HttpClient gets the HTTP client interface.
	 */
	HttpClient() ossys.HttpClient
	/*
	* FileSystem gets the local filesystem interface.
	 */
	FileSystem() ossys.LocalFilesystem
	/*
	* SetCacheData temporarily stores data at runtime.
	 */
	SetCacheData(key string, value string)
	/*
	* GetCacheData gets data temporarily stored at runtime.
	 */
	GetCacheData(key string) (string, bool)
	/*
	* StringVariable gets a string script variable value.
	 */
	StringVariable(name string) (string, bool)
	/*
	* IntegerVariable gets an integer script variable value.
	 */
	IntegerVariable(name string) (int, bool)
	/*
	* FloatVariable gets a float script variable value.
	 */
	FloatVariable(name string) (float64, bool)
	/*
	* BooleanVariable gets a boolean script variable value.
	 */
	BooleanVariable(name string) (bool, bool)
	/*
	* JsonVariable gets a JSON script variable value.
	 */
	JsonVariable(name string) (map[string]any, bool)
	/*
	* FilePathVariable gets a file path script variable value.
	 */
	FilePathVariable(name string) (string, bool)

	/*
	* InvokeWorker calls an exposed method from another local IPC worker.
	* The `module` identifies the target worker, which corresponds to the module name of the worker.
	* Note: The `module` does not include the version number, as the system determines the version based on the `go.mod` file.
	* The `method` is the name of the exposed method to call.
	* The `dto` (data transfer object) can be either a primitive type or a struct annotated with JSON tags.
	* The return value is always a JSON string representing the result of the method call.
	* If the invocation fails, an error is returned.
	 */
	InvokeWorker(module string, method string, dto any) (string, error)

	/*
	* Subscribe subscribes to an exposed event from another local IPC worker/yscript or the current worker/yscript.
	* The module identifies the target worker to subscribe to, and the topic is the name of the event topic to subscribe to.
	* The handler is the function that will be called when the event is published.
	* It returns a Subscriber object representing the subscription or an error if the subscription fails.
	 */
	Subscribe(module string, topic string, handler EventHandler) (Subscriber, error)
	/*
	* Unsubscribe cancels a subscription to an exposed event from another local IPC worker/yscript or the current worker/yscript.
	* The subscriber is the Subscriber object representing the subscription to cancel.
	 */
	Unsubscribe(subscriber Subscriber) error
	/*
	* Publish publishes an event to the local event bus.
	* The topic is the name of the event topic to publish, and the data is the event payload.
	* It returns an error if the publication fails.
	 */
	Publish(topic string, data any) error
	/*
	* VisionWorker gets the vision capability worker interface.
	 */
	VisionWorker() component.VisionWorker
	/*
	* OcrWorker gets the OCR capability worker interface.
	 */
	OcrWorker() component.OcrWorker
}
