package appgui

import (
	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/internal/gui"
)

/*
* AppGuiWindow defines the interface of a mobile app window (screen), providing methods
* for window information, app operations, and simulated touch/key input.
 */
type AppGuiWindow interface {
	gui.GuiWindow

	/*
	* Gets the package name (Android) or bundle ID (iOS) of the app.
	 */
	GetAppPackage() string
	/*
	* Gets the current activity name (Android) or view controller name (iOS).
	 */
	GetAppActivity() string

	/*
	* Gets the screen rectangle in device coordinates.
	 */
	GetScreenRect() basic.Rect

	/*
	* Gets the locator for the window content area.
	 */
	BodyLocator() (AppGuiLocator, error)
	/*
	* Gets a locator for the specified rectangle, clipped to the window bounds.
	 */
	RectLocator(rect basic.Rect) (AppGuiLocator, error)

	/*
	* Gets the current touch position relative to the window; returns nil if no active touch in the window.
	 */
	GetWindowTouchPos() *basic.Point

	/*
	* Simulates pressing a hardware or system key.
	* key: The key to press, as defined by the AppKey type.
	 */
	PressKey(key AppKey) error
}
