package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestBytes = []byte(`
	(module
	  (global $x (export "x") (mut i32) (i32.const 0))
	  (global $y (export "y") (mut i32) (i32.const 7))
	  (global $z (export "z") i32 (i32.const 42))

	  (func (export "get_x") (result i32)
	    (global.get $x))

	  (func (export "increment_x")
	    (global.set $x
	      (i32.add (global.get $x) (i32.const 1)))))
`)

func testGetGlobalInstance(t *testing.T) *Instance {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, TestBytes)
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	return instance
}

func TestGlobalGetType(t *testing.T) {
	x, err := testGetGlobalInstance(t).Exports.GetGlobal("x")
	assert.NoError(t, err)

	ty := x.Type()
	assert.Equal(t, ty.ValueType().Kind(), I32)
	assert.Equal(t, ty.Mutability(), MUTABLE)
}

func TestGlobalMutable(t *testing.T) {
	exports := testGetGlobalInstance(t).Exports

	x, err := exports.GetGlobal("x")
	assert.NoError(t, err)
	assert.Equal(t, x.Type().Mutability(), MUTABLE)

	y, err := exports.GetGlobal("y")
	assert.NoError(t, err)
	assert.Equal(t, y.Type().Mutability(), MUTABLE)

	z, err := exports.GetGlobal("z")
	assert.NoError(t, err)
	assert.Equal(t, z.Type().Mutability(), IMMUTABLE)
}

func TestGlobalReadWrite(t *testing.T) {
	y, err := testGetGlobalInstance(t).Exports.GetGlobal("y")
	assert.NoError(t, err)

	inititalValue, err := y.Get()
	assert.NoError(t, err)
	assert.Equal(t, int32(7), inititalValue)

	err = y.Set(8, I32)
	assert.NoError(t, err)

	newValue, err := y.Get()
	assert.NoError(t, err)
	assert.Equal(t, int32(8), newValue)
}

func TestGlobalReadWriteAndExportedFunctions(t *testing.T) {
	instance := testGetGlobalInstance(t)
	x, err := instance.Exports.GetGlobal("x")
	assert.NoError(t, err)

	value, err := x.Get()
	assert.NoError(t, err)
	assert.Equal(t, int32(0), value)

	err = x.Set(1, I32)
	assert.NoError(t, err)

	getX, err := instance.Exports.GetFunction("get_x")
	assert.NoError(t, err)

	result, err := getX()
	assert.NoError(t, err)
	assert.Equal(t, int32(1), result)

	incrX, err := instance.Exports.GetFunction("increment_x")
	assert.NoError(t, err)

	_, err = incrX()
	assert.NoError(t, err)

	result, err = getX()
	assert.NoError(t, err)
	assert.Equal(t, int32(2), result)
}

func TestGlobalReadWriteConstant(t *testing.T) {
	z, err := testGetGlobalInstance(t).Exports.GetGlobal("z")
	assert.NoError(t, err)

	value, err := z.Get()
	assert.NoError(t, err)
	assert.Equal(t, int32(42), value)

	err = z.Set(153, I32)
	assert.Error(t, err)

	value, err = z.Get()
	assert.NoError(t, err)
	assert.Equal(t, int32(42), value)
}
