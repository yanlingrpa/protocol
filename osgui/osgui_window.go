package osgui

import (
	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/ossys"
)

/*
* GuiWindow defines the interface of a GUI window, providing methods for window information,
* window operations, and simulated input.
 */
type GuiWindow interface {
	/*
	* Gets the unique identifier of the window.
	 */
	GetID() string
	/*
	* Gets the executable path of the process that launched the window.
	 */
	GetInitiatorPath() string
	/*
	* Gets the application path associated with the window itself.
	 */
	GetWindowSchema() string
	/*
	* Gets the native window handle.
	 */
	GetHwnd() uintptr
	/*
	* Gets the window title.
	 */
	GetWindowTitle() string
	/*
	* Exports current window information as key-value pairs.
	 */
	ToMap() map[string]any
	/*
	* Gets the window rectangle in screen coordinates.
	 */
	GetWindowRect() basic.Rect
	/*
	* Gets the client-area rectangle in screen coordinates.
	 */
	GetClientRect() basic.Rect

	/* Captures a snapshot of the current window area and returns image bytes.
	* gray: Whether to use grayscale.
	 */
	Snapshot(gray bool) ([]byte, error)

	/*
	* Moves the window to the specified position.
	 */
	MoveTo(global_x, global_y int) (bool, error)
	/*
	* Resizes the window.
	 */
	ResizeTo(width, height int) (bool, error)
	/*
	* Activates the window.
	 */
	Activate() (bool, error)
	/*
	* Deactivates the window.
	 */
	DeActivate() (bool, error)

	/*
	* Gets the locator for the window client area.
	 */
	BodyLocator() (Locator, error)
	/*
	* Gets a locator for the specified rectangle, clipped to the window bounds.
	 */
	RectLocator(rect basic.Rect) (Locator, error)

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
	* Gets the mouse position relative to the current window; returns nil if outside the window.
	 */
	GetWindowCursorPos() *basic.Point
	/*
	* Gets the current IME caret position relative to the window; returns nil if outside the window.
	 */
	GetWindowCaretPos() *basic.Point

	/*
	* Simulates keyboard input at the current window.
	* This method will unfocus any control within the window that currently has focus,
	* then set focus to the entire window, and finally send the specified key sequence.
	* keys: Key sequence to input (for example: ctrl + alt + del).
	* Supported keys are defined by the Keyboard type.
	 */
	PressKeys(keys ...Keyboard) error

	/*
	* Reads clipboard text.
	 */
	ReadClipboard() (string, error)
	/*
	* Writes text to the clipboard.
	 */
	WriteClipboard(text string) error

	/*
	* Gets information about the monitor that contains the window.
	 */
	GetMonitor() ossys.MonitorInfo

	/*
	* Closes the window.
	 */
	Close() error
}
