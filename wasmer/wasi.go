package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"unsafe"
	"runtime"
)

type WasiVersion C.wasi_version_t

const (
	WASI_VERSION_LATEST = WasiVersion(C.LATEST)
	WASI_VERSION_SNAPSHOT0 = WasiVersion(C.SNAPSHOT0)
	WASI_VERSION_SNAPSHOT1 = WasiVersion(C.SNAPSHOT1)
	WASI_VERSION_INVALID = WasiVersion(C.INVALID_VERSION)
)

func (self WasiVersion) String() string {
	switch self {
	case WASI_VERSION_LATEST:
		return "__latest__"
	case WASI_VERSION_SNAPSHOT0:
		return "wasi_unstable"
	case WASI_VERSION_SNAPSHOT1:
		return "wasi_snapshot_preview1"
	case WASI_VERSION_INVALID:
		return "__unknown__"
	}
	panic("Unknown WASI version")
}

func GetWasiVersion(module *Module) WasiVersion {
	return WasiVersion(C.wasi_get_wasi_version(module.inner()))
}

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

func (self *WasiStateBuilder) preopenDirectory(preopenDirectory string) *WasiStateBuilder {
	cPreopenDirectory := C.CString(preopenDirectory)
	defer C.free(unsafe.Pointer(cPreopenDirectory))

	C.wasi_config_preopen_dir(self.inner(), cPreopenDirectory)

	return self
}

func (self *WasiStateBuilder) mapDirectory(alias string, directory string) *WasiStateBuilder {
	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cDirectory := C.CString(directory)
	defer C.free(unsafe.Pointer(cDirectory))

	C.wasi_config_mapdir(self.inner(), cAlias, cDirectory)

	return self
}

func (self *WasiStateBuilder) captureStdin() *WasiStateBuilder {
	C.wasi_config_capture_stdin(self.inner())

	return self
}

func (self *WasiStateBuilder) inheritStdin() *WasiStateBuilder {
	C.wasi_config_inherit_stdin(self.inner())

	return self
}

func (self *WasiStateBuilder) captureStdout() *WasiStateBuilder {
	C.wasi_config_capture_stdout(self.inner())

	return self
}

func (self *WasiStateBuilder) inheritStdout() *WasiStateBuilder {
	C.wasi_config_inherit_stdout(self.inner())

	return self
}

func (self *WasiStateBuilder) captureStderr() *WasiStateBuilder {
	C.wasi_config_capture_stderr(self.inner())

	return self
}

func (self *WasiStateBuilder) inheritStderr() *WasiStateBuilder {
	C.wasi_config_inherit_stderr(self.inner())

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
