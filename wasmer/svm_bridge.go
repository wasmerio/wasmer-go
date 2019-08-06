package wasmer

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmer_runtime_c_api -lsvm_wasmer_c_api
// #include "./wasmer.h"
// #include "./svm_wasmer.h"
//
import "C"

import (
	"unsafe"
	"reflect"
)

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

/// Useful for debugging the `config.addr`
func castAddrPtrToBytes(u uintptr) []byte {
    return (*[32]byte)(unsafe.Pointer(u))[0:32]
}

func cWasmerSvmInstanceContextNodeDataGet(instanceContext *cWasmerInstanceContextT) unsafe.Pointer {
	return (unsafe.Pointer)(C.wasmer_svm_instance_context_node_data_get((*C.wasmer_instance_context_t)(instanceContext)))
}

func cWasmerSvmImportObject(
	importObject **cWasmerImportObjectT,
	imports *cWasmerImportT,
	importsLength uint,
	config *SvmInstanceConfig,
) cWasmerResultT {
	addrHeader := (*reflect.SliceHeader)(unsafe.Pointer(&config.Addr))
	addrPtr := (unsafe.Pointer)(addrHeader.Data)

	return (cWasmerResultT)(C.wasmer_svm_import_object(
		(**C.wasmer_import_object_t)(unsafe.Pointer(importObject)),
		addrPtr,
		(C.uint)(config.MaxPages),
		(C.uint)(config.MaxPagesSlices),
		config.NodeDataPtr,
		(*C.wasmer_import_t)(imports),
		(C.uint)(importsLength),
	))
}
