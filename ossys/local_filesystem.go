package ossys

import "time"

/*
* LocalFilesystem defines capabilities for local filesystem operations.
* This interface covers common scenarios such as path handling, file I/O, and directory management.
*
* Path and permission rules:
* 1) All relative paths are resolved from DataRoot.
* 2) Files and directories under DataRoot are granted full permissions by default.
* 3) Files and directories under ScriptRoot are read-only.
* 4) Paths outside DataRoot and ScriptRoot require explicit authorization by the implementation.
 */
type LocalFilesystem interface {
	/*
	* DataRoot gets the root directory for script data storage.
	* This directory is intended for storing script-generated data and files.
	* Relative paths in this interface are resolved from this root.
	 */
	DataRoot() string

	/*
	* ScriptRoot gets the root directory for script files.
	* This directory is intended for storing script files.
	* Paths under this root are read-only.
	 */
	ScriptRoot() string

	/*
	* JoinPath joins paths starting from the root directory.
	* Returns the joined full path string based on DataRoot.
	 */
	JoinPath(path ...string) string

	/*
	* PathExists checks whether a path exists.
	* Returns whether it exists and any possible error.
	* Relative path is resolved from DataRoot.
	 */
	PathExists(path string) (bool, error)

	/*
	* IsDir checks whether a path is a directory.
	* Returns whether it is a directory and any possible error.
	* Relative path is resolved from DataRoot.
	 */
	IsDir(path string) (bool, error)

	/*
	* IsFile checks whether a path is a file.
	* Returns whether it is a file and any possible error.
	* Relative path is resolved from DataRoot.
	 */
	IsFile(path string) (bool, error)

	/*
	* ListDir lists entries in a directory.
	* Returns entry names under the target directory.
	* Relative path is resolved from DataRoot.
	 */
	ListDir(path string) ([]string, error)

	/*
	* StatPath gets metadata for a path.
	* Returns whether the path exists, whether it is a directory, size in bytes, and last modified time.
	* Relative path is resolved from DataRoot.
	 */
	StatPath(path string) (exists bool, isDir bool, size int64, modifiedAt time.Time, err error)

	/*
	* Mkdir creates a directory.
	* path is the directory path to create.
	* Relative path is resolved from DataRoot.
	 */
	Mkdir(path string) error

	/*
	* MkdirAll creates directories recursively.
	* If directories already exist, behavior depends on the implementation.
	* Relative path is resolved from DataRoot.
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
	* Relative path is resolved from DataRoot.
	 */
	WriteFile(filePath string, data []byte) error

	/*
	* ReadFile reads file content.
	* Returns the read byte data and any possible error.
	* Relative path is resolved from DataRoot.
	 */
	ReadFile(filePath string) ([]byte, error)

	/*
	* Remove deletes a file or directory.
	* path is the target path to delete.
	* Relative path is resolved from DataRoot.
	 */
	Remove(path string) error

	/*
	* RemoveAll recursively deletes a directory and all its contents.
	* path is the directory path to delete recursively.
	* Relative path is resolved from DataRoot.
	 */
	RemoveAll(path string) error

	/*
	* Rename renames or moves a file or directory.
	* src is the source path, and dst is the destination path.
	* Relative paths are resolved from DataRoot.
	 */
	Rename(src, dst string) error

	/*
	* CopyFile copies a file.
	* src is the source file, and dst is the destination file.
	* Relative paths are resolved from DataRoot.
	 */
	CopyFile(src, dst string) error
}
