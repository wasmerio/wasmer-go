package wasmer

import (
	"errors"
	"unsafe"
)

// GetLastError returns the last error message if any, otherwise returns an error.
func GetLastError() (string, error) {
	var errorLength = C_wasmer_last_error_length()

	if errorLength == 0 {
		return "", nil
	}

	var errorMessage = make([]C_char, errorLength)
	var errorMessagePointer = (*C_char)(unsafe.Pointer(&errorMessage[0]))

	var errorResult = C_wasmer_last_error_message(errorMessagePointer, errorLength)

	if -1 == errorResult {
		return "", errors.New("Cannot read last error")
	}

	return C_GoString(errorMessagePointer), nil
}
