package wasmer

// #include <wasmer_wasm.h>
import "C"
import "unsafe"

type Error struct {
	message string
}

func newErrorWith(message string) *Error {
	return &Error{
		message: message,
	}
}

func newErrorFromWasmer() *Error {
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
		message: C.GoStringN(errorMessagePointer, errorLength-1),
	}
}

func (error *Error) Error() string {
	return error.message
}

type TrapError struct {
	message string
	origin  *Frame
	trace   []*Frame
}

func newErrorFromTrap(pointer *C.wasm_trap_t) *TrapError {
	trap := newTrap(pointer, nil)

	return &TrapError{
		message: trap.Message(),
		origin: trap.Origin(),
		trace: trap.Trace().frames,
	}
}

func (self *TrapError) Error() string {
	return self.message
}

func (self *TrapError) Origin() *Frame {
	return self.origin
}

func (self *TrapError) Trace() []*Frame {
	return self.trace
}
