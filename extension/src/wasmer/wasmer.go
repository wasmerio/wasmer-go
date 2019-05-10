package wasmer

/*
#cgo LDFLAGS: -L../../ -lwasmer_runtime_c_api
#include "../../wasmer.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"io/ioutil"
	"unsafe"
)

func ReadBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func Validate(bytes []byte) bool {
	return true == C.wasmer_validate((*C.uchar) (unsafe.Pointer(&bytes[0])), C.uint(len(bytes)))
}

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

func (self Instance) Call(function_name string, arguments []Value) (Value, error) {
	var wasm_arguments []C.wasmer_value_t = make([]C.wasmer_value_t, len(arguments))

	for index, value := range arguments {
		switch value.GetType() {
		case Type_I32:
			wasm_arguments[index].tag = C.WASM_I32
			wasm_arguments[index].value[C.WASM_I32] = byte(value.ToI32())
		case Type_I64:
			wasm_arguments[index].tag = C.WASM_I64
			wasm_arguments[index].value[C.WASM_I64] = byte(value.ToI64())
		case Type_F32:
			wasm_arguments[index].tag = C.WASM_F32
			wasm_arguments[index].value[C.WASM_F32] = byte(value.ToF32())
		case Type_F64:
			wasm_arguments[index].tag = C.WASM_F64
			wasm_arguments[index].value[C.WASM_F64] = byte(value.ToF64())
		default:
			panic("Unreachable")
		}
	}
	
	var wasm_result C.wasmer_value_t
	var wasm_results []C.wasmer_value_t = []C.wasmer_value_t{wasm_result}

	var wasm_function_name = C.CString(function_name)
	defer C.free(unsafe.Pointer(wasm_function_name))

	var wasm_argument_c_pointer *C.wasmer_value_t = nil

	if len(arguments) > 0 {
		wasm_argument_c_pointer = (*C.wasmer_value_t) (unsafe.Pointer(&wasm_arguments[0]))
	}

	var call_result C.wasmer_result_t = C.wasmer_instance_call(
		self.instance,
		wasm_function_name,
		wasm_argument_c_pointer,
		C.int(len(wasm_arguments)),
		(*C.wasmer_value_t) (unsafe.Pointer(&wasm_results[0])),
		C.int(len(wasm_results)),
	)

	if (C.WASMER_OK != call_result) {
		return ValueI32(0), NewExportedFunctionError(function_name, "Failed to call the `%s` exported function.")
	}

	var result = wasm_results[0];

	switch result.tag {
	case C.WASM_I32:
		return ValueI32(int32(result.value[C.WASM_I32])), nil
	case C.WASM_I64:
		return ValueI64(int64(result.value[C.WASM_I64])), nil
	case C.WASM_F32:
		return ValueF32(float32(result.value[C.WASM_F32])), nil
	case C.WASM_F64:
		return ValueF64(float64(result.value[C.WASM_F64])), nil
	default:
		panic("unreachable")
	}
}

func GetLastError() (string, error) {
	var error_length C.int = C.wasmer_last_error_length()
	var error_message *C.char = nil

	if -1 == C.wasmer_last_error_message(error_message, error_length) {
		return "", errors.New("Cannot read last error.")
	}

	return C.GoStringN(error_message, error_length), nil
}
