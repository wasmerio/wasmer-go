package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportTypeForFunctionType(t *testing.T) {
	params := NewValueTypes(I32, I64)
	results := NewValueTypes(F32)
	functionType := NewFunctionType(params, results)

	module := "foo"
	name := "bar"
	importType := NewImportType(module, name, functionType)
	assert.Equal(t, importType.Module(), module)
	assert.Equal(t, importType.Name(), name)

	externType := importType.Type()
	assert.Equal(t, externType.Kind(), FUNCTION)

	functionTypeAgain := externType.IntoFunctionType()
	assert.Equal(t, len(functionTypeAgain.Params()), len(params))
	assert.Equal(t, len(functionTypeAgain.Results()), len(results))
}

func TestImportTypeForGlobalType(t *testing.T) {
	valueType := NewValueType(I32)
	globalType := NewGlobalType(valueType, MUTABLE)

	module := "foo"
	name := "bar"
	importType := NewImportType(module, name, globalType)
	assert.Equal(t, importType.Module(), module)
	assert.Equal(t, importType.Name(), name)

	externType := importType.Type()
	assert.Equal(t, externType.Kind(), GLOBAL)

	globalTypeAgain := externType.IntoGlobalType()
	assert.Equal(t, globalTypeAgain.ValueType().Kind(), I32)
	assert.Equal(t, globalTypeAgain.Mutability(), MUTABLE)
}

func TestImportTypeForTableType(t *testing.T) {
	valueType := NewValueType(I32)

	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)
	assert.NoError(t, err)

	tableType := NewTableType(valueType, limits)

	module := "foo"
	name := "bar"
	importType := NewImportType(module, name, tableType)
	assert.Equal(t, importType.Module(), module)
	assert.Equal(t, importType.Name(), name)

	externType := importType.Type()
	assert.Equal(t, externType.Kind(), TABLE)

	tableTypeAgain := externType.IntoTableType()
	valueTypeAgain := tableTypeAgain.ValueType()
	assert.Equal(t, valueTypeAgain.Kind(), I32)

	limitsAgain := tableTypeAgain.Limits()
	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}

func TestImportTypeForMemoryType(t *testing.T) {
	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)
	assert.NoError(t, err)

	memoryType := NewMemoryType(limits)

	module := "foo"
	name := "bar"
	importType := NewImportType(module, name, memoryType)
	assert.Equal(t, importType.Module(), module)
	assert.Equal(t, importType.Name(), name)

	externType := importType.Type()
	assert.Equal(t, externType.Kind(), MEMORY)

	memoryTypeAgain := externType.IntoMemoryType()
	limitsAgain := memoryTypeAgain.Limits()
	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}
