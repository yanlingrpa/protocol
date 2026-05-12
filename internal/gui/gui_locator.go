package gui

import "yanlingrpa.com/yanling/protocol/basic"

/*
* GuiLocator defines the common abstract interface for GUI locators across platforms.
 */
type GuiLocator interface {
	/*
	* Converts to a map object.
	 */
	ToMap() map[string]any
	/*
	* Gets the size of the current locator.
	 */
	GetSize() basic.Size
	/*
	* Returns the absolute position and size of the current locator on the screen.
	 */
	GetScreenRect() basic.Rect
	/*
	* Gets the position and size of the current locator within its parent window.
	 */
	GetWindowRect() basic.Rect
	/*
	* Returns the position and size of the current locator relative to its body area.
	 */
	GetBodyRect() basic.Rect

	/* Captures a snapshot of the current locator area and returns image bytes.
	* gray: Whether to use grayscale.
	 */
	Snapshot(gray bool) ([]byte, error)

	/* Converts a locator-relative position to an absolute screen position.
	* locator_point: Relative position within the locator.
	* return: Absolute position on the screen.
	 */
	TransToScreen(locator_point *basic.Point) *basic.Point

	/*
	* Converts a locator-relative position to a relative position within the parent window.
	 */
	TransToWindow(locator_point *basic.Point) *basic.Point

	/*
	* Converts a locator-relative position to a relative position within the client/body area.
	 */
	TransToBody(locator_point *basic.Point) *basic.Point

	/*
	* Converts an absolute screen position to a relative position within the current locator.
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	/*
	* Converts a position relative to the parent window to a relative position within the current locator.
	 */
	TransFromWindow(window_pos *basic.Point) *basic.Point

	/*
	* Converts a position relative to the client/body area to a relative position within the current locator.
	 */
	TransFromBody(body_pos *basic.Point) *basic.Point

	/*
	* Performs OCR on the current locator area and returns the recognition result.
	 */
	OcrRead(confidence float64) (*basic.OcrResult, error)

	/*
	* Attempts to focus the current locator.
	 */
	Focus() error

	/*
	* Indicates whether the current locator is in a text-editable state.
	 */
	IsEditing() bool

	/*
	* Reads text content from the current locator.
	 */
	ReadText() (string, error)

	/*
	* Writes text into the current locator.
	 */
	WriteText(text string) error

	/*
	* Clears text content in the current locator.
	 */
	ClearText() error

	/*
	* Checks whether vertical scrolling is supported in the current locator area.
	 */
	CanScrollVertical() bool
	/*
	* Performs vertical scrolling in the current locator area.
	 */
	ScrollVertical(forward bool, distance int) error

	/*
	* Checks whether horizontal scrolling is supported in the current locator area.
	 */
	CanScrollHorizontal() bool
	/*
	* Performs horizontal scrolling in the current locator area.
	 */
	ScrollHorizontal(forward bool, distance int) error
}
