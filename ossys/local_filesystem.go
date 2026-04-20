package ossys

import "time"

/*
* LocalFilesystem 定义本地文件系统操作能力
* 该接口覆盖路径处理、文件读写与目录管理等常见场景
 */
type LocalFilesystem interface {
	/*
	* JoinDataPath 从数据目录开始拼接路径
	* 返回拼接后的完整路径字符串
	 */
	JoinDataPath(path ...string) string

	/*
	* PathExists 检查路径是否存在
	* 返回是否存在以及可能的错误
	 */
	PathExists(path string) (bool, error)

	/*
	* IsDir 检查路径是否为目录
	* 返回是否为目录以及可能的错误
	 */
	IsDir(path string) (bool, error)

	/*
	* IsFile 检查路径是否为文件
	* 返回是否为文件以及可能的错误
	 */
	IsFile(path string) (bool, error)

	/*
	* MkdirAll 递归创建目录
	* 若目录已存在，行为由具体实现决定
	 */
	MkdirAll(path string) error

	/*
	* CreateTmpFile 创建临时文件
	* expiredAt 表示临时文件过期时间，返回文件路径
	 */
	CreateTmpFile(expiredAt time.Time) (string, error)

	/*
	* WriteFile 写入文件内容
	* data 为待写入的字节数据
	 */
	WriteFile(filePath string, data []byte) error

	/*
	* ReadFile 读取文件内容
	* 返回读取到的字节数据和可能的错误
	 */
	ReadFile(filePath string) ([]byte, error)

	/*
	* Remove 删除文件或文件夹
	* path 为待删除的目标路径
	 */
	Remove(path string) error

	/*
	* RemoveAll 递归删除目录及其所有内容
	* path 为待递归删除的目录路径
	 */
	RemoveAll(path string) error

	/*
	* Rename 重命名或移动文件或目录
	* src 为源路径，dst 为目标路径
	 */
	Rename(src, dst string) error

	/*
	* CopyFile 复制文件
	* src 为源文件，dst 为目标文件
	 */
	CopyFile(src, dst string) error
}
