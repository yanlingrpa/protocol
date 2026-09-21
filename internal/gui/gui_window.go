package gui

import "yanlingrpa.com/yanling/protocol/basic"

/*
* GuiWindow 定义跨平台通用 GUI 窗口的抽象接口。
 */
type GuiWindow interface {
	/*
	* 获取窗口的唯一标识符。
	 */
	GetID() string
	/*
	* 获取窗口标题。
	 */
	GetWindowTitle() string
	/*
	* 将当前窗口信息导出为键值对。
	 */
	ToMap() map[string]any
	/*
	* 获取屏幕坐标系中的客户端区域矩形。
	 */
	GetClientRect() basic.Rect

	/* 捕获当前窗口区域的快照并返回图片字节数据。
	* gray: 是否使用灰度模式。
	 */
	Snapshot(gray bool) ([]byte, error)

	/*
	* 激活窗口。
	 */
	Activate() (bool, error)
	/*
	* 取消激活窗口。
	 */
	DeActivate() (bool, error)

	/* 将窗口相对位置转换为绝对屏幕位置。
	* window_pos: 窗口内的相对位置。
	* return: 屏幕上的绝对位置。
	 */
	TransToScreen(window_pos *basic.Point) *basic.Point

	/* 将绝对屏幕位置转换为窗口相对位置。
	* screen_pos: 屏幕上的绝对位置。
	* return: 窗口内的相对位置；如果该点在窗口外，则返回 nil。
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	/*
	* 读取剪贴板文本。
	 */
	ReadClipboard() (string, error)
	/*
	* 向剪贴板写入文本。
	 */
	WriteClipboard(text string) error

	/*
	* 关闭窗口。
	 */
	Close() error
}
