package elliecore

/*
#cgo LDFLAGS: -L${SRCDIR}/../rustmods/elliecore/target/debug -lelliecore
#include <stdlib.h>
char* run_cmd(const char* cmd);
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
