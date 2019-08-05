package wasmer

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmer_runtime_c_api -lsvm_wasmer_c_api
// #include "./wasmer.h"
// #include "./svm_wasmer.h"
//
import "C"
import "unsafe"

func cWasmerSvmImportObject(
	importObject **cWasmerImportObjectT,
	addr unsafe.Pointer,
	maxPages uint,
	maxPagesSlices uint,
	nodeDataPtr unsafe.Pointer,
	imports *cWasmerImportT,
	importsLength uint,
) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_svm_import_object(
		(**C.wasmer_import_object_t)(unsafe.Pointer(importObject)),
		addr,
		(C.uint)(maxPages),
		(C.uint)(maxPagesSlices),
		nodeDataPtr,
		(*C.wasmer_import_t)(imports),
		(C.uint)(importsLength),
	))
}
