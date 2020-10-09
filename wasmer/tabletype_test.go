package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableType(t *testing.T) {
	valueType := NewValueType(I32)

	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)

	assert.NoError(t, err)

	tableType := NewTableType(valueType, limits)

	valueTypeAgain := tableType.ValueType()

	assert.Equal(t, valueTypeAgain.Kind(), I32)

	limitsAgain := tableType.Limits()

	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}

func TestTableTypeIntoExternTypeAndBack(t *testing.T) {
	valueType := NewValueType(I32)

	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)

	assert.NoError(t, err)

	tableType := NewTableType(valueType, limits)
	externType := tableType.IntoExternType()

	assert.Equal(t, externType.Kind(), TABLE)

	tableTypeAgain := externType.IntoTableType()

	valueTypeAgain := tableTypeAgain.ValueType()

	assert.Equal(t, valueTypeAgain.Kind(), I32)

	limitsAgain := tableTypeAgain.Limits()

	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}
