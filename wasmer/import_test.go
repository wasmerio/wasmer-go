package wasmer_test

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
)

func TestImport(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "examples", "imported_function.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

	instance, err := wasm.NewInstance(bytes)
	defer instance.Close()

	assert.NoError(t, err)
	t.Log(wasm.GetLastError())

	add1, exists := instance.Exports["add1"]

	assert.Equal(t, true, exists)

	result, err := add1(1, 2)

	assert.NoError(t, err)
	t.Log(result)
	t.Log(wasm.GetLastError())
}
