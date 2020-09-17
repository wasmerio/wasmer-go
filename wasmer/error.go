package wasmer

// #include <wasmer_wasm.h>
import "C"
import "unsafe"

type Error struct {
	message string
}

func newError() *Error {
	var errorLength = C.wasmer_last_error_length()

	if errorLength == 0 {
		return &Error{
			message: "(no error from Wasmer)",
		}
	}

	var errorMessage = make([]C.char, errorLength)
	var errorMessagePointer = (*C.char)(unsafe.Pointer(&errorMessage[0]))

	var errorResult = C.wasmer_last_error_message(errorMessagePointer, errorLength)

	if errorResult == -1 {
		return &Error{
			message: "(failed to read last error from Wasmer)",
		}
	}

	return &Error{
		message: C.GoString(errorMessagePointer),
	}
}

func (error *Error) Error() string {
	return error.message
}
