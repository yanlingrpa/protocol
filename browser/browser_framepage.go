package browser

import (
	"net/url"
	"time"
)

/*
* BrowserFramePage 定义浏览器页面或框架页的通用接口
 */
type BrowserFramePage interface {
	/*
	* 判断当前页面是否为 iframe 或 frame
	 */
	IsFrame() bool
	/*
	* 获取当前页面唯一标识
	 */
	GetID() string
	/*
	* 获取当前页面 URL
	 */
	GetURL() *url.URL
	/*
	* 获取当前页面域名
	 */
	GetDomain() string
	/*
	* 获取当前页面标题
	 */
	GetTitle() string
	/*
	* 在当前页面上下文中执行 JavaScript 代码
	 */
	Evaluate(jsCode string, arg ...any) (any, error)
	/*
	* 重新加载当前页面，并等待到指定超时
	 */
	Reload(timeout time.Duration) error
	/*
	* 等待匹配指定 CSS 选择器的元素出现
	 */
	WaitSelector(selector string, timeout time.Duration) (BrowserElement, error)
	/*
	* 查询匹配指定 CSS 选择器的单个元素
	 */
	QuerySelector(selector string, timeout time.Duration) (BrowserElement, error)
	/*
	* 查询匹配指定 CSS 选择器的全部元素
	 */
	QuerySelectorAll(selector string, timeout time.Duration) ([]BrowserElement, error)
	/*
	* 查询匹配指定 XPath 的单个元素
	 */
	QueryXPath(xpath string, timeout time.Duration) (BrowserElement, error)
	/*
	* 查询匹配指定 XPath 的全部元素
	 */
	QueryXPathAll(xpath string, timeout time.Duration) ([]BrowserElement, error)
}
