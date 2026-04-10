package script

import (
	"time"

	"github.com/yanlingrpa/protocol/browser"
	"github.com/yanlingrpa/protocol/extention"
	"github.com/yanlingrpa/protocol/osgui"
	"github.com/yanlingrpa/protocol/ossys"
)

type Subscriber interface {
	GetSpecifier() string
	GetTopic() string
	IsActive() bool
}

type Event struct {
	Topic     string    // 事件主题
	Data      any       // 事件数据
	OccuredAt time.Time // 事件发生时间
}

type EventHandler func(event Event)

type ModuleRuntime interface {
	HostSpecifier() string                                                              // 获取启动引擎执行的入口脚本的规范说明
	CurrentSpecifier() string                                                           // 获取当前正在执行的脚本的规范说明
	GuiWindow(id string) (osgui.GuiWindow, bool)                                        // 根据窗口ID获取GUI窗口
	BrowserWindow(id string) (browser.BrowserWindow, bool)                              // 根据窗口ID获取浏览器窗口
	DeviceInfo() ossys.DeviceInfo                                                       // 获取设备信息
	Logger() ossys.ScriptLogger                                                         // 获取脚本日志记录器
	Storage() ossys.LocalStorage                                                        // 获取项目存储
	HttpClient() ossys.HttpClient                                                       // 获取HTTP客户端
	FileSystem() ossys.LocalFilesystem                                                  // 获取文件系统
	SetCacheData(key string, value any)                                                 // 在运行时临时存储任意数据
	GetCacheData(key string) (any, bool)                                                // 获取运行时临时存储的任意数据
	GetVariable(name string) (any, bool)                                                // 获取脚本变量值
	InvokeApi(specifier string, api string, args ...any) (any, error)                   // 调用本地其他模组的开放API
	Subscribe(specifier string, topic string, handler EventHandler) (Subscriber, error) // 订阅本地其他模组的开放事件
	Unsubscribe(subscriber Subscriber) error                                            // 取消订阅本地其他模组的开放事件
	VisionExtension() extention.VisionExtension                                         // 获取视觉能力扩展接口
}
