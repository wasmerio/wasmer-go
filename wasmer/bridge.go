package wasmer

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmer_runtime_c_api
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
type cWasmerExportDescriptorT C.wasmer_export_descriptor_t
type cWasmerExportDescriptorsT C.wasmer_export_descriptors_t
type cWasmerExportFuncT C.wasmer_export_func_t
type cWasmerExportT C.wasmer_export_t
type cWasmerExportsT C.wasmer_exports_t
type cWasmerImportDescriptorT C.wasmer_import_descriptor_t
type cWasmerImportDescriptorsT C.wasmer_import_descriptors_t
type cWasmerImportExportKind C.wasmer_import_export_kind
type cWasmerImportExportValue C.wasmer_import_export_value
type cWasmerImportFuncT C.wasmer_import_func_t
type cWasmerImportT C.wasmer_import_t
type cWasmerInstanceContextT C.wasmer_instance_context_t
type cWasmerInstanceT C.wasmer_instance_t
type cWasmerMemoryT C.wasmer_memory_t
type cWasmerModuleT C.wasmer_module_t
type cWasmerResultT C.wasmer_result_t
type cWasmerSerializedModuleT C.wasmer_serialized_module_t
type cWasmerValueT C.wasmer_value_t
type cWasmerValueTag C.wasmer_value_tag

const cWasmF32 = C.WASM_F32
const cWasmF64 = C.WASM_F64
const cWasmFunction = C.WASM_FUNCTION
const cWasmGlobal = C.WASM_GLOBAL
const cWasmI32 = C.WASM_I32
const cWasmI64 = C.WASM_I64
const cWasmMemory = C.WASM_MEMORY
const cWasmTable = C.WASM_TABLE
const cWasmerOk = C.WASMER_OK

func cNewWasmerImportT(moduleName string, importName string, function *cWasmerImportFuncT) cWasmerImportT {
	var importedFunction C.wasmer_import_t
	importedFunction.module_name = (C.wasmer_byte_array)(cGoStringToWasmerByteArray(moduleName))
	importedFunction.import_name = (C.wasmer_byte_array)(cGoStringToWasmerByteArray(importName))
	importedFunction.tag = cWasmFunction

	var pointer = (**C.wasmer_import_func_t)(unsafe.Pointer(&importedFunction.value))
	*pointer = (*C.wasmer_import_func_t)(function)

	return (cWasmerImportT)(importedFunction)
}

func cWasmerCompile(module **cWasmerModuleT, wasmBytes *cUchar, wasmBytesLength cUint) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_compile(
		(**C.wasmer_module_t)(unsafe.Pointer(module)),
		(*C.uchar)(wasmBytes),
		(C.uint)(wasmBytesLength),
	))
}


// TODO: is this an auto-gen file? how should I create this?
func cWasmerCompileWithLimit(module **cWasmerModuleT, wasmBytes *cUchar, wasmBytesLength cUint, gasLimit uint64) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_compile_with_limit(
		(**C.wasmer_module_t)(unsafe.Pointer(module)),
		(*C.uchar)(wasmBytes),
		(C.uint)(wasmBytesLength),
		(C.uint64_t)(gasLimit),
	))
}

func cWasmerInstanceGetPointsUsed(instance *cWasmerInstanceT) uint64 {
	return uint64(C.wasmer_instance_get_points_used(
		(*C.wasmer_instance_t)(instance),
	))
}

func cWasmerInstanceSetPointsUsed(instance *cWasmerInstanceT, points uint64) {
	C.wasmer_instance_set_points_used(
		(*C.wasmer_instance_t)(instance),
		(C.uint64_t)(points),
	)
}
// End TODO: autogen?

func cWasmerExportDescriptorKind(exportDescriptor *cWasmerExportDescriptorT) cWasmerImportExportKind {
	return (cWasmerImportExportKind)(C.wasmer_export_descriptor_kind(
		(*C.wasmer_export_descriptor_t)(exportDescriptor),
	))
}

func cWasmerExportDescriptorName(exportDescriptor *cWasmerExportDescriptorT) cWasmerByteArray {
	return (cWasmerByteArray)(C.wasmer_export_descriptor_name(
		(*C.wasmer_export_descriptor_t)(exportDescriptor),
	))
}

