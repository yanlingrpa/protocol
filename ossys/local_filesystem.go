package ossys

import "time"

/*
* LocalFilesystem defines capabilities for local filesystem operations.
* This interface covers common scenarios such as path handling, file I/O, and directory management.
 */
type LocalFilesystem interface {
	/*
	* JoinDataPath joins paths starting from the data directory.
	* Returns the joined full path string.
	 */
	JoinDataPath(path ...string) string

	/*
	* PathExists checks whether a path exists.
	* Returns whether it exists and any possible error.
	 */
	PathExists(path string) (bool, error)

	/*
	* IsDir checks whether a path is a directory.
	* Returns whether it is a directory and any possible error.
	 */
	IsDir(path string) (bool, error)

	/*
	* IsFile checks whether a path is a file.
	* Returns whether it is a file and any possible error.
	 */
	IsFile(path string) (bool, error)

	/*
	* MkdirAll creates directories recursively.
	* If directories already exist, behavior depends on the implementation.
	 */
	MkdirAll(path string) error

	/*
	* CreateTmpFile creates a temporary file.
	* expiredAt indicates the expiration time of the temporary file; returns the file path.
	 */
	CreateTmpFile(expiredAt time.Time) (string, error)

	/*
	* WriteFile writes content to a file.
	* data is the byte data to write.
	 */
	WriteFile(filePath string, data []byte) error

	/*
	* ReadFile reads file content.
	* Returns the read byte data and any possible error.
	 */
	ReadFile(filePath string) ([]byte, error)

	/*
	* Remove deletes a file or directory.
	* path is the target path to delete.
	 */
	Remove(path string) error

	/*
	* RemoveAll recursively deletes a directory and all its contents.
	* path is the directory path to delete recursively.
	 */
	RemoveAll(path string) error

	/*
	* Rename renames or moves a file or directory.
	* src is the source path, and dst is the destination path.
	 */
	Rename(src, dst string) error

	/*
	* CopyFile copies a file.
	* src is the source file, and dst is the destination file.
	 */
	CopyFile(src, dst string) error
}
