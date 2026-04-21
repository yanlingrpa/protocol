package osgui

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
)

/*
* Locator defines a locator interface that provides methods for retrieving locator information,
* operating on the locator, and simulating input.
 */
type Locator interface {
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
	* Returns the locator rectangle relative to its parent window.
	 */
	GetWindowRect() basic.Rect
	/*
	* Returns the position and size of the current locator relative to its body area.
	* Body: The client area of a window, excluding non-client parts such as title bar and borders,
	* usually where the application content is displayed.
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
	* locator_point: Relative position within the locator.
	* return: Relative position within the window.
	 */
	TransToWindow(locator_point *basic.Point) *basic.Point

	/*
	* Converts a locator-relative position to a relative position within the client area.
	* locator_point: Relative position within the locator.
	* return: Relative position within the client area.
	 */
	TransToBody(locator_point *basic.Point) *basic.Point

	/*
	* Converts an absolute screen position to a relative position within the current locator.
	* screen_pos: Absolute position on the screen.
	* return: Relative position within the locator; returns nil if the point is outside the locator.
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	/*
	* Converts a position relative to the parent window to a relative position within the current locator.
	* window_pos: Relative position within the window.
	* return: Relative position within the locator; returns nil if the point is outside the locator.
	 */
	TransFromWindow(window_pos *basic.Point) *basic.Point

	/*
	* Converts a position relative to the client area to a relative position within the current locator.
	* body_pos: Relative position within the client area.
	* return: Relative position within the locator; returns nil if the point is outside the locator.
	 */
	TransFromBody(body_pos *basic.Point) *basic.Point

	/*
	* Gets a sub-locator of the current locator, clipped to the intersection with current bounds.
	* locator_point: Relative position of the sub-locator within the current locator.
	* size: Size of the sub-locator.
	 */
	SubLocator(locator_point basic.Point, size basic.Size) Locator

	/*
	* Finds sub-locators containing visually similar images within the current locator area.
	* image: Path/URL/Base64 image string relative to the current project.
	* sim: Similarity threshold in range 0.1~1.0; higher means more similar.
	* return: List of matched sub-locators; returns nil if none are found.
	* Results are sorted by similarity descending, and items below the threshold are filtered out.
	 */
	ImageLocator(image string, sim float32) ([]Locator, error)

	/*
	* Waits for sub-locators containing visually similar images to appear in the current locator area.
	* timeout: Timeout duration.
	* image: Path/URL/Base64 image string relative to the current project.
	* sim: Similarity threshold in range 0.1~1.0; higher means more similar.
	* return: On success, returns matched sub-locators sorted by similarity descending;
	* results below sim are filtered out. Returns error on timeout or recognition failure.
	 */
	WaitForImage(timeout time.Duration, image string, sim float32) ([]Locator, error)

	/*
	* Finds sub-locators containing specified text within the current locator area.
	* texts: Text list to find. Multiple texts mean all must be present.
	* return: List of matched sub-locators; returns nil if none are found.
	* Only locators containing all texts are returned, sorted by area from small to large.
	 */
	TextLocator(texts ...string) ([]Locator, error)

	/*
	* Waits for sub-locators containing specified text to appear in the current locator area.
	* timeout: Timeout duration.
	* texts: Text list to wait for. Multiple texts mean all must be present.
	* return: On success, returns matched sub-locators that contain all texts,
	* sorted by area from small to large. Returns error on timeout or recognition failure.
	 */
	WaitForText(timeout time.Duration, texts ...string) ([]Locator, error)

	/*
	* Finds card-like sub-locators in the current locator area.
	* A card-like sub-locator is an independent card-shaped region that usually includes
	* graphics and text, commonly seen in list items and buttons.
	* min_size: Minimum size of sub-locators to find; nil means no minimum limit.
	* max_size: Maximum size of sub-locators to find; nil means no maximum limit.
	* return: List of matched sub-locators; returns nil if none are found.
	* Results are sorted by coordinates from left to right and top to bottom.
	 */
	CardLocator(min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* Waits for card-like sub-locators to appear in the current locator area.
	* A card-like sub-locator is an independent card-shaped region that usually includes
	* graphics and text, commonly seen in list items and buttons.
	* timeout: Timeout duration.
	* min_size: Minimum size of sub-locators to find; nil means no minimum limit.
	* max_size: Maximum size of sub-locators to find; nil means no maximum limit.
	* return: On success, returns matched sub-locators sorted by coordinates
	* from left to right and top to bottom. Returns error on timeout or recognition failure.
	 */
	WaitForCard(timeout time.Duration, min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* Uses the vision module to find sub-locators with specific visual shapes
	* in the current locator area.
	* description: Description of the target shape, for example: a red circular button,
	* or a blue rectangular input box.
	* min_size: Minimum size of sub-locators to find; nil means no minimum limit.
	* max_size: Maximum size of sub-locators to find; nil means no maximum limit.
	* return: List of matched sub-locators; returns nil if none are found.
	* Results are sorted by coordinates from left to right and top to bottom.
	 */
	VisionLocator(description string, min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* Waits for sub-locators with specific visual shapes to appear,
	* using the vision module in the current locator area.
	* timeout: Timeout duration.
	* description: Description of the target shape, for example: a red circular button,
	* or a blue rectangular input box.
	* min_size: Minimum size of sub-locators to find; nil means no minimum limit.
	* max_size: Maximum size of sub-locators to find; nil means no maximum limit.
	* return: On success, returns matched sub-locators sorted by coordinates
	* from left to right and top to bottom. Returns error on timeout or recognition failure.
	 */
	WaitForVision(timeout time.Duration, description string, min_size, max_size *basic.Size) ([]Locator, error)

	/*
	* Performs OCR on the current locator area and returns the recognition result.
	* confidence: Recognition confidence threshold in range 0.1~1.0; higher means more strict.
	 */
	OcrRead(confidence float64) (*basic.OcrResult, error)

	/*
	* Moves the mouse to a specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means the center position of the current locator.
	 */
	MouseMove(locator_point *basic.Point) error

	/*
	* Presses the mouse button down at a specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means the center position of the current locator.
	 */
	MouseDown(locator_point *basic.Point) error

	/*
	* Releases the mouse button at a specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means the center position of the current locator.
	 */
	MouseUp(locator_point *basic.Point) error

	/*
	* Drags from from_locator_point in the current locator to to_locator_point in the target locator.
	* from_locator_point: Relative start position in current locator; nil means current mouse
	* position in current locator. If the mouse is outside, the locator center is used.
	* to_locator_point: Relative end position in target locator; nil means current mouse
	* position in target locator. If the mouse is outside, the target locator center is used.
	 */
	DragTo(target Locator, from_locator_point *basic.Point, to_locator_point *basic.Point) error

	/*
	* Focuses the current locator.
	* If the mouse is already inside the locator, it is considered focused.
	* Otherwise, the mouse is moved to the locator center and clicked to obtain focus.
	 */
	Focus() error

	/*
	* Performs a mouse click at the specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means current mouse position in the locator.
	* If the mouse is not inside the locator, the locator center is used.
	 */
	Click(locator_point *basic.Point) error

	/*
	* Performs a mouse double-click at the specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means current mouse position in the locator.
	* If the mouse is not inside the locator, the locator center is used.
	 */
	DoubleClick(locator_point *basic.Point) error

	/*
	* Performs a mouse right-click at the specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means current mouse position in the locator.
	* If the mouse is not inside the locator, the locator center is used.
	 */
	RightClick(locator_point *basic.Point) error

	/*
	* Gets the mouse position relative to the current locator.
	* return: Relative mouse position in the current locator; returns nil if outside.
	 */
	GetLocatorCursorPos() *basic.Point

	/*
	* Gets the IME caret position inside the current locator; returns nil if unavailable.
	* Example: When an input box is focused, the IME caret usually appears inside it,
	* and this method can retrieve its relative position in the locator.
	* return: Relative IME caret position in the locator; returns nil if unavailable.
	 */
	GetLocatorCaretPos() *basic.Point

	/*
	* Simulates keyboard input.
	* keys: Key sequence to input (for example: ctrl + alt + del).
	* Supported keys are defined by the Keyboard type.
	 */
	PressKeys(keys ...Keyboard) error

	/*
	* Indicates whether the current locator is in a text-editable state.
	* Example: After clicking an input box, it usually enters editable state,
	* and keyboard input can be sent to it.
	* return: true if text input is available; otherwise false.
	 */
	IsEditing() bool

	/*
	* Waits until the current locator enters a text-editable state or times out.
	* timeout: Timeout duration.
	* return: Sub-locator containing the text input caret.
	* Returns error on timeout; returns nil if no caret exists in the locator.
	 */
	WaitForEditing(timeout time.Duration) (Locator, error)

	/*
	* Reads text content from the current locator.
	* Equivalent to clicking the locator, selecting all text, copying to clipboard,
	* and then reading text from the clipboard.
	* If clipboard text cannot be read, OCR is attempted directly on the locator area.
	* return: Text content in the locator; returns error if content cannot be obtained.
	 */
	ReadText() (string, error)

	/*
	* Writes text into the current locator.
	* Equivalent to clicking the locator and then entering text.
	* return: nil on success, otherwise an error.
	 */
	WriteText(text string) error

	/*
	* Clears text content in the current locator.
	* Equivalent to clicking the locator, selecting all text, and pressing Delete
	* or Backspace to clear content.
	* return: nil on success, otherwise an error.
	 */
	ClearText() error

	/*
	* Indicates whether vertical scrolling is available in the current locator area.
	* return: true if vertical scrolling is supported; otherwise false.
	 */
	IsVerticalScroller() bool
	/*
	* Performs vertical scrolling in the current locator area.
	* up: Whether to scroll up; true for up, false for down.
	* lines: Number of lines to scroll; larger values scroll farther.
	 */
	ScrollVertical(up bool, lines int) error

	/*
	* Indicates whether horizontal scrolling is available in the current locator area.
	* return: true if horizontal scrolling is supported; otherwise false.
	 */
	IsHorizontalScroller() bool
	/*
	* Performs horizontal scrolling in the current locator area.
	* left: Whether to scroll left; true for left, false for right.
	* lines: Number of lines to scroll; larger values scroll farther.
	 */
	ScrollHorizontal(left bool, lines int) error
}
