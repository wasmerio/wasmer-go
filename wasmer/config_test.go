package wasmer

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestConfig(t *testing.T) {
	config := NewConfig()

	engine := NewEngineWithConfig(config)
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

func TestConfig_UseJITEngine(t *testing.T) {
	config := NewConfig()
	config.UseJITEngine()

	engine := NewEngineWithConfig(config)
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

func TestConfig_UseNativeEngine(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		return
	}

	config := NewConfig()
	config.UseNativeEngine()

	engine := NewEngineWithConfig(config)
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

func TestConfig_UseCraneliftCompiler(t *testing.T) {
	config := NewConfig()
	config.UseCraneliftCompiler()

	engine := NewEngineWithConfig(config)
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

func TestConfig_UseLLVMCompiler(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		return
	}

	config := NewConfig()
	config.UseLLVMCompiler()

	engine := NewEngineWithConfig(config)
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

func TestConfig_JITWithCranelift(t *testing.T) {
	config := NewConfig()
	config.UseJITEngine()
	config.UseCraneliftCompiler()

	engine := NewEngineWithConfig(config)
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

func TestConfig_JITWithLLVM(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		return
	}

	config := NewConfig()
	config.UseJITEngine()
	config.UseLLVMCompiler()

	engine := NewEngineWithConfig(config)
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

func TestConfig_NativeWithCranelift(t *testing.T) {
	config := NewConfig()
	config.UseNativeEngine()
	config.UseCraneliftCompiler()

	engine := NewEngineWithConfig(config)
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

func TestConfig_NativeWithLLVM(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		return
	}

	config := NewConfig()
	config.UseNativeEngine()
	config.UseLLVMCompiler()

	engine := NewEngineWithConfig(config)
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
