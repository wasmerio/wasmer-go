package wasmer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestCompilerKind(t *testing.T) {
	assert.Equal(t, CRANELIFT.String(), "cranelift")
	assert.Equal(t, LLVM.String(), "llvm")
	assert.Equal(t, SINGLEPASS.String(), "singlepass")
}

func TestEngineKind(t *testing.T) {
	assert.Equal(t, JIT.String(), "jit")
	assert.Equal(t, NATIVE.String(), "native")
}

func TestConfig(t *testing.T) {
	config := NewConfig()

	engine := NewEngineWithConfig(config)
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

func TestConfig_AllCombinations(t *testing.T) {
	type Test struct {
		compilerName string
		engineName   string
		config       *Config
	}
	var configs = []Test{}

	is_amd64 := runtime.GOARCH == "amd64"
	//is_aarch64 := runtime.GOARCH == "arm64"
	is_linux := runtime.GOOS == "linux"
	is_darwin := runtime.GOOS == "darwin"
	//is_windows := runtime.GOOS == "windows"

	// Cranelift with the JIT engine works everywhere
	configs = append(configs, Test{"Cranelift", "JIT", NewConfig().UseCraneliftCompiler().UseJITEngine()})

	// Cranelift with the Native engine works on Linux+Darwin/amd64.
	if is_amd64 && (is_linux || is_darwin) {
		configs = append(configs, Test{"Cranelift", "Native", NewConfig().UseCraneliftCompiler().UseNativeEngine()})
	}

	/*
		LLVM isn't part of the prepacked `libwasmer` for the moment.

			// LLVM with the JIT engine works on Linux+Darwin/amd64.
			if is_amd64 && (is_linux || is_darwin) {
				configs = append(configs, Test{"LLVM", "JIT", NewConfig().UseLLVMCompiler().UseJITEngine()})
			}

			// LLVM with the Native engine works on Linux+Darwin/amd64.
			if is_amd64 && (is_linux || is_darwin) {
				configs = append(configs, Test{"LLVM", "Native", NewConfig().UseLLVMCompiler().UseNativeEngine()})
			}
	*/

	// Singlepass with the JIT engine works on Linux+Darwin/amd64.
	if is_amd64 && (is_linux || is_darwin) {
		configs = append(configs, Test{"Singlepass", "JIT", NewConfig().UseSinglepassCompiler().UseJITEngine()})
	}

	for _, test := range configs {
		t.Run(
			fmt.Sprintf("compiler=%s, engine=%s", test.compilerName, test.engineName),
			func(t *testing.T) {
				engine := NewEngineWithConfig(test.config)
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
			},
		)
	}
}
