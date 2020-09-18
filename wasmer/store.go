package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

// Store represents all global state that can be manipulated by
// WebAssembly programs. It consists of the runtime representation
// of all instances of functions, tables, memories, and globals that
// have been allocated during the lifetime of the abstract machine.
//
// The `Store` holds the engine (that is —amongst many things— used to compile
// the Wasm bytes into a valid module artifact), in addition to the
// [`Tunables`] (that are used to create the memories, tables and globals).
//
// Spec: https://webassembly.github.io/spec/core/exec/runtime.html#store
type Store struct {
	_inner *C.wasm_store_t
	Engine *Engine
}

func NewStore(engine *Engine) *Store {
	store := &Store{
		_inner: C.wasm_store_new(engine.inner()),
		Engine: engine,
	}

	runtime.SetFinalizer(store, func(store *Store) {
		C.wasm_store_delete(store.inner())
	})

	return store
}

func (store *Store) inner() *C.wasm_store_t {
	return store._inner
}
