package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"runtime"
	"unsafe"
)

type importTypes struct {
	_inner      C.wasm_importtype_vec_t
	importTypes []*ImportType
}

func newImportTypes(module *Module) *importTypes {
	self := &importTypes{}
	C.wasm_module_imports(module.inner(), &self._inner)

	runtime.KeepAlive(module)
	runtime.SetFinalizer(self, func(self *importTypes) {
		C.wasm_importtype_vec_delete(self.inner())
	})

	numberOfImportTypes := int(self.inner().size)
	types := make([]*ImportType, numberOfImportTypes)
	firstImportType := unsafe.Pointer(self.inner().data)
	sizeOfImportTypePointer := unsafe.Sizeof(firstImportType)

	var currentTypePointer *C.wasm_importtype_t

	for nth := 0; nth < numberOfImportTypes; nth++ {
		currentTypePointer = *(**C.wasm_importtype_t)(unsafe.Pointer(uintptr(firstImportType) + uintptr(nth)*sizeOfImportTypePointer))
		importType := newImportType(currentTypePointer, self)
		types[nth] = importType
	}

	self.importTypes = types

	return self
}

func (self *importTypes) inner() *C.wasm_importtype_vec_t {
	return &self._inner
}

type ImportType struct {
	_inner   *C.wasm_importtype_t
	_ownedBy interface{}
}

func newImportType(pointer *C.wasm_importtype_t, ownedBy interface{}) *ImportType {
	importType := &ImportType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(importType, func(importType *ImportType) {
			C.wasm_importtype_delete(importType.inner())
		})
	}

	return importType
}

func (self *ImportType) inner() *C.wasm_importtype_t {
	return self._inner
}

func (self *ImportType) Module() string {
	byteVec := C.wasm_importtype_module(self.inner())
	module := C.GoStringN(byteVec.data, C.int(byteVec.size))

	runtime.KeepAlive(self)

	return module
}

func (self *ImportType) Name() string {
	byteVec := C.wasm_importtype_name(self.inner())
	name := C.GoStringN(byteVec.data, C.int(byteVec.size))

	runtime.KeepAlive(self)

	return name
}
