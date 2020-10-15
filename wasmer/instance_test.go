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

	_, err = NewInstance(module)

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

	instance, err := NewInstance(module)

	assert.NoError(t, err)

	extern, err := instance.Exports.Get("function")

	assert.NoError(t, err)
	assert.Equal(t, extern.Kind(), FUNCTION)
	assert.Equal(t, extern.Type().Kind(), FUNCTION)

	function, err := instance.Exports.GetFunction("function")

	assert.NoError(t, err)
	assert.NotNil(t, function)
}
