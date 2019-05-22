package wasmer

import (
	"reflect"
	"unsafe"
)

// Memory represents an exported memory of a WebAssembly instance. To read
// and write data, please see the `Data` function.
type Memory struct {
	memory *cWasmerMemoryT
}

// Instantiates a new WebAssembly exported memory.
func newMemory(memory *cWasmerMemoryT) Memory {
	return Memory{memory}
}

// Length calculates the memory length (in bytes).
func (memory *Memory) Length() uint32 {
	return uint32(cWasmerMemoryDataLength(memory.memory))
}

// Data returns a slice of bytes over the WebAssembly memory.
func (memory *Memory) Data() []byte {
	var length = memory.Length()
	var data = (*uint8)(cWasmerMemoryData(memory.memory))

	var header reflect.SliceHeader
	header = *(*reflect.SliceHeader)(unsafe.Pointer(&header))

	header.Data = uintptr(unsafe.Pointer(data))
	header.Len = int(length)
	header.Cap = int(length)

	return *(*[]byte)(unsafe.Pointer(&header))
}
