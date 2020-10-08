package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type MemoryType struct {
	_inner   *C.wasm_memorytype_t
	_ownedBy interface{}
}

func newMemoryType(pointer *C.wasm_memorytype_t, ownedBy interface{}) *MemoryType {
	memoryType := &MemoryType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(memoryType, func(memoryType *MemoryType) {
			C.wasm_memorytype_delete(memoryType.inner())
		})
	}

	return memoryType
}

func NewMemoryType(limits *Limits) *MemoryType {
	pointer := C.wasm_memorytype_new(limits.inner())

	return newMemoryType(pointer, nil)
}

func (self *MemoryType) inner() *C.wasm_memorytype_t {
	return self._inner
}

func (self *MemoryType) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *MemoryType) Limits() *Limits {
	cLimits := C.wasm_memorytype_limits(self.inner())
	limits := newLimits(cLimits, self.ownedBy())

	runtime.KeepAlive(self)

	return limits
}

func (self *MemoryType) IntoExternType() *ExternType {
	pointer := C.wasm_memorytype_as_externtype_const(self.inner())

	return newExternType(pointer, self.ownedBy())
}
