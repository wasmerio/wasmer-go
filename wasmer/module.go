package wasmer

// #include <wasmer_wasm.h>
//
// #define own
//
// // We can't create a `wasm_byte_vec_t` directly in Go otherwise cgo
// // complains with “Go pointer to Go pointer”. The hack consists at
// // creating the `wasm_byte_vec_t` directly in C.
//
// own wasm_module_t* to_wasm_module_new(wasm_store_t *store, uint8_t *bytes, size_t bytes_length) {
//     wasm_byte_vec_t wasm_bytes;
//     wasm_bytes.size = bytes_length;
//     wasm_bytes.data = (wasm_byte_t*) bytes;
//
//     return wasm_module_new(store, &wasm_bytes);
// }
//
// bool to_wasm_module_validate(wasm_store_t *store, uint8_t *bytes, size_t bytes_length) {
//     wasm_byte_vec_t wasm_bytes;
//     wasm_bytes.size = bytes_length;
//     wasm_bytes.data = (wasm_byte_t*) bytes;
//
//     return wasm_module_validate(store, &wasm_bytes);
// }
//
// wasm_module_t* to_wasm_module_deserialize(wasm_store_t *store, uint8_t *bytes, size_t bytes_length) {
//     wasm_byte_vec_t serialized_bytes;
//     serialized_bytes.size = bytes_length;
//     serialized_bytes.data = (wasm_byte_t*) bytes;
//
//     return wasm_module_deserialize(store, &serialized_bytes);
// }
import "C"
import (
	"runtime"
	"unsafe"
)

// Module contains stateless WebAssembly code that has already been
// compiled and can be instantiated multiple times.
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
	store  *Store
}

func NewModule(store *Store, bytes []byte) (*Module, error) {
	wasmBytes, err := Wat2Wasm(string(bytes))

	if err != nil {
		return nil, err
	}

	var wasmBytesPtr *C.uint8_t
	wasmBytesLength := len(wasmBytes)

	if wasmBytesLength > 0 {
		wasmBytesPtr = (*C.uint8_t)(unsafe.Pointer(&wasmBytes[0]))
	}

	self := &Module{
		_inner: C.to_wasm_module_new(store.inner(), wasmBytesPtr, C.size_t(wasmBytesLength)),
		store:  store,
	}

	if self._inner == nil {
		return nil, newErrorFromWasmer()
	}

	runtime.KeepAlive(bytes)
	runtime.KeepAlive(wasmBytes)
	runtime.SetFinalizer(self, func(self *Module) {
		C.wasm_module_delete(self.inner())
	})

	return self, nil
}

func ValidateModule(store *Store, bytes []byte) error {
	wasmBytes, err := Wat2Wasm(string(bytes))

	if err != nil {
		return err
	}

	var wasmBytesPtr *C.uint8_t
	wasmBytesLength := len(wasmBytes)

	if wasmBytesLength > 0 {
		wasmBytesPtr = (*C.uint8_t)(unsafe.Pointer(&wasmBytes[0]))
	}

	isValid := C.to_wasm_module_validate(store.inner(), wasmBytesPtr, C.size_t(wasmBytesLength))

	runtime.KeepAlive(bytes)
	runtime.KeepAlive(wasmBytes)

	if !isValid {
		return newErrorFromWasmer()
	}

	return nil
}

func (self *Module) inner() *C.wasm_module_t {
	return self._inner
}

func (self *Module) Name() string {
	var name C.wasm_name_t

	C.wasm_module_name(self.inner(), &name)

	goName := nameToString(&name)

	C.wasm_name_delete(&name)

	return goName
}

func (self *Module) Imports() []*ImportType {
	return newImportTypes(self).importTypes
}

func (self *Module) Exports() []*ExportType {
	return newExportTypes(self).exportTypes
}

func (self *Module) Serialize() ([]byte, error) {
	var bytes C.wasm_byte_vec_t

	C.wasm_module_serialize(self.inner(), &bytes)

	if bytes.data == nil {
		return nil, newErrorFromWasmer()
	}

	goBytes := C.GoBytes(unsafe.Pointer(bytes.data), C.int(bytes.size))
	C.wasm_byte_vec_delete(&bytes)

	return goBytes, nil
}

func DeserializeModule(store *Store, bytes []byte) (*Module, error) {
	var bytesPtr *C.uint8_t
	bytesLength := len(bytes)

	if bytesLength > 0 {
		bytesPtr = (*C.uint8_t)(unsafe.Pointer(&bytes[0]))
	}

	self := &Module{
		_inner: C.to_wasm_module_deserialize(store.inner(), bytesPtr, C.size_t(bytesLength)),
	}

	if self._inner == nil {
		return nil, newErrorFromWasmer()
	}

	return self, nil
}

func (self *Module) Close() {
	if self._inner != nil {
		C.wasm_module_delete(self._inner)
	}
}
