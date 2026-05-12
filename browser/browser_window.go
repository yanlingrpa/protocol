package browser

import (
	"yanlingrpa.com/yanling/protocol/osgui"
)

/*
* BrowserWindow defines the browser window interface.
 */
type BrowserWindow interface {
	osgui.OSGuiWindow
	/*
	* Gets the default tab page.
	 */
	DefaultPage() BrowserTabPage
	/*
	* Gets the currently active tab page.
	 */
	CurrentPage() BrowserTabPage
	/*
	* Gets a tab page by identifier.
	 */
	IDTabPage() BrowserTabPage
	/*
	* Creates a new tab page and opens the specified URL.
	 */
	NewTabPage(id string, url string) (BrowserTabPage, error)
}
