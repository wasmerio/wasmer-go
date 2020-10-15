package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Global struct {
	_inner   *C.wasm_global_t
	_ownedBy interface{}
}

func newGlobal(pointer *C.wasm_global_t, ownedBy interface{}) *Global {
	global := &Global{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(global, func(global *Global) {
			C.wasm_global_delete(global.inner())
		})
	}

	return global
}

func (self *Global) inner() *C.wasm_global_t {
	return self._inner
}

func (self *Global) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *Global) IntoExtern() *Extern {
	pointer := C.wasm_global_as_extern(self.inner())

	return newExtern(pointer, self.ownedBy())
}
