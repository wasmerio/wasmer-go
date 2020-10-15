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

func (self *Memory) inner() *C.wasm_memory_t {
	return self._inner
}

func (self *Memory) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *Memory) IntoExtern() *Extern {
	pointer := C.wasm_memory_as_extern(self.inner())

	return newExtern(pointer, self.ownedBy())
}
