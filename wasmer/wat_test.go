package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWat2Wasm(t *testing.T) {
	wasm, err := Wat2Wasm("(module)")
	assert.NoError(t, err)
	assert.Equal(t, wasm, []byte("\x00asm\x01\x00\x00\x00"))
}

func TestWasm2Wasm(t *testing.T) {
	wasm, err := Wat2Wasm(string([]byte("\x00asm\x01\x00\x00\x00")))
	assert.NoError(t, err)
	assert.Equal(t, wasm, []byte("\x00asm\x01\x00\x00\x00"))
}

func TestBadWat2Wasm(t *testing.T) {
	_, err := Wat2Wasm("(module")
	assert.EqualError(t, err, "expected `)`\n     --> <anon>:1:8\n      |\n    1 | (module\n      |        ^")
}
