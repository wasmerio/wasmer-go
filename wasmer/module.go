package wasmer

// #include <wasmer_wasm.h>
import "C"
import "unsafe"

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
	var bytes_ptr *C.char

	if len(bytes) > 0 {
		bytes_ptr = (*C.char)(unsafe.Pointer(&bytes[0]))
	}

	var wasm_bytes C.wasm_byte_vec_t
	wasm_bytes.size = C.size_t(len(bytes))
	wasm_bytes.data = (*C.wasm_byte_t)(bytes_ptr)

	module := &Module{
		_inner: C.wasm_module_new(store.inner(), &wasm_bytes),
	}

	if module == nil {
		return nil, newError()
	}

	return module, nil
}
