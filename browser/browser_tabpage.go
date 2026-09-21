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
	* 等待具有指定 ID 的新标签页出现。
	 */
	WaitForNewTab(id string, timeout time.Duration) (BrowserTabPage, error)
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
