package browser

import (
	"net/url"
	"time"
)

/*
* BrowserFramePage defines the common interface for browser pages or frame pages.
 */
type BrowserFramePage interface {
	/*
	* Determines whether the current page is an iframe or frame.
	 */
	IsFrame() bool
	/*
	* Gets the unique identifier of the current page.
	 */
	GetID() string
	/*
	* Gets the URL of the current page.
	 */
	GetURL() *url.URL
	/*
	* Gets the domain of the current page.
	 */
	GetDomain() string
	/*
	* Gets the title of the current page.
	 */
	GetTitle() string
	/*
	* Executes JavaScript code in the context of the current page.
	 */
	Evaluate(jsCode string, arg ...any) (any, error)
	/*
	* Reloads the current page and waits up to the specified timeout.
	 */
	Reload(timeout time.Duration) error
	/*
	* Waits for an element matching the specified CSS selector to appear.
	 */
	WaitSelector(selector string, timeout time.Duration) (BrowserElement, error)
	/*
	* Queries a single element matching the specified CSS selector.
	 */
	QuerySelector(selector string, timeout time.Duration) (BrowserElement, error)
	/*
	* Queries all elements matching the specified CSS selector.
	 */
	QuerySelectorAll(selector string, timeout time.Duration) ([]BrowserElement, error)
	/*
	* Queries a single element matching the specified XPath.
	 */
	QueryXPath(xpath string, timeout time.Duration) (BrowserElement, error)
	/*
	* Queries all elements matching the specified XPath.
	 */
	QueryXPathAll(xpath string, timeout time.Duration) ([]BrowserElement, error)
}