func cWasmerExportDescriptors(module *cWasmerModuleT, exportDescriptors **cWasmerExportDescriptorsT) {
	C.wasmer_export_descriptors(
		(*C.wasmer_module_t)(module),
		(**C.wasmer_export_descriptors_t)(unsafe.Pointer(exportDescriptors)),
	)
}

func cWasmerExportDescriptorsDestroy(exportDescriptors *cWasmerExportDescriptorsT) {
	C.wasmer_export_descriptors_destroy(
		(*C.wasmer_export_descriptors_t)(exportDescriptors),
	)
}

func cWasmerExportDescriptorsGet(exportDescriptors *cWasmerExportDescriptorsT, index cInt) *cWasmerExportDescriptorT {
	return (*cWasmerExportDescriptorT)(C.wasmer_export_descriptors_get(
		(*C.wasmer_export_descriptors_t)(exportDescriptors),
		(C.int)(index),
	))
}

func cWasmerExportDescriptorsLen(exportDescriptors *cWasmerExportDescriptorsT) cInt {
	return (cInt)(C.wasmer_export_descriptors_len(
		(*C.wasmer_export_descriptors_t)(exportDescriptors),
	))
}

func cWasmerExportFuncParams(function *cWasmerExportFuncT, parameters *cWasmerValueTag, parametersLength cUint32T) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_export_func_params(
		(*C.wasmer_export_func_t)(function),
		(*C.wasmer_value_tag)(parameters),
		(C.uint32_t)(parametersLength),
	))
}

func cWasmerExportFuncParamsArity(function *cWasmerExportFuncT, result *cUint32T) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_export_func_params_arity(
		(*C.wasmer_export_func_t)(function),
		(*C.uint32_t)(result),
	))
}

func cWasmerExportFuncResultsArity(function *cWasmerExportFuncT, result *cUint32T) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_export_func_returns_arity(
		(*C.wasmer_export_func_t)(function),
		(*C.uint32_t)(result),
	))
}

func cWasmerExportKind(export *cWasmerExportT) cWasmerImportExportKind {
	return (cWasmerImportExportKind)(C.wasmer_export_kind(
		(*C.wasmer_export_t)(export),
	))
}

func cWasmerExportName(export *cWasmerExportT) cWasmerByteArray {
	return (cWasmerByteArray)(C.wasmer_export_name(
		(*C.wasmer_export_t)(export),
	))
}

func cWasmerExportToFunc(export *cWasmerExportT) *cWasmerExportFuncT {
	return (*cWasmerExportFuncT)(C.wasmer_export_to_func(
		(*C.wasmer_export_t)(export),
	))
}

func cWasmerExportToMemory(export *cWasmerExportT, memory **cWasmerMemoryT) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_export_to_memory(
		(*C.wasmer_export_t)(export),
		(**C.wasmer_memory_t)(unsafe.Pointer(memory)),
	))
}

func cWasmerExportsDestroy(exports *cWasmerExportsT) {
	C.wasmer_exports_destroy(
		(*C.wasmer_exports_t)(exports),
	)
}

func cWasmerExportsGet(exports *cWasmerExportsT, index cInt) *cWasmerExportT {
	return (*cWasmerExportT)(C.wasmer_exports_get(
		(*C.wasmer_exports_t)(exports),
		(C.int)(index),
	))
}

func cWasmerExportsLen(exports *cWasmerExportsT) cInt {
	return (cInt)(C.wasmer_exports_len(
		(*C.wasmer_exports_t)(exports),
	))
}

func cWasmerImportDescriptorKind(importDescriptor *cWasmerImportDescriptorT) cWasmerImportExportKind {
	return (cWasmerImportExportKind)(C.wasmer_import_descriptor_kind(
		(*C.wasmer_import_descriptor_t)(importDescriptor),
	))
}

func cWasmerImportDescriptorModuleName(importDescriptor *cWasmerImportDescriptorT) cWasmerByteArray {
	return (cWasmerByteArray)(C.wasmer_import_descriptor_module_name(
		(*C.wasmer_import_descriptor_t)(importDescriptor),
	))
}

func cWasmerImportDescriptorName(importDescriptor *cWasmerImportDescriptorT) cWasmerByteArray {
	return (cWasmerByteArray)(C.wasmer_import_descriptor_name(
		(*C.wasmer_import_descriptor_t)(importDescriptor),
	))
}

