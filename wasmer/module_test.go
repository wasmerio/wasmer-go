package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

// Almost immediate crash w/ memcheck:
// === RUN   TestModuleImports
// runtime: out of memory: cannot allocate 18446744073189457920-byte block (3702784 in use)
// fatal error: out of memory
//
// Or:
//
//	module_test.go:67:
//	      	Error Trace:	module_test.go:67
//	      	Error:      	Not equal:
//	      	            	expected: "function"
//	      	            	actual  : ""
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

	imports, release := module.ImportsSafe()
	defer release(module) // Keep module alive.

	assert.Equal(t, len(imports), 4)

	// 0
	assert.Equal(t, "ns", imports[0].Module())
	assert.Equal(t, "function", imports[0].Name())

	type0 := imports[0].Type()
	assert.Equal(t, FUNCTION, type0.Kind())

	functionType := type0.IntoFunctionType()
	assert.Equal(t, 0, len(functionType.Params()))
	assert.Equal(t, 0, len(functionType.Results()))

	// 1
	assert.Equal(t, "ns", imports[1].Module())
	assert.Equal(t, "global", imports[1].Name())

	type1 := imports[1].Type()
	assert.Equal(t, GLOBAL, type1.Kind())

	globalType := type1.IntoGlobalType()
	assert.Equal(t, F32, globalType.ValueType().Kind())
	assert.Equal(t, IMMUTABLE, globalType.Mutability())

	// 2
	assert.Equal(t, "ns", imports[2].Module())
	assert.Equal(t, "table", imports[2].Name())

	type2 := imports[2].Type()
	assert.Equal(t, TABLE, type2.Kind())

	tableType := type2.IntoTableType()
	tableLimits := tableType.Limits()
	assert.Equal(t, FuncRef, tableType.ValueType().Kind())
	assert.EqualValues(t, 1, tableLimits.Minimum())
	assert.EqualValues(t, 2, tableLimits.Maximum())

	// 3
	assert.Equal(t, "ns", imports[3].Module())
	assert.Equal(t, "memory", imports[3].Name())

	type3 := imports[3].Type()
	assert.Equal(t, MEMORY, type3.Kind())

	memoryType := type3.IntoMemoryType()
	memoryLimits := memoryType.Limits()
	assert.EqualValues(t, 3, memoryLimits.Minimum())
	assert.EqualValues(t, 4, memoryLimits.Maximum())
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

	exports, release := module.ExportsSafe()
	defer release(module) // Keep module alive.

	assert.Equal(t, 4, len(exports))

	// 0
	assert.Equal(t, "function", exports[0].Name())

	type0 := exports[0].Type()
	assert.Equal(t, FUNCTION, type0.Kind())

	functionType := type0.IntoFunctionType()
	assert.Equal(t, 2, len(functionType.Params()))
	assert.Equal(t, 0, len(functionType.Results()))

	// 1
	assert.Equal(t, "global", exports[1].Name())

	type1 := exports[1].Type()
	assert.Equal(t, GLOBAL, type1.Kind())

	globalType := type1.IntoGlobalType()
	assert.Equal(t, I32, globalType.ValueType().Kind())
	assert.Equal(t, IMMUTABLE, globalType.Mutability())

	// 2
	assert.Equal(t, "table", exports[2].Name())

	type2 := exports[2].Type()
	assert.Equal(t, TABLE, type2.Kind())

	tableType := type2.IntoTableType()
	tableLimits := tableType.Limits()
	assert.Equal(t, FuncRef, tableType.ValueType().Kind())
	assert.EqualValues(t, 0, tableLimits.Minimum())

	// 3
	assert.Equal(t, "memory", exports[3].Name())

	type3 := exports[3].Type()
	assert.Equal(t, MEMORY, type3.Kind())

	memoryType := type3.IntoMemoryType()
	memoryLimits := memoryType.Limits()
	assert.EqualValues(t, 1, memoryLimits.Minimum())
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

	// Note: We probably should use ExportsSafe (to be safe), but because
	// moduleAgain is used at the end of the test, moduleAgain is kept alive.
	exports := moduleAgain.Exports()
	assert.Equal(t, 1, len(exports))
	assert.Equal(t, "function", exports[0].Name())

	type0 := exports[0].Type()
	assert.Equal(t, FUNCTION, type0.Kind())

	functionType := type0.IntoFunctionType()
	assert.Equal(t, 2, len(functionType.Params()))
	assert.Equal(t, 0, len(functionType.Results()))

	_, err = NewInstance(moduleAgain, NewImportObject())
	assert.NoError(t, err)
}
