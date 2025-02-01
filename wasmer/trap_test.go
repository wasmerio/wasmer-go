package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
