package ossys

import "time"

/*
* LocalFilesystem 定义本地文件系统操作的能力接口。
* 该接口覆盖常见场景，例如路径处理、文件 I/O 和目录管理。
*
* 路径与权限规则：
* 1) 所有相对路径都从 DataRoot 解析。
* 2) DataRoot 下的文件和目录默认具备完整权限。
* 3) ScriptRoot 下的文件和目录为只读。
* 4) 位于 DataRoot 和 ScriptRoot 之外的路径需要由实现显式授权。
 */
type LocalFilesystem interface {
	/*
	* DataRoot 获取脚本数据存储根目录。
	* 该目录用于存放脚本生成的数据和文件。
	* 该接口中的相对路径均相对于此根目录解析。
	 */
	DataRoot() string

	/*
	* ScriptRoot 获取脚本文件根目录。
	* 该目录用于存放脚本文件。
	* 该根目录下的路径为只读。
	 */
	ScriptRoot() string

	/*
	* JoinPath 从根目录开始拼接路径。
	* 返回基于 DataRoot 的完整路径字符串。
	 */
	JoinPath(path ...string) string

	/*
	* PathExists 检查路径是否存在。
	* 返回是否存在以及可能的错误。
	* 相对路径从 DataRoot 解析。
	 */
	PathExists(path string) (bool, error)

	/*
	* IsDir 检查路径是否为目录。
	* 返回是否为目录以及可能的错误。
	* 相对路径从 DataRoot 解析。
	 */
	IsDir(path string) (bool, error)

	/*
	* IsFile 检查路径是否为文件。
	* 返回是否为文件以及可能的错误。
	* 相对路径从 DataRoot 解析。
	 */
	IsFile(path string) (bool, error)

	/*
	* ListDir 列出目录中的条目。
	* 返回目标目录下的条目名称。
	* 相对路径从 DataRoot 解析。
	 */
	ListDir(path string) ([]string, error)

	/*
	* StatPath 获取路径元数据。
	* 返回该路径是否存在、是否为目录、大小（字节数）和最后修改时间。
	* 相对路径从 DataRoot 解析。
	 */
	StatPath(path string) (exists bool, isDir bool, size int64, modifiedAt time.Time, err error)

	/*
	* Mkdir 创建目录。
	* path 是要创建的目录路径。
	* 相对路径从 DataRoot 解析。
	 */
	Mkdir(path string) error

	/*
	* MkdirAll 递归创建目录。
	* 如果目录已存在，行为取决于具体实现。
	* 相对路径从 DataRoot 解析。
	 */
	MkdirAll(path string) error

	/*
	* CreateTmpFile 创建临时文件。
	* expiredAt 表示临时文件的过期时间；返回文件路径。
	 */
	CreateTmpFile(expiredAt time.Time) (string, error)

	/*
	* WriteFile 向文件写入内容。
	* data 是要写入的字节数据。
	* 相对路径从 DataRoot 解析。
	 */
	WriteFile(filePath string, data []byte) error

	/*
	* ReadFile 读取文件内容。
	* 返回读取到的字节数据和可能的错误。
	* 相对路径从 DataRoot 解析。
	 */
	ReadFile(filePath string) ([]byte, error)

	/*
	* Remove 删除文件或目录。
	* path 是要删除的目标路径。
	* 相对路径从 DataRoot 解析。
	 */
	Remove(path string) error

	/*
	* RemoveAll 递归删除目录及其所有内容。
	* path 是要递归删除的目录路径。
	* 相对路径从 DataRoot 解析。
	 */
	RemoveAll(path string) error

	/*
	* Rename 重命名或移动文件或目录。
	* src 是源路径，dst 是目标路径。
	* 相对路径从 DataRoot 解析。
	 */
	Rename(src, dst string) error

	/*
	* CopyFile 复制文件。
	* src 是源文件，dst 是目标文件。
	* 相对路径从 DataRoot 解析。
	 */
	CopyFile(src, dst string) error
}
