package wasmer

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func testEngine(t *testing.T, engine *Engine) {
	store := NewStore(engine)
	module, err := NewModule(store, testGetBytes())
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
	triple, err := NewTriple("x86_64-apple-darwin")
	assert.NoError(t, err)

	cpuFeatures := NewCpuFeatures()
	err = cpuFeatures.Add("sse2")
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
