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
	globalType := NewGlobalType(valueType, MUTABLE)
	assert.Equal(t, globalType.ValueType().Kind(), I32)
	assert.Equal(t, globalType.Mutability(), MUTABLE)
}

func TestGlobalTypeIntoExternTypeAndBack(t *testing.T) {
	valueType := NewValueType(I32)

	globalType := NewGlobalType(valueType, MUTABLE)
	externType := globalType.IntoExternType()
	assert.Equal(t, externType.Kind(), GLOBAL)

	globalTypeAgain := externType.IntoGlobalType()
	assert.Equal(t, globalTypeAgain.ValueType().Kind(), I32)
	assert.Equal(t, globalTypeAgain.Mutability(), MUTABLE)
}
