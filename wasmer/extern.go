package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Extern struct {
	_inner   *C.wasm_extern_t
	_ownedBy interface{}
}

type IntoExtern interface {
	IntoExtern() *Extern
}

func newExtern(pointer *C.wasm_extern_t, ownedBy interface{}) *Extern {
	extern := &Extern{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(extern, func(extern *Extern) {
			C.wasm_extern_delete(extern.inner())
		})
	}

	return extern
}

func (self *Extern) inner() *C.wasm_extern_t {
	return self._inner
}

func (self *Extern) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *Extern) Kind() ExternKind {
	kind := ExternKind(C.wasm_extern_kind(self.inner()))

	runtime.KeepAlive(self)

	return kind
}

func (self *Extern) Type() *ExternType {
	ty := C.wasm_extern_type(self.inner())

	runtime.KeepAlive(self)

	return newExternType(ty, self.ownedBy())
}

func (self *Extern) IntoFunction() *Function {
	pointer := C.wasm_extern_as_func(self.inner())

	if pointer == nil {
		return nil
	}

	return newFunction(pointer, self.ownedBy())
}
