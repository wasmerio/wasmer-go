package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Engine struct {
	_inner *C.wasm_engine_t
}

func newEngine(engine *C.wasm_engine_t) *Engine {
	self := &Engine{
		_inner: engine,
	}

	runtime.SetFinalizer(self, func(self *Engine) {
		C.wasm_engine_delete(self.inner())
	})

	return self
}

func NewEngine() *Engine {
	return newEngine(C.wasm_engine_new())
}

func newConfig(engine C.wasmer_engine_t) *C.wasm_config_t {
	config := C.wasm_config_new()
	C.wasm_config_set_engine(config, engine)

	return config
}

func NewJITEngine() *Engine {
	return newEngine(
		C.wasm_engine_new_with_config(
			newConfig(C.JIT),
		),
	)
}

func NewNativeEngine() *Engine {
	return newEngine(
		C.wasm_engine_new_with_config(
			newConfig(C.NATIVE),
		),
	)
}

func (self *Engine) inner() *C.wasm_engine_t {
	return self._inner
}
