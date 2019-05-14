package wasmer

import (
	"fmt"
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

func TestValidate(t *testing.T) {
	output := Validate(GetBytes())

	assert.True(t, output)
}

func TestValidateInvalid(t *testing.T) {
	output := Validate(GetInvalidBytes())

	assert.False(t, output)
}

func TestInstantiate(t *testing.T) {
	instance, err := NewInstance(GetBytes())

	assert.NotNil(t, instance.instance)
	assert.NoError(t, err)
}

func TestInstantiateInvalidModule(t *testing.T) {
	instance, err := NewInstance(GetInvalidBytes())

	assert.EqualError(t, err, "Failed to compile the module.")
	assert.Nil(t, instance.instance)
}

func TestBasicSum(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("sum", []Value{ValueI32(1), ValueI32(2)})

	assert.Equal(t, ValueI32(3), output)
	assert.NoError(t, err)
}

func TestCallUndefinedFunction(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	function_name := "foo"
	output, err := instance.Call(function_name, []Value{})

	assert.Equal(t, ValueI32(0), output)
	assert.EqualError(t, err, fmt.Sprintf("The instance has no exported function named `%s`.", function_name))

	last_err, err := GetLastError()

	assert.Equal(t, "error instantiating", last_err)
	assert.NoError(t, err)
}

func TestCallMissingArguments(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("sum", []Value{ValueI32(1)})

	assert.Equal(t, ValueI32(0), output)
	assert.EqualError(t, err, "Missing 1 argument(s) when calling the `sum` exported function; Expect 2 argument(s), given 1.")
}

func TestCallExtraArguments(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("sum", []Value{ValueI32(1), ValueI32(2), ValueI32(3)})

	assert.Equal(t, ValueI32(0), output)
	assert.EqualError(t, err, "Given 1 extra argument(s) when calling the `sum` exported function; Expect 2 argument(s), given 3.")
}

func TestCallArity0(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("arity_0", []Value{})

	assert.Equal(t, ValueI32(42), output)
	assert.NoError(t, err)
}

func TestCallI32I32(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("i32_i32", []Value{ValueI32(7)})

	assert.Equal(t, ValueI32(7), output)
	assert.NoError(t, err)
}

func TestCallI64I64(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("i64_i64", []Value{ValueI64(7)})

	assert.Equal(t, ValueI64(7), output)
	assert.NoError(t, err)
}

func TestCallF32F32(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("f32_f32", []Value{ValueF32(7.42)})

	assert.Equal(t, ValueF32(7.42), output)
	assert.NoError(t, err)
}

func TestCallF64F64(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("f64_f64", []Value{ValueF64(7.42)})

	assert.Equal(t, ValueF64(7.42), output)
	assert.NoError(t, err)
}

func TestCallI32I64F32F64F64(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call(
		"i32_i64_f32_f64_f64",
		[]Value{
			ValueI32(1),
			ValueI64(2),
			ValueF32(3.4),
			ValueF64(5.6),
		},
	)

	assert.Equal(t, Type_F64, output.GetType())
	assert.Equal(t, 1 + 2 + 3.4 + 5.6, float64(int(output.ToF64() * 10000000)) / 10000000)
	assert.NoError(t, err)
}

func TestCallBoolCastToI32(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("bool_casted_to_i32", []Value{})

	assert.Equal(t, ValueI32(1), output)
	assert.NoError(t, err)
}

func TestCallString(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("string", []Value{})

	assert.Equal(t, ValueI32(1048576), output)
	assert.NoError(t, err)
}

func TestCallVoid(t *testing.T) {
	instance, _ := NewInstance(GetBytes())
	output, err := instance.Call("void", []Value{})

	assert.Equal(t, ValueVoid(), output)
	assert.NoError(t, err)
}
