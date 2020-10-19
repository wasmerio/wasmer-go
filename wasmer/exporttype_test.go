package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportTypeForFunctionType(t *testing.T) {
	params := NewValueTypes(I32, I64)
	results := NewValueTypes(F32)
	functionType := NewFunctionType(params, results)

	name := "foo"
	exportType := NewExportType(name, functionType)

	assert.Equal(t, exportType.Name(), name)

	externType := exportType.Type()

	assert.Equal(t, externType.Kind(), FUNCTION)

	functionTypeAgain := externType.IntoFunctionType()

	assert.Equal(t, len(functionTypeAgain.Params()), len(params))
	assert.Equal(t, len(functionTypeAgain.Results()), len(results))
}

func TestExportTypeForGlobalType(t *testing.T) {
	valueType := NewValueType(I32)
	globalType := NewGlobalType(valueType, MUTABLE)

	name := "foo"
	exportType := NewExportType(name, globalType)

	assert.Equal(t, exportType.Name(), name)

	externType := exportType.Type()

	assert.Equal(t, externType.Kind(), GLOBAL)

	globalTypeAgain := externType.IntoGlobalType()

	assert.Equal(t, globalTypeAgain.ValueType().Kind(), I32)
	assert.Equal(t, globalTypeAgain.Mutability(), MUTABLE)
}

func TestExportTypeForTableType(t *testing.T) {
	valueType := NewValueType(I32)

	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)

	assert.NoError(t, err)

	tableType := NewTableType(valueType, limits)

	name := "foo"
	exportType := NewExportType(name, tableType)

	assert.Equal(t, exportType.Name(), name)

	externType := exportType.Type()

	assert.Equal(t, externType.Kind(), TABLE)

	tableTypeAgain := externType.IntoTableType()
	valueTypeAgain := tableTypeAgain.ValueType()

	assert.Equal(t, valueTypeAgain.Kind(), I32)

	limitsAgain := tableTypeAgain.Limits()

	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}

func TestExportTypeForMemoryType(t *testing.T) {
	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)

	assert.NoError(t, err)

	memoryType := NewMemoryType(limits)

	name := "foo"
	exportType := NewExportType(name, memoryType)

	assert.Equal(t, exportType.Name(), name)

	externType := exportType.Type()

	assert.Equal(t, externType.Kind(), MEMORY)

	memoryTypeAgain := externType.IntoMemoryType()
	limitsAgain := memoryTypeAgain.Limits()

	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}
