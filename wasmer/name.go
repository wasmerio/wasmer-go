package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"runtime"
)

func newName(str string) C.wasm_name_t {
	name := C.wasm_name_t{}
	C.wasm_name_new_from_string(&name, C.CString(str))

	runtime.KeepAlive(str)

	return name
}
