package elliecore

/*
#cgo LDFLAGS: -L${SRCDIR}/../rustmods/elliecore/target/debug -lelliecore
#include <stdlib.h>
char* run_cmd(const char* cmd);
char* run_cmd_with_env(const char* cmd, const char* envs);
char* read_file(const char* path);
char* write_file(const char* path, const char* content);
char* append_file(const char* path, const char* content);
char* delete_file(const char* path);
char* list_dir(const char* path);
char* get_env(const char* key);
char* set_env(const char* key, const char* value);
char* get_cwd();
char* change_dir(const char* path);
int file_exists(const char* path);
char* file_metadata(const char* path);
*/
import "C"
import (
	"unsafe"
)

// RunCmd executes a shell command using the Rust FFI and returns the result as a string.
// This is a low level, cross platform abstraction for shell command execution.
func RunCmd(cmd string) string {
	input := C.CString(cmd)
	defer C.free(unsafe.Pointer(input))

	result := C.run_cmd(input)
	defer C.free(unsafe.Pointer(result))

	output := C.GoString(result)
	return output
}

// RunCmdWithEnv executes a shell command with environment variables using the Rust FFI.
// envs should be a semicolon-separated list of key=value pairs, e.g. "FOO=bar;BAZ=qux"
func RunCmdWithEnv(cmd string, envs string) string {
	cmdC := C.CString(cmd)
	defer C.free(unsafe.Pointer(cmdC))
	envsC := C.CString(envs)
	defer C.free(unsafe.Pointer(envsC))

	result := C.run_cmd_with_env(cmdC, envsC)
	defer C.free(unsafe.Pointer(result))

	output := C.GoString(result)
	return output
}

// ReadFile reads a file and returns its contents as a string using the Rust FFI.
func ReadFile(path string) string {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))

	result := C.read_file(pathC)
	defer C.free(unsafe.Pointer(result))

	output := C.GoString(result)
	return output
}

// WriteFile writes content to a file (overwrites if exists).
func WriteFile(path, content string) string {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	contentC := C.CString(content)
	defer C.free(unsafe.Pointer(contentC))
	result := C.write_file(pathC, contentC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// AppendFile appends content to a file (creates if not exists).
func AppendFile(path, content string) string {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	contentC := C.CString(content)
	defer C.free(unsafe.Pointer(contentC))
	result := C.append_file(pathC, contentC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// DeleteFile removes a file.
func DeleteFile(path string) string {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	result := C.delete_file(pathC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// ListDir lists files and directories in a path (newline separated).
func ListDir(path string) string {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	result := C.list_dir(pathC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// GetEnv retrieves the value of an environment variable.
func GetEnv(key string) string {
	keyC := C.CString(key)
	defer C.free(unsafe.Pointer(keyC))
	result := C.get_env(keyC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// SetEnv sets an environment variable for the current process.
func SetEnv(key, value string) string {
	keyC := C.CString(key)
	defer C.free(unsafe.Pointer(keyC))
	valueC := C.CString(value)
	defer C.free(unsafe.Pointer(valueC))
	result := C.set_env(keyC, valueC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// GetCwd returns the current working directory.
func GetCwd() string {
	result := C.get_cwd()
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// ChangeDir changes the current working directory.
func ChangeDir(path string) string {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	result := C.change_dir(pathC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// FileExists checks if a file or directory exists.
func FileExists(path string) bool {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	return C.file_exists(pathC) == 1
}

// FileMetadata returns file size, readonly, and modified time as JSON.
func FileMetadata(path string) string {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))
	result := C.file_metadata(pathC)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}
