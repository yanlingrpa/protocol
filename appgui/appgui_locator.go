package appgui

import (
	"time"

	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/internal/gui"
)

/*
* AppGuiLocator defines a locator interface for a rectangular region on a mobile app screen,
* providing methods for retrieving locator information, performing touch gestures,
* and simulating input.
 */
type AppGuiLocator interface {
	gui.GuiLocator

	/*
	* Gets a sub-locator of the current locator, clipped to the intersection with current bounds.
	* locator_point: Relative position of the sub-locator within the current locator.
	* size: Size of the sub-locator.
	 */
	SubLocator(locator_point basic.Point, size basic.Size) AppGuiLocator

	/*
	* Finds sub-locators containing visually similar images within the current locator area.
	* image: Path/URL/Base64 image string relative to the current project.
	* sim: Similarity threshold in range 0.1~1.0; higher means more similar.
	* return: List of matched sub-locators; returns nil if none are found.
	* Results are sorted by similarity descending, and items below the threshold are filtered out.
	 */
	ImageLocator(image string, sim float32) ([]AppGuiLocator, error)

	/*
	* Waits for sub-locators containing visually similar images to appear in the current locator area.
	* timeout: Timeout duration.
	* image: Path/URL/Base64 image string relative to the current project.
	* sim: Similarity threshold in range 0.1~1.0; higher means more similar.
	* return: On success, returns matched sub-locators sorted by similarity descending;
	* results below sim are filtered out. Returns error on timeout or recognition failure.
	 */
	WaitForImage(timeout time.Duration, image string, sim float32) ([]AppGuiLocator, error)

	/*
	* Finds sub-locators containing specified text within the current locator area.
	* texts: Text list to find. Multiple texts mean all must be present.
	* return: List of matched sub-locators; returns nil if none are found.
	* Only locators containing all texts are returned, sorted by area from small to large.
	 */
	TextLocator(texts ...string) ([]AppGuiLocator, error)

	/*
	* Waits for sub-locators containing specified text to appear in the current locator area.
	* timeout: Timeout duration.
	* texts: Text list to wait for. Multiple texts mean all must be present.
	* return: On success, returns matched sub-locators that contain all texts,
	* sorted by area from small to large. Returns error on timeout or recognition failure.
	 */
	WaitForText(timeout time.Duration, texts ...string) ([]AppGuiLocator, error)

	/*
	* Finds card-like sub-locators in the current locator area.
	* A card-like sub-locator is an independent card-shaped region that usually includes
	* graphics and text, commonly seen in list items and buttons.
	* min_size: Minimum size of sub-locators to find; nil means no minimum limit.
	* max_size: Maximum size of sub-locators to find; nil means no maximum limit.
	* return: List of matched sub-locators; returns nil if none are found.
	* Results are sorted by coordinates from left to right and top to bottom.
	 */
	CardLocator(min_size, max_size *basic.Size) ([]AppGuiLocator, error)

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
	WaitForCard(timeout time.Duration, min_size, max_size *basic.Size) ([]AppGuiLocator, error)

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
	VisionLocator(description string, min_size, max_size *basic.Size) ([]AppGuiLocator, error)

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
	WaitForVision(timeout time.Duration, description string, min_size, max_size *basic.Size) ([]AppGuiLocator, error)

	/*
	* Moves the touch point to a specified position inside the locator without lifting.
	* locator_point: Relative position within the locator.
	* nil means the center position of the current locator.
	 */
	TouchMove(locator_point *basic.Point) error

	/*
	* Presses a finger down at the current touch position within the locator.
	 */
	TouchDown() error

	/*
	* Lifts the finger at the current touch position within the locator.
	 */
	TouchUp() error

	/*
	* Performs a swipe gesture from a start position to an end position within the current locator.
	* Both positions are relative coordinates within the current locator.
	* from_locator_point: Relative start position; nil means the current touch position,
	* or the locator center if no active touch.
	* to_locator_point: Relative end position; nil means the current touch position,
	* or the locator center if no active touch.
	 */
	SwipeTo(from_locator_point *basic.Point, to_locator_point *basic.Point) error

	/*
	* Performs a pinch gesture in the current locator area.
	* spread: true to spread fingers (zoom in), false to pinch fingers (zoom out).
	* scale: Scale factor for the gesture distance, range 0.1~1.0; larger means more spread/pinch.
	 */
	Pinch(spread bool, scale float32) error

	/*
	* Performs a single tap at the specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means current touch position in the locator.
	* If there is no active touch inside the locator, the locator center is used.
	 */
	Tap(locator_point *basic.Point) error

	/*
	* Performs a double tap at the specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means current touch position in the locator.
	* If there is no active touch inside the locator, the locator center is used.
	 */
	DoubleTap(locator_point *basic.Point) error

	/*
	* Performs a long press at the specified position inside the locator.
	* locator_point: Relative position within the locator.
	* nil means current touch position in the locator.
	* If there is no active touch inside the locator, the locator center is used.
	 */
	LongPress(locator_point *basic.Point) error

	/*
	* Gets the current touch position relative to the locator.
	* return: Relative touch position in the locator; returns nil if no active touch inside.
	 */
	GetLocatorTouchPos() *basic.Point

	/*
	* Waits until the current locator enters a text-editable state or times out.
	* timeout: Timeout duration.
	* return: Sub-locator containing the text input caret.
	* Returns error on timeout; returns nil if no caret exists in the locator.
	 */
	WaitForEditing(timeout time.Duration) (AppGuiLocator, error)

	/*
	* Simulates pressing a hardware or system key while the locator is focused.
	* keys: Key sequence to press, as defined by the AppKey type.
	 */
	PressKeys(keys ...AppKey) error
}
