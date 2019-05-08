package wasmer

/*
#cgo LDFLAGS: -L../../ -lwasmer_runtime_c_api
#include "../../wasmer.h"
*/
import "C"
import (
	"fmt"
	"io/ioutil"
	"math"
	"unsafe"
)

func ReadBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func Validate(bytes []byte) bool {
	return true == C.wasmer_validate((*C.uchar) (unsafe.Pointer(&bytes[0])), C.uint(len(bytes)))
}

type ValueType int

const (
	Type_I32 ValueType = iota
	Type_I64
	Type_F32
	Type_F64
)

type Value struct {
	value uint64
	ty ValueType
}

func ValueI32(value int32) Value {
	return Value {
		value: uint64(value),
		ty: Type_I32,
	};
}

func ValueI64(value int64) Value {
	return Value {
		value: uint64(value),
		ty: Type_I64,
	};
}

func ValueF32(value float32) Value {
	return Value {
		value: uint64(math.Float32bits(value)),
		ty: Type_F32,
	};
}

func ValueF64(value float64) Value {
	return Value {
		value: math.Float64bits(value),
		ty: Type_F64,
	};
}

func (self Value) GetType() ValueType {
	return self.ty;
}

func (self Value) ToI32() int32 {
	return int32(self.value);
}

func (self Value) ToI64() int64 {
	return int64(self.value);
}

func (self Value) ToF32() float32 {
	return math.Float32frombits(uint32(self.value));
}

func (self Value) ToF64() float64 {
	return math.Float64frombits(self.value);
}

func (self Value) String() string {
	switch (self.ty) {
	case Type_I32:
		return fmt.Sprintf("%d", self.ToI32());
	case Type_I64:
		return fmt.Sprintf("%d", self.ToI64())
	case Type_F32:
		return fmt.Sprintf("%f", self.ToF32())
	case Type_F64:
		return fmt.Sprintf("%f", self.ToF64())
	default:
		return ""
	}
}

type Instance struct {
	instance *C.wasmer_instance_t
}

func NewInstance(bytes []byte) Instance {
	var imports []C.wasmer_import_t = []C.wasmer_import_t{}
	var instance *C.wasmer_instance_t = nil

	var compile_result C.wasmer_result_t = C.wasmer_instantiate(
		&instance,
		(*C.uchar) (unsafe.Pointer(&bytes[0])),
		C.uint(len(bytes)),
		(*C.wasmer_import_t) (unsafe.Pointer(&imports)),
		C.int(0),
	)

	if (C.WASMER_OK != compile_result) {
		panic("Failed to compile the module.")
	}

	return Instance { instance }
}

func (self Instance) Call(function_name string) Value {
	var parameter_one C.wasmer_value_t
	parameter_one.tag = C.WASM_I32
	parameter_one.value[C.WASM_I32] = 1

	var parameter_two C.wasmer_value_t
	parameter_two.tag = C.WASM_I32
	parameter_two.value[C.WASM_I32] = 2

	var parameters []C.wasmer_value_t = []C.wasmer_value_t{parameter_one, parameter_two}

	var result_one C.wasmer_value_t
	var results []C.wasmer_value_t = []C.wasmer_value_t{result_one}

	var export_name = C.CString("sum")
	//defer C.free(unsafe.Pointer(export_name))

	var call_result C.wasmer_result_t = C.wasmer_instance_call(
		self.instance,
		export_name,
		(*C.wasmer_value_t) (unsafe.Pointer(&parameters[0])),
		C.int(len(parameters)),
		(*C.wasmer_value_t) (unsafe.Pointer(&results[0])),
		C.int(len(results)),
	)

	if (C.WASMER_OK != call_result) {
		panic("Failed to call the `sum` function.")
	}

	switch results[0].tag {
	case C.WASM_I32:
		return ValueI32(int32(results[0].value[C.WASM_I32]));
	case C.WASM_I64:
		return ValueI64(int64(results[0].value[C.WASM_I64]));
	case C.WASM_F32:
		return ValueF32(float32(results[0].value[C.WASM_F32]));
	case C.WASM_F64:
		return ValueF64(float64(results[0].value[C.WASM_F64]));
	default:
		panic("unreachable")
	}
}
