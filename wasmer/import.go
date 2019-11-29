package wasmer

import (
	"fmt"
	"reflect"
	"unsafe"
)

// ImportedFunctionError represents any kind of errors related to a
// WebAssembly imported function. It is returned by `Import` or `Imports`
// functions only.
type ImportedFunctionError struct {
	functionName string
	message      string
}

// ImportObjectError represents errors related to `ImportObject`s.
type ImportObjectError struct {
	message string
}

// ImportObject owns a set of imports.
// It can be combined with a `Module` to create an `Instance`.
type ImportObject struct {
	inner *cWasmerImportObjectT
}

// NewImportObject creates an empty `ImportObject`.
func NewImportObject() *ImportObject {
	var inner = cNewWasmerImportObject()

	return &ImportObject{inner}
}

// Imports returns `*Imports` for a given `ImportObject`
func (importObject *ImportObject) Imports() (*Imports, error) {
	imports := cWasmerImportObjectGetFunctions(importObject.inner)

	if imports == nil {
		return nil, NewImportObjectError("Could not get functions from the given import object")
	}

	output := NewImports()

	for _, impoort := range imports {
		rawFunction := cGetFunctionFromImport(&impoort)

		if rawFunction == nil {
			// This is expected to never happen
			continue
		}

		namespaceName, importName := cGetInfoFromImport(&impoort)
		nextOutput, err := output.appendRaw(namespaceName, importName, rawFunction)

		if err != nil {
			return nil, err
		}

		output = nextOutput
	}

	return output, nil
}

// Extend adds the given imports to the existing import object
func (importObject *ImportObject) Extend(imports Imports) error {
	var numberOfImports = len(imports.imports)

	if numberOfImports == 0 {
		return nil
	}

	var cImports = make([]cWasmerImportT, numberOfImports)
	var importFunctionNth = 0

	for importName, importFunction := range imports.imports {
		cImports[importFunctionNth] = *getCWasmerImport(importName, importFunction)
		importFunctionNth++
	}

	if importFunctionNth == 0 {
		return nil
	}

	var extendResult = cWasmerImportObjectExtend(
		importObject.inner,
		(*cWasmerImportT)(unsafe.Pointer(&cImports[0])),
		(cUint)(len(imports.imports)),
	)

	if extendResult != cWasmerOk {
		return NewImportObjectError("Could not extend import object with the given imports")
	}

	return nil
}

// Close frees the `ImportObject`
func (importObject *ImportObject) Close() {
	cWasmerImportObjectDestroy(importObject.inner)
}

// NewImportObjectError constructs a new `ImportObjectError`
func NewImportObjectError(message string) *ImportObjectError {
	return &ImportObjectError{message}
}

// ImportObjectError is an actual error. The `Error` function
// returns the error message.
func (error *ImportObjectError) Error() string {
	return fmt.Sprintf(error.message)
}

// NewImportedFunctionError constructs a new `ImportedFunctionError`,
// where `functionName` is the name of the imported function, and
// `message` is the error message. If the error message contains `%s`,
// then this parameter will be replaced by `functionName`.
func NewImportedFunctionError(functionName string, message string) *ImportedFunctionError {
	return &ImportedFunctionError{functionName, message}
}

// ImportedFunctionError is an actual error. The `Error` function
// returns the error message.
func (error *ImportedFunctionError) Error() string {
	return fmt.Sprintf(error.message, error.functionName)
}

// Import represents an WebAssembly instance imported function.
type Import struct {
	// An implementation must be of type:
	// `func(context unsafe.Pointer, arguments ...interface{}) interface{}`.
	// It represents the real function implementation written in Go.
	implementation interface{}

	// The pointer to the cgo function implementation, something
	// like `C.foo`.
	cgoPointer unsafe.Pointer

	// The pointer to the Wasmer imported function.
	importedFunctionPointer *cWasmerImportFuncT

	// The function implementation signature as a WebAssembly signature.
	wasmInputs []cWasmerValueTag

	// The function implementation signature as a WebAssembly signature.
	wasmOutputs []cWasmerValueTag

	// The namespace of the imported function.
	namespace string
}

// Imports represents a set of imported functions for a WebAssembly instance.
type Imports struct {
	// All imports.
	imports map[string]Import

	// Current namespace where to register the import.
	currentNamespace string
}

// NewImports constructs a new empty `Imports`.
func NewImports() *Imports {
	var imports = make(map[string]Import)
	var currentNamespace = "env"

	return &Imports{imports, currentNamespace}
}

// Namespace changes the current namespace of the next imported functions.
func (imports *Imports) Namespace(namespace string) *Imports {
	imports.currentNamespace = namespace

	return imports
}

