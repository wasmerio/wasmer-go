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

	assert.Error(t, err, "Failed to compile the module.")
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
	assert.Errorf(t, err, "Failed to call the `%s` exported function", function_name)
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
