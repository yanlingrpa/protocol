package ossys

/*
* DeviceInfo 定义用于查询设备信息的能力接口。
* 该接口用于检索与操作系统、硬件、用户和显示器相关的信息。
 */
type BrokerInfo interface {
	/*
	* OS 获取操作系统名称。
	* 示例返回值：Windows、Linux、Darwin。
	 */
	OS() string

	/*
	* OSVersion 获取操作系统版本。
	* 返回值通常是系统版本字符串。
	 */
	OSVersion() string

	/*
	* DeviceId 获取唯一设备标识。
	* 返回值可用于区分设备身份。
	 */
	DeviceId() string

	/*
	* DeviceName 获取设备名称。
	* 返回值通常是系统设置中配置的设备名称。
	 */
	DeviceName() string

	/*
	* NumLogicCPU 获取逻辑 CPU 核心数量。
	* 返回值是逻辑处理器的数量。
	 */
	NumLogicCPU() int

	/*
	* HasNvidiaGPU 检查是否存在 Nvidia GPU。
	* 若检测到 Nvidia 显卡，则返回 true。
	 */
	HasNvidiaGPU() bool

	/*
	* GetGpuMemoryMB 获取 GPU 显存大小。
	* 返回值单位为 MB。
	 */
	GetGpuMemoryMB() int

	/*
	* GetComputerName 获取计算机名称。
	* 返回值是主机名或系统计算机名称。
	 */
	GetComputerName() string

	/*
	* GetUserName 获取当前用户名。
	* 返回值是当前登录账户的用户名。
	 */
	GetUserName() string

	/*
	* GetMonitors 获取所有显示器的信息。
	* 返回值是显示器信息列表。
	 */
	GetMonitors() []MonitorInfo

	/*
	* GetPrimaryMonitor 获取主显示器的信息。
	* 返回值是当前系统的主显示器。
	 */
	GetPrimaryMonitor() MonitorInfo
}
