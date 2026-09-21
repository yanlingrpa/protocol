package osgui

import (
	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/internal/gui"
	"yanlingrpa.com/yanling/protocol/ossys"
)

/*
* OSGuiWindow 定义 GUI 窗口的接口，提供窗口信息、窗口操作以及模拟输入相关方法。
 */
type OSGuiWindow interface {
	gui.GuiWindow

	/*
	* 获取启动该窗口的进程可执行文件路径。
	 */
	GetInitiatorPath() string
	/*
	* 获取与该窗口本身关联的应用路径。
	 */
	GetWindowSchema() string
	/*
	* 获取原生窗口句柄。
	 */
	GetHwnd() uintptr
	/*
	* 获取屏幕坐标系中的窗口矩形区域。
	 */
	GetWindowRect() basic.Rect

	/*
	* 将窗口移动到指定位置。
	 */
	MoveTo(global_x, global_y int) (bool, error)
	/*
	* 调整窗口大小。
	 */
	ResizeTo(width, height int) (bool, error)

	/*
	* 获取窗口客户端区域的定位器。
	 */
	BodyLocator() (OSGuiLocator, error)
	/*
	* 获取指定矩形区域的定位器，并裁剪到窗口边界内。
	 */
	RectLocator(rect basic.Rect) (OSGuiLocator, error)

	/*
	* 获取相对于当前窗口的鼠标位置；若鼠标位于窗口外，则返回 nil。
	 */
	GetWindowCursorPos() *basic.Point
	/*
	* 获取相对于当前窗口的 IME 光标位置；若位于窗口外，则返回 nil。
	 */
	GetWindowCaretPos() *basic.Point

	/*
	* 模拟在当前窗口中输入键盘事件。
	* 该方法会先取消当前窗口中已获得焦点的控件焦点，
	* 然后将焦点设置到整个窗口，最后发送指定的按键序列。
	* keys: 要输入的按键序列（例如：ctrl + alt + del）。
	* 支持的按键定义于 Keyboard 类型。
	 */
	PressKeys(keys ...Keyboard) error

	/*
	* 获取包含该窗口的显示器信息。
	 */
	GetMonitor() ossys.MonitorInfo
}
