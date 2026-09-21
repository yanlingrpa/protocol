package gui

import "yanlingrpa.com/yanling/protocol/basic"

/*
* GuiLocator 定义跨平台通用 GUI 定位器的抽象接口。
 */
type GuiLocator interface {
	/*
	* 获取当前定位器所属的窗口。
	 */
	GuiWindow() GuiWindow
	/*
	* 转换为 map 对象。
	 */
	ToMap() map[string]any
	/*
	* 获取当前定位器的大小。
	 */
	GetSize() basic.Size
	/*
	* 返回当前定位器在屏幕上的绝对位置和尺寸。
	 */
	GetScreenRect() basic.Rect
	/*
	* 获取当前定位器相对于其父窗口的位置和尺寸。
	 */
	GetWindowRect() basic.Rect
	/*
	* 返回当前定位器相对于其 body 区域的位置和尺寸。
	 */
	GetBodyRect() basic.Rect

	/* 捕获当前定位器区域的快照并返回图片字节数据。
	* gray: 是否使用灰度模式。
	 */
	Snapshot(gray bool) ([]byte, error)

	/* 将定位器相对位置转换为绝对屏幕位置。
	* locator_point: 定位器内的相对位置。
	* return: 屏幕上的绝对位置。
	 */
	TransToScreen(locator_point *basic.Point) *basic.Point

	/*
	* 将定位器相对位置转换为相对于父窗口的相对位置。
	 */
	TransToWindow(locator_point *basic.Point) *basic.Point

	/*
	* 将定位器相对位置转换为相对于 client/body 区域的相对位置。
	 */
	TransToBody(locator_point *basic.Point) *basic.Point

	/*
	* 将绝对屏幕位置转换为当前定位器内的相对位置。
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	/*
	* 将相对于父窗口的位置转换为当前定位器内的相对位置。
	 */
	TransFromWindow(window_pos *basic.Point) *basic.Point

	/*
	* 将相对于 client/body 区域的位置转换为当前定位器内的相对位置。
	 */
	TransFromBody(body_pos *basic.Point) *basic.Point

	/*
	* 对当前定位器区域执行 OCR，并返回识别结果。
	 */
	OcrRead(confidence float64) (*basic.OcrResult, error)

	/*
	* 尝试聚焦当前定位器。
	 */
	Focus() error

	/*
	* 指示当前定位器是否处于可文本编辑状态。
	 */
	IsEditing() bool

	/*
	* 读取当前定位器中的文本内容。
	 */
	ReadText() (string, error)

	/*
	* 向当前定位器写入文本。
	 */
	WriteText(text string) error

	/*
	* 清除当前定位器中的文本内容。
	 */
	ClearText() error

	/*
	* 检查当前定位器区域是否支持垂直滚动。
	 */
	CanScrollVertical() bool
	/*
	* 在当前定位器区域执行垂直滚动。
	 */
	ScrollVertical(forward bool, distance int) error

	/*
	* 检查当前定位器区域是否支持水平滚动。
	 */
	CanScrollHorizontal() bool
	/*
	* 在当前定位器区域执行水平滚动。
	 */
	ScrollHorizontal(forward bool, distance int) error
}
