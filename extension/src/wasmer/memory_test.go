package wasmer_test

import (
	"github.com/stretchr/testify/assert"
	"path"
	"runtime"
	"testing"
	wasm "wasmer"
)

func TestMemoryIsAbsent(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "no_memory.wasm")

	bytes, _ := wasm.ReadBytes(module_path)
	_, err := wasm.NewInstance(bytes)

	assert.EqualError(t, err, "No memory exported.")
}

func TestMemoryLength(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	var memory_length uint32 = instance.Memory.Length()

	assert.Equal(t, uint32(0x110000), memory_length)
}

func TestMemoryDataIsASlice(t *testing.T) {
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	memory_length := instance.Memory.Length()
	var memory []byte = instance.Memory.Data()

	assert.Equal(t, memory_length, uint32(len(memory)))
	assert.Equal(t, memory_length, uint32(cap(memory)))
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

	memory_1 := instance.Memory.Data()
	memory_2 := instance.Memory.Data()

	assert.Equal(t, "Hello, World!", string(memory_1[pointer:pointer+13]))
	assert.Equal(t, "Hello, World!", string(memory_2[pointer:pointer+13]))

	memory_1[pointer] = 'A'

	assert.Equal(t, "Aello, World!", string(memory_1[pointer:pointer+13]))
	assert.Equal(t, "Aello, World!", string(memory_2[pointer:pointer+13]))

	memory_3 := instance.Memory.Data()

	assert.Equal(t, "Aello, World!", string(memory_3[pointer:pointer+13]))
}
