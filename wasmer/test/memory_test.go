package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/spacemeshos/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
)

func TestMemoryIsAbsent(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "no_memory.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)
	_, err := wasm.NewInstance(bytes)

	assert.EqualError(t, err, "No memory exported.")
}

func TestMemoryLength(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	var memoryLength uint32 = instance.Memory.Length()

	assert.Equal(t, uint32(0x110000), memoryLength)
}

func TestMemoryDataIsASlice(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	memoryLength := instance.Memory.Length()
	var memory []byte = instance.Memory.Data()

	assert.Equal(t, memoryLength, uint32(len(memory)))
	assert.Equal(t, memoryLength, uint32(cap(memory)))
}

func TestMemoryData(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	result, _ := instance.Exports["string"]()
	pointer := result.ToI32()

	memory := instance.Memory.Data()

	assert.Equal(t, "Hello, World!", string(memory[pointer:pointer+13]))
}

func TestMemoryDataReadWrite(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	result, _ := instance.Exports["string"]()
	pointer := result.ToI32()

	memory1 := instance.Memory.Data()
	memory2 := instance.Memory.Data()

	assert.Equal(t, "Hello, World!", string(memory1[pointer:pointer+13]))
	assert.Equal(t, "Hello, World!", string(memory2[pointer:pointer+13]))

	memory1[pointer] = 'A'

	assert.Equal(t, "Aello, World!", string(memory1[pointer:pointer+13]))
	assert.Equal(t, "Aello, World!", string(memory2[pointer:pointer+13]))

	memory3 := instance.Memory.Data()

	assert.Equal(t, "Aello, World!", string(memory3[pointer:pointer+13]))
}

func TestMemoryGrow(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	memory := instance.Memory
	oldMemoryLength := memory.Length()

	assert.Equal(t, uint32(1114112), oldMemoryLength)

	err := memory.Grow(1)

	assert.NoError(t, err)

	memoryLength := memory.Length()

	assert.Equal(t, uint32(1179648), memoryLength)
	assert.Equal(t, uint32(65536), memoryLength-oldMemoryLength)
}

func TestMemoryGrowTooMuch(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	err := instance.Memory.Grow(100000)

	assert.EqualError(t, err, "Failed to grow the memory:\n    Grow Error: Failed to add pages because would exceed maximum number of pages. Left: 17, Right: 100000, Pages added: 100017")
}