func cWasmerImportDescriptors(module *cWasmerModuleT, importDescriptors **cWasmerImportDescriptorsT) {
	C.wasmer_import_descriptors(
		(*C.wasmer_module_t)(module),
		(**C.wasmer_import_descriptors_t)(unsafe.Pointer(importDescriptors)),
	)
}

func cWasmerImportDescriptorsDestroy(importDescriptors *cWasmerImportDescriptorsT) {
	C.wasmer_import_descriptors_destroy(
		(*C.wasmer_import_descriptors_t)(importDescriptors),
	)
}

func cWasmerImportDescriptorsGet(importDescriptors *cWasmerImportDescriptorsT, index cInt) *cWasmerImportDescriptorT {
	return (*cWasmerImportDescriptorT)(C.wasmer_import_descriptors_get(
		(*C.wasmer_import_descriptors_t)(importDescriptors),
		(C.uint)(index),
	))
}

func cWasmerImportDescriptorsLen(importDescriptors *cWasmerImportDescriptorsT) cInt {
	return (cInt)(C.wasmer_import_descriptors_len(
		(*C.wasmer_import_descriptors_t)(importDescriptors),
	))
}

func cWasmerImportFuncDestroy(function *cWasmerImportFuncT) {
	C.wasmer_import_func_destroy(
		(*C.wasmer_import_func_t)(function),
	)
}

func cWasmerImportFuncNew(
	function unsafe.Pointer,
	parametersSignature *cWasmerValueTag,
	parametersLength cUint,
	resultsSignature *cWasmerValueTag,
	resultsLength cUint,
) *cWasmerImportFuncT {
	return (*cWasmerImportFuncT)(C.wasmer_import_func_new(
		(*[0]byte)(function),
		(*C.wasmer_value_tag)(parametersSignature),
		(C.uint)(parametersLength),
		(*C.wasmer_value_tag)(resultsSignature),
		(C.uint)(resultsLength),
	))
}

func cWasmerInstanceCall(
	instance *cWasmerInstanceT,
	name *cChar,
	parameters *cWasmerValueT,
	parametersLength cUint32T,
	results *cWasmerValueT,
	resultsLength cUint32T,
) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_instance_call(
		(*C.wasmer_instance_t)(instance),
		(*C.char)(name),
		(*C.wasmer_value_t)(parameters),
		(C.uint32_t)(parametersLength),
		(*C.wasmer_value_t)(results),
		(C.uint32_t)(resultsLength),
	))
}

func cWasmerInstanceContextDataGet(instanceContext *cWasmerInstanceContextT) unsafe.Pointer {
	return unsafe.Pointer(C.wasmer_instance_context_data_get(
		(*C.wasmer_instance_t)(instanceContext),
	))
}

func cWasmerInstanceContextDataSet(instance *cWasmerInstanceT, dataPointer unsafe.Pointer) {
	C.wasmer_instance_context_data_set(
		(*C.wasmer_instance_t)(instance),
		dataPointer,
	)
}

func cWasmerInstanceContextMemory(instanceContext *cWasmerInstanceContextT) *cWasmerMemoryT {
	return (*cWasmerMemoryT)(C.wasmer_instance_context_memory(
		(*C.wasmer_instance_context_t)(instanceContext),
		0,
	))
}

func cWasmerInstanceDestroy(instance *cWasmerInstanceT) {
	C.wasmer_instance_destroy(
		(*C.wasmer_instance_t)(instance),
	)
}

func cWasmerInstanceExports(instance *cWasmerInstanceT, exports **cWasmerExportsT) {
	C.wasmer_instance_exports(
		(*C.wasmer_instance_t)(instance),
		(**C.wasmer_exports_t)(unsafe.Pointer(exports)),
	)
}

func cWasmerInstantiate(
	instance **cWasmerInstanceT,
	wasmBytes *cUchar,
	wasmBytesLength cUint,
	imports *cWasmerImportT,
	importsLength cInt,
) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_instantiate(
		(**C.wasmer_instance_t)(unsafe.Pointer(instance)),
		(*C.uchar)(wasmBytes),
		(C.uint)(wasmBytesLength),
		(*C.wasmer_import_t)(imports),
		(C.int)(importsLength),
	))
}

