package wasmer

// #include <wasmer.h>
import "C"
import "runtime"

// Engine is used by the Store to drive the compilation and the
// execution of a WebAssembly module.
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

// NewEngine instantiates and returns a new Engine with the default configuration.
//
//   engine := NewEngine()
//
func NewEngine() *Engine {
	return newEngine(C.wasm_engine_new())
}

// NewEngineWithConfig instantiates and returns a new Engine with the given configuration.
//
//   config := NewConfig()
//   engine := NewEngineWithConfig(config)
//
func NewEngineWithConfig(config *Config) *Engine {
	return newEngine(C.wasm_engine_new_with_config(config.inner()))
}

// NewUniversalEngine instantiates and returns a new Universal engine.
//
//   engine := NewUniversalEngine()
//
func NewUniversalEngine() *Engine {
	config := NewConfig()
	config.UseUniversalEngine()

	return NewEngineWithConfig(config)
}

// NewDylibEngine instantiates and returns a new Dylib engine.
//
//   engine := NewDylibEngine()
//
func NewDylibEngine() *Engine {
	config := NewConfig()
	config.UseDylibEngine()

	return NewEngineWithConfig(config)
}

func (self *Engine) inner() *C.wasm_engine_t {
	return self._inner
}

// NewJITEngine is a deprecated function. Please use NewUniversalEngine instead.
func NewJITEngine() *Engine {
	return NewUniversalEngine()
}

// NewNativeEngine is a deprecated function. Please use NewDylibEngine instead.
func NewNativeEngine() *Engine {
	return NewDylibEngine()
}
