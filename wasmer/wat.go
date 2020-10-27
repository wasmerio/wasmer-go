package wasmer

// #include <wasmer_wasm.h>
//
// wasm_byte_vec_t *to_wat2wasm(uint8_t *wat_data, size_t wat_length) {
//     wasm_byte_vec_t wat;
//     wat.size = wat_length;
//     wat.data = (wasm_byte_t*) wat_data;
//
//     return wat2wasm(&wat);
// }
import "C"
import (
	"runtime"
	"unsafe"
)

func Wat2Wasm(wat string) ([]byte, error) {
	watAsBytes := []byte(wat)
	var watPtr *C.uint8_t
	var watLength = len(wat)

	if watLength > 0 {
		watPtr = (*C.uint8_t)(unsafe.Pointer(&watAsBytes[0]))
	}

	var wasm *C.wasm_byte_vec_t = C.to_wat2wasm(watPtr, C.size_t(watLength))
	runtime.KeepAlive(wat)

	if wasm == nil {
		return nil, newErrorFromWasmer()
	}

	wasmBytes := C.GoBytes(unsafe.Pointer(wasm.data), C.int(wasm.size))
	C.wasm_byte_vec_delete(wasm)

	return wasmBytes, nil
}
