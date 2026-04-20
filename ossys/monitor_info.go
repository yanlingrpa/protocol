package ossys

import (
	"yanlingrpa.com/yanling/protocol/basic"
)

/*
* MonitorInfo 定义显示器信息查询能力
* 该接口用于获取边界、工作区、DPI 与主屏标识
 */
type MonitorInfo interface {
	/*
	* GetBounds 获取屏幕边界位置和大小
	* 返回值为显示器完整边界矩形
	 */
	GetBounds() basic.Rect

	/*
	* GetWorkArea 获取屏幕工作区位置和大小
	* 工作区通常不包含任务栏等系统占用区域
	 */
	GetWorkArea() basic.Rect

	/*
	* GetDPI 获取屏幕 DPI 缩放信息
	* 返回值为水平 DPI 或缩放相关数值
	 */
	GetDPI() uint32

	/*
	* IsPrimary 判断是否为主屏幕
	* 返回 true 表示当前显示器为主显示器
	 */
	IsPrimary() bool
}
