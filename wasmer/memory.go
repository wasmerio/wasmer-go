package wasmer

/*
#cgo LDFLAGS: -L./ -lwasmer_runtime_c_api
#include "./wasmer.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// Represents an exported memory of a WebAssembly instance. To read
// and write data, please see the `Data` function.
type Memory struct {
	memory *C.wasmer_memory_t
}

// Instantiates a new WebAssembly exported memory.
func newMemory(memory *C.wasmer_memory_t) Memory {
	return Memory{memory}
}

// Calculates the memory length (in bytes).
func (self *Memory) Length() uint32 {
	return uint32(C.wasmer_memory_data_length(self.memory))
}

// Returns a slice of bytes over the WebAssembly memory.
func (self *Memory) Data() []byte {
	var length uint32 = self.Length()
	var data *uint8 = (*uint8)(C.wasmer_memory_data(self.memory))

	var header reflect.SliceHeader
	header = *(*reflect.SliceHeader)(unsafe.Pointer(&header))

	header.Data = uintptr(unsafe.Pointer(data))
	header.Len = int(length)
	header.Cap = int(length)

	return *(*[]byte)(unsafe.Pointer(&header))
}
