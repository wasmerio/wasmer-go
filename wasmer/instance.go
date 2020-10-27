package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type Instance struct {
	_inner  *C.wasm_instance_t
	Exports *Exports
}

func NewInstance(module *Module, imports *ImportObject) (*Instance, error) {
	var traps *C.wasm_trap_t

	instance := C.wasm_instance_new(
		module.store.inner(),
		module.inner(),
		imports.intoInner(),
		&traps,
	)

	runtime.KeepAlive(module)
	runtime.KeepAlive(module.store)
	runtime.KeepAlive(imports)

	if instance == nil {
		return nil, newErrorFromWasmer()
	}

	if traps != nil {
		return nil, newErrorWith("trapped! to do")
	}

	output := &Instance{
		_inner:  instance,
		Exports: newExports(instance, module),
	}

	runtime.SetFinalizer(output, func(self *Instance) {
		C.wasm_instance_delete(self.inner())
	})

	return output, nil
}

func (self *Instance) inner() *C.wasm_instance_t {
	return self._inner
}
