package wasmer

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func testEngine(t *testing.T, engine *Engine) {
	store := NewStore(engine)
	module, err := NewModule(store, testGetBytes("tests.wasm"))
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	sum, err := instance.Exports.GetFunction("sum")
	assert.NoError(t, err)

	result, err := sum(37, 5)
	assert.NoError(t, err)
	assert.Equal(t, result, int32(42))
}

func TestEngine(t *testing.T) {
	testEngine(t, NewEngine())
}

func TestJITEngine(t *testing.T) {
	testEngine(t, NewJITEngine())
}

func TestNativeEngine(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		return
	}

	testEngine(t, NewNativeEngine())
}

func TestEngineWithTarget(t *testing.T) {
	triple, err := NewTriple("aarch64-unknown-linux-gnu")
	assert.NoError(t, err)

	cpuFeatures := NewCpuFeatures()
	assert.NoError(t, err)

	target := NewTarget(triple, cpuFeatures)

	config := NewConfig()
	config.UseTarget(target)

	engine := NewEngineWithConfig(config)
	store := NewStore(engine)

	module, err := NewModule(store, testGetBytes())
	assert.NoError(t, err)

	_ = module
}
