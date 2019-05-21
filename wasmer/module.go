package wasmer

/*
#cgo LDFLAGS: -L./ -lwasmer_runtime_c_api
#include "./wasmer.h"
*/
import "C"
import (
	"io/ioutil"
	"unsafe"
)

// ReadBytes reads a `.wasm` file and returns its content as an array of bytes.
func ReadBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Validate validates a sequence of bytes that is supposed to represent a valid
// WebAssembly module.
func Validate(bytes []byte) bool {
	return true == C.wasmer_validate((*C.uchar)(unsafe.Pointer(&bytes[0])), C.uint(len(bytes)))
}
