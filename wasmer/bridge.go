package wasmer

// #cgo LDFLAGS: -L./ -lwasmer_runtime_c_api
// #include "./wasmer.h"
//
import "C"
import "unsafe"

type cBool C.bool
type cChar C.char
type cInt C.int
type cUchar C.uchar
type cUint C.uint
type cUint32T C.uint32_t
type cUint8T C.uint8_t
type cWasmerByteArray C.wasmer_byte_array
type cWasmerExportFuncT C.wasmer_export_func_t
type cWasmerExportT C.wasmer_export_t
type cWasmerExportsT C.wasmer_exports_t
type cWasmerImportExportKind C.wasmer_import_export_kind
type cWasmerImportExportValue C.wasmer_import_export_value
type cWasmerImportFuncT C.wasmer_import_func_t
type cWasmerImportT C.wasmer_import_t
type cWasmerInstanceT C.wasmer_instance_t
type cWasmerMemoryT C.wasmer_memory_t
type cWasmerResultT C.wasmer_result_t
type cWasmerValueT C.wasmer_value_t
type cWasmerValueTag C.wasmer_value_tag

const cWasmF32 = C.WASM_F32
const cWasmF64 = C.WASM_F64
const cWasmFunction = C.WASM_FUNCTION
const cWasmI32 = C.WASM_I32
const cWasmI64 = C.WASM_I64
const cWasmMemory = C.WASM_MEMORY
const cWasmerOk = C.WASMER_OK

func cWasmerLastErrorLength() cInt {
	return cInt(C.wasmer_last_error_length())
}

func cWasmerLastErrorMessage(buffer *cChar, length cInt) cInt {
	return cInt(C.wasmer_last_error_message((*C.char)(buffer), (C.int)(length)))
}

func cWasmerMemoryDataLength(memory *cWasmerMemoryT) cUint32T {
	return cUint32T(C.wasmer_memory_data_length((*C.wasmer_memory_t)(memory)))
}

func cWasmerMemoryData(memory *cWasmerMemoryT) *cUint8T {
	return (*cUint8T)(C.wasmer_memory_data((*C.wasmer_memory_t)(memory)))
}

func cWasmerValidate(wasmBytes *cUchar, wasmBytesLength cUint) cBool {
	return cBool(C.wasmer_validate((*C.uchar)(wasmBytes), (C.uint)(wasmBytesLength)))
}

func cWasmerInstantiate(instance **cWasmerInstanceT, wasmBytes *cUchar, wasmBytesLength cUint, imports *cWasmerImportT, importsLength cInt) cWasmerResultT {
	return cWasmerResultT(C.wasmer_instantiate((**C.wasmer_instance_t)(unsafe.Pointer(instance)), (*C.uchar)(wasmBytes), (C.uint)(wasmBytesLength), (*C.wasmer_import_t)(imports), (C.int)(importsLength)))
}

func cWasmerInstanceExports(instance *cWasmerInstanceT, exports **cWasmerExportsT) {
	C.wasmer_instance_exports((*C.wasmer_instance_t)(instance), (**C.wasmer_exports_t)(unsafe.Pointer(exports)))
}

func cWasmerExportsDestroy(exports *cWasmerExportsT) {
	C.wasmer_exports_destroy((*C.wasmer_exports_t)(exports))
}

func cWasmerExportsLen(exports *cWasmerExportsT) cInt {
	return cInt(C.wasmer_exports_len((*C.wasmer_exports_t)(exports)))
}

func cWasmerExportsGet(exports *cWasmerExportsT, index cInt) *cWasmerExportT {
	return (*cWasmerExportT)(C.wasmer_exports_get((*C.wasmer_exports_t)(exports), (C.int)(index)))
}

func cWasmerExportKind(export *cWasmerExportT) cWasmerImportExportKind {
	return (cWasmerImportExportKind)(C.wasmer_export_kind((*C.wasmer_export_t)(export)))
}