// Append adds a new imported function to the current set.
func (imports *Imports) Append(importName string, implementation interface{}, cgoPointer unsafe.Pointer) (*Imports, error) {
	var importType = reflect.TypeOf(implementation)

	if importType.Kind() != reflect.Func {
		return nil, NewImportedFunctionError(importName, fmt.Sprintf("Imported function `%%s` must be a function; given `%s`.", importType.Kind()))
	}

	var importInputsArity = importType.NumIn()

	if importInputsArity < 1 {
		return nil, NewImportedFunctionError(importName, "Imported function `%s` must at least have one argument for the instance context.")
	}

	if importType.In(0).Kind() != reflect.UnsafePointer {
		return nil, NewImportedFunctionError(importName, fmt.Sprintf("The instance context of the `%%s` imported function must be of kind `unsafe.Pointer`; given `%s`; is it missing?", importType.In(0).Kind()))
	}

	importInputsArity--
	var importOutputsArity = importType.NumOut()
	var wasmInputs = make([]cWasmerValueTag, importInputsArity)
	var wasmOutputs = make([]cWasmerValueTag, importOutputsArity)

	for nth := 0; nth < importInputsArity; nth++ {
		var importInput = importType.In(nth + 1)

		switch importInput.Kind() {
		case reflect.Int32:
			wasmInputs[nth] = cWasmI32
		case reflect.Int64:
			wasmInputs[nth] = cWasmI64
		case reflect.Float32:
			wasmInputs[nth] = cWasmF32
		case reflect.Float64:
			wasmInputs[nth] = cWasmF64
		default:
			return nil, NewImportedFunctionError(importName, fmt.Sprintf("Invalid input type for the `%%s` imported function; given `%s`; only accept `int32`, `int64`, `float32`, and `float64`.", importInput.Kind()))
		}
	}

	if importOutputsArity > 1 {
		return nil, NewImportedFunctionError(importName, "The `%s` imported function must have at most one output value.")
	} else if importOutputsArity == 1 {
		switch importType.Out(0).Kind() {
		case reflect.Int32:
			wasmOutputs[0] = cWasmI32
		case reflect.Int64:
			wasmOutputs[0] = cWasmI64
		case reflect.Float32:
			wasmOutputs[0] = cWasmF32
		case reflect.Float64:
			wasmOutputs[0] = cWasmF64
		default:
			return nil, NewImportedFunctionError(importName, fmt.Sprintf("Invalid output type for the `%%s` imported function; given `%s`; only accept `int32`, `int64`, `float32`, and `float64`.", importType.Out(0).Kind()))
		}
	}

	var importedFunctionPointer *cWasmerImportFuncT
	var namespace = imports.currentNamespace

	imports.imports[importName] = Import{
		implementation,
		cgoPointer,
		importedFunctionPointer,
		wasmInputs,
		wasmOutputs,
		namespace,
	}

	return imports, nil
}

// Like Append but not for Go imports.
func (imports *Imports) appendRaw(
	namespace string,
	importName string,
	wasmerImportFunc *cWasmerImportFuncT,
) (*Imports, error) {
	wasmInputs := cGetParamsForImportFunc(wasmerImportFunc)

	if wasmInputs == nil {
		return imports, NewImportedFunctionError(importName, fmt.Sprintf("Could not get the inputs for `%%s` in namespace `%s`", namespace))
	}

	wasmOutputs := cGetReturnsForImportFunc(wasmerImportFunc)

	if wasmOutputs == nil {
		return imports, NewImportedFunctionError(importName, fmt.Sprintf("Could not get the outputs for `%%s` in namespace `%s`", namespace))
	}

	imports.imports[importName] = Import{
		nil,
		unsafe.Pointer(wasmerImportFunc),
		wasmerImportFunc,
		wasmInputs,
		wasmOutputs,
		namespace,
	}

	return imports, nil
}

// Close closes/frees all imported functions that have been registered by Wasmer.
func (imports *Imports) Close() {
	for _, importFunction := range imports.imports {
		if nil != importFunction.importedFunctionPointer {
			cWasmerImportFuncDestroy(importFunction.importedFunctionPointer)
		}
	}
}

// Helper function: Get a C import for a given import
func getCWasmerImport(importName string, importFunction Import) *cWasmerImportT {
	var wasmInputsArity = len(importFunction.wasmInputs)
	var wasmOutputsArity = len(importFunction.wasmOutputs)

	var importFunctionInputsCPointer *cWasmerValueTag
	var importFunctionOutputsCPointer *cWasmerValueTag

	if wasmInputsArity > 0 {
		importFunctionInputsCPointer = (*cWasmerValueTag)(unsafe.Pointer(&importFunction.wasmInputs[0]))
	}

	if wasmOutputsArity > 0 {
		importFunctionOutputsCPointer = (*cWasmerValueTag)(unsafe.Pointer(&importFunction.wasmOutputs[0]))
	}

	importFunction.importedFunctionPointer = cWasmerImportFuncNew(
		importFunction.cgoPointer,
		importFunctionInputsCPointer,
		cUint(wasmInputsArity),
		importFunctionOutputsCPointer,
		cUint(wasmOutputsArity),
	)
	var newImport = cNewWasmerImportT(
		importFunction.namespace,
		importName,
		importFunction.importedFunctionPointer,
	)

	return &newImport
}

// InstanceContext represents a way to access instance API from within
// an imported context.
type InstanceContext struct {
	context *cWasmerInstanceContextT
	memory  Memory
}

// IntoInstanceContext casts the first `context unsafe.Pointer`
// argument of an imported function into an `InstanceContext`.
func IntoInstanceContext(instanceContext unsafe.Pointer) InstanceContext {
	context := (*cWasmerInstanceContextT)(instanceContext)
	memory := newMemory(cWasmerInstanceContextMemory(context))

	return InstanceContext{context, memory}
}

// Memory returns the current instance memory.
func (instanceContext *InstanceContext) Memory() *Memory {
	return &instanceContext.memory
}

// Data returns the instance context data as an `interface{}`. It's
// up to the user to cast it appropriately.
func (instanceContext *InstanceContext) Data() interface{} {
	dataPtr := *(*unsafe.Pointer)(cWasmerInstanceContextDataGet(instanceContext.context))

	return *(*interface{})(dataPtr)
}
