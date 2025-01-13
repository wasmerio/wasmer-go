package wasmer

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, []byte("(module)"))
	assert.NoError(t, err)

	_, err = NewInstance(module, NewImportObject())
	assert.NoError(t, err)
}

func TestInstanceExports(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, moduledone, err := NewModuleSafe(
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
	defer moduledone(module)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	extern, err := instance.Exports.Get("function")
	assert.NoError(t, err)
	assert.Equal(t, extern.Kind(), FUNCTION)

	function, err := instance.Exports.GetFunction("function")
	assert.NoError(t, err)
	assert.NotNil(t, function)

	global, err := instance.Exports.GetGlobal("global")
	assert.NoError(t, err)
	assert.NotNil(t, global)

	table, err := instance.Exports.GetTable("table")
	assert.NoError(t, err)
	assert.NotNil(t, table)

	memory, err := instance.Exports.GetMemory("memory")
	assert.NoError(t, err)
	assert.NotNil(t, memory)
}

func TestInstanceMissingImports(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, moduledone, err := NewModuleSafe(
		store,
		[]byte(`
			(module
			  (func (import "missing" "function"))
			  (func (import "exists" "function")))
		`),
	)
	assert.NoError(t, err)
	defer moduledone(module)

	function := NewFunction(
		store,
		NewFunctionType(NewValueTypes(), NewValueTypes()),
		func(args []Value) ([]Value, error) {
			return []Value{}, nil
		},
	)

	importObject := NewImportObject()
	importObject.Register(
		"exists",
		map[string]IntoExtern{
			"function": function,
		},
	)

	_, err = NewInstance(module, importObject)
	assert.Error(t, err)
}

func TestInstanceTraps(t *testing.T) {
	if runtime.GOOS == "darwin" {
		// This test fails on darwin (when executed repeated) with the following error:
		//   signal 16 received but handler not on signal stack
		//   mp.gsignal stack ....
		//   fatal error: non-Go code set up signal handler without SA_ONSTACK flag
		//
		// This is a bit strange since
		// https://github.com/wasmerio/wasmer/blob/cfb9413a670a0123f4f403ecf1897257fb681e72/lib/vm/src/trap/traphandlers.rs#L163
		// appears to set this flag; but maybe it's not applied to darwin?
		t.Skip("skipping test on non-linux platforms")
	}

	engine := NewEngine()
	store := NewStore(engine)
	module, moduledone, err := NewModuleSafe(
		store,
		[]byte(`
			(module
			  (start $start_f)
			  (type $start_t (func))
			  (func $start_f (type $start_t)
			    unreachable))
		`),
	)
	assert.NoError(t, err)
	defer moduledone(module)

	_, err = NewInstance(module, NewImportObject())
	assert.Error(t, err)
	assert.Equal(t, "unreachable", err.Error())
}