func cWasmerExportToMemory(export *cWasmerExportT, memory **cWasmerMemoryT) cWasmerResultT {
	return cWasmerResultT(C.wasmer_export_to_memory((*C.wasmer_export_t)(export), (**C.wasmer_memory_t)(unsafe.Pointer(memory))))
}

func cWasmerExportName(export *cWasmerExportT) cWasmerByteArray {
	return (cWasmerByteArray)(C.wasmer_export_name((*C.wasmer_export_t)(export)))
}

func cWasmerExportToFunc(export *cWasmerExportT) *cWasmerExportFuncT {
	return (*cWasmerExportFuncT)(C.wasmer_export_to_func((*C.wasmer_export_t)(export)))
}

func cWasmerExportFuncParamsArity(function *cWasmerExportFuncT, result *cUint32T) cWasmerResultT {
	return cWasmerResultT(C.wasmer_export_func_params_arity((*C.wasmer_export_func_t)(function), (*C.uint32_t)(result)))
}

func cWasmerExportFuncParams(function *cWasmerExportFuncT, parameters *cWasmerValueTag, parametersLength cUint32T) cWasmerResultT {
	return cWasmerResultT(C.wasmer_export_func_params((*C.wasmer_export_func_t)(function), (*C.wasmer_value_tag)(parameters), (C.uint32_t)(parametersLength)))
}

func cWasmerExportFuncResultsArity(function *cWasmerExportFuncT, result *cUint32T) cWasmerResultT {
	return cWasmerResultT(C.wasmer_export_func_returns_arity((*C.wasmer_export_func_t)(function), (*C.uint32_t)(result)))
}

func cWasmerInstanceCall(instance *cWasmerInstanceT, name *cChar, parameters *cWasmerValueT, parametersLength cUint32T, results *cWasmerValueT, resultsLength cUint32T) cWasmerResultT {
	return cWasmerResultT(C.wasmer_instance_call((*C.wasmer_instance_t)(instance), (*C.char)(name), (*C.wasmer_value_t)(parameters), (C.uint32_t)(parametersLength), (*C.wasmer_value_t)(results), (C.uint32_t)(resultsLength)))
}

func cWasmerInstanceDestroy(instance *cWasmerInstanceT) {
	C.wasmer_instance_destroy((*C.wasmer_instance_t)(instance))
}

func cWasmerImportFuncNew(function unsafe.Pointer, parametersSignature *cWasmerValueTag, parametersLength cUint, resultsSignature *cWasmerValueTag, resultsLength cUint) *cWasmerImportFuncT {
	return (*cWasmerImportFuncT)(C.wasmer_import_func_new((*[0]byte)(function), (*C.wasmer_value_tag)(parametersSignature), (C.int)(parametersLength), (*C.wasmer_value_tag)(resultsSignature), (C.int)(resultsLength)))
}

func cGoStringToWasmerByteArray(string string) cWasmerByteArray {
	var cString = cCString(string)

	var byteArray cWasmerByteArray
	byteArray.bytes = (*C.uchar)(unsafe.Pointer(cString))
	byteArray.bytes_len = (C.uint)(len(string))

	return byteArray
}

func cNewWasmerImportT(moduleName string, importName string, function *cWasmerImportFuncT) cWasmerImportT {
	var importedFunction C.wasmer_import_t
	importedFunction.module_name = (C.wasmer_byte_array)(cGoStringToWasmerByteArray(moduleName))
	importedFunction.import_name = (C.wasmer_byte_array)(cGoStringToWasmerByteArray(importName))
	importedFunction.tag = cWasmFunction

	var pointer = (**C.wasmer_import_func_t)(unsafe.Pointer(&importedFunction.value))
	*pointer = (*C.wasmer_import_func_t)(function)

	return (cWasmerImportT)(importedFunction)
}

func cGoString(string *cChar) string {
	return C.GoString((*C.char)(string))
}

func cGoStringN(string *cChar, length cInt) string {
	return C.GoStringN((*C.char)(string), (C.int)(length))
}

func cCString(string string) *cChar {
	return (*cChar)(C.CString(string))
}

func cFree(pointer unsafe.Pointer) {
	C.free(pointer)
}
