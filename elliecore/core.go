package elliecore

/*
Package elliecore provides Go bindings for the Rust elliecore library.

This package bridges Go and Rust through CGO, allowing Go code to call 
high-performance Rust implementations for system operations like file I/O,
command execution, and environment management.

The Rust library is built as a static library (libelliecore.a) and linked
during Go compilation. All C string memory is properly managed to prevent leaks.
*/

/*
#cgo LDFLAGS: -L../rustmods/elliecore/target/debug -lelliecore
#include <stdlib.h>

// Rust FFI function declarations
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

// cStringToGoString converts a C string to Go string and frees the C string
func cStringToGoString(cstr *C.char) string {
	if cstr == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// RunCmd executes a shell command and returns the output or error.
func RunCmd(cmd string) string {
	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))
	
	result := C.run_cmd(cCmd)
	return cStringToGoString(result)
}

// RunCmdWithEnv executes a shell command with additional environment variables.
func RunCmdWithEnv(cmd string, envs string) string {
	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))
	cEnvs := C.CString(envs)
	defer C.free(unsafe.Pointer(cEnvs))
	
	result := C.run_cmd_with_env(cCmd, cEnvs)
	return cStringToGoString(result)
}

// ReadFile reads the content of a file.
func ReadFile(path string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	
	result := C.read_file(cPath)
	return cStringToGoString(result)
}

// WriteFile overwrites a file with given content.
func WriteFile(path, content string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cContent := C.CString(content)
	defer C.free(unsafe.Pointer(cContent))
	
	result := C.write_file(cPath, cContent)
	return cStringToGoString(result)
}

// AppendFile appends content to a file, creates it if not exists.
func AppendFile(path, content string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cContent := C.CString(content)
	defer C.free(unsafe.Pointer(cContent))
	
	result := C.append_file(cPath, cContent)
	return cStringToGoString(result)
}

// DeleteFile removes a file.
func DeleteFile(path string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	
	result := C.delete_file(cPath)
	return cStringToGoString(result)
}

// ListDir lists files and directories in a path.
func ListDir(path string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	
	result := C.list_dir(cPath)
	return cStringToGoString(result)
}

// GetEnv returns the value of an environment variable.
func GetEnv(key string) string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	
	result := C.get_env(cKey)
	return cStringToGoString(result)
}

// SetEnv sets an environment variable in the current process.
func SetEnv(key, value string) string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	
	result := C.set_env(cKey, cValue)
	return cStringToGoString(result)
}

// GetCwd returns the current working directory.
func GetCwd() string {
	result := C.get_cwd()
	return cStringToGoString(result)
}

// ChangeDir changes the current working directory.
func ChangeDir(path string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	
	result := C.change_dir(cPath)
	return cStringToGoString(result)
}

// FileExists checks if a file or dir exists.
func FileExists(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	
	result := C.file_exists(cPath)
	return result == 1
}

// FileMetadata returns JSON with file size, readonly, modified time.
func FileMetadata(path string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	
	result := C.file_metadata(cPath)
	return cStringToGoString(result)
}
