package ossys

import (
	"github.com/yanlingrpa/protocol/basic"
)

type MonitorInfo interface {
	GetBounds() basic.Rect   // 获取屏幕的边界位置和大小
	GetWorkArea() basic.Rect // 获取屏幕的工作区位置和大小
	GetDPI() uint32          // 获取屏幕的DPI缩放比例 (水平DPI)
	IsPrimary() bool         // 是否为主屏幕
}
