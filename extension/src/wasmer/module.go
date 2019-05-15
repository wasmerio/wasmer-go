package wasmer

/*
#cgo LDFLAGS: -L../../ -lwasmer_runtime_c_api
#include "../../wasmer.h"
*/
import "C"
import (
	"io/ioutil"
	"unsafe"
)

func ReadBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func Validate(bytes []byte) bool {
	return true == C.wasmer_validate((*C.uchar)(unsafe.Pointer(&bytes[0])), C.uint(len(bytes)))
}
