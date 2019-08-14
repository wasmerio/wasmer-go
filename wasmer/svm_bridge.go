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

func cSliceUnsafePointer(slice []byte) unsafe.Pointer {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	return (unsafe.Pointer)(header.Data)
}

func cWasmerSvmImportObject(
	importObject **cWasmerImportObjectT,
	imports *cWasmerImportT,
	importsLength uint,
	config *SvmInstanceConfig,
) cWasmerResultT {
	addr := cSliceUnsafePointer(config.Addr)
	state := cSliceUnsafePointer(config.State)

	return (cWasmerResultT)(C.wasmer_svm_import_object(
		(**C.wasmer_import_object_t)(unsafe.Pointer(importObject)),
		addr,
		state,
		(C.uint)(config.MaxPages),
		(C.uint)(config.MaxPagesSlices),
		config.NodeDataPtr,
		(*C.wasmer_import_t)(imports),
		(C.uint)(importsLength),
	))
}

func cWasmerSvmRegisterGet(regBuf *unsafe.Pointer, instanceContext *cWasmerInstanceContextT, regBits int, regIndex int) cWasmerResultT {
	*regBuf = (unsafe.Pointer)(C.wasmer_svm_register_get(
		(*C.wasmer_instance_context_t)(instanceContext),
		(C.int)(regBits),
		(C.int)(regIndex),
		))

	return cWasmerOk
}

func cWasmerSvmRegisterSet(instanceContext *cWasmerInstanceContextT, regBits int, regIndex int, buf []byte) cWasmerResultT {
	C.wasmer_svm_register_set(
		(*C.wasmer_instance_context_t)(instanceContext),
		(C.int)(regBits),
		(C.int)(regIndex),
		cSliceUnsafePointer(buf),
		(C.uint8_t)(len(buf)),
	)

	return cWasmerOk
}
