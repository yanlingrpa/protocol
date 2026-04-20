package ossys

/*
* DeviceInfo 定义设备信息查询能力
* 该接口用于获取操作系统、硬件、用户与显示器相关信息
 */
type DeviceInfo interface {
	/*
	* OS 获取操作系统名称
	* 返回值示例：Windows、Linux、Darwin
	 */
	OS() string

	/*
	* OSVersion 获取操作系统版本
	* 返回值通常为系统版本字符串
	 */
	OSVersion() string

	/*
	* DeviceId 获取设备唯一标识符
	* 返回值可用于设备维度的身份区分
	 */
	DeviceId() string

	/*
	* DeviceName 获取设备名称
	* 返回值通常为系统设置中的设备名
	 */
	DeviceName() string

	/*
	* NumLogicCPU 获取逻辑 CPU 核心数
	* 返回值为逻辑处理器数量
	 */
	NumLogicCPU() int

	/*
	* HasNvidiaGPU 判断是否存在 Nvidia GPU
	* 返回 true 表示检测到 Nvidia 显卡
	 */
	HasNvidiaGPU() bool

	/*
	* GetGpuMemoryMB 获取 GPU 显存大小
	* 返回值单位为 MB
	 */
	GetGpuMemoryMB() int

	/*
	* GetComputerName 获取计算机名称
	* 返回值为主机名或系统计算机名
	 */
	GetComputerName() string

	/*
	* GetUserName 获取当前用户名称
	* 返回值为当前登录账户名
	 */
	GetUserName() string

	/*
	* GetMonitors 获取所有显示器信息
	* 返回值为显示器信息列表
	 */
	GetMonitors() []MonitorInfo

	/*
	* GetPrimaryMonitor 获取主显示器信息
	* 返回值为当前系统主显示器
	 */
	GetPrimaryMonitor() MonitorInfo
}
