package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExternTypeToString(t *testing.T) {
	assert.Equal(t, FUNCTION.String(), "func")
	assert.Equal(t, GLOBAL.String(), "global")
	assert.Equal(t, TABLE.String(), "table")
	assert.Equal(t, MEMORY.String(), "memory")
}
