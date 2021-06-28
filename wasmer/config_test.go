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
	assert.Equal(t, UNIVERSAL.String(), "universal")
	assert.Equal(t, DYLIB.String(), "dylib")

	// Deprecated.
	assert.Equal(t, JIT.String(), "universal")
	assert.Equal(t, NATIVE.String(), "dylib")
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
	has_universal := IsEngineAvailable(UNIVERSAL)
	has_dylib := IsEngineAvailable(DYLIB)

	if IsCompilerAvailable(CRANELIFT) {
		// Cranelift with the Universal engine works everywhere
		if has_universal {
			configs = append(configs, Test{"Cranelift", "Universal", NewConfig().UseCraneliftCompiler().UseUniversalEngine()})
		}

		// Cranelift with the Dylib engine works on Linux+Darwin/amd64.
		if has_dylib && is_amd64 && (is_linux || is_darwin) {
			configs = append(configs, Test{"Cranelift", "Dylib", NewConfig().UseCraneliftCompiler().UseDylibEngine()})
		}
	}

	if IsCompilerAvailable(LLVM) {
		// LLVM with the Universal engine works on Linux+Darwin/amd64.
		if has_universal && is_amd64 && (is_linux || is_darwin) {
			configs = append(configs, Test{"LLVM", "Universal", NewConfig().UseLLVMCompiler().UseUniversalEngine()})
		}

		// LLVM with the Dylib engine works on Linux+Darwin/amd64+aarch64.
		if has_dylib && (is_linux || is_darwin) {
			configs = append(configs, Test{"LLVM", "Dylib", NewConfig().UseLLVMCompiler().UseDylibEngine()})
		}
	}

	if IsCompilerAvailable(SINGLEPASS) {
		// Singlepass with the Universal engine works on Linux+Darwin/amd64.
		if has_universal && is_amd64 && (is_linux || is_darwin) {
			configs = append(configs, Test{"Singlepass", "Universal", NewConfig().UseSinglepassCompiler().UseUniversalEngine()})
		}
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
