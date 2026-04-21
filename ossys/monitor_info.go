package ossys

import (
	"yanlingrpa.com/yanling/protocol/basic"
)

/*
* MonitorInfo defines capabilities for querying monitor information.
* This interface is used to get bounds, work area, DPI, and primary-screen status.
 */
type MonitorInfo interface {
	/*
	* GetBounds gets the monitor bounds position and size.
	* Returns the full monitor bounds rectangle.
	 */
	GetBounds() basic.Rect

	/*
	* GetWorkArea gets the monitor work-area position and size.
	* The work area usually excludes system-occupied areas such as the taskbar.
	 */
	GetWorkArea() basic.Rect

	/*
	* GetDPI gets DPI scaling information of the monitor.
	* The return value is horizontal DPI or a related scaling value.
	 */
	GetDPI() uint32

	/*
	* IsPrimary indicates whether this is the primary monitor.
	* Returns true when the current monitor is the primary monitor.
	 */
	IsPrimary() bool
}
