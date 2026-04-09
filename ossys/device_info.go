package ossys

type DeviceInfo interface {
	OS() string                     // 获取操作系统名称
	OSVersion() string              // 获取操作系统版本
	DeviceId() string               // 获取设备唯一标识符
	DeviceName() string             // 获取设备名称
	NumLogicCPU() int               // 获取逻辑CPU核心数
	HasNvidiaGPU() bool             // 是否有Nvidia GPU
	GetGpuMemoryMB() int            // 获取GPU内存大小（MB）
	GetComputerName() string        // 获取计算机名称
	GetUserName() string            // 获取当前用户名称
	GetMonitors() []MonitorInfo     // 获取所有显示器信息
	GetPrimaryMonitor() MonitorInfo // 获取主显示器信息
}
