package script

import (
	"time"

	"yanlingrpa.com/yanling/protocol/browser"
	"yanlingrpa.com/yanling/protocol/extension"
	"yanlingrpa.com/yanling/protocol/osgui"
	"yanlingrpa.com/yanling/protocol/ossys"
)

/*
* Subscriber 定义事件订阅者信息
* 用于标识订阅来源、主题与状态
 */
type Subscriber interface {
	/*
	* GetSpecifier 获取Topic所属模块标识
	 */
	GetSpecifier() string
	/*
	* GetTopic 获取订阅的事件主题
	 */
	GetTopic() string
	/*
	* IsActive 判断订阅是否仍然有效
	 */
	IsActive() bool
}

/*
* Event 定义本地事件总线的事件数据
 */
type Event struct {
	/*
	* Topic 事件主题
	 */
	Topic string
	/*
	* Data 事件数据
	 */
	Data any
	/*
	* OccuredAt 事件发生时间
	 */
	OccuredAt time.Time
}

/*
* EventHandler 定义事件处理函数签名
 */
type EventHandler func(event Event)

/*
* ModuleRuntime 定义脚本运行时上下文能力
* 该接口提供窗口访问、系统能力、缓存、变量与事件总线访问能力
 */
type ModuleRuntime interface {
	/*
	* HostSpecifier 获取启动引擎执行的入口脚本规范说明
	 */
	HostSpecifier() string
	/*
	* CurrentSpecifier 获取当前正在执行脚本的规范说明
	 */
	CurrentSpecifier() string
	/*
	* GuiWindow 根据窗口 ID 获取 GUI 窗口
	 */
	GuiWindow(id string) (osgui.GuiWindow, bool)
	/*
	* BrowserWindow 根据窗口 ID 获取浏览器窗口
	 */
	BrowserWindow(id string) (browser.BrowserWindow, bool)
	/*
	* DeviceInfo 获取设备信息能力接口
	 */
	DeviceInfo() ossys.DeviceInfo
	/*
	* Logger 获取脚本日志记录器
	 */
	Logger() ossys.ScriptLogger
	/*
	* Storage 获取项目本地存储接口
	 */
	Storage() ossys.LocalStorage
	/*
	* HttpClient 获取 HTTP 客户端接口
	 */
	HttpClient() ossys.HttpClient
	/*
	* FileSystem 获取本地文件系统接口
	 */
	FileSystem() ossys.LocalFilesystem
	/*
	* SetCacheData 在运行时临时存储任意数据
	 */
	SetCacheData(key string, value any)
	/*
	* GetCacheData 获取运行时临时存储数据
	 */
	GetCacheData(key string) (any, bool)
	/*
	* GetVariable 获取脚本变量值
	 */
	GetVariable(name string) (any, bool)
	/*
	* InvokeApi 调用本地其他模组的开放 API
	 */
	InvokeApi(specifier string, api string, args ...any) (any, error)
	/*
	* Subscribe 订阅本地其他模组的开放事件
	 */
	Subscribe(specifier string, topic string, handler EventHandler) (Subscriber, error)
	/*
	* Unsubscribe 取消订阅本地其他模组的开放事件
	 */
	Unsubscribe(subscriber Subscriber) error
	/*
	* Publish 发布事件到本地事件总线
	 */
	Publish(topic string, data any) error
	/*
	* Vision 获取视觉能力扩展接口
	 */
	Vision() extension.VisionExtension
}
