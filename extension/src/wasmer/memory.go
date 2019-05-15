package wasmer

/*
#cgo LDFLAGS: -L../../ -lwasmer_runtime_c_api
#include "../../wasmer.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type Memory struct {
	memory *C.wasmer_memory_t
}

func NewMemory(memory *C.wasmer_memory_t) Memory {
	return Memory{memory}
}

func (self *Memory) Length() uint32 {
	return uint32(C.wasmer_memory_data_length(self.memory))
}

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
