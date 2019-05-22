package wasmer

/*
#cgo LDFLAGS: -L./ -lwasmer_runtime_c_api
#include "./wasmer.h"
*/
import "C"
import "unsafe"

type C_wasmer_instance_t C.wasmer_instance_t
type C_wasmer_memory_t C.wasmer_memory_t
type C_wasmer_result_t C.wasmer_result_t
type C_wasmer_exports_t C.wasmer_exports_t
type C_wasmer_export_t C.wasmer_export_t
type C_wasmer_import_t C.wasmer_import_t
type C_wasmer_import_export_kind C.wasmer_import_export_kind
type C_wasmer_export_func_t C.wasmer_export_func_t
type C_wasmer_value_tag C.wasmer_value_tag
type C_wasmer_value_t C.wasmer_value_t
type C_wasmer_byte_array C.wasmer_byte_array
type C_uchar C.uchar
type C_char C.char
type C_int C.int
type C_bool C.bool
type C_uint C.uint
type C_uint8_t C.uint8_t
type C_uint32_t C.uint32_t

const C_WASMER_OK = C.WASMER_OK
const C_WASM_I32 = C.WASM_I32
const C_WASM_I64 = C.WASM_I64
const C_WASM_F32 = C.WASM_F32
const C_WASM_F64 = C.WASM_F64
const C_WASM_MEMORY = C.WASM_MEMORY
const C_WASM_FUNCTION = C.WASM_FUNCTION

func C_wasmer_last_error_length() C_int {
	return C_int(C.wasmer_last_error_length())
}

func C_wasmer_last_error_message(buffer *C_char, length C_int) C_int {
	return C_int(C.wasmer_last_error_message((*C.char)(buffer), (C.int)(length)))
}

func C_wasmer_memory_data_length(memory *C_wasmer_memory_t) C_uint32_t {
	return C_uint32_t(C.wasmer_memory_data_length((*C.wasmer_memory_t)(memory)))
}

func C_wasmer_memory_data(memory *C_wasmer_memory_t) *C_uint8_t {
	return (*C_uint8_t)(C.wasmer_memory_data((*C.wasmer_memory_t)(memory)))
}

func C_wasmer_validate(wasm_bytes *C_uchar, wasm_bytes_length C_uint) C_bool {
	return C_bool(C.wasmer_validate((*C.uchar)(wasm_bytes), (C.uint)(wasm_bytes_length)))
}

func C_wasmer_instantiate(instance **C_wasmer_instance_t, wasm_bytes *C_uchar, wasm_bytes_length C_uint, imports *C_wasmer_import_t, imports_length C_int) C_wasmer_result_t {
	return C_wasmer_result_t(C.wasmer_instantiate((**C.wasmer_instance_t)(unsafe.Pointer(instance)), (*C.uchar)(wasm_bytes), (C.uint)(wasm_bytes_length), (*C.wasmer_import_t)(imports), (C.int)(imports_length)))
}

func C_wasmer_instance_exports(instance *C_wasmer_instance_t, exports **C_wasmer_exports_t) {
	C.wasmer_instance_exports((*C.wasmer_instance_t)(instance), (**C.wasmer_exports_t)(unsafe.Pointer(exports)))
}

func C_wasmer_exports_destroy(exports *C_wasmer_exports_t) {
	C.wasmer_exports_destroy((*C.wasmer_exports_t)(exports))
}

func C_wasmer_exports_len(exports *C_wasmer_exports_t) C_int {
	return C_int(C.wasmer_exports_len((*C.wasmer_exports_t)(exports)))
}

func C_wasmer_exports_get(exports *C_wasmer_exports_t, index C_int) *C_wasmer_export_t {
	return (*C_wasmer_export_t)(C.wasmer_exports_get((*C.wasmer_exports_t)(exports), (C.int)(index)))
}

func C_wasmer_export_kind(export *C_wasmer_export_t) C_wasmer_import_export_kind {
	return (C_wasmer_import_export_kind)(C.wasmer_export_kind((*C.wasmer_export_t)(export)))
}

func C_wasmer_export_to_memory(export *C_wasmer_export_t, memory **C_wasmer_memory_t) C_wasmer_result_t {
	return C_wasmer_result_t(C.wasmer_export_to_memory((*C.wasmer_export_t)(export), (**C.wasmer_memory_t)(unsafe.Pointer(memory))))
}

func C_wasmer_export_name(export *C_wasmer_export_t) C_wasmer_byte_array {
	return (C_wasmer_byte_array)(C.wasmer_export_name((*C.wasmer_export_t)(export)))
}

func C_wasmer_export_to_func(export *C_wasmer_export_t) *C_wasmer_export_func_t {
	return (*C_wasmer_export_func_t)(C.wasmer_export_to_func((*C.wasmer_export_t)(export)))
}

func C_wasmer_export_func_params_arity(function *C_wasmer_export_func_t, result *C_uint32_t) C_wasmer_result_t {
	return C_wasmer_result_t(C.wasmer_export_func_params_arity((*C.wasmer_export_func_t)(function), (*C.uint32_t)(result)))
}

func C_wasmer_export_func_params(function *C_wasmer_export_func_t, parameters *C_wasmer_value_tag, parametersLength C_uint32_t) C_wasmer_result_t {
	return C_wasmer_result_t(C.wasmer_export_func_params((*C.wasmer_export_func_t)(function), (*C.wasmer_value_tag)(parameters), (C.uint32_t)(parametersLength)))
}

func C_wasmer_export_func_returns_arity(function *C_wasmer_export_func_t, result *C_uint32_t) C_wasmer_result_t {
	return C_wasmer_result_t(C.wasmer_export_func_returns_arity((*C.wasmer_export_func_t)(function), (*C.uint32_t)(result)))
}

func C_wasmer_instance_call(instance *C_wasmer_instance_t, name *C_char, parameters *C_wasmer_value_t, parametersLength C_uint32_t, results *C_wasmer_value_t, resultsLength C_uint32_t) C_wasmer_result_t {
	return C_wasmer_result_t(C.wasmer_instance_call((*C.wasmer_instance_t)(instance), (*C.char)(name), (*C.wasmer_value_t)(parameters), (C.uint32_t)(parametersLength), (*C.wasmer_value_t)(results), (C.uint32_t)(resultsLength)))
}

func C_wasmer_instance_destroy(instance *C_wasmer_instance_t) {
	C.wasmer_instance_destroy((*C.wasmer_instance_t)(instance))
}

func C_GoString(string *C_char) string {
	return C.GoString((*C.char)(string))
}

func C_GoStringN(string *C_char, length C_int) string {
	return C.GoStringN((*C.char)(string), (C.int)(length))
}

func C_CString(string string) *C_char {
	return (*C_char)(C.CString(string))
}

func C_free(pointer unsafe.Pointer) {
	C.free(pointer)
}
