package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestWasiWithCapturedStdout(t *testing.T) {
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

	importObject, err := wasiEnv.GenerateImportObject(store, module)

	instance, err := NewInstance(module, importObject)
	assert.NoError(t, err)

	start, err := instance.Exports.GetWasiStartFunction()
	assert.NoError(t, err)

	start()

	stdout := string(wasiEnv.readStdout())

	assert.Equal(
		t,
		stdout,
		"Found program name: `test-program`\n"+
			"Found 1 arguments: --foo\n"+
			"Found 2 environment variables: ABC=DEF, X=ZY\n"+
			"Found 1 preopened directories: DirEntry(\"/the_host_current_directory\")\n",
	)
}
