package wasmer

import (
	"path"
	"runtime"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "tests.wasm")

	bytes, _ := ReadBytes(module_path)
	output := Validate(bytes)

	assert.True(t, output)
}

func TestValidateInvalid(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "invalid.wasm")

	bytes, _ := ReadBytes(module_path)
	output := Validate(bytes)

	assert.False(t, output)
}
