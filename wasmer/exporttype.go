package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"runtime"
	"unsafe"
)

type exportTypes struct {
	_inner      C.wasm_exporttype_vec_t
	exportTypes []*ExportType
}

func newExportTypes(module *Module) *exportTypes {
	self := &exportTypes{}
	C.wasm_module_exports(module.inner(), &self._inner)

	runtime.KeepAlive(module)
	runtime.SetFinalizer(self, func(self *exportTypes) {
		C.wasm_exporttype_vec_delete(self.inner())
	})

	numberOfExportTypes := int(self.inner().size)
	types := make([]*ExportType, numberOfExportTypes)
	firstExportType := unsafe.Pointer(self.inner().data)
	sizeOfExportTypePointer := unsafe.Sizeof(firstExportType)

	var currentTypePointer *C.wasm_exporttype_t

	for nth := 0; nth < numberOfExportTypes; nth++ {
		currentTypePointer = *(**C.wasm_exporttype_t)(unsafe.Pointer(uintptr(firstExportType) + uintptr(nth)*sizeOfExportTypePointer))
		exportType := newExportType(currentTypePointer, self)
		types[nth] = exportType
	}

	self.exportTypes = types

	return self
}

func (self *exportTypes) inner() *C.wasm_exporttype_vec_t {
	return &self._inner
}

type ExportType struct {
	_inner   *C.wasm_exporttype_t
	_ownedBy interface{}
}

func newExportType(pointer *C.wasm_exporttype_t, ownedBy interface{}) *ExportType {
	exportType := &ExportType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(exportType, func(exportType *ExportType) {
			C.wasm_exporttype_delete(exportType.inner())
		})
	}

	return exportType
}

func (self *ExportType) inner() *C.wasm_exporttype_t {
	return self._inner
}

func (self *ExportType) Name() string {
	byteVec := C.wasm_exporttype_name(self.inner())
	name := C.GoStringN(byteVec.data, C.int(byteVec.size))

	runtime.KeepAlive(self)

	return name
}
