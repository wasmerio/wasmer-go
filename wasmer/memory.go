package wasmer

import (
	"fmt"
	"reflect"
	"unsafe"
)

// MemoryError represents any kind of errors related to a WebAssembly memory. It
// is returned by `Memory` functions only.
type MemoryError struct {
	// Error message.
	message string
}

// NewMemoryError constructs a new `MemoryError`.
func NewMemoryError(message string) *MemoryError {
	return &MemoryError{message}
}

// `MemoryError` is an actual error. The `Error` function returns
// the error message.
func (error *MemoryError) Error() string {
	return error.message
}

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
	if nil == memory.memory {
		return 0
	}

	return uint32(cWasmerMemoryDataLength(memory.memory))
}

// Data returns a slice of bytes over the WebAssembly memory.
func (memory *Memory) Data() []byte {
	if nil == memory.memory {
		return make([]byte, 0)
	}

	var length = memory.Length()
	var data = (*uint8)(cWasmerMemoryData(memory.memory))

	var header reflect.SliceHeader
	header = *(*reflect.SliceHeader)(unsafe.Pointer(&header))

	header.Data = uintptr(unsafe.Pointer(data))
	header.Len = int(length)
	header.Cap = int(length)

	return *(*[]byte)(unsafe.Pointer(&header))
}

// Grow the memory by a number of pages (65kb each).
func (memory *Memory) Grow(numberOfPages uint32) error {
	if nil == memory.memory {
		return nil
	}

	var growResult = cWasmerMemoryGrow(memory.memory, cUint32T(numberOfPages))

	if growResult != cWasmerOk {
		var lastError, err = GetLastError()
		var errorMessage = "Failed to grow the memory:\n    %s"

		if err != nil {
			errorMessage = fmt.Sprintf(errorMessage, "(unknown details)")
		} else {
			errorMessage = fmt.Sprintf(errorMessage, lastError)
		}

		return NewMemoryError(errorMessage)
	}

	return nil
}
