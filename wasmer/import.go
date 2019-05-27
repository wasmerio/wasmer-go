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
}

// Imports represents a set of imported functions for a WebAssembly instance.
type Imports struct {
	imports map[string]Import
}

// NewImports constructs a new empty `Imports`.
func NewImports() *Imports {
	var imports = make(map[string]Import)

	return &Imports{imports}
}

// Append adds a new imported function to the current set.
func (imports *Imports) Append(importName string, implementation interface{}, cgoPointer unsafe.Pointer) (*Imports, error) {
	var importType = reflect.TypeOf(implementation)

	if importType.Kind() != reflect.Func {
		return nil, NewImportedFunctionError(importName, fmt.Sprintf("Import function `%%s` must be a function; given %s.", importType.Kind()))
	}

	var importInputsArity = importType.NumIn() - 1 // Skip the first input of kind `InstanceContext`
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
			return nil, NewImportedFunctionError(importName, "Invalid input type for the `%s` imported function.")
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
			return nil, NewImportedFunctionError(importName, "Invalid output type for the `%s` imported function.")
		}
	}

	var importedFunctionPointer *cWasmerImportFuncT

	imports.imports[importName] = Import{
		implementation,
		cgoPointer,
		importedFunctionPointer,
		wasmInputs,
		wasmOutputs,
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
