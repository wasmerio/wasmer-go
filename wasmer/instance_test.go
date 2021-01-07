package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (func (import "missing" "function"))
			  (func (import "exists" "function")))
		`),
	)
	assert.NoError(t, err)

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
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
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

	_, err = NewInstance(module, NewImportObject())
	assert.Error(t, err)
	assert.Equal(t, "unreachable", err.Error())
}
