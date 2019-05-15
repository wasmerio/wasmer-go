package wasmer

import (
	"reflect"
	"unsafe"
)

type Memory struct {
	data   *uint8
	length uint32
}

func NewMemory(data *uint8, length uint32) Memory {
	return Memory{data, length}
}

func (self *Memory) Length() uint32 {
	return self.length
}

func (self *Memory) Data() []byte {
	var header reflect.SliceHeader
	header = *(*reflect.SliceHeader)(unsafe.Pointer(&header))

	header.Data = uintptr(unsafe.Pointer(self.data))
	header.Len = int(self.length)
	header.Cap = int(self.length)

	return *(*[]byte)(unsafe.Pointer(&header))
}
