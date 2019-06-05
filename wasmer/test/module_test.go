package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
)

func TestValidate(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "tests.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)
	output := wasm.Validate(bytes)

	assert.True(t, output)
}

func TestValidateInvalid(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "invalid.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)
	output := wasm.Validate(bytes)

	assert.False(t, output)
}

func TestCompile(t *testing.T) {
	module, err := wasm.Compile(GetBytes())
	defer module.Close()

	assert.NoError(t, err)
}

func TestCompileInvalidModule(t *testing.T) {
	module, err := wasm.Compile(GetInvalidBytes())
	defer module.Close()

	assert.EqualError(t, err, "Failed to compile the module.")
}

func TestModuleInstantiate(t *testing.T) {
	module, err := wasm.Compile(GetBytes())
	defer module.Close()

	assert.NoError(t, err)

	instance, err := module.Instantiate()
	defer instance.Close()

	assert.NoError(t, err)

	result, _ := instance.Exports["sum"](1, 2)

	assert.Equal(t, wasm.I32(3), result)
}
