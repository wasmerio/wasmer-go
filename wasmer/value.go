package wasmer

import (
	"fmt"
	"math"
)

// Represents the `Value` type.
type ValueType int

const (
	// Represents the WebAssembly `i32` type.
	Type_I32 ValueType = iota

	// Represents the WebAssembly `i64` type.
	Type_I64

	// Represents the WebAssembly `f32` type.
	Type_F32

	// Represents the WebAssembly `f64` type.
	Type_F64

	// WebAssembly doesn't have “void” type, but it is introduced
	// here to represent the returned value of a WebAssembly exported
	// function that returns nothing.
	Type_Void
)

// Represents a WebAssembly value of a particular type.
type Value struct {
	// The WebAssembly value (as bits).
	value uint64

	// The WebAssembly value type.
	ty ValueType
}

// Constructs a WebAssembly value of type `i32`.
func I32(value int32) Value {
	return Value{
		value: uint64(value),
		ty:    Type_I32,
	}
}

// Constructs a WebAssembly value of type `i64`.
func I64(value int64) Value {
	return Value{
		value: uint64(value),
		ty:    Type_I64,
	}
}

// Constructs a WebAssembly value of type `f32`.
func F32(value float32) Value {
	return Value{
		value: uint64(math.Float32bits(value)),
		ty:    Type_F32,
	}
}

// Constructs a WebAssembly value of type `f64`.
func F64(value float64) Value {
	return Value{
		value: math.Float64bits(value),
		ty:    Type_F64,
	}
}

// Constructs an empty WebAssembly value.
func void() Value {
	return Value{
		value: 0,
		ty:    Type_Void,
	}
}

// Gets the type of the WebAssembly value.
func (self Value) GetType() ValueType {
	return self.ty
}

// Reads the WebAssembly value bits as an `int32`. The WebAssembly
// value type is ignored.
func (self Value) ToI32() int32 {
	return int32(self.value)
}

// Reads the WebAssembly value bits as an `int64`. The WebAssembly
// value type is ignored.
func (self Value) ToI64() int64 {
	return int64(self.value)
}

// Reads the WebAssembly value bits as a `float32`. The WebAssembly
// value type is ignored.
func (self Value) ToF32() float32 {
	return math.Float32frombits(uint32(self.value))
}

// Reads the WebAssembly value bits as a `float64`. The WebAssembly
// value type is ignored.
func (self Value) ToF64() float64 {
	return math.Float64frombits(self.value)
}

// Reads the WebAssembly value bits as a `nil`. The WebAssembly
// value type is ignored.
func (self Value) ToVoid() interface{} {
	return nil
}

// Formats the WebAssembly value as a Go string.
func (self Value) String() string {
	switch self.ty {
	case Type_I32:
		return fmt.Sprintf("%d", self.ToI32())
	case Type_I64:
		return fmt.Sprintf("%d", self.ToI64())
	case Type_F32:
		return fmt.Sprintf("%f", self.ToF32())
	case Type_F64:
		return fmt.Sprintf("%f", self.ToF64())
	case Type_Void:
		return "void"
	default:
		return ""
	}
}
