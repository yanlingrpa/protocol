package browser

import "time"

type BrowserTabPage interface {
	BrowserFramePage
	Activate() error
	Destroy() error
	WaitForNewTab(id string, timeout time.Duration) error
	SaveCookies() error
	LoadCookies() error
	ClearCookies() error
	SaveLocalStorage() error
	LoadLocalStorage() error
	ClearLocalStorage() error
	ClearAllIndexDBFiles() error
}
