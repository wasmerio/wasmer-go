package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type Exports struct {
	_inner  C.wasm_extern_vec_t
	exports map[string]*Extern
}

func newExports(instance *C.wasm_instance_t, module *Module) *Exports {
	self := &Exports{}
	C.wasm_instance_exports(instance, &self._inner)

	runtime.SetFinalizer(self, func(exports *Exports) {
		C.wasm_extern_vec_delete(exports.inner())
	})

	numberOfExports := int(self.inner().size)
	exports := make(map[string]*Extern, numberOfExports)
	firstExport := unsafe.Pointer(self.inner().data)
	sizeOfExportPointer := unsafe.Sizeof(firstExport)

	var currentExportPointer *C.wasm_extern_t

	moduleExports := module.Exports()

	for nth := 0; nth < numberOfExports; nth++ {
		currentExportPointer = *(**C.wasm_extern_t)(unsafe.Pointer(uintptr(firstExport) + uintptr(nth)*sizeOfExportPointer))
		export := newExtern(currentExportPointer, self)
		exports[moduleExports[nth].Name()] = export
	}

	self.exports = exports

	return self
}

func (self *Exports) inner() *C.wasm_extern_vec_t {
	return &self._inner
}

func (self *Exports) Get(name string) (*Extern, error) {
	export, exists := self.exports[name]

	if exists == false {
		return nil, newErrorWith(fmt.Sprintf("Export `%s` does not exist", name))
	}

	return export, nil
}

func (self *Exports) GetFunction(name string) (*Function, error) {
	exports, err := self.Get(name)

	if err != nil {
		return nil, err
	}

	return exports.IntoFunction(), nil
}
