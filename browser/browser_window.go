package browser

import (
	"yanlingrpa.com/yanling/protocol/osgui"
)

type BrowserWindow interface {
	osgui.GuiWindow
	DefaultPage() BrowserTabPage
	CurrentPage() BrowserTabPage
	IDTabPage() BrowserTabPage
	NewTabPage(id string, url string) (BrowserTabPage, error)
}
