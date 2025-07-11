package elliecore

/*
#cgo LDFLAGS: -L${SRCDIR}/../rustmods/elliecore/target/debug -lelliecore
#include <stdlib.h>
char* run_cmd(const char* cmd);
char* run_cmd_with_env(const char* cmd, const char* envs);
char* read_file(const char* path);
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
