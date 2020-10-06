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

func TestValidateModule(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	err := ValidateModule(store, []byte("(module)"))

	assert.NoError(t, err)
}

func TestModuleImports(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
                	(module
                	  (import "ns" "function" (func))
                	  (import "ns" "global" (global f32))
                	  (import "ns" "table" (table 1 2 anyfunc))
                	  (import "ns" "memory" (memory 3 4)))
                `),
	)

	assert.NoError(t, err)

	imports := module.Imports()

	assert.Equal(t, len(imports), 4)

	assert.Equal(t, imports[0].Module(), "ns")
	assert.Equal(t, imports[0].Name(), "function")

	assert.Equal(t, imports[1].Module(), "ns")
	assert.Equal(t, imports[1].Name(), "global")

	assert.Equal(t, imports[2].Module(), "ns")
	assert.Equal(t, imports[2].Name(), "table")

	assert.Equal(t, imports[3].Module(), "ns")
	assert.Equal(t, imports[3].Name(), "memory")
}
