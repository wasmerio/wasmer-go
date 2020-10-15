package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Function struct {
	_inner   *C.wasm_func_t
	_ownedBy interface{}
}

func newFunction(pointer *C.wasm_func_t, ownedBy interface{}) *Function {
	function := &Function{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(function, func(function *Function) {
			C.wasm_func_delete(function.inner())
		})
	}

	return function
}

func (self *Function) inner() *C.wasm_func_t {
	return self._inner
}

func (self *Function) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *Function) IntoExtern() *Extern {
	pointer := C.wasm_func_as_extern(self.inner())

	return newExtern(pointer, self.ownedBy())
}
