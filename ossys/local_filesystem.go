package ossys

import "time"

type LocalFilesystem interface {
	JoinDataPath(path ...string) string                // 从数据目录开始拼接路径
	PathExists(path string) (bool, error)              // 检查路径是否存在，返回是否存在和错误
	IsDir(path string) (bool, error)                   // 检查路径是否为目录
	IsFile(path string) (bool, error)                  // 检查路径是否为文件
	MkdirAll(path string) error                        // 创建目录，返回错误
	CreateTmpFile(expiredAt time.Time) (string, error) // 创建临时文件，返回文件路径和错误
	WriteFile(filePath string, data []byte) error      // 写入文件，返回错误
	ReadFile(filePath string) ([]byte, error)          // 读取文件，返回文件内容和错误
	Remove(path string) error                          // 删除文件或文件夹，返回错误
	RemoveAll(path string) error                       // 递归删除文件夹及其内容，返回错误
	Rename(src, dst string) error                      // 重命名或移动文件或目录，返回错误
	CopyFile(src, dst string) error                    // 复制文件，返回错误
}
