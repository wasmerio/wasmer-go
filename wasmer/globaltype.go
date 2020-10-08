package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type GlobalMutability C.wasm_mutability_t

const (
	IMMUTABLE = GlobalMutability(C.WASM_CONST)
	MUTABLE   = GlobalMutability(C.WASM_VAR)
)

func (self GlobalMutability) String() string {
	switch self {
	case IMMUTABLE:
		return "const"
	case MUTABLE:
		return "var"
	}
	panic("Unknown mutability") // unreachable
}

type GlobalType struct {
	_inner   *C.wasm_globaltype_t
	_ownedBy interface{}
}

func newGlobalType(pointer *C.wasm_globaltype_t, ownedBy interface{}) *GlobalType {
	globalType := &GlobalType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(globalType, func(globalType *GlobalType) {
			C.wasm_globaltype_delete(globalType.inner())
		})
	}

	return globalType
}

func NewGlobalType(valueType *ValueType, mutability GlobalMutability) *GlobalType {
	pointer := C.wasm_globaltype_new(valueType.inner(), C.wasm_mutability_t(mutability))

	return newGlobalType(pointer, nil)
}

func (self *GlobalType) inner() *C.wasm_globaltype_t {
	return self._inner
}

func (self *GlobalType) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *GlobalType) ValueType() *ValueType {
	pointer := C.wasm_globaltype_content(self.inner())

	return newValueType(pointer, self.ownedBy())
}

func (self *GlobalType) Mutability() GlobalMutability {
	mutability := GlobalMutability(C.wasm_globaltype_mutability(self.inner()))

	runtime.KeepAlive(self)

	return mutability
}

func (self *GlobalType) IntoExternType() *ExternType {
	pointer := C.wasm_globaltype_as_externtype_const(self.inner())

	return newExternType(pointer, self.ownedBy())
}
