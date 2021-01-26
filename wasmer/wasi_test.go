package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWasi(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, testGetBytes("wasi.wasm"))
	assert.NoError(t, err)

	wasiEnv, err := NewWasiStateBuilder("test-program").
	    argument("--foo").
	    environment("ABC", "DEF").
	    environment("X", "ZY").
	    mapDirectory("the_host_current_directory", ".").
	    finalize()
	assert.NoError(t, err)

	importObject, err := wasiEnv.generateImportObject(store, module)

	instance, err := NewInstance(module, importObject)
	assert.NoError(t, err)

	start, err := instance.Exports.GetFunction("_start")
	assert.NoError(t, err)

	start()
}
