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
