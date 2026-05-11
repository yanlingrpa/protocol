package script

import (
	"time"

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
	* GetSpecifier gets the module specifier that owns the topic.
	 */
	GetSpecifier() string
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
	* HostSpecifier gets the specifier of the entry script launched by the engine.
	 */
	HostSpecifier() string
	/*
	* CurrentSpecifier gets the specifier of the currently executing script.
	 */
	CurrentSpecifier() string
	/*
	* GuiWindow gets a GUI window by window ID.
	 */
	GuiWindow(id string) (osgui.GuiWindow, bool)
	/*
	* BrowserWindow gets a browser window by window ID.
	 */
	BrowserWindow(id string) (browser.BrowserWindow, bool)
	/*
	* DeviceInfo gets the device information capability interface.
	 */
	DeviceInfo() ossys.DeviceInfo
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
	* SetCacheData temporarily stores arbitrary data at runtime.
	 */
	SetCacheData(key string, value any)
	/*
	* GetCacheData gets data temporarily stored at runtime.
	 */
	GetCacheData(key string) (any, bool)
	/*
	* GetVariable gets a script variable value.
	 */
	GetVariable(name string) (any, bool)

	/*
	* InvokeWorker calls an exposed method from another local IPC worker.
	* The specifier identifies the target worker, and the method is the name of the exposed method to call.
	* The dto is the data transfer object passed as an argument to the method.
	* It returns the result from the method call or an error if the invocation fails.
	 */
	InvokeWorker(specifier string, method string, dto any) (any, error)

	/*
	* Subscribe subscribes to an exposed event from another local IPC worker/yscript or the current worker/yscript.
	* The specifier identifies the target worker to subscribe to, and the topic is the name of the event topic to subscribe to.
	* The handler is the function that will be called when the event is published.
	* It returns a Subscriber object representing the subscription or an error if the subscription fails.
	 */
	Subscribe(specifier string, topic string, handler EventHandler) (Subscriber, error)
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
