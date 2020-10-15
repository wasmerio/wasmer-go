package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"fmt"
	"runtime"
)

type Instance struct {
	_inner  *C.wasm_instance_t
	exports map[string]*Extern
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

	exports := make(map[string]*Extern)

	return &Instance{
		_inner:  instance,
		exports: exports,
	}, nil
}

func (self *Instance) inner() *C.wasm_instance_t {
	return self._inner
}

func (self *Instance) Exports(name string) (*Extern, error) {
	export, exists := self.exports[name]

	if exists == false {
		return nil, newErrorWith(fmt.Sprintf("Export `%s` does not exist", name))
	}

	return export, nil
}
