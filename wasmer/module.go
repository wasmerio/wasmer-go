package wasmer

// #include <wasmer_wasm.h>
//
// // We can't create a `wasm_byte_vec_t` directly in Go otherwise cgo
// // complains with “Go pointer to Go pointer”. The hack consists at
// // creating the `wasm_byte_vec_t` directly in C.
// own wasm_module_t* go_wasm_module_new(wasm_store_t *store, uint8_t *bytes, size_t bytes_length) {
//     wasm_byte_vec_t wasm_bytes;
//     wasm_bytes.size = bytes_length;
//     wasm_bytes.data = (wasm_byte_t*) bytes;
//
//    return wasm_module_new(store, &wasm_bytes);
// }
import "C"
import (
	"runtime"
	"unsafe"
)

// A WebAssembly module contains stateless WebAssembly code that has
// already been compiled and can be instantiated multiple times.
//
// Creates a new WebAssembly Module given the configuration
// in the store.
//
// If the provided bytes are not WebAssembly-like (start with
// `b"\0asm"`), this function will try to to convert the bytes
// assuming they correspond to the WebAssembly text format.
//
// ## Security
//
// Before the code is compiled, it will be validated using the store
// features.
type Module struct {
	_inner *C.wasm_module_t
}

func NewModule(store *Store, bytes []byte) (*Module, error) {
	var bytesPtr *C.uint8_t

	if len(bytes) > 0 {
		bytesPtr = (*C.uint8_t)(unsafe.Pointer(&bytes[0]))
	}

	module := &Module{
		_inner: C.go_wasm_module_new(store.inner(), bytesPtr, C.size_t(len(bytes))),
	}

	runtime.KeepAlive(bytes)
	runtime.SetFinalizer(module, func(module *Module) {
		C.wasm_module_delete(module.inner())
	})

	if module == nil {
		return nil, newError()
	}

	return module, nil
}

func (module *Module) inner() *C.wasm_module_t {
	return module._inner
}
