package wasmertest

// #include <stdlib.h>
//
// extern int32_t sum(void *ctx, int32_t x, int32_t y);
// extern int32_t missingContext();
// extern int32_t badInstanceContext(int32_t x);
// extern int32_t badInput(void *ctx, char x);
// extern char badOutput(void *ctx);
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

//export missingContext
func missingContext() int32 {
	return 7
}

//export badInstanceContext
func badInstanceContext(x int32) int32 {
	return x + 7
}

//export badInput
func badInput(context unsafe.Pointer, x C.char) int32 {
	return 7
}

//export badOutput
func badOutput(context unsafe.Pointer) C.char {
	return 'a'
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

func testImportMissingInstanceContext(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", missingContext, C.missingContext)
	assert.EqualError(t, err, "Imported function `foo` must at least have one argument for the instance context.")
}

func testImportBadTypeForInstanceContext(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", badInstanceContext, C.badInstanceContext)
	assert.EqualError(t, err, "The instance context of the `foo` imported function must be of kind `unsafe.Pointer`; given `int32`; is it missing?")
}

func testImportBadInput(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", badInput, C.badInput)
	assert.EqualError(t, err, "Invalid input type for the `foo` imported function; given `int8`; only accept `int32`, `int64`, `float32`, and `float64`.")
}

func testImportBadOutput(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", badOutput, C.badOutput)
	assert.EqualError(t, err, "Invalid output type for the `foo` imported function; given `int8`; only accept `int32`, `int64`, `float32`, and `float64`.")
}
