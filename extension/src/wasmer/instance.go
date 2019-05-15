package wasmer

/*
#cgo LDFLAGS: -L../../ -lwasmer_runtime_c_api
#include "../../wasmer.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type InstanceError struct {
	message string
}

func NewInstanceError(message string) error {
	return &InstanceError{message}
}

func (self *InstanceError) Error() string {
	return self.message
}

type ExportedFunctionError struct {
	function_name string
	message       string
}

func NewExportedFunctionError(function_name string, message string) error {
	return &ExportedFunctionError{function_name, message}
}

func (self *ExportedFunctionError) Error() string {
	return fmt.Sprintf(self.message, self.function_name)
}

type Instance struct {
	instance *C.wasmer_instance_t
	Exports  map[string]func(...interface{}) (Value, error)
	Memory   Memory
}

func NewInstance(bytes []byte) (Instance, error) {
	var imports []C.wasmer_import_t = []C.wasmer_import_t{}
	var instance *C.wasmer_instance_t = nil

	var compile_result C.wasmer_result_t = C.wasmer_instantiate(
		&instance,
		(*C.uchar)(unsafe.Pointer(&bytes[0])),
		C.uint(len(bytes)),
		(*C.wasmer_import_t)(unsafe.Pointer(&imports)),
		C.int(0),
	)

	var memory Memory
	var empty_instance Instance = Instance{instance: nil, Exports: nil, Memory: memory}

	if C.WASMER_OK != compile_result {
		return empty_instance, NewInstanceError("Failed to compile the module.")
	}

	var exports map[string]func(...interface{}) (Value, error) = make(map[string]func(...interface{}) (Value, error))

	var wasm_exports *C.wasmer_exports_t = nil
	var has_memory bool = false

	C.wasmer_instance_exports(instance, &wasm_exports)
	defer C.wasmer_exports_destroy(wasm_exports)

	var number_of_exports int = int(C.wasmer_exports_len(wasm_exports))

	for nth := 0; nth < number_of_exports; nth++ {
		var wasm_export *C.wasmer_export_t = C.wasmer_exports_get(wasm_exports, C.int(nth))
		var wasm_export_kind C.wasmer_import_export_kind = C.wasmer_export_kind(wasm_export)

		switch wasm_export_kind {
		case C.WASM_MEMORY:
			var wasm_memory *C.wasmer_memory_t = nil

			if C.wasmer_export_to_memory(wasm_export, &wasm_memory) != C.WASMER_OK {
				return empty_instance, NewInstanceError("Failed when extracting the exported memory.")
			}

			var wasm_memory_data *C.uint8_t = C.wasmer_memory_data(wasm_memory)
			var wasm_memory_data_length C.uint32_t = C.wasmer_memory_data_length(wasm_memory)

			memory = NewMemory((*uint8)(wasm_memory_data), uint32(wasm_memory_data_length))
			has_memory = true

		case C.WASM_FUNCTION:
			var wasm_export_name C.wasmer_byte_array = C.wasmer_export_name(wasm_export)
			var exported_function_name = C.GoStringN((*C.char)(unsafe.Pointer(wasm_export_name.bytes)), (C.int)(wasm_export_name.bytes_len))
			var wasm_function *C.wasmer_export_func_t = C.wasmer_export_to_func(wasm_export)
			var wasm_function_inputs_arity C.uint

			if C.wasmer_export_func_params_arity(wasm_function, &wasm_function_inputs_arity) != C.WASMER_OK {
				return empty_instance, NewExportedFunctionError(exported_function_name, "Failed to read the input arity of the `%s` exported function.")
			}

			var wasm_function_input_signatures []C.wasmer_value_tag = make([]C.wasmer_value_tag, int(wasm_function_inputs_arity))

			if wasm_function_inputs_arity > 0 {
				var wasm_function_input_signatures_c_pointer *C.wasmer_value_tag = (*C.wasmer_value_tag)(unsafe.Pointer(&wasm_function_input_signatures[0]))

				if C.wasmer_export_func_params(wasm_function, wasm_function_input_signatures_c_pointer, C.int(wasm_function_inputs_arity)) != C.WASMER_OK {
					return empty_instance, NewExportedFunctionError(exported_function_name, "Failed to read the signature of the `%s` exported function.")
				}
			}

			var wasm_function_outputs_arity C.uint

			if C.wasmer_export_func_returns_arity(wasm_function, &wasm_function_outputs_arity) != C.WASMER_OK {
				return empty_instance, NewExportedFunctionError(exported_function_name, "Failed to read the output arity of the `%s` exported function.")
			}

			var number_of_expected_arguments int = int(wasm_function_inputs_arity)

			exports[exported_function_name] = func(arguments ...interface{}) (Value, error) {
				var number_of_given_arguments int = len(arguments)
				var diff int = number_of_expected_arguments - number_of_given_arguments

				if diff > 0 {
					return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Missing %d argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", diff, number_of_expected_arguments, number_of_given_arguments))
				} else if diff < 0 {
					return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Given %d extra argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", -diff, number_of_expected_arguments, number_of_given_arguments))
				}

				var wasm_inputs []C.wasmer_value_t = make([]C.wasmer_value_t, wasm_function_inputs_arity)

				for nth, value := range arguments {
					var wasm_input_type = wasm_function_input_signatures[nth]

					switch wasm_input_type {
					case C.WASM_I32:
						wasm_inputs[nth].tag = C.WASM_I32
						var pointer *int32 = (*int32)(unsafe.Pointer(&wasm_inputs[nth].value))

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

							if value.GetType() != Type_I32 {
								return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i32`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToI32()
						default:
							return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i32`, cannot cast given value to this type.", nth+1))
						}
					case C.WASM_I64:
						wasm_inputs[nth].tag = C.WASM_I64
						var pointer *int64 = (*int64)(unsafe.Pointer(&wasm_inputs[nth].value))

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

							if value.GetType() != Type_I64 {
								return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i64`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToI64()
						default:
							return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i64`, cannot cast given value to this type.", nth+1))
						}
					case C.WASM_F32:
						wasm_inputs[nth].tag = C.WASM_F32
						var pointer *float32 = (*float32)(unsafe.Pointer(&wasm_inputs[nth].value))

						switch value.(type) {
						case float32:
							*pointer = value.(float32)
						case Value:
							var value = value.(Value)

							if value.GetType() != Type_F32 {
								return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f32`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToF32()
						default:
							return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f32`, cannot cast given value to this type.", nth+1))
						}
					case C.WASM_F64:
						wasm_inputs[nth].tag = C.WASM_F64
						var pointer *float64 = (*float64)(unsafe.Pointer(&wasm_inputs[nth].value))

						switch value.(type) {
						case float32:
							*pointer = float64(value.(float32))
						case float64:
							*pointer = value.(float64)
						case Value:
							var value = value.(Value)

							if value.GetType() != Type_F64 {
								return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f64`, cannot cast given value to this type.", nth+1))
							}

							*pointer = value.ToF64()
						default:
							return I32(0), NewExportedFunctionError(exported_function_name, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f64`, cannot cast given value to this type.", nth+1))
						}
					default:
						panic(fmt.Sprintf("Invalid arguments type when calling the `%s` exported function.", exported_function_name))
					}
				}

				var wasm_outputs []C.wasmer_value_t = make([]C.wasmer_value_t, wasm_function_outputs_arity)

				var wasm_function_name = C.CString(exported_function_name)
				defer C.free(unsafe.Pointer(wasm_function_name))

				var wasm_inputs_c_pointer *C.wasmer_value_t = nil

				if wasm_function_inputs_arity > 0 {
					wasm_inputs_c_pointer = (*C.wasmer_value_t)(unsafe.Pointer(&wasm_inputs[0]))
				} else {
					wasm_inputs_c_pointer = (*C.wasmer_value_t)(unsafe.Pointer(&wasm_inputs))
				}

				var wasm_outputs_c_pointer *C.wasmer_value_t = nil

				if wasm_function_outputs_arity > 0 {
					wasm_outputs_c_pointer = (*C.wasmer_value_t)(unsafe.Pointer(&wasm_outputs[0]))
				} else {
					wasm_outputs_c_pointer = (*C.wasmer_value_t)(unsafe.Pointer(&wasm_outputs))
				}

				var call_result C.wasmer_result_t = C.wasmer_instance_call(
					instance,
					wasm_function_name,
					wasm_inputs_c_pointer,
					C.int(wasm_function_inputs_arity),
					wasm_outputs_c_pointer,
					C.int(wasm_function_outputs_arity),
				)

				if C.WASMER_OK != call_result {
					return I32(0), NewExportedFunctionError(exported_function_name, "Failed to call the `%s` exported function.")
				}

				if wasm_function_outputs_arity > 0 {
					var result = wasm_outputs[0]

					switch result.tag {
					case C.WASM_I32:
						pointer := (*int32)(unsafe.Pointer(&result.value))

						return I32(*pointer), nil
					case C.WASM_I64:
						pointer := (*int64)(unsafe.Pointer(&result.value))

						return I64(*pointer), nil
					case C.WASM_F32:
						pointer := (*float32)(unsafe.Pointer(&result.value))

						return F32(*pointer), nil
					case C.WASM_F64:
						pointer := (*float64)(unsafe.Pointer(&result.value))

						return F64(*pointer), nil
					default:
						panic("unreachable")
					}
				} else {
					return Void(), nil
				}
			}
		}
	}

	if has_memory == false {
		return empty_instance, NewInstanceError("No memory exported.")
	}

	return Instance{instance: instance, Exports: exports, Memory: memory}, nil
}

func (self *Instance) Close() {
	if self.instance != nil {
		C.wasmer_instance_destroy(self.instance)
	}
}
