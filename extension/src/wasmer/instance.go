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
	return &InstanceError { message }
}

func (self *InstanceError) Error() string {
	return self.message
}

type ExportedFunctionError struct {
	function_name string
	message string
}

func NewExportedFunctionError(function_name string, message string) error {
	return &ExportedFunctionError { function_name, message }
}

func (self *ExportedFunctionError) Error() string {
	return fmt.Sprintf(self.message, self.function_name)
}

type Instance struct {
	instance *C.wasmer_instance_t
}

func NewInstance(bytes []byte) (Instance, error) {
	var imports []C.wasmer_import_t = []C.wasmer_import_t{}
	var instance *C.wasmer_instance_t = nil

	var compile_result C.wasmer_result_t = C.wasmer_instantiate(
		&instance,
		(*C.uchar) (unsafe.Pointer(&bytes[0])),
		C.uint(len(bytes)),
		(*C.wasmer_import_t) (unsafe.Pointer(&imports)),
		C.int(0),
	)

	if C.WASMER_OK != compile_result {
		return Instance { instance: nil }, NewInstanceError("Failed to compile the module.")
	}

	return Instance { instance }, nil
}

func (self Instance) Call(function_name string, arguments ...Value) (Value, error) {
	var wasm_exports *C.wasmer_exports_t = nil
	C.wasmer_instance_exports(self.instance, &wasm_exports)

	defer C.wasmer_exports_destroy(wasm_exports)

	var number_of_exports int = int(C.wasmer_exports_len(wasm_exports))

	if number_of_exports == 0 {
		return I32(0), NewExportedFunctionError(function_name, "The instance has no exports, cannot call the function `%s`.")
	}

	var wasm_function *C.wasmer_export_func_t = nil

	for nth := 0; nth < number_of_exports; nth++ {
		var wasm_export *C.wasmer_export_t = C.wasmer_exports_get(wasm_exports, C.int(nth))
		var wasm_export_kind C.wasmer_import_export_kind = C.wasmer_export_kind(wasm_export)

		if wasm_export_kind != C.WASM_FUNCTION {
			continue
		}

		var wasm_export_name C.wasmer_byte_array = C.wasmer_export_name(wasm_export)

		if int(wasm_export_name.bytes_len) != len(function_name) {
			continue
		}

		if function_name == C.GoStringN((*C.char) (unsafe.Pointer(wasm_export_name.bytes)), (C.int) (wasm_export_name.bytes_len)) {
			wasm_function = C.wasmer_export_to_func(wasm_export)

			break
		}
	}

	if wasm_function == nil {
		return I32(0), NewExportedFunctionError(function_name, "The instance has no exported function named `%s`.")
	}

	var wasm_function_inputs_arity C.uint

	if C.wasmer_export_func_params_arity(wasm_function, &wasm_function_inputs_arity) != C.WASMER_OK {
		return I32(0), NewExportedFunctionError(function_name, "Failed to read the input arity of the `%s` exported function.")
	}

	var wasm_function_outputs_arity C.uint

	if C.wasmer_export_func_returns_arity(wasm_function, &wasm_function_outputs_arity) != C.WASMER_OK {
		return I32(0), NewExportedFunctionError(function_name, "Failed to read the output arity of the `%s` exported function.")
	}

	var number_of_expected_arguments int = int(wasm_function_inputs_arity)
	var number_of_given_arguments int = len(arguments)
	var diff int = number_of_expected_arguments - number_of_given_arguments

	if diff > 0 {
		return I32(0), NewExportedFunctionError(function_name, fmt.Sprintf("Missing %d argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", diff, number_of_expected_arguments, number_of_given_arguments))
	} else if diff < 0 {
		return I32(0), NewExportedFunctionError(function_name, fmt.Sprintf("Given %d extra argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", -diff, number_of_expected_arguments, number_of_given_arguments))
	}

	var wasm_inputs []C.wasmer_value_t = make([]C.wasmer_value_t, wasm_function_inputs_arity)

	for index, value := range arguments {
		switch value.GetType() {
		case Type_I32:
			wasm_inputs[index].tag = C.WASM_I32
			pointer := (*int32) (unsafe.Pointer(&wasm_inputs[index].value))
			*pointer = value.ToI32()
		case Type_I64:
			wasm_inputs[index].tag = C.WASM_I64
			pointer := (*int64) (unsafe.Pointer(&wasm_inputs[index].value))
			*pointer = value.ToI64()
		case Type_F32:
			wasm_inputs[index].tag = C.WASM_F32
			pointer := (*float32) (unsafe.Pointer(&wasm_inputs[index].value))
			*pointer = value.ToF32()
		case Type_F64:
			wasm_inputs[index].tag = C.WASM_F64
			pointer := (*float64) (unsafe.Pointer(&wasm_inputs[index].value))
			*pointer = value.ToF64()
		default:
			panic(fmt.Sprintf("Invalid arguments type when calling the `%s` exported function.", function_name))
		}
	}
	
	var wasm_outputs []C.wasmer_value_t = make([]C.wasmer_value_t, wasm_function_outputs_arity)

	var wasm_function_name = C.CString(function_name)
	defer C.free(unsafe.Pointer(wasm_function_name))

	var wasm_input_c_pointer *C.wasmer_value_t = nil

	if wasm_function_inputs_arity > 0 {
		wasm_input_c_pointer = (*C.wasmer_value_t) (unsafe.Pointer(&wasm_inputs[0]))
	} else {
		wasm_input_c_pointer = (*C.wasmer_value_t) (unsafe.Pointer(&wasm_inputs))
	}

	var wasm_output_c_pointer *C.wasmer_value_t = nil

	if wasm_function_outputs_arity > 0 {
		wasm_output_c_pointer = (*C.wasmer_value_t) (unsafe.Pointer(&wasm_outputs[0]))
	} else {
		wasm_output_c_pointer = (*C.wasmer_value_t) (unsafe.Pointer(&wasm_outputs))
	}

	var call_result C.wasmer_result_t = C.wasmer_instance_call(
		self.instance,
		wasm_function_name,
		wasm_input_c_pointer,
		C.int(wasm_function_inputs_arity),
		wasm_output_c_pointer,
		C.int(wasm_function_outputs_arity),
	)

	if (C.WASMER_OK != call_result) {
		return I32(0), NewExportedFunctionError(function_name, "Failed to call the `%s` exported function.")
	}

	if wasm_function_outputs_arity > 0 {
		var result = wasm_outputs[0];

		switch result.tag {
		case C.WASM_I32:
			pointer := (*int32) (unsafe.Pointer(&result.value))

			return I32(*pointer), nil
		case C.WASM_I64:
			pointer := (*int64) (unsafe.Pointer(&result.value))

			return I64(*pointer), nil
		case C.WASM_F32:
			pointer := (*float32) (unsafe.Pointer(&result.value))

			return F32(*pointer), nil
		case C.WASM_F64:
			pointer := (*float64) (unsafe.Pointer(&result.value))

			return F64(*pointer), nil
		default:
			panic("unreachable")
		}
	} else {
		return Void(), nil
	}
}
