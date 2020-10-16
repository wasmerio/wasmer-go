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
import "unsafe"

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

func fromGoValue(value interface{}, kind ValueKind) (C.wasm_val_t, error) {
	output := C.wasm_val_t{}

	switch kind {
	case I32:
		output.kind = kind.inner()

		var of = (*int32)(unsafe.Pointer(&output.of))

		switch value.(type) {
		case int8:
			*of = int32(value.(int8))
		case uint8:
			*of = int32(value.(uint8))
		case int16:
			*of = int32(value.(int16))
		case uint16:
			*of = int32(value.(uint16))
		case int32:
			*of = value.(int32)
		case int:
			*of = int32(value.(int))
		case uint:
			*of = int32(value.(uint))
		default:
			return output, newErrorWith("i32")
		}
	case I64:
		output.kind = kind.inner()

		var of = (*int64)(unsafe.Pointer(&output.of))

		switch value.(type) {
		case int8:
			*of = int64(value.(int8))
		case uint8:
			*of = int64(value.(uint8))
		case int16:
			*of = int64(value.(int16))
		case uint16:
			*of = int64(value.(uint16))
		case int32:
			*of = int64(value.(int32))
		case uint32:
			*of = int64(value.(int64))
		case int64:
			*of = value.(int64)
		case int:
			*of = int64(value.(int))
		case uint:
			*of = int64(value.(uint))
		default:
			return output, newErrorWith("i64")
		}
	case F32:
		output.kind = kind.inner()

		var of = (*float32)(unsafe.Pointer(&output.of))

		switch value.(type) {
		case float32:
			*of = value.(float32)
		default:
			return output, newErrorWith("f32")
		}
	case F64:
		output.kind = kind.inner()

		var of = (*float64)(unsafe.Pointer(&output.of))

		switch value.(type) {
		case float32:
			*of = float64(value.(float32))
		case float64:
			*of = value.(float64)
		default:
			return output, newErrorWith("f64")
		}
	default:
		panic("To do, `fromGoValue`!")
	}

	return output, nil
}
