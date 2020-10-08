package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionType(t *testing.T) {
	params := []*ValueType{NewValueType(I32), NewValueType(I64)}
	results := []*ValueType{NewValueType(F32)}

	functionType := NewFunctionType(params, results)

	assert.Equal(t, len(functionType.Params()), len(params))
	assert.Equal(t, len(functionType.Results()), len(results))
}

func TestFunctionTypeIntoExternTypeAndBack(t *testing.T) {
	params := []*ValueType{NewValueType(I32), NewValueType(I64)}
	results := []*ValueType{NewValueType(F32)}

	functionType := NewFunctionType(params, results)
	externType := functionType.IntoExternType()

	assert.Equal(t, externType.Kind(), FUNCTION)

	functionTypeAgain := externType.IntoFunctionType()

	assert.Equal(t, len(functionTypeAgain.Params()), len(params))
	assert.Equal(t, len(functionTypeAgain.Results()), len(results))
}
