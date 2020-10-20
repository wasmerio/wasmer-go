package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Memory struct {
	_inner   *C.wasm_memory_t
	_ownedBy interface{}
}

func newMemory(pointer *C.wasm_memory_t, ownedBy interface{}) *Memory {
	memory := &Memory{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(memory, func(memory *Memory) {
			C.wasm_memory_delete(memory.inner())
		})
	}

	return memory
}

func NewMemory(store *Store, ty *MemoryType) *Memory {
	pointer := C.wasm_memory_new(store.inner(), ty.inner())

	runtime.KeepAlive(store)

	return newMemory(pointer, nil)
}

func (self *Memory) inner() *C.wasm_memory_t {
	return self._inner
}

func (self *Memory) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *Memory) Type() *MemoryType {
	ty := C.wasm_memory_type(self.inner())

	runtime.KeepAlive(self)

	return newMemoryType(ty, self.ownedBy())
}

func (self *Memory) Size() Pages {
	return Pages(C.wasm_memory_size(self.inner()))
}

func (self *Memory) IntoExtern() *Extern {
	pointer := C.wasm_memory_as_extern(self.inner())

	return newExtern(pointer, self.ownedBy())
}
