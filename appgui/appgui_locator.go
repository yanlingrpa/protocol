package appgui

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/internal/gui"
)

/*
* AppGuiLocator 定义移动应用屏幕矩形区域的定位器接口，
* 提供获取定位器信息、执行触摸手势以及模拟输入等方法。
 */
type AppGuiLocator interface {
	gui.GuiLocator

	/*
	* 获取当前定位器所属的App映射窗口。
	 */
	GetWindow() AppGuiWindow

	/*
	* 获取当前定位器的子定位器，并裁剪到当前边界的交集范围内。
	* locator_point: 子定位器在当前定位器中的相对位置。
	* size: 子定位器的大小。
	 */
	SubLocator(locator_point basic.Point, size basic.Size) AppGuiLocator

	/*
	* 在当前定位器区域中查找与视觉相似的图片对应的子定位器。
	* image: 当前项目相对路径/URL/Base64 图片字符串。
	* sim: 相似度阈值，范围 0.1~1.0；值越高表示越相似。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 结果按相似度从高到低排序，低于阈值的项会被过滤掉。
	 */
	ImageLocator(image string, sim float32) ([]AppGuiLocator, error)

	/*
	* 等待当前定位器区域中出现与视觉相似图片对应的子定位器。
	* timeout: 超时时间。
	* image: 当前项目相对路径/URL/Base64 图片字符串。
	* sim: 相似度阈值，范围 0.1~1.0；值越高表示越相似。
	* return: 成功时返回按相似度降序排序的匹配子定位器；
	* 低于 sim 的结果会被过滤掉。超时或识别失败时返回错误。
	 */
	WaitForImage(timeout time.Duration, image string, sim float32) ([]AppGuiLocator, error)

	/*
	* 在当前定位器区域中查找包含指定文本的子定位器。
	* texts: 要查找的文本列表。多个文本表示必须全部存在。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 仅返回包含所有文本的定位器，并按区域从小到大排序。
	 */
	TextLocator(texts ...string) ([]AppGuiLocator, error)

	/*
	* 等待当前定位器区域中出现包含指定文本的子定位器。
	* timeout: 超时时间。
	* texts: 要等待的文本列表。多个文本表示必须全部存在。
	* return: 成功时返回包含所有文本的匹配子定位器，
	* 按区域从小到大排序。超时或识别失败时返回错误。
	 */
	WaitForText(timeout time.Duration, texts ...string) ([]AppGuiLocator, error)

	/*
	* 在当前定位器区域中查找卡片式子定位器。
	* 卡片式子定位器是一个独立的卡片形状区域，通常包含图形和文本，常见于列表项和按钮。
	* min_size: 要查找的子定位器最小尺寸；nil 表示不设最小限制。
	* max_size: 要查找的子定位器最大尺寸；nil 表示不设最大限制。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 结果按坐标从左到右、从上到下排序。
	 */
	CardLocator(min_size, max_size *basic.Size) ([]AppGuiLocator, error)

	/*
	* 等待当前定位器区域中出现卡片式子定位器。
	* 卡片式子定位器是一个独立的卡片形状区域，通常包含图形和文本，常见于列表项和按钮。
	* timeout: 超时时间。
	* min_size: 要查找的子定位器最小尺寸；nil 表示不设最小限制。
	* max_size: 要查找的子定位器最大尺寸；nil 表示不设最大限制。
	* return: 成功时返回按坐标从左到右、从上到下排序的匹配子定位器。
	* 超时或识别失败时返回错误。
	 */
	WaitForCard(timeout time.Duration, min_size, max_size *basic.Size) ([]AppGuiLocator, error)

	/*
	* 使用视觉模块在当前定位器区域中查找具有特定视觉形状的子定位器。
	* description: 目标形状的描述，例如：红色圆形按钮，
	* 或蓝色矩形输入框。
	* min_size: 要查找的子定位器最小尺寸；nil 表示不设最小限制。
	* max_size: 要查找的子定位器最大尺寸；nil 表示不设最大限制。
	* return: 匹配到的子定位器列表；未找到时返回 nil。
	* 结果按坐标从左到右、从上到下排序。
	 */
	VisionLocator(description string, min_size, max_size *basic.Size) ([]AppGuiLocator, error)

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
	WaitForVision(timeout time.Duration, description string, min_size, max_size *basic.Size) ([]AppGuiLocator, error)

	/*
	* 将触摸点移动到定位器内的指定位置，但不抬起手指。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前定位器的中心位置。
	 */
	TouchMove(locator_point *basic.Point) error

	/*
	* 在定位器内当前触摸位置按下手指。
	 */
	TouchDown() error

	/*
	* 在定位器内当前触摸位置抬起手指。
	 */
	TouchUp() error

	/*
	* 在当前定位器内执行从起始位置到结束位置的滑动手势。
	* 两个位置都使用当前定位器内的相对坐标。
	* from_locator_point: 相对起始位置；nil 表示当前触摸位置，
	* 若没有活动触摸，则使用定位器中心。
	* to_locator_point: 相对结束位置；nil 表示当前触摸位置，
	* 若没有活动触摸，则使用定位器中心。
	 */
	SwipeTo(from_locator_point *basic.Point, to_locator_point *basic.Point) error

	/*
	* 在当前定位器区域执行捏合手势。
	* spread: true 表示展开手指（放大），false 表示收拢手指（缩小）。
	* scale: 手势距离的缩放因子，范围 0.1~1.0；值越大表示展开/收拢越明显。
	 */
	Pinch(spread bool, scale float32) error

	/*
	* 在定位器内指定位置执行单击操作。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前触摸位置。
	* 若定位器内没有活动触摸，则使用定位器中心。
	 */
	Tap(locator_point *basic.Point) error

	/*
	* 在定位器内指定位置执行双击操作。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前触摸位置。
	* 若定位器内没有活动触摸，则使用定位器中心。
	 */
	DoubleTap(locator_point *basic.Point) error

	/*
	* 在定位器内指定位置执行长按操作。
	* locator_point: 定位器内的相对位置。
	* nil 表示当前触摸位置。
	* 若定位器内没有活动触摸，则使用定位器中心。
	 */
	LongPress(locator_point *basic.Point) error

	/*
	* 获取相对于当前定位器的当前触摸位置。
	* return: 定位器内的相对触摸位置；如果没有活动触摸则返回 nil。
	 */
	GetLocatorTouchPos() *basic.Point

	/*
	* 等待当前定位器进入可文本编辑状态，或在超时后返回。
	* timeout: 超时时间。
	* return: 包含文本输入光标的子定位器。
	* 超时时返回错误；若定位器中不存在光标，则返回 nil。
	 */
	WaitForEditing(timeout time.Duration) (AppGuiLocator, error)

	/*
	* 在定位器获得焦点时模拟按下硬件或系统按键。
	* keys: 要按下的按键序列，定义于 AppKey 类型。
	 */
	PressKeys(keys ...AppKey) error
}
