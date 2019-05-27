package wasmer

import (
	"fmt"
	"reflect"
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

// ExportedFunctionError represents any kind of errors related to a
// WebAssembly exported function. It is returned by `Instance`
// functions only.
type ExportedFunctionError struct {
	functionName string
	message      string
}

// NewExportedFunctionError constructs a new `ExportedFunctionError`,
// where `functionName` is the name of the exported function, and
// `message` is the error message. If the error message contains `%s`,
// then this parameter will be replaced by `functionName`.
func NewExportedFunctionError(functionName string, message string) *ExportedFunctionError {
	return &ExportedFunctionError{functionName, message}
}

// ExportedFunctionError is an actual error. The `Error` function
// returns the error message.
func (error *ExportedFunctionError) Error() string {
	return fmt.Sprintf(error.message, error.functionName)
}

type Import struct {
	implementation interface{}
	pointer        unsafe.Pointer
}

func NewImport(implementation interface{}, pointer unsafe.Pointer) Import {
	return Import{implementation, pointer}
}

// Instance represents a WebAssembly instance.
type Instance struct {
	// The underlying WebAssembly instance.
	instance *cWasmerInstanceT

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
	Exports map[string]func(...interface{}) (Value, error)

	// The exported memory of a WebAssembly instance.
	Memory Memory
}

// InstanceContext foobar
type InstanceContext = unsafe.Pointer

// NewInstance constructs a new `Instance`.
func NewInstance(bytes []byte) (Instance, error) {
	var imports map[string]Import

	return NewInstanceWithImports(bytes, imports)
}

func NewInstanceWithImports(bytes []byte, imports map[string]Import) (Instance, error) {
	var numberOfImports = len(imports)
	var wasmImports = make([]cWasmerImportT, numberOfImports)
	var importFunctionNth = 0

	for importName, importFunction := range imports {
		var importType = reflect.TypeOf(importFunction.implementation)

		if importType.Kind() != reflect.Func {
			panic(fmt.Sprintf("Import function must be a function."))
		}

		var importFunctionInputsArity = importType.NumIn() - 1 // Skip the first input of kind `InstanceContext`
		var importFunctionOutputsArity = importType.NumOut()
		var importFunctionInputs = make([]cWasmerValueTag, importFunctionInputsArity)
		var importFunctionOutputs = make([]cWasmerValueTag, importFunctionOutputsArity)

		for nth := 0; nth < importFunctionInputsArity; nth++ {
			var importFunctionInput = importType.In(nth + 1)

			switch importFunctionInput.Kind() {
			case reflect.Int32:
				importFunctionInputs[nth] = cWasmI32
			case reflect.Int64:
				importFunctionInputs[nth] = cWasmI64
			case reflect.Float32:
				importFunctionInputs[nth] = cWasmF32
			case reflect.Float64:
				importFunctionInputs[nth] = cWasmF64
			default:
				panic(fmt.Sprintf("Invalid input type for the `%s` imported function.", importName))
			}
		}

		if importFunctionOutputsArity > 1 {
			panic(fmt.Sprintf("The `%s` imported function must have at most one output value.", importName))
		} else if importFunctionOutputsArity == 1 {
			switch importType.Out(0).Kind() {
			case reflect.Int32:
				importFunctionOutputs[0] = cWasmI32
			case reflect.Int64:
				importFunctionOutputs[0] = cWasmI64
			case reflect.Float32:
				importFunctionOutputs[0] = cWasmF32
			case reflect.Float64:
				importFunctionOutputs[0] = cWasmF64
			default:
				panic(fmt.Sprintf("Invalid output type for the `%s` imported function.", importName))
			}
		}

		var importFunctionInputsCPointer *cWasmerValueTag
		var importFunctionOutputsCPointer *cWasmerValueTag

		if importFunctionInputsArity > 0 {
			importFunctionInputsCPointer = (*cWasmerValueTag)(unsafe.Pointer(&importFunctionInputs[0]))
		}

		if importFunctionOutputsArity > 0 {
			importFunctionOutputsCPointer = (*cWasmerValueTag)(unsafe.Pointer(&importFunctionOutputs[0]))
		}

		var function = cWasmerImportFuncNew(
			importFunction.pointer,
			importFunctionInputsCPointer,
			cUint(importFunctionInputsArity),
			importFunctionOutputsCPointer,
			cUint(importFunctionOutputsArity),
		)

		var importedFunction = cNewWasmerImportT("env", importName, function)

		wasmImports[importFunctionNth] = importedFunction

		importFunctionNth++
	}

	var wasmImportsCPointer *cWasmerImportT

	if numberOfImports > 0 {
		wasmImportsCPointer = (*cWasmerImportT)(unsafe.Pointer(&wasmImports[0]))
	}

	var instance *cWasmerInstanceT

	var compileResult = cWasmerInstantiate(
		&instance,
		(*cUchar)(unsafe.Pointer(&bytes[0])),
		cUint(len(bytes)),
		wasmImportsCPointer,
		cInt(numberOfImports),
	)

	var memory Memory
	var emptyInstance = Instance{instance: nil, Exports: nil, Memory: memory}

	if compileResult != cWasmerOk {
		return emptyInstance, NewInstanceError("Failed to compile the module.")
	}

	var exports = make(map[string]func(...interface{}) (Value, error))

	var wasmExports *cWasmerExportsT
	var hasMemory = false

	cWasmerInstanceExports(instance, &wasmExports)
	defer cWasmerExportsDestroy(wasmExports)

	var numberOfExports = int(cWasmerExportsLen(wasmExports))

	for nth := 0; nth < numberOfExports; nth++ {
		var wasmExport = cWasmerExportsGet(wasmExports, cInt(nth))
		var wasmExportKind = cWasmerExportKind(wasmExport)

		switch wasmExportKind {
		case cWasmMemory:
			var wasmMemory *cWasmerMemoryT

			if cWasmerExportToMemory(wasmExport, &wasmMemory) != cWasmerOk {
				return emptyInstance, NewInstanceError("Failed to extract the exported memory.")
			}

			memory = newMemory(wasmMemory)
			hasMemory = true

		case cWasmFunction:
			var wasmExportName = cWasmerExportName(wasmExport)
			var exportedFunctionName = cGoStringN((*cChar)(unsafe.Pointer(wasmExportName.bytes)), (cInt)(wasmExportName.bytes_len))
			var wasmFunction = cWasmerExportToFunc(wasmExport)
			var wasmFunctionInputsArity cUint32T

			if cWasmerExportFuncParamsArity(wasmFunction, &wasmFunctionInputsArity) != cWasmerOk {
				return emptyInstance, NewExportedFunctionError(exportedFunctionName, "Failed to read the input arity of the `%s` exported function.")
			}

			var wasmFunctionInputSignatures = make([]cWasmerValueTag, int(wasmFunctionInputsArity))

			if wasmFunctionInputsArity > 0 {
				var wasmFunctionInputSignaturesCPointer = (*cWasmerValueTag)(unsafe.Pointer(&wasmFunctionInputSignatures[0]))

				if cWasmerExportFuncParams(wasmFunction, wasmFunctionInputSignaturesCPointer, wasmFunctionInputsArity) != cWasmerOk {
					return emptyInstance, NewExportedFunctionError(exportedFunctionName, "Failed to read the signature of the `%s` exported function.")
				}
			}

			var wasmFunctionOutputsArity cUint32T

			if cWasmerExportFuncResultsArity(wasmFunction, &wasmFunctionOutputsArity) != cWasmerOk {
				return emptyInstance, NewExportedFunctionError(exportedFunctionName, "Failed to read the output arity of the `%s` exported function.")
			}

			var numberOfExpectedArguments = int(wasmFunctionInputsArity)

			exports[exportedFunctionName] = func(arguments ...interface{}) (Value, error) {
				var numberOfGivenArguments = len(arguments)
				var diff = numberOfExpectedArguments - numberOfGivenArguments

				if diff > 0 {
					return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Missing %d argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", diff, numberOfExpectedArguments, numberOfGivenArguments))
				} else if diff < 0 {
					return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Given %d extra argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", -diff, numberOfExpectedArguments, numberOfGivenArguments))
				}

				var wasmInputs = make([]cWasmerValueT, wasmFunctionInputsArity)

				for nth, value := range arguments {
					var wasmInputType = wasmFunctionInputSignatures[nth]

					switch wasmInputType {
					case cWasmI32:
						wasmInputs[nth].tag = cWasmI32
						var pointer = (*int32)(unsafe.Pointer(&wasmInputs[nth].value))

						switch value.(type) {
						case int8:
							*pointer = int32(value.(int8))
						case uint8:
							*pointer = int32(value.(uint8))
						case int16:
							*pointer = int32(value.(int16))
						case uint16:
							*pointer = int32(value.(uint16))
						case int32:
							*pointer = int32(value.(int32))
						case int:
							*pointer = int32(value.(int))
						case uint:
							*pointer = int32(value.(uint))
						case Value:
							var value = value.(Value)

							if value.GetType() != TypeI32 {
								return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i32`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToI32()
						default:
							return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i32`, cannot cast given value to this type.", nth+1))
						}
					case cWasmI64:
						wasmInputs[nth].tag = cWasmI64
						var pointer = (*int64)(unsafe.Pointer(&wasmInputs[nth].value))

						switch value.(type) {
						case int8:
							*pointer = int64(value.(int8))
						case uint8:
							*pointer = int64(value.(uint8))
						case int16:
							*pointer = int64(value.(int16))
						case uint16:
							*pointer = int64(value.(uint16))
						case int32:
							*pointer = int64(value.(int32))
						case uint32:
							*pointer = int64(value.(uint32))
						case int64:
							*pointer = int64(value.(int64))
						case int:
							*pointer = int64(value.(int))
						case uint:
							*pointer = int64(value.(uint))
						case Value:
							var value = value.(Value)

							if value.GetType() != TypeI64 {
								return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i64`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToI64()
						default:
							return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i64`, cannot cast given value to this type.", nth+1))
						}
					case cWasmF32:
						wasmInputs[nth].tag = cWasmF32
						var pointer = (*float32)(unsafe.Pointer(&wasmInputs[nth].value))

						switch value.(type) {
						case float32:
							*pointer = value.(float32)
						case Value:
							var value = value.(Value)

							if value.GetType() != TypeF32 {
								return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f32`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToF32()
						default:
							return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f32`, cannot cast given value to this type.", nth+1))
						}
					case cWasmF64:
						wasmInputs[nth].tag = cWasmF64
						var pointer = (*float64)(unsafe.Pointer(&wasmInputs[nth].value))

						switch value.(type) {
						case float32:
							*pointer = float64(value.(float32))
						case float64:
							*pointer = value.(float64)
						case Value:
							var value = value.(Value)

							if value.GetType() != TypeF64 {
								return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f64`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToF64()
						default:
							return I32(0), NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f64`, cannot cast given value to this type.", nth+1))
						}
					default:
						panic(fmt.Sprintf("Invalid arguments type when calling the `%s` exported function.", exportedFunctionName))
					}
				}

				var wasmOutputs = make([]cWasmerValueT, wasmFunctionOutputsArity)

				var wasmFunctionName = cCString(exportedFunctionName)
				defer cFree(unsafe.Pointer(wasmFunctionName))

				var wasmInputsCPointer *cWasmerValueT

				if wasmFunctionInputsArity > 0 {
					wasmInputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmInputs[0]))
				} else {
					wasmInputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmInputs))
				}

				var wasmOutputsCPointer *cWasmerValueT

				if wasmFunctionOutputsArity > 0 {
					wasmOutputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmOutputs[0]))
				} else {
					wasmOutputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmOutputs))
				}

				var callResult = cWasmerInstanceCall(
					instance,
					wasmFunctionName,
					wasmInputsCPointer,
					wasmFunctionInputsArity,
					wasmOutputsCPointer,
					wasmFunctionOutputsArity,
				)

				if callResult != cWasmerOk {
					return I32(0), NewExportedFunctionError(exportedFunctionName, "Failed to call the `%s` exported function.")
				}

				if wasmFunctionOutputsArity > 0 {
					var result = wasmOutputs[0]

					switch result.tag {
					case cWasmI32:
						pointer := (*int32)(unsafe.Pointer(&result.value))

						return I32(*pointer), nil
					case cWasmI64:
						pointer := (*int64)(unsafe.Pointer(&result.value))

						return I64(*pointer), nil
					case cWasmF32:
						pointer := (*float32)(unsafe.Pointer(&result.value))

						return F32(*pointer), nil
					case cWasmF64:
						pointer := (*float64)(unsafe.Pointer(&result.value))

						return F64(*pointer), nil
					default:
						panic("unreachable")
					}
				} else {
					return void(), nil
				}
			}
		}
	}

	if hasMemory == false {
		return emptyInstance, NewInstanceError("No memory exported.")
	}

	return Instance{instance: instance, Exports: exports, Memory: memory}, nil
}

// Close closes/frees an `Instance`.
func (instance *Instance) Close() {
	if instance.instance != nil {
		cWasmerInstanceDestroy(instance.instance)
	}
}
