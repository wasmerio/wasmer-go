package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"unsafe"
	"runtime"
)

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

func (self *WasiStateBuilder) environment(key string, value string) *WasiStateBuilder {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.wasi_config_env(self.inner(), cKey, cValue)

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

	runtime.SetFinalizer(environment, func(environment *C.wasi_env_t) {
		C.wasi_env_delete(environment)
	})

	return &WasiEnvironment{
		_inner: environment,
	}, nil
}

func (self *WasiEnvironment) inner() *C.wasi_env_t {
	return self._inner
}

func (self *WasiEnvironment) generateImportObject(store *Store, module *Module) (*ImportObject, error) {
	var wasiNamedExterns C.wasm_named_extern_vec_t
	C.wasm_named_extern_vec_new_empty(&wasiNamedExterns)

	var success = C.wasi_get_unordered_imports(store.inner(), module.inner(), self.inner(), &wasiNamedExterns)

	if success == false {
		return nil, newErrorFromWasmer()
	}

	importObject := NewImportObject()

	numberOfNamedExterns := int(wasiNamedExterns.size)
	firstNamedExtern := unsafe.Pointer(wasiNamedExterns.data)
	sizeOfNamedExtern := unsafe.Sizeof(firstNamedExtern)

	var currentNamedExtern *C.wasm_named_extern_t

	for nth := 0; nth < numberOfNamedExterns; nth++ {
		currentNamedExtern = *(**C.wasm_named_extern_t)(unsafe.Pointer(uintptr(firstNamedExtern) + uintptr(nth)*sizeOfNamedExtern))
		module := nameToString(C.wasm_named_extern_module(currentNamedExtern))
		name := nameToString(C.wasm_named_extern_name(currentNamedExtern))
		extern := newExtern(C.wasm_extern_copy(C.wasm_named_extern_unwrap(currentNamedExtern)), nil)

		_, exists := importObject.externs[module]

		if exists == false {
			importObject.externs[module] = make(map[string]IntoExtern)
		}

		importObject.externs[module][name] = extern
	}

	C.wasm_named_extern_vec_delete(&wasiNamedExterns)

	return importObject, nil
}
