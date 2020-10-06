package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type ExternType struct {
	_inner   *C.wasm_externtype_t
	_ownedBy interface{}
}

type IntoExternType interface {
	IntoExternType() *ExternType
}

func newExternType(pointer *C.wasm_externtype_t, ownedBy interface{}) *ExternType {
	externType := &ExternType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(externType, func(externType *ExternType) {
			C.wasm_externtype_delete(externType.inner())
		})
	}

	return externType
}

func (self *ExternType) inner() *C.wasm_externtype_t {
	return self._inner
}
