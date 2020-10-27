package wasmer

// #include <wasmer_wasm.h>
import "C"
import "unsafe"

type WasiStateBuilder struct {
	_inner *C.wasi_config_t
}

func NewWasiStateBuilder(programName string) *WasiStateBuilder {
	cProgramName := C.CString(programName)
	defer C.free(unsafe.Pointer(cProgramName))
	wasiConfig := C.wasi_config_new(cProgramName)

	stateBuilder := &WasiStateBuilder{
		_inner: wasiConfig,
	}

	return stateBuilder
}

func (self *WasiStateBuilder) argument(argument string) *WasiStateBuilder {
	cArgument := C.CString(argument)
	defer C.free(unsafe.Pointer(cArgument))
	C.wasi_config_arg(self.inner(), cArgument)

	return self
}

func (self *WasiStateBuilder) finalize() (*WasiEnvironment, error) {
	return newWasiEnvironment(self)
}

func (self *WasiStateBuilder) inner() *C.wasi_config_t {
	return self._inner
}

type WasiEnvironment struct {
	_inner *C.wasi_env_t
}

func newWasiEnvironment(stateBuilder *WasiStateBuilder) (*WasiEnvironment, error) {
	environment := C.wasi_env_new(stateBuilder.inner())

	if environment == nil {
		return nil, newErrorFromWasmer()
	}

	return &WasiEnvironment{
		_inner: environment,
	}, nil
}

func (self *WasiEnvironment) generateImportObject() *ImportObject {
	importObject := NewImportObject()

	/*
		var wasiImports C.wasm_extern_vec_t
		C.wasm_extern_vec_new_uninitialized(&wasiImports
	*/
}
