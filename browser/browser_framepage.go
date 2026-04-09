package browser

import (
	"net/url"
	"time"
)

type BrowserFramePage interface {
	IsFrame() bool
	GetID() string
	GetURL() *url.URL
	GetDomain() string
	GetTitle() string
	Evaluate(jsCode string, arg ...any) (any, error)
	Reload(timeout time.Duration) error
	WaitSelector(selector string, timeout time.Duration) (BrowserElement, error)
	QuerySelector(selector string, timeout time.Duration) (BrowserElement, error)
	QuerySelectorAll(selector string, timeout time.Duration) ([]BrowserElement, error)
	QueryXPath(xpath string, timeout time.Duration) (BrowserElement, error)
	QueryXPathAll(xpath string, timeout time.Duration) ([]BrowserElement, error)
}
