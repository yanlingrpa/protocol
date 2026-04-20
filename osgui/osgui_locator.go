package osgui

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
)

/*
* Locator 定义了一个定位器接口，提供了获取定位器信息、操作定位器、模拟输入等功能
 */
type Locator interface {
	/*
	* 转换为map对象
	 */
	ToMap() map[string]any
	/*
	* 获取当前定位器的尺寸
	 */
	GetSize() basic.Size
	/*
	* 返回当前定位器在屏幕上的绝对位置和大小
	 */
	GetScreenRect() basic.Rect
	/*
	* 获取当前定位器在所属窗体内的位置和大小
	* 返回当前定位器相对于其所属窗体内的相对位置和大小
	 */
	GetWindowRect() basic.Rect
	/*
	* 返回当前定位器相对于其所属body的位置和大小
	* Body: 客户区，指的是窗口中除去标题栏、边框等非客户区部分的区域，通常是应用程序实际显示内容的区域
	 */
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

	/*
	* 将定位器内的相对位置转换为所属窗体内的相对位置
	* locator_point: 定位器内的相对位置
	* return: 窗体内的相对位置
	 */
	TransToWindow(locator_point *basic.Point) *basic.Point

	/*
	* 将定位器内的相对位置转换为客户区内的相对位置
	* locator_point: 定位器内的相对位置
	* return: 客户区内的相对位置
	 */
	TransToBody(locator_point *basic.Point) *basic.Point

	/*
	* 将屏幕上的绝对位置转换为当前定位器内的相对位置
	* screen_pos: 屏幕上的绝对位置
	* return: 定位器内的相对位置, 如果点不在定位器内则返回nil
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	/*
	* 将所属窗体内的相对位置转换为当前定位器内的相对位置
	* window_pos: 窗体内的相对位置
	* return: 定位器内的相对位置, 如果点不在定位器内则返回nil
	 */
	TransFromWindow(window_pos *basic.Point) *basic.Point

	/*
	* 将客户区内的相对位置转换为当前定位器内的相对位置
	* body_pos: 客户区内的相对位置
	* return: 定位器内的相对位置, 如果点不在定位器内则返回nil
	 */
	TransFromBody(body_pos *basic.Point) *basic.Point

	/*
	* 获取当前定位器的子定位器，仅取得该区域与当前定位器边界的交集部分
	* locator_point: 子定位器在当前定位器内的相对位置
	* size: 子定位器的尺寸
	 */
	SubLocator(locator_point basic.Point, size basic.Size) Locator

	/*
	* 在当前定位器所在区域内查找包含相似图片的子定位器
	* image: 图片相对与当前project的路径/URL/Base64编码字符串
	* sim: 相似度，取值范围0.1~1.0，值越大表示越相似
	* return: 找到的子定位器列表，如果没有找到则返回nil，相似度最高的定位器排在前面，达不到相似度要求的会被过滤掉
	 */
	ImageLocator(image string, sim float32) ([]Locator, error)

	/*
	* 在当前定位器所在区域内等待包含相似图片的子定位器出现
	* timeout: 超时时间
	* image: 图片相对与当前project的路径/URL/Base64编码字符串
	* sim: 相似度，取值范围0.1~1.0，值越大表示越相似
	* return: 成功时返回匹配的子定位器列表（按相似度从高到低排序，低于sim的结果会被过滤）；超时或识别异常时返回error
	 */
	WaitForImage(timeout time.Duration, image string, sim float32) ([]Locator, error)

	/*
	* 在当前定位器所在区域内查找包含指定文本的子定位器
	* texts: 需要查找的文本列表，支持多个文本, 多个文本时表示同时包含
	* return: 找到的子定位器列表，如果没有找到则返回nil，必须同时包含所有文本的定位器才会被返回，定位器按面积从小到大排序返回
	 */
	TextLocator(texts ...string) ([]Locator, error)

	/*
	* 在当前定位器所在区域内等待包含指定文本的子定位器出现
	* timeout: 超时时间
	* texts: 需要等待的文本列表，支持多个文本, 多个文本时表示同时包含
	* return: 成功时返回匹配的子定位器列表（仅返回同时包含所有texts的定位器，按面积从小到大排序）；超时或识别异常时返回error
	 */
	WaitForText(timeout time.Duration, texts ...string) ([]Locator, error)

	/*
	* 在当前定位器所在区域内查找卡片式子定位器，卡片式子定位器是指一个独立的类似卡片的区域，通常包含一个图形和一些文本，常见于列表项、按钮等
	* min_size: 查找的子定位器的最小尺寸, nil 表示不限制最小尺寸
	* max_size: 查找的子定位器的最大尺寸, nil 表示不限制最大尺寸
	* return: 找到的子定位器列表，如果没有找到则返回nil，子定位器按坐标从左到右、从上到下排序返回
	 */
	CardLocator(min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* 等待当前定位器所在区域内出现卡片式子定位器，卡片式子子定位器是指一个独立的类似卡片的区域，通常包含一个图形和一些文本，常见于列表项、按钮等
	* timeout: 超时时间
	* min_size: 查找的子定位器的最小尺寸, nil 表示不限制最小尺寸
	* max_size: 查找的子定位器的最大尺寸, nil 表示不限制最大尺寸
	* return: 成功时返回匹配的子定位器列表（按坐标从左到右、从上到下排序）；超时或识别异常时返回error
	 */
	WaitForCard(timeout time.Duration, min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* 在当前定位器所在区域内用视觉模块查找指定图形形态的子定位器
	* description: 对图形形态的描述，例如：红色的圆形按钮，蓝色的矩形输入框等
	* min_size: 查找的子定位器的最小尺寸, nil 表示不限制最小尺寸
	* max_size: 查找的子定位器的最大尺寸, nil 表示不限制最大尺寸
	* return: 找到的子定位器列表，如果没有找到则返回nil，子定位器按坐标从左到右、从上到下排序返回
	 */
	VisionLocator(description string, min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* 等待当前定位器所在区域内用视觉模块查找指定图形形态的子定位器
	* timeout: 超时时间
	* description: 对图形形态的描述，例如：红色的圆形按钮，蓝色的矩形输入框等
	* min_size: 查找的子定位器的最小尺寸, nil 表示不限制最小尺寸
	* max_size: 查找的子定位器的最大尺寸, nil 表示不限制最大尺寸
	* return: 成功时返回匹配的子定位器列表（按坐标从左到右、从上到下排序）；超时或识别异常时返回error
	 */
	WaitForVision(timeout time.Duration, description string, min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* 对当前定位器所在区域进行OCR识别，返回识别结果
	* confidence: 识别的置信度，取值范围0.1~1.0，值越大表示越准确
	 */
	OcrRead(confidence float64) (*basic.OcrResult, error)

	/*
	* 将鼠标移动到定位器内的指定位置
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器的中心位置
	 */
	MouseMove(locator_point *basic.Point) error

	/*
	* 将鼠标按下到定位器内的指定位置
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器的中心位置
	 */
	MouseDown(locator_point *basic.Point) error

	/*
	* 将鼠标弹起到定位器内的指定位置
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器的中心位置
	 */
	MouseUp(locator_point *basic.Point) error

	/*
	* 在定位器内从from_locator_point位置拖拽到target locator中的to_locator_point位置
	* from_locator_point: 拖拽的起始位置在当前定位器内的相对位置，nil表示当前定位器内的鼠标位置，如果鼠标不在当前定位器内，则取定位器中心点
	* to_locator_point: 拖拽的结束位置在目标定位器内的相对位置，nil表示目标定位器内的鼠标位置，如果鼠标不在目标定位器内，则取目标定位器中心点
	 */
	DragTo(target Locator, from_locator_point *basic.Point, to_locator_point *basic.Point) error

	/*
	* 当前定位器获取焦点
	* 如果鼠标在当前定位器内则认为已经获取焦点
	* 否则将鼠标移动到当前定位器中心并点击以获取焦点
	 */
	Focus() error

	/*
	* 在定位器内的指定位置进行鼠标点击
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器内的鼠标位置，
	* 如果鼠标不在当前定位器内，则取定位器中心点
	 */
	Click(locator_point *basic.Point) error

	/*
	* 在定位器内的指定位置进行鼠标双击
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器内的鼠标位置，
	* 如果鼠标不在当前定位器内，则取定位器中心点
	 */
	DoubleClick(locator_point *basic.Point) error

	/*
	* 在定位器内的指定位置进行鼠标右键点击
	* locator_point: 位置在定位器内的相对位置
	* nil表示当前定位器内的鼠标位置，
	* 如果鼠标不在当前定位器内，则取定位器中心点
	 */
	RightClick(locator_point *basic.Point) error

	/*
	* 获取鼠标在当前定位器内的相对位置
	* return: 鼠标在当前定位器内的相对位置, 如果鼠标不在定位器内则返回nil
	 */
	GetLocatorCursorPos() *basic.Point

	/*
	* 获取当前定位器内输入法光标的位置, 没有则返回nil
	* 例如：在输入框内时，输入法光标通常会显示在输入框内，此时可以通过该方法获取输入法光标在定位器内的相对位置
	* return: 输入法光标在定位器内的相对位置, 没有则返回nil
	 */
	GetLocatorCaretPos() *basic.Point

	/*
	* 模拟键盘输入
	* keys: 要输入的按键序列 （例如：ctrl + alt + del）
	* 支持的按键见Keyboard类型定义
	 */
	PressKeys(keys ...Keyboard) error

	/*
	* 当前定位器内是否处于可以输入文本的状态
	* 例如：输入框被点击后通常会进入可以输入文本的状态，此时可以通过键盘输入文本到该输入框
	* return: true表示可以输入文本，false表示不可以输入文本
	 */
	IsEditing() bool

	/*
	* 等待当前定位器进入可以输入文本的状态，直到超时
	* timeout: 超时时间
	* return: 文本输入光标所在的子定位器，如果超时则返回错误，如果定位器内没有输入光标则返回nil
	 */
	WaitForEditing(timeout time.Duration) (Locator, error)

	/*
	* 获取当前定位器内的文本内容，
	* 等同于先点击定位器，然后全选文本，最后复制文本到剪贴板，再从剪贴板读取文本内容
	* 如果读取不到文本内容，则尝试直接使用OCR识别当前定位器内的文本内容并返回
	* return: 定位器内的文本内容，如果无法获取文本内容则返回错误
	 */
	ReadText() (string, error)

	/*
	* 向当前定位器内输入文本，
	* 等同于先点击定位器，然后输入文本
	* return: 如果输入成功则返回nil，否则返回错误
	 */
	WriteText(text string) error

	/*
	* 清空当前定位器内的文本内容，
	* 等同于先点击定位器，然后全选文本，最后按下删除键或退格键清空文本内容
	* return: 如果清空成功则返回nil，否则返回错误
	 */
	ClearText() error

	/*
	* 当前定位器所在区域内可以垂直滚动
	* return: true表示可以垂直滚动，false表示不可以垂直滚动
	 */
	IsVerticalScroller() bool
	/*
	* 在当前定位器所在区域内垂直滚动
	* up: 是否向上滚动，true表示向上滚动，false表示向下滚动
	* lines: 滚动的行数，值越大滚动幅度越大
	 */
	ScrollVertical(up bool, lines int) error

	/*
	* 当前定位器所在区域内可以水平滚动
	* return: true表示可以水平滚动，false表示不可以水平滚动
	 */
	IsHorizontalScroller() bool
	/*
	* 在当前定位器所在区域内水平滚动
	* left: 是否向左滚动，true表示向左滚动，false表示向右滚动
	* lines: 滚动的行数，值越大滚动幅度越大
	 */
	ScrollHorizontal(left bool, lines int) error
}
