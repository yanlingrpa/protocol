package gui

import "yanlingrpa.com/yanling/protocol/basic"

/*
* GuiWindow defines the common abstract interface for GUI windows across platforms.
 */
type GuiWindow interface {
	/*
	* Gets the unique identifier of the window.
	 */
	GetID() string
	/*
	* Gets the window title.
	 */
	GetWindowTitle() string
	/*
	* Exports current window information as key-value pairs.
	 */
	ToMap() map[string]any
	/*
	* Gets the client-area rectangle in screen coordinates.
	 */
	GetClientRect() basic.Rect

	/* Captures a snapshot of the current window area and returns image bytes.
	* gray: Whether to use grayscale.
	 */
	Snapshot(gray bool) ([]byte, error)

	/*
	* Activates the window.
	 */
	Activate() (bool, error)
	/*
	* Deactivates the window.
	 */
	DeActivate() (bool, error)

	/* Converts a window-relative position to an absolute screen position.
	* window_pos: Relative position inside the window.
	* return: Absolute position on the screen.
	 */
	TransToScreen(window_pos *basic.Point) *basic.Point

	/* Converts an absolute screen position to a window-relative position.
	* screen_pos: Absolute position on the screen.
	* return: Relative position inside the window; returns nil if the point is outside the window.
	 */
	TransFromScreen(screen_pos *basic.Point) *basic.Point

	/*
	* Reads clipboard text.
	 */
	ReadClipboard() (string, error)
	/*
	* Writes text to the clipboard.
	 */
	WriteClipboard(text string) error

	/*
	* Closes the window.
	 */
	Close() error
}
