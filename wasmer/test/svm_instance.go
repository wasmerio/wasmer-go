package wasmertest

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
// extern void inc(void *context, int32_t x);
// extern int32_t get(void *context);
// extern void copy_from_reg(void *context, int32_t regBits, int32_t regIdx);
// extern void copy_to_reg(void *context, int32_t regBits, int32_t regIdx);
import "C"

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/spacemeshos/go-ext-wasm/wasmer"
	"testing"
	"unsafe"
	"encoding/binary"
)

type DummyNodeData struct {
}

type Counter struct {
	Value int32
}

func NewCounter(initValue int32) *Counter {
	return &Counter { Value: initValue }
}

//export inc
func inc(context unsafe.Pointer, x int32) {
	var instanceContext = wasm.IntoInstanceContext(context)
	counter := (*Counter)(instanceContext.NodeData())
	counter.Value += x
}

//export get
func get(context unsafe.Pointer) int32 {
	var instanceContext = wasm.IntoInstanceContext(context)
	counter := (*Counter)(instanceContext.NodeData())
	return counter.Value
}

//export copy_from_reg
func copy_from_reg(context unsafe.Pointer, regBits int32, regIdx int32) {
	var instanceContext = wasm.IntoInstanceContext(context)
	counter := (*Counter)(instanceContext.NodeData())

	slice := wasm.WasmerSvmRegisterGet(instanceContext, regBits, regIdx)
	value := binary.LittleEndian.Uint32(slice)
	counter.Value = int32(value);
}

//export copy_to_reg
func copy_to_reg(context unsafe.Pointer, regBits int32, regIdx int32) {
	var instanceContext = wasm.IntoInstanceContext(context)
	counter := (*Counter)(instanceContext.NodeData())

	regBuf := make([]byte, 4)
    binary.LittleEndian.PutUint32(regBuf, uint32(counter.Value))

	wasm.WasmerSvmRegisterSet(instanceContext, regBits, regIdx, regBuf)
}

func NewDummyNodeData() unsafe.Pointer {
	return unsafe.Pointer(&DummyNodeData{})
}

func NewTestSvmConfig(nodeDataPtr unsafe.Pointer) *wasm.SvmInstanceConfig {
	addr := [32]byte{0}
	addr[0], addr[1], addr[2] = 10, 20, 30

	state := [32]byte{0}
	state[0], state[1], state[2] = 0xA, 0xB, 0xC

	maxPages := uint(5)
	maxPagesSlices := uint(100)
	config := wasm.NewSvmInstanceConfig(addr[:], state[:], maxPages, maxPagesSlices, nodeDataPtr)

	return config
}

func compileModule(t *testing.T, dir string, file string) *wasm.Module {
	bytes := getImportedFunctionBytes(dir, file)

	module, err := wasm.Compile(bytes)
	assert.NoError(t, err)

	return &module
}

func svmInstantiate(t *testing.T, module *wasm.Module, imports *wasm.Imports, config *wasm.SvmInstanceConfig) *wasm.Instance {
	instance, err := wasm.NewSvmInstance(module, imports, config)
	assert.NoError(t, err)

	return &instance
}

func testNewSvmInstanceNoImports(t *testing.T) {
	module := compileModule(t, "examples", "simple.wasm")
	instance := svmInstantiate(t, module, wasm.NewImports(), NewTestSvmConfig(NewDummyNodeData()))

	sum := instance.Exports["sum"]
	output, err := sum(10, 20)
	assert.Equal(t, wasm.I32(30), output)
	assert.NoError(t, err)

	module.Close()
	instance.Close()
}

func testNewSvmInstanceWithImports(t *testing.T) {
	imports, err := wasm.NewImports().Namespace("env").Append("sum", sum, C.sum)
	imports, err = imports.Append("copy_from_reg", copy_from_reg, C.copy_from_reg)
	assert.NoError(t, err)

	module := compileModule(t, "examples", "imported_function.wasm")
	instance := svmInstantiate(t, module, imports, NewTestSvmConfig(NewDummyNodeData()))

	add1 := instance.Exports["add1"]
	result, err := add1(1, 2)

	assert.Equal(t, wasm.I32(4), result)
	assert.NoError(t, err)
}

func testNewSvmInstanceWithImportsAndNodeData(t *testing.T) {
	imports, err := wasm.NewImports().Namespace("env").Append("inc", inc, C.inc)
	imports, err = imports.Append("get", get, C.get)
	imports, err = imports.Append("copy_from_reg", copy_from_reg, C.copy_from_reg)
	imports, err = imports.Append("copy_to_reg", copy_to_reg, C.copy_to_reg)
	assert.NoError(t, err)

	module := compileModule(t, "examples", "counter.wasm")
	counter := NewCounter(2)

	nodeDataPtr := unsafe.Pointer(counter)
	instance := svmInstantiate(t, module, imports, NewTestSvmConfig(nodeDataPtr))

	inc_and_get := instance.Exports["inc_and_get"]
	result, err := inc_and_get(5)
	assert.Equal(t, int32(7), counter.Value)
	assert.Equal(t, wasm.I32(7), result)
	assert.NoError(t, err)

	inc_and_get = instance.Exports["inc_and_get"]
	result, err = inc_and_get(5)
	assert.Equal(t, int32(12), counter.Value)
	assert.Equal(t, wasm.I32(12), result)
	assert.NoError(t, err)
}

func testNewSvmInstanceWithRegisters(t *testing.T) {
	imports, err := wasm.NewImports().Namespace("env").Append("inc", inc, C.inc)
	imports, err = imports.Append("get", get, C.get)
	imports, err = imports.Append("copy_from_reg", copy_from_reg, C.copy_from_reg)
	imports, err = imports.Append("copy_to_reg", copy_to_reg, C.copy_to_reg)
	assert.NoError(t, err)

	module := compileModule(t, "examples", "counter.wasm")
	counter := NewCounter(2)

	nodeDataPtr := unsafe.Pointer(counter)
	instance := svmInstantiate(t, module, imports, NewTestSvmConfig(nodeDataPtr))

	regBits := int32(64)
	regIndex := int32(2)
	context := instance.Context()
	wasm.WasmerSvmContextRegisterSet(context, regBits, regIndex, []byte{10})

	slice := wasm.WasmerSvmContextRegisterGet(instance.Context(), regBits, regIndex)
	assert.Equal(t, uint32(10), binary.LittleEndian.Uint32(slice))

	copy_reg := instance.Exports["copy_from_reg_and_get"]
	result, err := copy_reg(regBits, regIndex)
	assert.Equal(t, int32(10), counter.Value)
	assert.Equal(t, wasm.I32(10), result)
	assert.NoError(t, err)

	instance.Exports["inc_and_get"](5)

	result, err = instance.Exports["copy_to_reg_and_get"](regIndex)
	assert.Equal(t, int32(15), counter.Value)
	assert.Equal(t, wasm.I32(15), result)
	assert.NoError(t, err)

	slice = wasm.WasmerSvmContextRegisterGet(instance.Context(), regBits, regIndex)
	assert.Equal(t, uint32(15), binary.LittleEndian.Uint32(slice))
}
