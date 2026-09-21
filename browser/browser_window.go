package browser

import (
	"yanlingrpa.com/yanling/protocol/osgui"
)

/*
* BrowserWindow 定义浏览器窗口接口。
 */
type BrowserWindow interface {
	osgui.OSGuiWindow
	/*
	* 获取默认标签页。
	 */
	DefaultPage() BrowserTabPage
	/*
	* 获取当前激活的标签页。
	 */
	CurrentPage() BrowserTabPage
	/*
	* 根据标识符获取标签页。
	 */
	GetTabPage(id string) BrowserTabPage
	/*
	* 获取所有标签页。
	 */
	ListTabPages() []BrowserTabPage
	/*
	* 创建一个新标签页，并打开指定 URL。
	 */
	NewTabPage(id string, url string) (BrowserTabPage, error)
}
