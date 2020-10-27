package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExternTypeToString(t *testing.T) {
	assert.Equal(t, FUNCTION.String(), "func")
	assert.Equal(t, GLOBAL.String(), "global")
	assert.Equal(t, TABLE.String(), "table")
	assert.Equal(t, MEMORY.String(), "memory")
}
