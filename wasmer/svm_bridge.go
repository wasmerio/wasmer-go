package wasmer

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmer_runtime_c_api -lsvm_wasmer_c_api
// #include "./wasmer.h"
// #include "./svm_wasmer.h"
//
import "C"
import "unsafe"

func cWasmerSvmImportObject(
	importObject **cWasmerImportObjectT,
	imports *cWasmerImportT,
	importsLength uint,
	config *SvmInstanceConfig,
) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_svm_import_object(
		(**C.wasmer_import_object_t)(unsafe.Pointer(importObject)),
		nil,
		// config.addr,
		(C.uint)(config.maxPages),
		(C.uint)(config.maxPagesSlices),
		config.nodeDataPtr,
		(*C.wasmer_import_t)(imports),
		(C.uint)(importsLength),
	))
}
