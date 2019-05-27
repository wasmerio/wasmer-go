package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"testing"
)

func TestValueI32(t *testing.T) {
	value := int32(7)
	wasmValue := wasm.I32(value)

	assert.Equal(t, value, wasmValue.ToI32())
	assert.Equal(t, wasm.TypeI32, wasmValue.GetType())
	assert.Equal(t, "7", wasmValue.String())
}

func TestValueI64(t *testing.T) {
	value := int64(7)
	wasmValue := wasm.I64(value)

	assert.Equal(t, value, wasmValue.ToI64())
	assert.Equal(t, wasm.TypeI64, wasmValue.GetType())
	assert.Equal(t, "7", wasmValue.String())
}

func TestValueF32(t *testing.T) {
	value := float32(7.42)
	wasmValue := wasm.F32(value)

	assert.Equal(t, value, wasmValue.ToF32())
	assert.Equal(t, wasm.TypeF32, wasmValue.GetType())
	assert.Equal(t, "7.420000", wasmValue.String())
}

func TestValueF64(t *testing.T) {
	value := float64(7.42)
	wasmValue := wasm.F64(value)

	assert.Equal(t, value, wasmValue.ToF64())
	assert.Equal(t, wasm.TypeF64, wasmValue.GetType())
	assert.Equal(t, "7.420000", wasmValue.String())
}
