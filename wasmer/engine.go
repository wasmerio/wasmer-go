package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Engine struct {
	_inner *C.wasm_engine_t
}

func NewEngine() *Engine {
	engine := &Engine{
		_inner: C.wasm_engine_new(),
	}

	runtime.SetFinalizer(engine, func(engine *Engine) {
		C.wasm_engine_delete(engine.inner())
	})

	return engine
}

func (engine *Engine) inner() *C.wasm_engine_t {
	return engine._inner
}
