package wasmer

// #include <wasmer.h>
import "C"

import (
	"runtime"
	"unsafe"
)

func newName(str string) C.wasm_name_t {
	var name C.wasm_name_t

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	C.wasm_name_new_from_string(&name, cStr)

	runtime.KeepAlive(str)

	return name
}

func nameToString(name *C.wasm_name_t) string {
	if name.data == nil {
		return ""
	}

	return C.GoStringN(name.data, C.int(name.size))
}
