package wasmer

import (
	"path"
	"runtime"
	"testing"
	"github.com/stretchr/testify/assert"
)

func GetBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "tests.wasm")

	bytes, _ := ReadBytes(module_path)

	return bytes
}

func GetInvalidBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "invalid.wasm")

	bytes, _ := ReadBytes(module_path)

	return bytes
}

func TestInstantiate(t *testing.T) {
	instance, err := NewInstance(GetBytes())
	defer instance.Close()

	assert.NotNil(t, instance.instance)
	assert.NoError(t, err)
}

func TestInstantiateInvalidModule(t *testing.T) {
	instance, err := NewInstance(GetInvalidBytes())
	defer instance.Close()

	assert.EqualError(t, err, "Failed to compile the module.")
	assert.Nil(t, instance.instance)
}

func TestBasicSum(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	sum := instance.Exports["sum"]
	output, err := sum(1, 2)

	assert.Equal(t, I32(3), output)
	assert.NoError(t, err)

	output, err = sum(I32(1), I32(2))

	assert.Equal(t, I32(3), output)
	assert.NoError(t, err)
}

func TestCallUndefinedFunction(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	function_name := "foo"
	_, export_exists := instance.Exports[function_name]

	assert.Equal(t, false, export_exists)
}

func TestCallMissingArguments(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["sum"](1)

	assert.Equal(t, I32(0), output)
	assert.EqualError(t, err, "Missing 1 argument(s) when calling the `sum` exported function; Expect 2 argument(s), given 1.")
}

func TestCallExtraArguments(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["sum"](1, 2, 3)

	assert.Equal(t, I32(0), output)
	assert.EqualError(t, err, "Given 1 extra argument(s) when calling the `sum` exported function; Expect 2 argument(s), given 3.")
}

func TestCallArity0(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["arity_0"]()

	assert.Equal(t, I32(42), output)
	assert.NoError(t, err)
}

func TestCallI32I32(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	i32_i32 := instance.Exports["i32_i32"]

	output, err := i32_i32(7)
	assert.Equal(t, I32(7), output)
	assert.NoError(t, err)

	output, err = i32_i32(int8(7))
	assert.Equal(t, I32(7), output)

	output, err = i32_i32(uint8(7))
	assert.Equal(t, I32(7), output)

	output, err = i32_i32(int16(7))
	assert.Equal(t, I32(7), output)

	output, err = i32_i32(uint16(7))
	assert.Equal(t, I32(7), output)

	output, err = i32_i32(int32(7))
	assert.Equal(t, I32(7), output)

	output, err = i32_i32(int(7))
	assert.Equal(t, I32(7), output)

	output, err = i32_i32(uint(7))
	assert.Equal(t, I32(7), output)

	output, err = i32_i32(I32(7))
	assert.Equal(t, I32(7), output)

	_, err = i32_i32(int64(7))
	assert.EqualError(t, err, "Argument #1 of the `i32_i32` exported function must be of type `i32`, cannot cast given value to this type.")
}

func TestCallI64I64(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	i64_i64 := instance.Exports["i64_i64"]

	output, err := i64_i64(7)
	assert.Equal(t, I64(7), output)
	assert.NoError(t, err)

	output, err = i64_i64(int8(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(uint8(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(int16(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(uint16(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(int32(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(uint32(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(int64(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(int(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(uint(7))
	assert.Equal(t, I64(7), output)

	output, err = i64_i64(I64(7))
	assert.Equal(t, I64(7), output)

	_, err = i64_i64("foo")
	assert.EqualError(t, err, "Argument #1 of the `i64_i64` exported function must be of type `i64`, cannot cast given value to this type.")
}

func TestCallF32F32(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	f32_f32 := instance.Exports["f32_f32"]

	output, err := f32_f32(float32(7.42))
	assert.Equal(t, F32(7.42), output)
	assert.NoError(t, err)

	output, err = f32_f32(F32(7.42))
	assert.Equal(t, F32(7.42), output)

	_, err = f32_f32("foo")
	assert.EqualError(t, err, "Argument #1 of the `f32_f32` exported function must be of type `f32`, cannot cast given value to this type.")
}

func TestCallF64F64(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	f64_f64 := instance.Exports["f64_f64"]

	output, err := f64_f64(7.42)
	assert.Equal(t, F64(7.42), output)
	assert.NoError(t, err)

	output, err = f64_f64(float32(7.42))
	assert.Equal(t, F64(float64(float32(7.42))), output)

	output, err = f64_f64(float64(7.42))
	assert.Equal(t, F64(7.42), output)

	output, err = f64_f64(F64(7.42))
	assert.Equal(t, F64(7.42), output)

	_, err = f64_f64("foo")
	assert.EqualError(t, err, "Argument #1 of the `f64_f64` exported function must be of type `f64`, cannot cast given value to this type.")
}

func TestCallI32I64F32F64F64(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["i32_i64_f32_f64_f64"](
		1,
		2,
		float32(3.4),
		5.6,
	)

	assert.Equal(t, Type_F64, output.GetType())
	assert.Equal(t, 1 + 2 + 3.4 + 5.6, float64(int(output.ToF64() * 10000000)) / 10000000)
	assert.NoError(t, err)
}

func TestCallBoolCastToI32(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["bool_casted_to_i32"]()

	assert.Equal(t, I32(1), output)
	assert.NoError(t, err)
}

func TestCallString(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["string"]()

	assert.Equal(t, I32(1048576), output)
	assert.NoError(t, err)
}

func TestCallVoid(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["void"]()

	assert.Equal(t, Void(), output)
	assert.NoError(t, err)
}
