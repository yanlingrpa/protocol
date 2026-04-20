package browser

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
)

/*
* BrowserElement 定义浏览器元素的通用操作接口
 */
type BrowserElement interface {
	/*
	* 让当前元素获得输入焦点，通常用于输入框、按钮等可交互元素
	 */
	Focus()
	/*
	* 将当前元素滚动到浏览器窗口的可视区域内
	 */
	ScrollIntoView() error
	/*
	* 将鼠标悬停到当前元素上
	 */
	Hover() error
	/*
	* 将鼠标从当前元素移出
	 */
	MoveMouseOut() error
	/*
	* 对当前元素执行单击操作
	 */
	Click() error
	/*
	* 对当前元素执行右键单击操作
	 */
	RightClick() error
	/*
	* 对当前元素执行双击操作
	 */
	DoubleClick() error
	/*
	* 对当前元素执行触屏点击操作
	 */
	Tap() error
	/*
	* 返回当前元素可交互的位置坐标
	 */
	Interactable() (basic.FPoint, error)
	/*
	* 按正则规则选择当前元素中的文本
	 */
	SelectText(regex string) error
	/*
	* 选择当前元素中的全部文本
	 */
	SelectAllText() error
	/*
	* 向当前元素输入文本
	 */
	Input(text string) error
	/*
	* 让当前元素失去焦点
	 */
	Blur() error
	/*
	* 按显示文本选择或取消选择子项
	 */
	SelectByText(texts []string, selected bool) error
	/*
	* 按正则规则选择或取消选择子项
	 */
	SelectByRegex(regexes []string, selected bool) error
	/*
	* 按 CSS 选择器选择或取消选择子项
	 */
	SelectByCss(selectors []string, selected bool) error
	/*
	* 检查当前元素是否匹配指定 CSS 选择器
	 */
	MatchByCss(selector string) (bool, error)
	/*
	* 获取当前元素的属性值
	 */
	GetAttribute(name string) (string, error)
	/*
	* 设置当前元素的属性值
	 */
	SetAttribute(name, value string) error
	/*
	* 获取当前元素的属性对象值
	 */
	GetProperty(name string) (any, error)
	/*
	* 设置当前元素的属性对象值
	 */
	SetProperty(name string, value any) error
	/*
	* 检查当前元素是否处于禁用状态
	 */
	Disabled() (bool, error)
	/*
	* 为文件输入类元素设置上传文件列表
	 */
	SetFiles(filePaths []string) error
	/*
	* 返回当前元素关联的内嵌框架页面
	 */
	FramePage() (BrowserFramePage, error)
	/*
	* 检查当前元素是否包含目标元素
	 */
	ContainsElement(target BrowserElement) (bool, error)
	/*
	* 获取当前元素的文本内容
	 */
	Text() (string, error)
	/*
	* 获取当前元素的 HTML 内容
	 */
	Html() (string, error)
	/*
	* 检查当前元素是否可见
	 */
	Visible() (bool, error)
	/*
	* 等待当前元素在指定时间内保持稳定
	 */
	WaitStable(stableTime time.Duration) error
	/*
	* 等待当前元素在 requestAnimationFrame 维度上稳定
	 */
	WaitStableRAF() error
	/*
	* 等待当前元素变为可交互，并返回可交互位置
	 */
	WaitInteractable() (basic.FPoint, error)
	/*
	* 等待当前元素变为可见
	 */
	WaitVisible() error
	/*
	* 等待当前元素变为可用
	 */
	WaitEnabled() error
	/*
	* 等待当前元素变为可输入
	 */
	WaitWritable() error
	/*
	* 等待当前元素变为不可见
	 */
	WaitInvisible() error
	/*
	* 在当前元素上下文中执行 JavaScript 代码
	 */
	Evalute(jsCode string, arg ...any) (any, error)
	/*
	* 获取当前元素的 XPath 表达式
	 */
	GetXPath(optimized bool) (string, error)
}
