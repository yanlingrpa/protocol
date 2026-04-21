package ossys

/*
* DeviceInfo defines capabilities for querying device information.
* This interface is used to retrieve information related to the OS, hardware, user, and monitors.
 */
type DeviceInfo interface {
	/*
	* OS gets the operating system name.
	* Example return values: Windows, Linux, Darwin.
	 */
	OS() string

	/*
	* OSVersion gets the operating system version.
	* The return value is usually a system version string.
	 */
	OSVersion() string

	/*
	* DeviceId gets the unique device identifier.
	* The return value can be used to distinguish device identity.
	 */
	DeviceId() string

	/*
	* DeviceName gets the device name.
	* The return value is usually the device name configured in system settings.
	 */
	DeviceName() string

	/*
	* NumLogicCPU gets the number of logical CPU cores.
	* The return value is the number of logical processors.
	 */
	NumLogicCPU() int

	/*
	* HasNvidiaGPU checks whether an Nvidia GPU exists.
	* Returns true if an Nvidia graphics card is detected.
	 */
	HasNvidiaGPU() bool

	/*
	* GetGpuMemoryMB gets GPU memory size.
	* The return value is in MB.
	 */
	GetGpuMemoryMB() int

	/*
	* GetComputerName gets the computer name.
	* The return value is the hostname or system computer name.
	 */
	GetComputerName() string

	/*
	* GetUserName gets the current user name.
	* The return value is the name of the currently logged-in account.
	 */
	GetUserName() string

	/*
	* GetMonitors gets information for all monitors.
	* The return value is a list of monitor information.
	 */
	GetMonitors() []MonitorInfo

	/*
	* GetPrimaryMonitor gets information for the primary monitor.
	* The return value is the current system primary monitor.
	 */
	GetPrimaryMonitor() MonitorInfo
}
