package wasmertest

// #include <stdlib.h>
//
// extern int32_t sum(void *ctx, int32_t x, int32_t y);
import "C"
import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
	"unsafe"
)

//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

func getImportedFunctionBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "examples", "imported_function.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

	return bytes
}

func testImport(t *testing.T) {
	imports, err := wasm.NewImports().Append("sum", sum, C.sum)
	assert.NoError(t, err)

	instance, err := wasm.NewInstanceWithImports(getImportedFunctionBytes(), imports)
	assert.NoError(t, err)

	defer instance.Close()

	add1, exists := instance.Exports["add1"]
	assert.Equal(t, true, exists)

	result, err := add1(1, 2)

	assert.Equal(t, wasm.I32(4), result)
	assert.NoError(t, err)
}
