package wasmer

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Import struct {
	implementation          interface{}
	pointer                 unsafe.Pointer
	importedFunctionPointer *cWasmerImportFuncT
	wasmInputs              []cWasmerValueTag
	wasmOutputs             []cWasmerValueTag
}

type Imports struct {
	imports map[string]Import
}

func NewImports() *Imports {
	var imports = make(map[string]Import)

	return &Imports{imports}
}

func (imports *Imports) Append(importName string, implementation interface{}, pointer unsafe.Pointer) *Imports {
	var importType = reflect.TypeOf(implementation)

	if importType.Kind() != reflect.Func {
		panic(fmt.Sprintf("Import function must be a function."))
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
			panic(fmt.Sprintf("Invalid input type for the `%s` imported function.", importName))
		}
	}

	if importOutputsArity > 1 {
		panic(fmt.Sprintf("The `%s` imported function must have at most one output value.", importName))
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
			panic(fmt.Sprintf("Invalid output type for the `%s` imported function.", importName))
		}
	}

	var importedFunctionPointer *cWasmerImportFuncT = nil

	imports.imports[importName] = Import{
		implementation,
		pointer,
		importedFunctionPointer,
		wasmInputs,
		wasmOutputs,
	}

	return imports
}

func (imports *Imports) Close() {
	for _, importFunction := range imports.imports {
		if nil != importFunction.importedFunctionPointer {
			cWasmerImportFuncDestroy(importFunction.importedFunctionPointer)
		}
	}
}
