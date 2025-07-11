package elliecore

/*
#cgo LDFLAGS: -L${SRCDIR}/../rustmods/elliecore/target/debug -lelliecore
#include <stdlib.h>
char* run_cmd(const char* cmd);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// RunCmd executes a shell command using the Rust FFI and prints the result.
func RunCmd(cmd string) {
	input := C.CString(cmd)
	defer C.free(unsafe.Pointer(input))

	result := C.run_cmd(input)
	defer C.free(unsafe.Pointer(result))

	fmt.Println("Rust says:\n", C.GoString(result))
}

func main() {
	fmt.Println("Testing Rust FFI call with 'echo Hello from Rust!'")
	RunCmd("echo Hello from Rust!")
}
