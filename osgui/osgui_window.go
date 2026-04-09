package osgui

import (
	"github.com/yanlingrpa/protocol/basic"
	"github.com/yanlingrpa/protocol/ossys"
)

type GuiWindow interface {
	// 获取窗口唯一标识符
	GetID() string
	// 启动窗口的程序路径
	GetInitiatorPath() string
	// 窗口自身的程序路径
	GetWindowSchema() string
	// 获取窗口句柄
	GetHwnd() uintptr
	// 获取窗口标题
	GetWindowTitle() string
	// 输出当前窗口的信息为键值对
	ToMap() map[string]any
	// 获取窗口自身在屏幕上的位置和大小
	GetWindowRect() basic.Rect
	// 获取窗口客户区在屏幕上的位置和大小
	GetClientRect() basic.Rect

	/* 对当前窗体所在区域进行截图，返回图片的字节数组
	* gray: 是否为灰度图
	 */
	Snapshot(gray bool) ([]byte, error)

	// 移动窗口到指定位置
	MoveTo(global_x, global_y int) (bool, error)
	// 调整窗口大小
	ResizeTo(width, height int) (bool, error)
	// 激活窗口
	Activate() (bool, error)
	// 取消激活窗口
	DeActivate() (bool, error)

	// 获取窗口的客户区定位器
	BodyLocator() (Locator, error)
	// 获取指定区域的定位器，仅取得该区域在窗体内的部分
	RectLocator(rect basic.Rect) (Locator, error)

	/* 将窗体内的相对位置转换为屏幕上的绝对位置
	* window_pos: 窗体内的相对位置
	* return: 屏幕上的绝对位置
	 */
	TransToScreen(window_pos *basic.Point) *basic.Point

	/* 将屏幕上的绝对位置转换为窗体内的相对位置
	* screen_pos: 屏幕上的绝对位置
	* return: 窗体内的相对位置, 如果点不在窗体内则返回nil
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	// 获取鼠标在当前窗体内的相对位置, 如果鼠标不在窗体内则返回nil
	GetWindowCursorPos() *basic.Point
	// 获取当前输入法光标在当前窗体内的相对位置, 如果光标不在窗体内则返回nil
	GetWindowCaretPos() *basic.Point

	/**
	* 模拟键盘输入
	* keys: 要输入的按键序列 （例如：ctrl + alt + del）
	* 支持的按键见Keyboard类型定义
	 */
	PressKeys(keys ...Keyboard) error

	// 读取剪贴板内容
	ReadClipboard() (string, error)
	// 写入剪贴板内容
	WriteClipboard(text string) error

	// 获取窗口所在的显示器信息
	GetMonitor() ossys.MonitorInfo

	// 关闭窗口
	Close() error
}
