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
	    captureStdout().
	    finalize()
	assert.NoError(t, err)

	importObject, err := wasiEnv.generateImportObject(store, module)

	instance, err := NewInstance(module, importObject)
	assert.NoError(t, err)

	start, err := instance.Exports.GetFunction("_start")
	assert.NoError(t, err)

	start()

	stdout := wasiEnv.readStdout()

	t.Log(stdout)
}

func TestWasiVersion(t *testing.T) {
	assert.Equal(t, WASI_VERSION_LATEST.String(), "__latest__")
	assert.Equal(t, WASI_VERSION_SNAPSHOT0.String(), "wasi_unstable")
	assert.Equal(t, WASI_VERSION_SNAPSHOT1.String(), "wasi_snapshot_preview1")
	assert.Equal(t, WASI_VERSION_INVALID.String(), "__unknown__")
}

func TestWasiGetVersion(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, testGetBytes("wasi.wasm"))
	assert.NoError(t, err)

	assert.Equal(t, GetWasiVersion(module), WASI_VERSION_SNAPSHOT1)
}
