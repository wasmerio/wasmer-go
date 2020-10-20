package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrap(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	message := "Hello"
	trap := NewTrap(store, message)

	assert.NotNil(t, trap)
	assert.Equal(t, message, trap.Message())
	assert.Nil(t, trap.Origin())
	assert.NotNil(t, trap.Trace())
	assert.Len(t, trap.Trace().frames, 0)
}
