package appgui

import (
	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/internal/gui"
)

/*
* AppGuiWindow 定义移动应用窗口（屏幕）的接口，提供窗口信息、应用操作以及模拟触摸/按键输入相关方法。
 */
type AppGuiWindow interface {
	gui.GuiWindow

	/*
	* 获取应用的包名（Android）或绑定标识（iOS）。
	 */
	GetAppPackage() string
	/*
	* 获取当前活动名称（Android）或视图控制器名称（iOS）。
	 */
	GetAppActivity() string

	/*
	* 获取设备坐标系中的屏幕矩形区域。
	 */
	GetScreenRect() basic.Rect

	/*
	* 获取窗口内容区域的定位器。
	 */
	BodyLocator() (AppGuiLocator, error)
	/*
	* 获取指定矩形区域的定位器，并裁剪到窗口边界内。
	 */
	RectLocator(rect basic.Rect) (AppGuiLocator, error)

	/*
	* 获取相对于窗口的当前触摸位置；如果窗口内没有活动触摸，则返回 nil。
	 */
	GetWindowTouchPos() *basic.Point

	/*
	* 模拟按下硬件或系统按键。
	* key: 要按下的键，定义于 AppKey 类型中。
	 */
	PressKeys(keys ...AppKey) error
}
