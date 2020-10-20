package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemory(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)

	assert.NoError(t, err)

	_ = NewMemory(store, NewMemoryType(limits))
}

/*func TestMemoryGetType(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)

	assert.NoError(t, err)

	memory := NewMemory(store, NewMemoryType(limits))

	limitsAgain := memory.Type().Limits()

	assert.Equal(t, limitsAgain.Minimum(), 1)
	assert.Equal(t, limitsAgain.Maximum(), 7)
}
*/

func TestMemorySize(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 3)

	assert.NoError(t, err)

	memory := NewMemory(store, NewMemoryType(limits))
	size := memory.Size()

	assert.Equal(t, size, Pages(1))

	success := memory.Grow(Pages(2))

	assert.True(t, success)

	size = memory.Size()

	assert.Equal(t, size, Pages(3))

	success = memory.Grow(Pages(1))

	assert.False(t, success)
}

func TestMemoryDataSize(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)

	assert.NoError(t, err)

	memory := NewMemory(store, NewMemoryType(limits))
	dataSize := memory.DataSize()

	assert.Equal(t, dataSize, uint(0x10000))
}

func TestMemoryData(t *testing.T) {
	instance := testGetInstance(t)

	getString, err := instance.Exports.GetFunction("string")

	assert.NoError(t, err)

	ptr, err := getString()

	assert.NoError(t, err)
	assert.Equal(t, ptr, int32(1048576))

	pointer := ptr.(int32)

	memory1, err := instance.Exports.GetMemory("memory")

	assert.NoError(t, err)

	memory2, err := instance.Exports.GetMemory("memory")

	data1 := memory1.Data()
	data2 := memory2.Data()

	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", string(data1[pointer:pointer+13]))
	assert.Equal(t, "Hello, World!", string(data2[pointer:pointer+13]))

	data1[pointer] = 'A'

	assert.Equal(t, "Aello, World!", string(data1[pointer:pointer+13]))
	assert.Equal(t, "Aello, World!", string(data2[pointer:pointer+13]))

}
