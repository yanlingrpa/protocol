package osgui

import (
	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/internal/gui"
	"yanlingrpa.com/yanling/protocol/ossys"
)

/*
* OSGuiWindow defines the interface of a GUI window, providing methods for window information,
* window operations, and simulated input.
 */
type OSGuiWindow interface {
	gui.GuiWindow

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
	* Gets the window rectangle in screen coordinates.
	 */
	GetWindowRect() basic.Rect

	/*
	* Moves the window to the specified position.
	 */
	MoveTo(global_x, global_y int) (bool, error)
	/*
	* Resizes the window.
	 */
	ResizeTo(width, height int) (bool, error)

	/*
	* Gets the locator for the window client area.
	 */
	BodyLocator() (OSGuiLocator, error)
	/*
	* Gets a locator for the specified rectangle, clipped to the window bounds.
	 */
	RectLocator(rect basic.Rect) (OSGuiLocator, error)

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
	* Gets information about the monitor that contains the window.
	 */
	GetMonitor() ossys.MonitorInfo
}
