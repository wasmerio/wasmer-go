package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"testing"
	"unsafe"
)

type DummyNodeData struct {
}

func TestNewSvmInstance(t *testing.T) {
	bytes := getImportedFunctionBytes("examples", "greet.wasm")
	// bytes := getImportedFunctionBytes("examples", "imported_function.wasm")

	module, err := wasm.Compile(bytes)
	assert.NoError(t, err)

	addr := [32]byte{}
	addr[0] = 10
	addr[1] = 20
	addr[2] = 30

	maxPages := uint(5)
	maxPagesSlices := uint(100)
	nodeDataPtr := unsafe.Pointer(&DummyNodeData{})
	config := wasm.NewSvmInstanceConfig(addr[:], maxPages, maxPagesSlices, nodeDataPtr)
	imports := wasm.NewImports()

	// wasm.NewInstanceFromModule(&module, imports)
	wasm.NewSvmInstance(&module, imports, config)
}
