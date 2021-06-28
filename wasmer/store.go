package wasmer

// #include <wasmer.h>
import "C"
import "runtime"

// Store represents all global state that can be manipulated by
// WebAssembly programs. It consists of the runtime representation of
// all instances of functions, tables, memories, and globals that have
// been allocated during the life time of the abstract machine.
//
// The Store holds the Engine (that is — amongst many things — used to
// compile the Wasm bytes into a valid module artifact).
//
// See also
//
// Specification: https://webassembly.github.io/spec/core/exec/runtime.html#store
type Store struct {
	_inner *C.wasm_store_t
	Engine *Engine
}

// NewStore instantiates a new Store with an Engine.
//
//   engine := NewEngine()
//   store := NewStore(engine)
func NewStore(engine *Engine) *Store {
	self := &Store{
		_inner: C.wasm_store_new(engine.inner()),
		Engine: engine,
	}

	runtime.SetFinalizer(self, func(self *Store) {
		self.Close()
	})

	return self
}

func (self *Store) inner() *C.wasm_store_t {
	return self._inner
}

// Force to close the Store.
//
// A runtime finalizer is registered on the Store, but it is possible
// to force the destruction of the Store by calling Close manually.
func (self *Store) Close() {
	runtime.SetFinalizer(self, nil)
	C.wasm_store_delete(self.inner())
}
