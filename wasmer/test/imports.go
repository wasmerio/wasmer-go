package wasmertest

// #include <stdlib.h>
//
// extern int32_t sum(void *ctx, int32_t x, int32_t y);
import "C"
import (
	"fmt"
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
	"unsafe"
)

//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	fmt.Println(x)
	fmt.Println(y)

	return x + y
}

func getImportedFunctionBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "examples", "imported_function.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

	return bytes
}

func testImport(t *testing.T) {
	instance, err := wasm.NewInstanceWithImports(
		getImportedFunctionBytes(),
		wasm.NewImports().
			Append("sum", sum, C.sum),
	)
	defer instance.Close()

	assert.NoError(t, err)

	add1, exists := instance.Exports["add1"]

	assert.Equal(t, true, exists)

	result, err := add1(1, 2)

	assert.Equal(t, wasm.I32(4), result)
	assert.NoError(t, err)
}
