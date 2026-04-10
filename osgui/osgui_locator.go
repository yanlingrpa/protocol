package osgui

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
)

type Locator interface {
	// 输出当前定位器的信息为键值对
	ToMap() map[string]any
	// 获取当前定位器的尺寸
	GetSize() basic.Size
	// 获取当前定位器在屏幕上的位置和大小
	GetScreenRect() basic.Rect
	// 获取当前定位器在所属窗体内的位置和大小
	GetWindowRect() basic.Rect
	// 获取当前定位器在客户区内的位置和大小
	GetBodyRect() basic.Rect
	/* 对当前定位器所在区域进行截图，返回图片的字节数组
	* gray: 是否为灰度图
	 */
	Snapshot(gray bool) ([]byte, error)

	/* 将定位器内的相对位置转换为屏幕上的绝对位置
	* locator_point: 定位器内的相对位置
	* return: 屏幕上的绝对位置
	 */
	TransToScreen(locator_point *basic.Point) *basic.Point

	/* 将定位器内的相对位置转换为所属窗体内的相对位置
	* locator_point: 定位器内的相对位置
	* return: 窗体内的相对位置
	 */
	TransToWindow(locator_point *basic.Point) *basic.Point

	/* 将定位器内的相对位置转换为客户区内的相对位置
	* locator_point: 定位器内的相对位置
	* return: 客户区内的相对位置
	 */
	TransToBody(locator_point *basic.Point) *basic.Point

	/* 将屏幕上的绝对位置转换为当前定位器内的相对位置
	* screen_pos: 屏幕上的绝对位置
	* return: 定位器内的相对位置, 如果点不在定位器内则返回nil
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	/* 将所属窗体内的相对位置转换为当前定位器内的相对位置
	* window_pos: 窗体内的相对位置
	* return: 定位器内的相对位置, 如果点不在定位器内则返回nil
	 */
	TransFromWindow(window_pos *basic.Point) *basic.Point

	/* 将客户区内的相对位置转换为当前定位器内的相对位置
	* body_pos: 客户区内的相对位置
	* return: 定位器内的相对位置, 如果点不在定位器内则返回nil
	 */
	TransFromBody(body_pos *basic.Point) *basic.Point

	/* 获取当前定位器的子定位器，仅取得该区域与当前定位器边界的交集部分
	* locator_point: 子定位器在当前定位器内的相对位置
	* size: 子定位器的尺寸
	 */
	SubLocator(locator_point basic.Point, size basic.Size) Locator

	/* 在当前定位器所在区域内查找包含相似图片的子定位器
	* image: 图片相对与当前project的路径/URL/Base64编码字符串
	* sim: 相似度，取值范围0.1~1.0，值越大表示越相似
	 */
	ImageLocator(image string, sim float32) (Locator, error)

	/* 在当前定位器所在区域内查找包含相似图片的子定位器
	* image: 图片相对与当前project的路径/URL/Base64编码字符串
	* sim: 相似度，取值范围0.1~1.0，值越大表示越相似
	 */
	ImageLocators(image string, sim float32) ([]Locator, error)

	/* 在当前定位器所在区域内查找包含指定文本的子定位器
	* texts: 需要查找的文本列表，支持多个文本, 多个文本时表示同时包含
	 */
	TextLocator(texts ...string) (Locator, error)

	/* 在当前定位器所在区域内查找包含指定文本的子定位器
	* texts: 需要查找的文本列表，支持多个文本, 多个文本时表示同时包含
	 */
	TextLocators(texts ...string) ([]Locator, error)

	/* 在当前定位器所在区域内查找指定图形形态的子定位器
	* shape: 指定的图形形态
	* min_size: 查找的子定位器的最小尺寸
	* max_size: 查找的子定位器的最大尺寸
	 */
	ShapeLocator(shape GraphicShape, min_size basic.Size, max_size basic.Size) (Locator, error)

	/* 在当前定位器所在区域内查找指定图形形态的子定位器
	* shape: 指定的图形形态
	* min_size: 查找的子定位器的最小尺寸
	* max_size: 查找的子定位器的最大尺寸
	 */
	ShapeLocators(shape GraphicShape, min_size basic.Size, max_size basic.Size) ([]Locator, error)

	/* 对当前定位器所在区域进行OCR识别，返回识别结果
	* confidence: 识别的置信度，取值范围0.1~1.0，值越大表示越准确
	 */
	Ocr(confidence float64) (*basic.OcrResult, error)

	/* 将鼠标移动到定位器内的指定位置
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器的中心位置
	 */
	MouseMove(locator_point *basic.Point) error
	/**
	* 当前定位器获取焦点
	* 如果鼠标在当前定位器内则认为已经获取焦点
	* 否则将鼠标移动到当前定位器中心并点击以获取焦点
	 */
	Focus() error
	/* 在定位器内的指定位置进行鼠标点击
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器内的鼠标位置，
	* 如果鼠标不在当前定位器内，则取定位器中心点
	 */
	Click(locator_point *basic.Point) error
	/* 在定位器内的指定位置进行鼠标双击
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器内的鼠标位置，
	* 如果鼠标不在当前定位器内，则取定位器中心点
	 */
	DoubleClick(locator_point *basic.Point) error
	/* 在定位器内的指定位置进行鼠标右键点击
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器内的鼠标位置，
	* 如果鼠标不在当前定位器内，则取定位器中心点
	 */
	RightClick(locator_point *basic.Point) error

	// 获取鼠标在当前定位器内的相对位置, 没有则返回nil
	GetLocatorCursorPos() *basic.Point
	// 获取当前输入法光标在当前定位器内的相对位置，没有则返回nil
	GetLocatorCaretPos() *basic.Point

	// 当前定位器包含正在闪烁的输入光标，表示可以输入文本
	IsEditing() bool
	/* 等待当前定位器进入可以输入文本的状态，直到超时
	* timeout: 超时时间
	* return: 如果在超时时间内进入可以输入文本的状态则返回true，否则返回false
	 */
	WaitForEditing(timeout time.Duration) bool
	// 读取当前定位器内的文本内容
	ReadText() (string, error)
	// 向当前定位器内输入文本内容，等同于先点击定位器，然后输入文本
	WriteText(text string) error
	// 清空当前定位器内的文本内容, 等同于先点击定位器，然后全选文本，最后删除文本
	ClearText() error

	// 当前定位器所在区域可以垂直滚动
	IsVerticalScroller() bool
	/**
	* 在当前定位器所在区域内垂直滚动
	* up: 是否向上滚动，true表示向上滚动，false表示向下滚动
	* lines: 滚动的行数，值越大滚动幅度越大
	 */
	ScrollVertical(up bool, lines int) error

	// 当前定位器所在区域可以水平滚动
	IsHorizontalScroller() bool
	/**
	* 在当前定位器所在区域内水平滚动
	* left: 是否向左滚动，true表示向左滚动，false表示向右滚动
	* lines: 滚动的行数，值越大滚动幅度越大
	 */
	ScrollHorizontal(left bool, lines int) error
}
