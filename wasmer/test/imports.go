package wasmertest

// #include <stdlib.h>
//
// extern int32_t foo(void *ctx, int32_t x, int32_t y);
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

//export foo
func foo(context unsafe.Pointer, x int32, y int32) int32 {
	fmt.Println(x)
	fmt.Println(y)

	return x + y
}

func testImport(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "examples", "imported_function.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

	instance, err := wasm.NewInstanceWithImports(
		bytes,
		map[string]wasm.Import{
			"sum": wasm.NewImport(foo, C.foo),
		},
	)
	defer instance.Close()

	assert.NoError(t, err)

	add1, exists := instance.Exports["add1"]

	assert.Equal(t, true, exists)

	result, err := add1(1, 2)

	assert.Equal(t, wasm.I32(4), result)
	assert.NoError(t, err)
}
