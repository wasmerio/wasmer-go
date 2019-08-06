package wasmer

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmer_runtime_c_api -lsvm_wasmer_c_api
// #include "./wasmer.h"
// #include "./svm_wasmer.h"
//
import "C"

import (
	"unsafe"
	"reflect"
	"fmt"
)

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

/// Useful for debugging the `config.addr`
func castAddrPtrToBytes(u uintptr) []byte {
    return (*[32]byte)(unsafe.Pointer(u))[0:32]
}

func cWasmerSvmImportObject(
	importObject **cWasmerImportObjectT,
	imports *cWasmerImportT,
	importsLength uint,
	config *SvmInstanceConfig,
) cWasmerResultT {
	addrHeader := (*reflect.SliceHeader)(unsafe.Pointer(&config.addr))
	addrPtr := (unsafe.Pointer)(addrHeader.Data)

	return (cWasmerResultT)(C.wasmer_svm_import_object(
		(**C.wasmer_import_object_t)(unsafe.Pointer(importObject)),
		addrPtr,
		(C.uint)(config.maxPages),
		(C.uint)(config.maxPagesSlices),
		config.nodeDataPtr,
		(*C.wasmer_import_t)(imports),
		(C.uint)(importsLength),
	))
}
