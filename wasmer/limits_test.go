package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLimitMaxDefault(t *testing.T) {
	assert.Equal(t, LimitMaxUnbound(), uint32(0xffffffff))
}

func TestLimits(t *testing.T) {
	var minimum uint32 = 1
	var maximum uint32 = 7

	limits, err := NewLimits(minimum, maximum)
	assert.NoError(t, err)
	assert.Equal(t, limits.Minimum(), minimum)
	assert.Equal(t, limits.Maximum(), maximum)
}
