package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstanceFunctionCall(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (type $sum_t (func (param i32 i32) (result i32)))
			  (func $sum_f (type $sum_t) (param $x i32) (param $y i32) (result i32)
			    local.get $x
			    local.get $y
			    i32.add)
			  (export "sum" (func $sum_f)))
		`),
	)

	assert.NoError(t, err)

	instance, err := NewInstance(module)

	assert.NoError(t, err)

	sum, err := instance.Exports.GetFunction("sum")

	assert.NoError(t, err)

	result, err := sum(1, 2)

	assert.NoError(t, err)
	assert.Equal(t, result, int32(3))
}

func TestInstanceFunctionCallReturnZeroValue(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (type $test_t (func (param i32 i32)))
			  (func $test_f (type $test_t) (param $x i32) (param $y i32))
			  (export "test" (func $test_f)))
		`),
	)

	assert.NoError(t, err)

	instance, err := NewInstance(module)

	assert.NoError(t, err)

	test, err := instance.Exports.GetFunction("test")

	assert.NoError(t, err)

	result, err := test(1, 2)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestInstanceFunctionCallReturnMultipleValues(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (type $swap_t (func (param i32 i64) (result i64 i32)))
			  (func $swap_f (type $swap_t) (param $x i32) (param $y i64) (result i64 i32)
			    local.get $y
			    local.get $x)
			  (export "swap" (func $swap_f)))
		`),
	)

	assert.NoError(t, err)

	instance, err := NewInstance(module)

	assert.NoError(t, err)

	swap, err := instance.Exports.GetFunction("swap")

	assert.NoError(t, err)

	results, err := swap(41, 42)

	assert.NoError(t, err)
	assert.Equal(t, results, []interface{}{int64(42), int32(41)})
}
