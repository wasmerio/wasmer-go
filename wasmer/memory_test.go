package wasmer_test

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
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
