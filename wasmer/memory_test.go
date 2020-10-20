package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemory(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)

	assert.NoError(t, err)

	_ = NewMemory(store, NewMemoryType(limits))
}

/*func TestMemoryGetType(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)

	assert.NoError(t, err)

	memory := NewMemory(store, NewMemoryType(limits))

	limitsAgain := memory.Type().Limits()

	assert.Equal(t, limitsAgain.Minimum(), 1)
	assert.Equal(t, limitsAgain.Maximum(), 7)
}
*/

func TestMemorySize(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)

	assert.NoError(t, err)

	memory := NewMemory(store, NewMemoryType(limits))
	size := memory.Size()

	assert.Equal(t, size, Pages(1))
}
