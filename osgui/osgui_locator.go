package osgui

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/internal/gui"
)

/*
* OSGuiLocator 定义一个定位器接口，提供获取定位器信息、操作定位器以及模拟输入的方法。
 */
type OSGuiLocator interface {
	gui.GuiLocator

	/*
	* 获取当前定位器所属的窗口。
	 */
	GetWindow() OSGuiWindow

	/*
	* 获取当前定位器的子定位器，并裁剪到与当前边界的交集范围内。
	* locator_point: 子定位器在当前定位器中的相对位置。
	* size: 子定位器的大小。
	 */
	SubLocator(locator_point basic.Point, size basic.Size) OSGuiLocator

	/*
	* 在当前定位器区域中查找与视觉相似的图片对应的子定位器。
	* image: 当前项目相对路径/URL/Base64 图片字符串。
	* sim: 相似度阈值，范围 0.1~1.0；值越高表示越相似。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 结果按相似度从高到低排序，低于阈值的项会被过滤掉。
	 */
	ImageLocator(image string, sim float32) ([]OSGuiLocator, error)

	/*
	* 等待当前定位器区域中出现与视觉相似图片对应的子定位器。
	* timeout: 超时时间。
	* image: 当前项目相对路径/URL/Base64 图片字符串。
	* sim: 相似度阈值，范围 0.1~1.0；值越高表示越相似。
	* return: 成功时返回按相似度降序排序的匹配子定位器；
	* 低于 sim 的结果会被过滤掉。超时或识别失败时返回错误。
	 */
	WaitForImage(timeout time.Duration, image string, sim float32) ([]OSGuiLocator, error)

	/*
	* 在当前定位器区域中查找包含指定文本的子定位器。
	* texts: 要查找的文本列表。多个文本表示必须全部存在。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 仅返回包含所有文本的定位器，并按区域从小到大排序。
	 */
	TextLocator(texts ...string) ([]OSGuiLocator, error)

	/*
	* 等待当前定位器区域中出现包含指定文本的子定位器。
	* timeout: 超时时间。
	* texts: 要等待的文本列表。多个文本表示必须全部存在。
	* return: 成功时返回包含所有文本的匹配子定位器，
	* 按区域从小到大排序。超时或识别失败时返回错误。
	 */
	WaitForText(timeout time.Duration, texts ...string) ([]OSGuiLocator, error)

	/*
	* 在当前定位器区域中查找卡片式子定位器。
	* 卡片式子定位器是一个独立的卡片形状区域，通常包含图形和文本，常见于列表项和按钮。
	* min_size: 要查找的子定位器最小尺寸；nil 表示不设最小限制。
	* max_size: 要查找的子定位器最大尺寸；nil 表示不设最大限制。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 结果按坐标从左到右、从上到下排序。
	 */
	CardLocator(min_size, max_size *basic.Size) ([]OSGuiLocator, error)

	/*
	* 等待当前定位器区域中出现卡片式子定位器。
	* 卡片式子定位器是一个独立的卡片形状区域，通常包含图形和文本，常见于列表项和按钮。
	* timeout: 超时时间。
	* min_size: 要查找的子定位器最小尺寸；nil 表示不设最小限制。
	* max_size: 要查找的子定位器最大尺寸；nil 表示不设最大限制。
	* return: 成功时返回按坐标从左到右、从上到下排序的匹配子定位器。
	* 超时或识别失败时返回错误。
	 */
	WaitForCard(timeout time.Duration, min_size, max_size *basic.Size) ([]OSGuiLocator, error)

	/*
	* 使用视觉模块在当前定位器区域中查找具有特定视觉形状的子定位器。
	* description: 目标形状的描述，例如：红色圆形按钮，
	* 或蓝色矩形输入框。
	* min_size: 要查找的子定位器最小尺寸；nil 表示不设最小限制。
	* max_size: 要查找的子定位器最大尺寸；nil 表示不设最大限制。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 结果按坐标从左到右、从上到下排序。
	 */
	VisionLocator(description string, min_size, max_size *basic.Size) ([]OSGuiLocator, error)

	/*
	* 等待当前定位器区域中出现具有特定视觉形状的子定位器，
	* 使用视觉模块进行识别。
	* timeout: 超时时间。
	* description: 目标形状的描述，例如：红色圆形按钮，
	* 或蓝色矩形输入框。
	* min_size: 要查找的子定位器最小尺寸；nil 表示不设最小限制。
	* max_size: 要查找的子定位器最大尺寸；nil 表示不设最大限制。
	* return: 成功时返回按坐标从左到右、从上到下排序的匹配子定位器。
	* 超时或识别失败时返回错误。
	 */
	WaitForVision(timeout time.Duration, description string, min_size, max_size *basic.Size) ([]OSGuiLocator, error)

	/*
	* 将鼠标移动到定位器内的指定位置。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前定位器的中心位置。
	 */
	MouseMove(locator_point *basic.Point) error

	/*
	 * 在定位器内当前鼠标位置按下鼠标按钮。
	 * right: true 表示右键，false 表示左键。
	 */
	MouseDown(right bool) error

	/*
	 * 在定位器内当前鼠标位置释放鼠标按钮。
	 * right: true 表示右键，false 表示左键。
	 */
	MouseUp(right bool) error

	/*
	* 从起始位置拖动到结束位置，范围在当前定位器内。
	* 两个位置都使用当前定位器内的相对坐标。
	* from_locator_point: 定位器内的起始相对位置；nil 表示当前鼠标在定位器中的位置，
	* 如果鼠标不在定位器内，则使用定位器中心。
	* to_locator_point: 定位器内的结束相对位置；nil 表示当前鼠标在定位器中的位置，
	* 如果鼠标不在定位器内，则使用定位器中心。
	 */
	DragTo(from_locator_point *basic.Point, to_locator_point *basic.Point) error

	/*
	* 在定位器内指定位置执行鼠标点击。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前鼠标在定位器中的位置。
	* 如果鼠标不在定位器内，则使用定位器中心。
	 */
	Click(locator_point *basic.Point) error

	/*
	* 在定位器内指定位置执行鼠标双击。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前鼠标在定位器中的位置。
	* 如果鼠标不在定位器内，则使用定位器中心。
	 */
	DoubleClick(locator_point *basic.Point) error

	/*
	* 在定位器内指定位置执行鼠标右键点击。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前鼠标在定位器中的位置。
	* 如果鼠标不在定位器内，则使用定位器中心。
	 */
	RightClick(locator_point *basic.Point) error

	/*
	* 获取相对于当前定位器的鼠标位置。
	* return: 当前定位器中的相对鼠标位置；若鼠标在外部则返回 nil。
	 */
	GetLocatorCursorPos() *basic.Point

	/*
	* 获取当前定位器内 IME 光标的位置；不可用时返回 nil。
	* 示例：当输入框被聚焦时，IME 光标通常出现在其中，
	* 该方法可获取其在定位器中的相对位置。
	* return: 定位器中的相对 IME 光标位置；不可用时返回 nil。
	 */
	GetLocatorCaretPos() *basic.Point

	/*
	* 模拟在当前定位器中输入键盘事件。
	* 该方法会优先尝试为定位器获取焦点，
	* 然后发送指定的按键序列。
	* keys: 要输入的按键序列（例如：ctrl + alt + del）。
	* 支持的按键定义于 Keyboard 类型。
	 */
	PressKeys(keys ...Keyboard) error

	/*
	* 等待当前定位器进入可文本编辑状态，或在超时后返回。
	* timeout: 超时时间。
	* return: 包含文本输入光标的子定位器。
	* 超时时返回错误；若定位器中不存在光标，则返回 nil。
	 */
	WaitForEditing(timeout time.Duration) (OSGuiLocator, error)
}
