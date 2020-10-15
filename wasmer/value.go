package wasmer

// #include <wasmer_wasm.h>
//
// int32_t to_int32(wasm_val_t *value) {
//     return value->of.i32;
// }
//
// int64_t to_int64(wasm_val_t *value) {
//     return value->of.i64;
// }
//
// float32_t to_float32(wasm_val_t *value) {
//     return value->of.f32;
// }
//
// float64_t to_float64(wasm_val_t *value) {
//     return value->of.f64;
// }
import "C"

func toGoValue(pointer *C.wasm_val_t) interface{} {
	switch ValueKind(pointer.kind) {
	case I32:
		return int32(C.to_int32(pointer))
	case I64:
		return int64(C.to_int64(pointer))
	case F32:
		return float32(C.to_float32(pointer))
	case F64:
		return float64(C.to_float64(pointer))
	default:
		panic("to do `newValue`")
	}
}
