package browser

import "time"

/*
* BrowserTabPage defines the browser tab page interface.
 */
type BrowserTabPage interface {
	BrowserFramePage
	/*
	* Activates the current tab page.
	 */
	Activate() error
	/*
	* Destroys the current tab page.
	 */
	Destroy() error
	/*
	* Waits for a new tab page with the specified ID to appear.
	 */
	WaitForNewTab(id string, timeout time.Duration) error
	/*
	* Saves cookies for the current tab page.
	 */
	SaveCookies() error
	/*
	* Loads cookies for the current tab page.
	 */
	LoadCookies() error
	/*
	* Clears cookies for the current tab page.
	 */
	ClearCookies() error
	/*
	* Saves LocalStorage for the current tab page.
	 */
	SaveLocalStorage() error
	/*
	* Loads LocalStorage for the current tab page.
	 */
	LoadLocalStorage() error
	/*
	* Clears LocalStorage for the current tab page.
	 */
	ClearLocalStorage() error
	/*
	* Clears all IndexedDB files for the current tab page.
	 */
	ClearAllIndexDBFiles() error
}
