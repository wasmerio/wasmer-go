package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstance(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, []byte("(module)"))

	assert.NoError(t, err)

	_, err = NewInstance(module)

	assert.NoError(t, err)
}
