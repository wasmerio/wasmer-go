package wasmer

/*
#cgo LDFLAGS: -L../../ -lwasmer_runtime_c_api
#include "../../wasmer.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func GetLastError() (string, error) {
	var error_length C.int = C.wasmer_last_error_length()

	if error_length == 0 {
		return "", nil
	}

	var error_message []C.char = make([]C.char, error_length)
	var error_message_pointer *C.char = (*C.char)(unsafe.Pointer(&error_message[0]))

	var error_result = C.wasmer_last_error_message(error_message_pointer, error_length)

	if -1 == error_result {
		return "", errors.New("Cannot read last error.")
	}

	return C.GoString(error_message_pointer), nil
}
