package browser

import (
	"yanlingrpa.com/yanling/protocol/osgui"
)

/*
* BrowserWindow 定义浏览器窗口接口
 */
type BrowserWindow interface {
	osgui.GuiWindow
	/*
	* 获取默认标签页
	 */
	DefaultPage() BrowserTabPage
	/*
	* 获取当前活动标签页
	 */
	CurrentPage() BrowserTabPage
	/*
	* 按标识获取标签页
	 */
	IDTabPage() BrowserTabPage
	/*
	* 新建标签页并打开指定 URL
	 */
	NewTabPage(id string, url string) (BrowserTabPage, error)
}
