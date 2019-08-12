package wasmer

import (
	"fmt"
	"unsafe"
)

// InstanceError represents any kind of errors related to a WebAssembly instance. It
// is returned by `Instance` functions only.
type InstanceError struct {
	// Error message.
	message string
}

// NewInstanceError constructs a new `InstanceError`.
func NewInstanceError(message string) *InstanceError {
	return &InstanceError{message}
}

// `InstanceError` is an actual error. The `Error` function returns
// the error message.
func (error *InstanceError) Error() string {
	return error.message
}

// Instance represents a WebAssembly instance.
type Instance struct {
	// The underlying WebAssembly instance.
	instance *cWasmerInstanceT

	// The imported functions. Use the `NewInstanceWithImports`
	// constructor to set it.
	imports *Imports

	// All functions exported by the WebAssembly instance, indexed
	// by their name as a string. An exported function is a
	// regular variadic Go closure. Arguments are untyped. Since
	// WebAssembly only supports: `i32`, `i64`, `f32` and `f64`
	// types, the accepted Go types are: `int8`, `uint8`, `int16`,
	// `uint16`, `int32`, `uint32`, `int64`, `int`, `uint`, `float32`
	// and `float64`. In addition to those types, the `Value` type
	// (from this project) is accepted. The conversion from a Go
	// value to a WebAssembly value is done automatically except for
	// the `Value` type (where type is coerced, that's the intent
	// here). The WebAssembly type is automatically inferred. Note
	// that the returned value is of kind `Value`, and not a
	// standard Go type.
	Exports map[string]ExportFunctionT

	// The exported memory of a WebAssembly instance.
	Memory Memory
}

// NewInstance constructs a new `Instance` with no imported functions.
func NewInstance(bytes []byte) (Instance, error) {
	return NewInstanceWithImports(bytes, NewImports())
}

// NewInstanceWithImports constructs a new `Instance` with imported functions.
func NewInstanceWithImports(bytes []byte, imports *Imports) (Instance, error) {
	return newInstanceWithImports(
		imports,
		func(wasmImportsCPointer *cWasmerImportT, numberOfImports uint) (*cWasmerInstanceT, error) {
			var instance *cWasmerInstanceT

			var compileResult = cWasmerInstantiate(
				&instance,
				(*cUchar)(unsafe.Pointer(&bytes[0])),
				cUint(len(bytes)),
				wasmImportsCPointer,
				cInt(numberOfImports),
			)

			if compileResult != cWasmerOk {
				return nil, buildInstantiateError()
			}

			return instance, nil
		},
	)
}

func NewInstanceFromModule(
	module *Module,
	imports *Imports,
) (Instance, error) {
	return NewInstanceWithModuleAndImportObject(module, imports, DefaultImportObjectBuilder)
}

func NewInstanceWithModuleAndImportObject(
	module *Module,
	imports *Imports,
	importObjectBuilder func(*cWasmerImportT, uint) (*cWasmerImportObjectT, error),
) (Instance, error) {
	return newInstanceWithImports(
		imports,
		func(wasmImportsCPointer *cWasmerImportT, numberOfImports uint) (*cWasmerInstanceT, error) {
			var instance *cWasmerInstanceT

			importObject, err := importObjectBuilder(wasmImportsCPointer, numberOfImports)
			if err != nil {
				return nil, err
			}

			var compileResult = cWasmerModuleImportInstantiate(&instance, module.module, importObject)
			if compileResult != cWasmerOk {
				return nil, buildInstantiateError()
			}

			return instance, nil
		})
}

func newInstanceWithImports(
	imports *Imports,
	instanceBuilder func(*cWasmerImportT, uint) (*cWasmerInstanceT, error),
) (Instance, error) {
	wasmImportsCPointer, numberOfImports := imports.ToWasmerImports()

	var emptyInstance = Instance{instance: nil, imports: nil, Exports: nil, Memory: Memory{}}

	instance, err := instanceBuilder(wasmImportsCPointer, numberOfImports)
	if err != nil {
		return emptyInstance, err
	}

	exports, err := instanceExports(instance)
	if err != nil {
		return emptyInstance, err
	}

	if exports.memory == nil {
		return emptyInstance, NewInstanceError("No memory exported.")
	}

	return Instance{instance: instance, imports: imports, Exports: exports.functions, Memory: *exports.memory}, nil
}

// SetContextData assigns a data that can be used by all imported
// functions. Indeed, each imported function receives as its first
// argument an instance context (see `InstanceContext`). An instance
// context can hold a pointer to any kind of data. It is important to
// understand that this data is shared by all imported function, it's
// global to the instance.
func (instance *Instance) SetContextData(data unsafe.Pointer) {
	cWasmerInstanceContextDataSet(instance.instance, data)
}

/// Returns an `Unsafe.Pointer`	to the underlying instance context.
func (instance *Instance) Context() unsafe.Pointer {
	return (unsafe.Pointer)(cInstanceContextGet(instance.instance))
}

// Close closes/frees an `Instance`.
func (instance *Instance) Close() {
	if instance.imports != nil {
		instance.imports.Close()
	}

	if instance.instance != nil {
		cWasmerInstanceDestroy(instance.instance)
	}
}

func buildInstantiateError() error {
	var lastError, err = GetLastError()
	var errorMessage = "Failed to instantiate the module:\n    %s"

	if err != nil {
		errorMessage = fmt.Sprintf(errorMessage, "(unknown details)")
	} else {
		errorMessage = fmt.Sprintf(errorMessage, lastError)
	}

	return NewInstanceError(errorMessage)
}
