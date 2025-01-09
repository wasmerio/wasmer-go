package wasmer

import (
	"embed"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata
var testData embed.FS

func TestMemory(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)
	assert.NoError(t, err)

	_ = NewMemory(store, NewMemoryType(limits))
}

func TestMemoryGetType(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	limits, err := NewLimits(1, 7)
	assert.NoError(t, err)

	memory := NewMemory(store, NewMemoryType(limits))

	limitsAgain := memory.Type().Limits()
	assert.Equal(t, limitsAgain.Minimum(), uint32(1))
	assert.Equal(t, limitsAgain.Maximum(), uint32(7))
}

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

// This test exercises proper memory management.
// To do so, we run many iterations of "sum" function, while
// invoking GC.
//
// Errors observed prior to this fix include: SIGSEGV, SIGBUS as
// well as simply test failure, followed by panic, where the sum
// function points to an incorrect function
// (Parameters of type [I32, I32] did not match signature [] -> [])
//
// https://github.com/wasmerio/wasmer-go/issues/391
// https://github.com/wasmerio/wasmer-go/issues/364
func TestSumLoop(t *testing.T) {
	//debug.SetGCPercent(1) -- This also reproduces the issue,
	//but having explicit runtime.GC call below does that too.
	e := NewEngineWithConfig(NewConfig().UseCraneliftCompiler())
	s := NewStore(e)

	src, err := testData.ReadFile("testdata/sum.wasm")
	require.NoError(t, err)

	mod, err := NewModule(s, src)
	require.NoError(t, err)

	// Configure WASI_VERSION_SNAPSHOT1 environment
	we, err := NewWasiStateBuilder(mod.Name()).
		CaptureStdout().
		CaptureStderr().
		Finalize()
	require.NoError(t, err)

	imp, err := we.GenerateImportObject(s, mod)
	require.NoError(t, err)

	// Let's instantiate the WebAssembly module.
	instance, err := NewInstance(mod, imp)
	require.NoError(t, err)

	sum, release, err := instance.GetFunctionSafe("sum")
	require.NoError(t, err)
	defer release(instance)

	// This causes the issue to reproduce on the very first iteration
	// if KeepAlive call removed.
	runtime.GC()

	hi := 10240
	n := int32(0)
	for i := range hi {
		res, err := sum(n, i+1)
		//runtime.GC()
		require.NoError(t, err)
		n = res.(int32)
	}
	require.EqualValues(t, hi*(hi+1)/2, n)
}
