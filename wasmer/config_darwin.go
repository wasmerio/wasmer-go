//go:build darwin
// +build darwin

package wasmer

// #include <wasmer.h>
// extern uint64_t metering_delegate(enum wasmer_parser_operator_t op);
import "C"

func getPlatformLong(v uint64) C.ulonglong {
	return C.ulonglong(v)
}
