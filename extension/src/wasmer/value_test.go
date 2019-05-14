package wasmer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValueI32(t *testing.T) {
	value := int32(7)
	wasm_value := I32(value)

	assert.Equal(t, value, wasm_value.ToI32())
	assert.Equal(t, Type_I32, wasm_value.GetType())
	assert.Equal(t, "7", wasm_value.String())
}

func TestValueI64(t *testing.T) {
	value := int64(7)
	wasm_value := I64(value)

	assert.Equal(t, value, wasm_value.ToI64())
	assert.Equal(t, Type_I64, wasm_value.GetType())
	assert.Equal(t, "7", wasm_value.String())
}

func TestValueF32(t *testing.T) {
	value := float32(7.42)
	wasm_value := F32(value)

	assert.Equal(t, value, wasm_value.ToF32())
	assert.Equal(t, Type_F32, wasm_value.GetType())
	assert.Equal(t, "7.420000", wasm_value.String())
}

func TestValueF64(t *testing.T) {
	value := float64(7.42)
	wasm_value := F64(value)

	assert.Equal(t, value, wasm_value.ToF64())
	assert.Equal(t, Type_F64, wasm_value.GetType())
	assert.Equal(t, "7.420000", wasm_value.String())
}

func TestValueVoid(t *testing.T) {
	wasm_value := Void()

	assert.Equal(t, nil, wasm_value.ToVoid())
	assert.Equal(t, Type_Void, wasm_value.GetType())
	assert.Equal(t, "void", wasm_value.String())
}
