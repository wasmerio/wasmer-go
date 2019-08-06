package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/spacemeshos/go-ext-wasm/wasmer"
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

func TestModuleSerialize(t *testing.T) {
	module1, err := wasm.Compile(GetBytes())
	defer module1.Close()

	assert.NoError(t, err)

	bytes, err := module1.Serialize()

	assert.NoError(t, err)

	module2, err := wasm.DeserializeModule(bytes)
	defer module2.Close()

	assert.NoError(t, err)

	instance, err := module2.Instantiate()
	defer instance.Close()

	assert.NoError(t, err)

	result, _ := instance.Exports["sum"](1, 2)

	assert.Equal(t, wasm.I32(3), result)
}

func TestModuleDeserializeModuleWithEmptyBytes(t *testing.T) {
	_, err := wasm.DeserializeModule([]byte{})

	assert.EqualError(t, err, "Serialized module bytes are empty.")
}

func TestModuleDeserializeModuleWithRandomBytes(t *testing.T) {
	_, err := wasm.DeserializeModule([]byte("random"))

	assert.EqualError(t, err, "Failed to deserialize the module.")
}

func TestModuleExports(t *testing.T) {
	module, _ := wasm.Compile(GetBytes())
	defer module.Close()

	assert.Equal(
		t,
		[]wasm.ExportDescriptor{
			wasm.ExportDescriptor{
				Name: "void",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "i32_i64_f32_f64_f64",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "sum",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "__heap_base",
				Kind: wasm.ImportExportKindGlobal,
			},
			wasm.ExportDescriptor{
				Name: "arity_0",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "i32_i32",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "memory",
				Kind: wasm.ImportExportKindMemory,
			},
			wasm.ExportDescriptor{
				Name: "bool_casted_to_i32",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "__data_end",
				Kind: wasm.ImportExportKindGlobal,
			},
			wasm.ExportDescriptor{
				Name: "f32_f32",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "f64_f64",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "string",
				Kind: wasm.ImportExportKindFunction,
			},
			wasm.ExportDescriptor{
				Name: "i64_i64",
				Kind: wasm.ImportExportKindFunction,
			},
		},
		module.Exports,
	)
}

func TestModuleImports(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "log.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

	module, _ := wasm.Compile(bytes)
	defer module.Close()

	assert.Equal(
		t,
		[]wasm.ImportDescriptor{
			wasm.ImportDescriptor{
				Name:      "log_message",
				Namespace: "env",
				Kind:      wasm.ImportExportKindFunction,
			},
		},
		module.Imports,
	)
}

func TestModuleImportsNone(t *testing.T) {
	module, _ := wasm.Compile(GetBytes())
	defer module.Close()

	assert.Equal(t, []wasm.ImportDescriptor{}, module.Imports)
}
