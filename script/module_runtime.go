package script

import (
	"time"

	"yanlingrpa.com/yanling/protocol/browser"
	"yanlingrpa.com/yanling/protocol/extension"
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
	* IPCInvoke calls an exposed API from another local IPC service.
	 */
	IPCInvoke(specifier string, api string, args ...any) (any, error)
	/*
	* Subscribe subscribes to an exposed event from another local IPC service or the current service.
	 */
	Subscribe(specifier string, topic string, handler EventHandler) (Subscriber, error)
	/*
	* Unsubscribe cancels a subscription to an exposed event from another local IPC service or the current service.
	 */
	Unsubscribe(subscriber Subscriber) error
	/*
	* Publish publishes an event to the local event bus.
	 */
	Publish(topic string, data any) error
	/*
	* Vision gets the vision capability extension interface.
	 */
	Vision() extension.VisionExtension
}