func cWasmerLastErrorLength() cInt {
	return (cInt)(C.wasmer_last_error_length())
}

func cWasmerLastErrorMessage(buffer *cChar, length cInt) cInt {
	return (cInt)(C.wasmer_last_error_message(
		(*C.char)(buffer),
		(C.int)(length),
	))
}

func cWasmerMemoryData(memory *cWasmerMemoryT) *cUint8T {
	return (*cUint8T)(C.wasmer_memory_data(
		(*C.wasmer_memory_t)(memory),
	))
}

func cWasmerMemoryDataLength(memory *cWasmerMemoryT) cUint32T {
	return (cUint32T)(C.wasmer_memory_data_length(
		(*C.wasmer_memory_t)(memory),
	))
}

func cWasmerMemoryGrow(memory *cWasmerMemoryT, numberOfPages cUint32T) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_memory_grow(
		(*C.wasmer_memory_t)(memory),
		(C.uint32_t)(numberOfPages),
	))
}

func cWasmerModuleDeserialize(module **cWasmerModuleT, serializedModule *cWasmerSerializedModuleT) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_module_deserialize(
		(**C.wasmer_module_t)(unsafe.Pointer(module)),
		(*C.wasmer_serialized_module_t)(serializedModule),
	))
}

func cWasmerModuleDestroy(module *cWasmerModuleT) {
	C.wasmer_module_destroy((*C.wasmer_module_t)(module))
}

func cWasmerModuleInstantiate(
	module *cWasmerModuleT,
	instance **cWasmerInstanceT,
	imports *cWasmerImportT,
	importsLength cInt,
) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_module_instantiate(
		(*C.wasmer_module_t)(module),
		(**C.wasmer_instance_t)(unsafe.Pointer(instance)),
		(*C.wasmer_import_t)(imports),
		(C.int)(importsLength),
	))
}

func cWasmerModuleSerialize(serializedModule **cWasmerSerializedModuleT, module *cWasmerModuleT) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_module_serialize(
		(**C.wasmer_serialized_module_t)(unsafe.Pointer(serializedModule)),
		(*C.wasmer_module_t)(module),
	))
}

func cWasmerSerializedModuleBytes(serializedModule *cWasmerSerializedModuleT) []byte {
	var byteArray = C.wasmer_serialized_module_bytes(
		(*C.wasmer_serialized_module_t)(serializedModule),
	)

	return C.GoBytes(unsafe.Pointer(byteArray.bytes), (C.int)(byteArray.bytes_len))
}

func cWasmerSerializedModuleDestroy(serializedModule *cWasmerSerializedModuleT) {
	C.wasmer_serialized_module_destroy(
		(*C.wasmer_serialized_module_t)(serializedModule),
	)
}

func cWasmerSerializedModuleFromBytes(
	serializedModule **cWasmerSerializedModuleT,
	serializedModuleBytes *cUint8T,
	serializedModuleBytesLength cInt,
) cWasmerResultT {
	return (cWasmerResultT)(C.wasmer_serialized_module_from_bytes(
		(**C.wasmer_serialized_module_t)(unsafe.Pointer(serializedModule)),
		(*C.uint8_t)(serializedModuleBytes),
		(C.uint)(serializedModuleBytesLength),
	))
}

func cWasmerValidate(wasmBytes *cUchar, wasmBytesLength cUint) cBool {
	return (cBool)(C.wasmer_validate(
		(*C.uchar)(wasmBytes),
		(C.uint)(wasmBytesLength),
	))
}

func cCString(string string) *cChar {
	return (*cChar)(C.CString(string))
}

func cFree(pointer unsafe.Pointer) {
	C.free(pointer)
}

func cGoString(string *cChar) string {
	return C.GoString((*C.char)(string))
}

func cGoStringN(string *cChar, length cInt) string {
	return C.GoStringN((*C.char)(string), (C.int)(length))
}

func cGoStringToWasmerByteArray(string string) cWasmerByteArray {
	var cString = cCString(string)

	var byteArray cWasmerByteArray
	byteArray.bytes = (*C.uchar)(unsafe.Pointer(cString))
	byteArray.bytes_len = (C.uint)(len(string))

	return byteArray
}
