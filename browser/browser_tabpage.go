package browser

import "time"

/*
* BrowserTabPage 定义浏览器标签页接口
 */
type BrowserTabPage interface {
	BrowserFramePage
	/*
	* 激活当前标签页
	 */
	Activate() error
	/*
	* 销毁当前标签页
	 */
	Destroy() error
	/*
	* 等待指定 ID 的新标签页出现
	 */
	WaitForNewTab(id string, timeout time.Duration) error
	/*
	* 保存当前标签页 Cookies
	 */
	SaveCookies() error
	/*
	* 加载当前标签页 Cookies
	 */
	LoadCookies() error
	/*
	* 清空当前标签页 Cookies
	 */
	ClearCookies() error
	/*
	* 保存当前标签页 LocalStorage
	 */
	SaveLocalStorage() error
	/*
	* 加载当前标签页 LocalStorage
	 */
	LoadLocalStorage() error
	/*
	* 清空当前标签页 LocalStorage
	 */
	ClearLocalStorage() error
	/*
	* 清空当前标签页全部 IndexedDB 文件
	 */
	ClearAllIndexDBFiles() error
}
