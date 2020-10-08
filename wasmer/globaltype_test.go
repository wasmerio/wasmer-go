package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlobalMutability(t *testing.T) {
	assert.Equal(t, IMMUTABLE.String(), "const")
	assert.Equal(t, MUTABLE.String(), "var")
}

func TestGlobalType(t *testing.T) {
	valueType := NewValueType(I32)
	globalType := NewGlobalType(valueType, IMMUTABLE)

	assert.Equal(t, globalType.ValueType().Kind(), I32)
	assert.Equal(t, globalType.Mutability(), IMMUTABLE)
}
