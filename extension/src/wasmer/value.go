package wasmer

import (
	"fmt"
	"math"
)

type ValueType int

const (
	Type_I32 ValueType = iota
	Type_I64
	Type_F32
	Type_F64
	Type_Void
)

type Value struct {
	value uint64
	ty    ValueType
}

func I32(value int32) Value {
	return Value{
		value: uint64(value),
		ty:    Type_I32,
	}
}

func I64(value int64) Value {
	return Value{
		value: uint64(value),
		ty:    Type_I64,
	}
}

func F32(value float32) Value {
	return Value{
		value: uint64(math.Float32bits(value)),
		ty:    Type_F32,
	}
}

func F64(value float64) Value {
	return Value{
		value: math.Float64bits(value),
		ty:    Type_F64,
	}
}

func Void() Value {
	return Value{
		value: 0,
		ty:    Type_Void,
	}
}

func (self Value) GetType() ValueType {
	return self.ty
}

func (self Value) ToI32() int32 {
	return int32(self.value)
}

func (self Value) ToI64() int64 {
	return int64(self.value)
}

func (self Value) ToF32() float32 {
	return math.Float32frombits(uint32(self.value))
}

func (self Value) ToF64() float64 {
	return math.Float64frombits(self.value)
}

func (self Value) ToVoid() interface{} {
	return nil
}

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
