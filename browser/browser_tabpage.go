package browser

import "time"

/*
* BrowserTabPage 定义浏览器标签页接口。
 */
type BrowserTabPage interface {
	BrowserFramePage
	/*
	* 激活当前标签页。
	 */
	Activate() error
	/*
	* 销毁当前标签页。
	 */
	Destroy() error
	/*
	* 执行会触发新标签页打开的操作，等待新标签页出现，并设置新标签页的 ID。
	 */
	OpenNewTab(action func(), tabId string, timeout time.Duration) (BrowserTabPage, error)
	/*
	* 保存当前标签页的 cookies。
	 */
	SaveCookies() error
	/*
	* 加载当前标签页的 cookies。
	 */
	LoadCookies() error
	/*
	* 清除当前标签页的 cookies。
	 */
	ClearCookies() error
	/*
	* 保存当前标签页的 LocalStorage。
	 */
	SaveLocalStorage() error
	/*
	* 加载当前标签页的 LocalStorage。
	 */
	LoadLocalStorage() error
	/*
	* 清除当前标签页的 LocalStorage。
	 */
	ClearLocalStorage() error
	/*
	* 清除当前标签页中的所有 IndexedDB 文件。
	 */
	ClearAllIndexDBFiles() error
}
