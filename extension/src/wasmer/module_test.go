package wasmer_test

import (
	"github.com/stretchr/testify/assert"
	"path"
	"runtime"
	"testing"
	wasm "wasmer"
)

func TestValidate(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "tests.wasm")

	bytes, _ := wasm.ReadBytes(module_path)
	output := wasm.Validate(bytes)

	assert.True(t, output)
}

func TestValidateInvalid(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "invalid.wasm")

	bytes, _ := wasm.ReadBytes(module_path)
	output := wasm.Validate(bytes)

	assert.False(t, output)
}
