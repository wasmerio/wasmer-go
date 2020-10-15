package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Instance struct {
	_inner  *C.wasm_instance_t
	Exports *Exports
}

func NewInstance(module *Module) (*Instance, error) {
	var imports C.wasm_extern_vec_t
	C.wasm_extern_vec_new_empty(&imports)

	var traps *C.wasm_trap_t

	instance := C.wasm_instance_new(module.store.inner(), module.inner(), &imports, &traps)

	runtime.KeepAlive(module)
	runtime.KeepAlive(module.store)
	runtime.KeepAlive(imports)

	if instance == nil {
		return nil, newErrorFromWasmer()
	}

	if traps != nil {
		return nil, newErrorWith("trapped! to do")
	}

	return &Instance{
		_inner:  instance,
		Exports: newExports(instance, module),
	}, nil
}

func (self *Instance) inner() *C.wasm_instance_t {
	return self._inner
}
