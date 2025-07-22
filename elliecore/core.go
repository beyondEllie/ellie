package elliecore

import "github.com/tacheraSasi/ellie/utils"

// RunCmd executes a shell command using the Rust FFI and returns the result as a string.
// This is a low level, cross platform abstraction for shell command execution.
func RunCmd(cmd string) string {
	utils.RunCommand([]string{"sh", "-c", cmd}, "Error running command:")
	return ""
}

// RunCmdWithEnv executes a shell command with environment variables using the Rust FFI.
// envs should be a semicolon-separated list of key=value pairs, e.g. "FOO=bar;BAZ=qux"
func RunCmdWithEnv(cmd string, envs string) string {
	utils.RunCommand([]string{"sh", "-c", cmd}, "Error running command:")
	return ""
}

// ReadFile reads a file and returns its contents as a string using the Rust FFI.
func ReadFile(path string) string {
	return ""
}

// WriteFile writes content to a file (overwrites if exists).
func WriteFile(path, content string) string {
	return ""
}

// AppendFile appends content to a file (creates if not exists).
func AppendFile(path, content string) string {
	return ""
}

// DeleteFile removes a file.
func DeleteFile(path string) string {
	return ""
}

// ListDir lists files and directories in a path (newline separated).
func ListDir(path string) string {
	return ""
}

// GetEnv retrieves the value of an environment variable.
func GetEnv(key string) string {
	return ""
}

// SetEnv sets an environment variable for the current process.
func SetEnv(key, value string) string {
	return ""
}

// GetCwd returns the current working directory.
func GetCwd() string {
	return ""
}

// ChangeDir changes the current working directory.
func ChangeDir(path string) string {
	return ""
}

// FileExists checks if a file or directory exists.
func FileExists(path string) bool {
	return false
}

// FileMetadata returns file size, readonly, and modified time as JSON.
func FileMetadata(path string) string {
	return ""
}
