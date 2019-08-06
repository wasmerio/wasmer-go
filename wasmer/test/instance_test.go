package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/spacemeshos/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
)

func GetBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "tests.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

	return bytes
}

func GetInvalidBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "invalid.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

	return bytes
}

func TestInstantiate(t *testing.T) {
	instance, err := wasm.NewInstance(GetBytes())
	defer instance.Close()

	assert.NoError(t, err)
}

func TestInstantiateInvalidModule(t *testing.T) {
	instance, err := wasm.NewInstance(GetInvalidBytes())
	defer instance.Close()

	assert.EqualError(t, err, "Failed to instantiate the module:\n    compile error: Validation error \"Invalid type\"")
}

func TestBasicSum(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	sum := instance.Exports["sum"]
	output, err := sum(1, 2)

	assert.Equal(t, wasm.I32(3), output)
	assert.NoError(t, err)

	output, err = sum(wasm.I32(1), wasm.I32(2))

	assert.Equal(t, wasm.I32(3), output)
	assert.NoError(t, err)
}

func TestCallUndefinedFunction(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	functionName := "foo"
	_, exportExists := instance.Exports[functionName]

	assert.Equal(t, false, exportExists)
}

func TestCallMissingArguments(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["sum"](1)

	assert.Equal(t, wasm.I32(0), output)
	assert.EqualError(t, err, "Missing 1 argument(s) when calling the `sum` exported function; Expect 2 argument(s), given 1.")
}

func TestCallExtraArguments(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["sum"](1, 2, 3)

	assert.Equal(t, wasm.I32(0), output)
	assert.EqualError(t, err, "Given 1 extra argument(s) when calling the `sum` exported function; Expect 2 argument(s), given 3.")
}

func TestCallArity0(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["arity_0"]()

	assert.Equal(t, wasm.I32(42), output)
	assert.NoError(t, err)
}

func TestCallI32I32(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	i32_i32 := instance.Exports["i32_i32"]

	output, err := i32_i32(7)
	assert.Equal(t, wasm.I32(7), output)
	assert.NoError(t, err)

	output, _ = i32_i32(int8(7))
	assert.Equal(t, wasm.I32(7), output)

	output, _ = i32_i32(uint8(7))
	assert.Equal(t, wasm.I32(7), output)

	output, _ = i32_i32(int16(7))
	assert.Equal(t, wasm.I32(7), output)

	output, _ = i32_i32(uint16(7))
	assert.Equal(t, wasm.I32(7), output)

	output, _ = i32_i32(int32(7))
	assert.Equal(t, wasm.I32(7), output)

	output, _ = i32_i32(int(7))
	assert.Equal(t, wasm.I32(7), output)

	output, _ = i32_i32(uint(7))
	assert.Equal(t, wasm.I32(7), output)

	output, _ = i32_i32(wasm.I32(7))
	assert.Equal(t, wasm.I32(7), output)

	_, err = i32_i32(int64(7))
	assert.EqualError(t, err, "Argument #1 of the `i32_i32` exported function must be of type `i32`, cannot cast given value to this type.")
}

func TestCallI64I64(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	i64_i64 := instance.Exports["i64_i64"]

	output, err := i64_i64(7)
	assert.Equal(t, wasm.I64(7), output)
	assert.NoError(t, err)

	output, _ = i64_i64(int8(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(uint8(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(int16(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(uint16(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(int32(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(uint32(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(int64(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(int(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(uint(7))
	assert.Equal(t, wasm.I64(7), output)

	output, _ = i64_i64(wasm.I64(7))
	assert.Equal(t, wasm.I64(7), output)

	_, err = i64_i64("foo")
	assert.EqualError(t, err, "Argument #1 of the `i64_i64` exported function must be of type `i64`, cannot cast given value to this type.")
}

func TestCallF32F32(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	f32_f32 := instance.Exports["f32_f32"]

	output, err := f32_f32(float32(7.42))
	assert.Equal(t, wasm.F32(7.42), output)
	assert.NoError(t, err)

	output, _ = f32_f32(wasm.F32(7.42))
	assert.Equal(t, wasm.F32(7.42), output)

	_, err = f32_f32("foo")
	assert.EqualError(t, err, "Argument #1 of the `f32_f32` exported function must be of type `f32`, cannot cast given value to this type.")
}

func TestCallF64F64(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	f64_f64 := instance.Exports["f64_f64"]

	output, err := f64_f64(7.42)
	assert.Equal(t, wasm.F64(7.42), output)
	assert.NoError(t, err)

	output, _ = f64_f64(float32(7.42))
	assert.Equal(t, wasm.F64(float64(float32(7.42))), output)

	output, _ = f64_f64(float64(7.42))
	assert.Equal(t, wasm.F64(7.42), output)

	output, _ = f64_f64(wasm.F64(7.42))
	assert.Equal(t, wasm.F64(7.42), output)

	_, err = f64_f64("foo")
	assert.EqualError(t, err, "Argument #1 of the `f64_f64` exported function must be of type `f64`, cannot cast given value to this type.")
}

func TestCallI32I64F32F64F64(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["i32_i64_f32_f64_f64"](
		1,
		2,
		float32(3.4),
		5.6,
	)

	assert.Equal(t, wasm.TypeF64, output.GetType())
	assert.Equal(t, 1+2+3.4+5.6, float64(int(output.ToF64()*10000000))/10000000)
	assert.NoError(t, err)
}

func TestCallBoolCastToI32(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["bool_casted_to_i32"]()

	assert.Equal(t, wasm.I32(1), output)
	assert.NoError(t, err)
}

func TestCallString(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["string"]()

	assert.Equal(t, wasm.I32(1048576), output)
	assert.NoError(t, err)
}

func TestCallVoid(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	output, err := instance.Exports["void"]()

	assert.Equal(t, wasm.TypeVoid, output.GetType())
	assert.NoError(t, err)
}
