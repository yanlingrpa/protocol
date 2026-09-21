package ossys

import (
	"yanlingrpa.com/yanling/protocol/basic"
)

/*
* MonitorInfo 定义查询显示器信息的能力接口。
* 该接口用于获取边界、工作区、DPI 以及主屏幕状态等信息。
 */
type MonitorInfo interface {
	/*
	* GetBounds 获取显示器的边界位置和尺寸。
	* 返回完整的显示器边界矩形区域。
	 */
	GetBounds() basic.Rect

	/*
	* GetWorkArea 获取显示器的工作区位置和尺寸。
	* 工作区通常不包含任务栏等系统占用区域。
	 */
	GetWorkArea() basic.Rect

	/*
	* GetDPI 获取显示器的 DPI 缩放信息。
	* 返回值为横向 DPI 或相关缩放值。
	 */
	GetDPI() uint32

	/*
	* IsPrimary 表示当前显示器是否为主显示器。
	* 当当前显示器为主显示器时返回 true。
	 */
	IsPrimary() bool
}
