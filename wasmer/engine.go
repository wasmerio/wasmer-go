package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Engine struct {
	_inner *C.wasm_engine_t
}

func NewEngine() *Engine {
	self := &Engine{
		_inner: C.wasm_engine_new(),
	}

	runtime.SetFinalizer(self, func(self *Engine) {
		C.wasm_engine_delete(self.inner())
	})

	return self
}

func (self *Engine) inner() *C.wasm_engine_t {
	return self._inner
}
