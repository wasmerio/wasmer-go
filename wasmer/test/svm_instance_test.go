package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"testing"
	"unsafe"
)

type DummyNodeData struct {
}

func TestNewSvmInstanceNoImports(t *testing.T) {
	bytes := getImportedFunctionBytes("examples", "simple.wasm")

	module, err := wasm.Compile(bytes)
	assert.NoError(t, err)
	defer module.Close()

	addr := [32]byte {0}
	addr[0], addr[1], addr[2] = 10, 20, 30

	maxPages := uint(5)
	maxPagesSlices := uint(100)
	nodeDataPtr := unsafe.Pointer(&DummyNodeData{})
	config := wasm.NewSvmInstanceConfig(addr[:], maxPages, maxPagesSlices, nodeDataPtr)
	imports := wasm.NewImports()

	instance, err := wasm.NewSvmInstance(&module, imports, config)
	assert.NoError(t, err)
	defer instance.Close()

	sum := instance.Exports["sum"]
	output, err := sum(10, 20)
	assert.Equal(t, wasm.I32(30), output)
	assert.NoError(t, err)
}
