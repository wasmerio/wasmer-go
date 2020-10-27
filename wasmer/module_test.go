package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModule(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	_, err := NewModule(store, []byte("(module)"))
	assert.NoError(t, err)
}

func TestValidateModule(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	err := ValidateModule(store, []byte("(module)"))
	assert.NoError(t, err)
}

func TestModuleNameSome(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, []byte("(module $moduleName)"))
	assert.NoError(t, err)

	name := module.Name()
	assert.Equal(t, name, "moduleName")
}

func TestModuleNameNone(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, []byte("(module)"))
	assert.NoError(t, err)

	name := module.Name()
	assert.Equal(t, name, "")
}

func TestModuleImports(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
                	(module
                	  (import "ns" "function" (func))
                	  (import "ns" "global" (global f32))
                	  (import "ns" "table" (table 1 2 anyfunc))
                	  (import "ns" "memory" (memory 3 4)))
                `),
	)
	assert.NoError(t, err)

	imports := module.Imports()
	assert.Equal(t, len(imports), 4)

	// 0
	assert.Equal(t, imports[0].Module(), "ns")
	assert.Equal(t, imports[0].Name(), "function")

	type0 := imports[0].Type()
	assert.Equal(t, type0.Kind(), FUNCTION)

	functionType := type0.IntoFunctionType()
	assert.Equal(t, len(functionType.Params()), 0)
	assert.Equal(t, len(functionType.Results()), 0)

	// 1
	assert.Equal(t, imports[1].Module(), "ns")
	assert.Equal(t, imports[1].Name(), "global")

	type1 := imports[1].Type()
	assert.Equal(t, type1.Kind(), GLOBAL)

	globalType := type1.IntoGlobalType()
	assert.Equal(t, globalType.ValueType().Kind(), F32)
	assert.Equal(t, globalType.Mutability(), IMMUTABLE)

	// 2
	assert.Equal(t, imports[2].Module(), "ns")
	assert.Equal(t, imports[2].Name(), "table")

	type2 := imports[2].Type()
	assert.Equal(t, type2.Kind(), TABLE)

	tableType := type2.IntoTableType()
	tableLimits := tableType.Limits()
	assert.Equal(t, tableType.ValueType().Kind(), FuncRef)
	assert.Equal(t, tableLimits.Minimum(), uint32(1))
	assert.Equal(t, tableLimits.Maximum(), uint32(2))

	// 3
	assert.Equal(t, imports[3].Module(), "ns")
	assert.Equal(t, imports[3].Name(), "memory")

	type3 := imports[3].Type()
	assert.Equal(t, type3.Kind(), MEMORY)

	memoryType := type3.IntoMemoryType()
	memoryLimits := memoryType.Limits()
	assert.Equal(t, memoryLimits.Minimum(), uint32(3))
	assert.Equal(t, memoryLimits.Maximum(), uint32(4))
}

func TestModuleExports(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (func (export "function") (param i32 i64))
			  (global (export "global") i32 (i32.const 7))
			  (table (export "table") 0 funcref)
			  (memory (export "memory") 1))
		`),
	)
	assert.NoError(t, err)

	exports := module.Exports()
	assert.Equal(t, len(exports), 4)

	// 0
	assert.Equal(t, exports[0].Name(), "function")

	type0 := exports[0].Type()
	assert.Equal(t, type0.Kind(), FUNCTION)

	functionType := type0.IntoFunctionType()
	assert.Equal(t, len(functionType.Params()), 2)
	assert.Equal(t, len(functionType.Results()), 0)

	// 1
	assert.Equal(t, exports[1].Name(), "global")

	type1 := exports[1].Type()
	assert.Equal(t, type1.Kind(), GLOBAL)

	globalType := type1.IntoGlobalType()
	assert.Equal(t, globalType.ValueType().Kind(), I32)
	assert.Equal(t, globalType.Mutability(), IMMUTABLE)

	// 2
	assert.Equal(t, exports[2].Name(), "table")

	type2 := exports[2].Type()
	assert.Equal(t, type2.Kind(), TABLE)

	tableType := type2.IntoTableType()
	tableLimits := tableType.Limits()
	assert.Equal(t, tableType.ValueType().Kind(), FuncRef)
	assert.Equal(t, tableLimits.Minimum(), uint32(0))

	// 3
	assert.Equal(t, exports[3].Name(), "memory")

	type3 := exports[3].Type()
	assert.Equal(t, type3.Kind(), MEMORY)

	memoryType := type3.IntoMemoryType()
	memoryLimits := memoryType.Limits()
	assert.Equal(t, memoryLimits.Minimum(), uint32(1))
}

func TestModuleSerialize(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	_, err := NewModule(store, []byte("(module)"))
	assert.NoError(t, err)
}

func TestModuleDeserialize(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (func (export "function") (param i32 i64)))
		`),
	)
	assert.NoError(t, err)

	serializedModule, err := module.Serialize()
	assert.NoError(t, err)

	moduleAgain, err := DeserializeModule(store, serializedModule)
	assert.NoError(t, err)

	exports := moduleAgain.Exports()
	assert.Equal(t, len(exports), 1)
	assert.Equal(t, exports[0].Name(), "function")

	type0 := exports[0].Type()
	assert.Equal(t, type0.Kind(), FUNCTION)

	functionType := type0.IntoFunctionType()
	assert.Equal(t, len(functionType.Params()), 2)
	assert.Equal(t, len(functionType.Results()), 0)
}
