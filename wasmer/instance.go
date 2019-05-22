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

// Instance represents a WebAssembly instance.
type Instance struct {
	// The underlying WebAssembly instance.
	instance *C_wasmer_instance_t

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

// NewInstance constructs a new `Instance`.
func NewInstance(bytes []byte) (Instance, error) {
	var imports = []C_wasmer_import_t{}
	var instance *C_wasmer_instance_t

	var compileResult = C_wasmer_instantiate(
		&instance,
		(*C_uchar)(unsafe.Pointer(&bytes[0])),
		C_uint(len(bytes)),
		(*C_wasmer_import_t)(unsafe.Pointer(&imports)),
		C_int(0),
	)

	var memory Memory
	var emptyInstance = Instance{instance: nil, Exports: nil, Memory: memory}

	if compileResult != C_WASMER_OK {
		return emptyInstance, NewInstanceError("Failed to compile the module.")
	}

	var exports = make(map[string]func(...interface{}) (Value, error))

	var wasmExports *C_wasmer_exports_t
	var hasMemory = false

	C_wasmer_instance_exports(instance, &wasmExports)
	defer C_wasmer_exports_destroy(wasmExports)

	var numberOfExports = int(C_wasmer_exports_len(wasmExports))

	for nth := 0; nth < numberOfExports; nth++ {
		var wasmExport = C_wasmer_exports_get(wasmExports, C_int(nth))
		var wasmExportKind = C_wasmer_export_kind(wasmExport)

		switch wasmExportKind {
		case C_WASM_MEMORY:
			var wasmMemory *C_wasmer_memory_t

			if C_wasmer_export_to_memory(wasmExport, &wasmMemory) != C_WASMER_OK {
				return emptyInstance, NewInstanceError("Failed to extract the exported memory.")
			}

			memory = newMemory(wasmMemory)
			hasMemory = true

		case C_WASM_FUNCTION:
			var wasmExportName = C_wasmer_export_name(wasmExport)
			var exportedFunctionName = C_GoStringN((*C_char)(unsafe.Pointer(wasmExportName.bytes)), (C_int)(wasmExportName.bytes_len))
			var wasmFunction = C_wasmer_export_to_func(wasmExport)
			var wasmFunctionInputsArity C_uint32_t

			if C_wasmer_export_func_params_arity(wasmFunction, &wasmFunctionInputsArity) != C_WASMER_OK {
				return emptyInstance, NewExportedFunctionError(exportedFunctionName, "Failed to read the input arity of the `%s` exported function.")
			}

			var wasmFunctionInputSignatures = make([]C_wasmer_value_tag, int(wasmFunctionInputsArity))

			if wasmFunctionInputsArity > 0 {
				var wasmFunctionInputSignaturesCPointer = (*C_wasmer_value_tag)(unsafe.Pointer(&wasmFunctionInputSignatures[0]))

				if C_wasmer_export_func_params(wasmFunction, wasmFunctionInputSignaturesCPointer, wasmFunctionInputsArity) != C_WASMER_OK {
					return emptyInstance, NewExportedFunctionError(exportedFunctionName, "Failed to read the signature of the `%s` exported function.")
				}
			}

			var wasmFunctionOutputsArity C_uint32_t

			if C_wasmer_export_func_returns_arity(wasmFunction, &wasmFunctionOutputsArity) != C_WASMER_OK {
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

				var wasmInputs = make([]C_wasmer_value_t, wasmFunctionInputsArity)

				for nth, value := range arguments {
					var wasmInputType = wasmFunctionInputSignatures[nth]

					switch wasmInputType {
					case C_WASM_I32:
						wasmInputs[nth].tag = C_WASM_I32
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
					case C_WASM_I64:
						wasmInputs[nth].tag = C_WASM_I64
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
					case C_WASM_F32:
						wasmInputs[nth].tag = C_WASM_F32
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
					case C_WASM_F64:
						wasmInputs[nth].tag = C_WASM_F64
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

				var wasmOutputs = make([]C_wasmer_value_t, wasmFunctionOutputsArity)

				var wasmFunctionName = C_CString(exportedFunctionName)
				defer C_free(unsafe.Pointer(wasmFunctionName))

				var wasmInputsCPointer *C_wasmer_value_t

				if wasmFunctionInputsArity > 0 {
					wasmInputsCPointer = (*C_wasmer_value_t)(unsafe.Pointer(&wasmInputs[0]))
				} else {
					wasmInputsCPointer = (*C_wasmer_value_t)(unsafe.Pointer(&wasmInputs))
				}

				var wasmOutputsCPointer *C_wasmer_value_t

				if wasmFunctionOutputsArity > 0 {
					wasmOutputsCPointer = (*C_wasmer_value_t)(unsafe.Pointer(&wasmOutputs[0]))
				} else {
					wasmOutputsCPointer = (*C_wasmer_value_t)(unsafe.Pointer(&wasmOutputs))
				}

				var callResult = C_wasmer_instance_call(
					instance,
					wasmFunctionName,
					wasmInputsCPointer,
					wasmFunctionInputsArity,
					wasmOutputsCPointer,
					wasmFunctionOutputsArity,
				)

				if callResult != C_WASMER_OK {
					return I32(0), NewExportedFunctionError(exportedFunctionName, "Failed to call the `%s` exported function.")
				}

				if wasmFunctionOutputsArity > 0 {
					var result = wasmOutputs[0]

					switch result.tag {
					case C_WASM_I32:
						pointer := (*int32)(unsafe.Pointer(&result.value))

						return I32(*pointer), nil
					case C_WASM_I64:
						pointer := (*int64)(unsafe.Pointer(&result.value))

						return I64(*pointer), nil
					case C_WASM_F32:
						pointer := (*float32)(unsafe.Pointer(&result.value))

						return F32(*pointer), nil
					case C_WASM_F64:
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
		C_wasmer_instance_destroy(instance.instance)
	}
}
