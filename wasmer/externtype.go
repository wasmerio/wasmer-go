package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type ExternKind C.wasm_externkind_t

const (
	FUNCTION = ExternKind(C.WASM_EXTERN_FUNC)
	GLOBAL   = ExternKind(C.WASM_EXTERN_GLOBAL)
	TABLE    = ExternKind(C.WASM_EXTERN_TABLE)
	MEMORY   = ExternKind(C.WASM_EXTERN_MEMORY)
)

func (self ExternKind) String() string {
	switch self {
	case FUNCTION:
		return "func"
	case GLOBAL:
		return "global"
	case TABLE:
		return "table"
	case MEMORY:
		return "memory"
	}
	panic("Unknown extern kind") // unreachable
}

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

func (self *ExternType) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *ExternType) Kind() ExternKind {
	kind := ExternKind(C.wasm_externtype_kind(self.inner()))

	runtime.KeepAlive(self)

	return kind
}

func (self *ExternType) IntoFunctionType() *FunctionType {
	pointer := C.wasm_externtype_as_functype_const(self.inner())

	if pointer == nil {
		return nil
	}

	return newFunctionType(pointer, self.ownedBy())
}

func (self *ExternType) IntoGlobalType() *GlobalType {
	pointer := C.wasm_externtype_as_globaltype_const(self.inner())

	if pointer == nil {
		return nil
	}

	return newGlobalType(pointer, self.ownedBy())
}

func (self *ExternType) IntoMemoryType() *MemoryType {
	pointer := C.wasm_externtype_as_memorytype_const(self.inner())

	if pointer == nil {
		return nil
	}

	return newMemoryType(pointer, self.ownedBy())
}
