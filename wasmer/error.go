package wasmer

/*
#cgo LDFLAGS: -L./ -lwasmer_runtime_c_api
#include "./wasmer.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// GetLastError returns the last error message if any, otherwise returns an error.
func GetLastError() (string, error) {
	var errorLength = C.wasmer_last_error_length()

	if errorLength == 0 {
		return "", nil
	}

	var errorMessage = make([]C.char, errorLength)
	var errorMessagePointer = (*C.char)(unsafe.Pointer(&errorMessage[0]))

	var errorResult = C.wasmer_last_error_message(errorMessagePointer, errorLength)

	if -1 == errorResult {
		return "", errors.New("Cannot read last error")
	}

	return C.GoString(errorMessagePointer), nil
}
