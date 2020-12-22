package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"unsafe"
)

// Wat2Wasm parsers a string as either WAT code or a binary Wasm module.
//
// ⚠️ This is not part of the standard Wasm C API. It is Wasmer specific.
//
//   wat := "(module)"
//   wasm, _ := Wat2Wasm(wat)
//   engine := wasmer.NewEngine()
//	 store := wasmer.NewStore(engine)
//	 module, _ := wasmer.NewModule(store, wasmBytes)
//
func Wat2Wasm(wat string) ([]byte, error) {
	var watBytes C.wasm_byte_vec_t
	var watLength = len(wat)
	C.wasm_byte_vec_new(&watBytes, C.size_t(watLength), C.CString(wat))

	var wasm C.wasm_byte_vec_t
	C.wat2wasm(
		&watBytes,
		&wasm,
	)
	C.wasm_byte_vec_delete(&watBytes)

	if wasm.data == nil {
		return nil, newErrorFromWasmer()
	}

	wasmBytes := C.GoBytes(unsafe.Pointer(wasm.data), C.int(wasm.size))
	C.wasm_byte_vec_delete(&wasm)

	return wasmBytes, nil
}
