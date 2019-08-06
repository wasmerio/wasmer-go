package wasmertest

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
import "C"

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"testing"
	"unsafe"
)

type DummyNodeData struct {
}

func NewDummyNodeData() *DummyNodeData {
	return &DummyNodeData{}
}

func NewTestSvmConfig() *wasm.SvmInstanceConfig  {
	addr := [32]byte {0}
	addr[0], addr[1], addr[2] = 10, 20, 30

	maxPages := uint(5)
	maxPagesSlices := uint(100)
	nodeDataPtr := unsafe.Pointer(NewDummyNodeData())
	return wasm.NewSvmInstanceConfig(addr[:], maxPages, maxPagesSlices, nodeDataPtr)
}

func compileModule(t *testing.T, dir string, file string) *wasm.Module {
	bytes := getImportedFunctionBytes(dir, file)

	module, err := wasm.Compile(bytes)
	assert.NoError(t, err)

	return &module
}

func svmInstantiate(t *testing.T, module *wasm.Module, imports *wasm.Imports, cofig *wasm.SvmInstanceConfig) *wasm.Instance {
	instance, err := wasm.NewSvmInstance(module, imports, NewTestSvmConfig())
	assert.NoError(t, err)

	return &instance
}

func testNewSvmInstanceNoImports(t *testing.T) {
	module := compileModule(t, "examples", "simple.wasm")

	instance := svmInstantiate(t, module, wasm.NewImports(), NewTestSvmConfig())

	sum := instance.Exports["sum"]
	output, err := sum(10, 20)
	assert.Equal(t, wasm.I32(30), output)
	assert.NoError(t, err)

	module.Close()
	instance.Close()
}

func testNewSvmInstanceWithImports(t *testing.T) {
	imports, err := wasm.NewImports().Namespace("env").Append("sum", sum, C.sum)
	assert.NoError(t, err)

	module := compileModule(t, "examples", "imported_function.wasm")
	instance := svmInstantiate(t, module, imports, NewTestSvmConfig())

	add1 := instance.Exports["add1"]
	result, err := add1(1, 2)

	assert.Equal(t, wasm.I32(4), result)
	assert.NoError(t, err)
}
