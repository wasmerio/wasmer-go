package wasmer

// #include <wasmer.h>
// extern uint64_t metering_delegate_alt(enum wasmer_parser_operator_t op);
import "C"
import "unsafe"

func getInternalCPointer() unsafe.Pointer {
	return unsafe.Pointer(C.metering_delegate_alt)
}

//export metering_delegate_alt
func metering_delegate_alt(op C.wasmer_parser_operator_t) C.uint64_t {
	// a simple alogorithm for now just map from opcode to cost directly
	// all the responsibility is placed on the caller of PushMeteringMiddleware
	v, b := opCodeMap[Opcode(op)]
	if !b {
		return 0 // no value means no cost
	}
	return C.uint64_t(v)
}
