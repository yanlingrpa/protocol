package browser

import (
	"net/url"
	"time"
)

/*
* BrowserFramePage 定义浏览器页面或 frame 页面通用接口。
 */
type BrowserFramePage interface {
	/*
	* 判断当前页面是否为 iframe 或 frame。
	 */
	IsFrame() bool
	/*
	* 获取当前页面的唯一标识符。
	 */
	GetID() string
	/*
	* 获取当前页面的 URL。
	 */
	GetURL() *url.URL
	/*
	* 获取当前页面的域名。
	 */
	GetDomain() string
	/*
	* 获取当前页面的标题。
	 */
	GetTitle() string
	/*
	* 在当前页面上下文中执行 JavaScript 代码。
	 */
	Evaluate(jsCode string, arg ...any) (any, error)
	/*
	* 刷新当前页面，并等待至指定超时时间。
	 */
	Reload(timeout time.Duration) error
	/*
	* 等待匹配指定 CSS 选择器的元素出现。
	 */
	WaitSelector(selector string, timeout time.Duration) (BrowserElement, error)
	/*
	* 查询单个匹配指定 CSS 选择器的元素。
	 */
	QuerySelector(selector string, timeout time.Duration) (BrowserElement, error)
	/*
	* 查询所有匹配指定 CSS 选择器的元素。
	 */
	QuerySelectorAll(selector string, timeout time.Duration) ([]BrowserElement, error)
	/*
	* 查询单个匹配指定 XPath 的元素。
	 */
	QueryXPath(xpath string, timeout time.Duration) (BrowserElement, error)
	/*
	* 查询所有匹配指定 XPath 的元素。
	 */
	QueryXPathAll(xpath string, timeout time.Duration) ([]BrowserElement, error)
}
