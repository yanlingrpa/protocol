package browser

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
)

/*
* BrowserElement defines the common operation interface for browser elements.
 */
type BrowserElement interface {
	/*
	* Gives the current element input focus, usually for interactive elements such as input boxes and buttons.
	 */
	Focus()
	/*
	* Scrolls the current element into the visible area of the browser window.
	 */
	ScrollIntoView() error
	/*
	* Moves the mouse cursor over the current element.
	 */
	Hover() error
	/*
	* Moves the mouse cursor out of the current element.
	 */
	MoveMouseOut() error
	/*
	* Performs a single-click on the current element.
	 */
	Click() error
	/*
	* Performs a right-click on the current element.
	 */
	RightClick() error
	/*
	* Performs a double-click on the current element.
	 */
	DoubleClick() error
	/*
	* Performs a tap action on the current element.
	 */
	Tap() error
	/*
	* Returns an interactable position coordinate for the current element.
	 */
	Interactable() (basic.FPoint, error)
	/*
	* Selects text within the current element by a regular-expression rule.
	 */
	SelectText(regex string) error
	/*
	* Selects all text within the current element.
	 */
	SelectAllText() error
	/*
	* Inputs text into the current element.
	 */
	Input(text string) error
	/*
	* Removes focus from the current element.
	 */
	Blur() error
	/*
	* Selects or deselects options by displayed text.
	 */
	SelectByText(texts []string, selected bool) error
	/*
	* Selects or deselects options by regular-expression rules.
	 */
	SelectByRegex(regexes []string, selected bool) error
	/*
	* Selects or deselects options by CSS selectors.
	 */
	SelectByCss(selectors []string, selected bool) error
	/*
	* Checks whether the current element matches the specified CSS selector.
	 */
	MatchByCss(selector string) (bool, error)
	/*
	* Gets the attribute value of the current element.
	 */
	GetAttribute(name string) (string, error)
	/*
	* Sets the attribute value of the current element.
	 */
	SetAttribute(name, value string) error
	/*
	* Gets the property value of the current element.
	 */
	GetProperty(name string) (any, error)
	/*
	* Sets the property value of the current element.
	 */
	SetProperty(name string, value any) error
	/*
	* Checks whether the current element is disabled.
	 */
	Disabled() (bool, error)
	/*
	* Sets the upload file list for file-input elements.
	 */
	SetFiles(filePaths []string) error
	/*
	* Returns the embedded frame page associated with the current element.
	 */
	FramePage() (BrowserFramePage, error)
	/*
	* Checks whether the current element contains the target element.
	 */
	ContainsElement(target BrowserElement) (bool, error)
	/*
	* Gets the text content of the current element.
	 */
	Text() (string, error)
	/*
	* Gets the HTML content of the current element.
	 */
	Html() (string, error)
	/*
	* Checks whether the current element is visible.
	 */
	Visible() (bool, error)
	/*
	* Waits until the current element remains stable for the specified duration.
	 */
	WaitStable(stableTime time.Duration) error
	/*
	* Waits until the current element is stable across requestAnimationFrame ticks.
	 */
	WaitStableRAF() error
	/*
	* Waits until the current element becomes interactable and returns an interactable position.
	 */
	WaitInteractable() (basic.FPoint, error)
	/*
	* Waits until the current element becomes visible.
	 */
	WaitVisible() error
	/*
	* Waits until the current element becomes enabled.
	 */
	WaitEnabled() error
	/*
	* Waits until the current element becomes writable.
	 */
	WaitWritable() error
	/*
	* Waits until the current element becomes invisible.
	 */
	WaitInvisible() error
	/*
	* Executes JavaScript code in the context of the current element.
	 */
	Evaluate(jsCode string, arg ...any) (any, error)
	/*
	* Gets the XPath expression of the current element.
	 */
	GetXPath(optimized bool) (string, error)
}
