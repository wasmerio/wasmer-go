package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"runtime"
)

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

func NewGlobal(store *Store, ty *GlobalType, value Value) *Global {
	pointer := C.wasm_global_new(
		store.inner(),
		ty.inner(),
		value.inner(),
	)

	return newGlobal(pointer, nil)
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

func (self *Global) Type() *GlobalType {
	ty := C.wasm_global_type(self.inner())

	runtime.KeepAlive(self)

	return newGlobalType(ty, self.ownedBy())
}

func (self *Global) Set(value interface{}, kind ValueKind) error {
	if self.Type().Mutability() == IMMUTABLE {
		return newErrorWith("The global variable is not mutable, cannot set a new value")
	}

	result, err := fromGoValue(value, kind)

	if err != nil {
		//TODO: Make this error explicit
		panic(err.Error())
	}

	C.wasm_global_set(self.inner(), &result)

	return nil
}

func (self *Global) Get() (interface{}, error) {
	var value C.wasm_val_t

	C.wasm_global_get(self.inner(), &value)

	return toGoValue(&value), nil
}
