package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryType(t *testing.T) {
	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)
	assert.NoError(t, err)

	memoryType := NewMemoryType(limits)
	limitsAgain := memoryType.Limits()
	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}

func TestMemoryTypeIntoExternTypeAndBack(t *testing.T) {
	var minimum uint32 = 1
	var maximum uint32 = 7
	limits, err := NewLimits(minimum, maximum)
	assert.NoError(t, err)

	memoryType := NewMemoryType(limits)
	externType := memoryType.IntoExternType()
	assert.Equal(t, externType.Kind(), MEMORY)

	memoryTypeAgain := externType.IntoMemoryType()
	limitsAgain := memoryTypeAgain.Limits()
	assert.Equal(t, limitsAgain.Minimum(), minimum)
	assert.Equal(t, limitsAgain.Maximum(), maximum)
}
