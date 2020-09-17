package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModule(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	_, err := NewModule(store, []byte("(module)"))

	assert.NoError(t, err)
}
