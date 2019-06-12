package wasmertest

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
// extern int32_t missingContext();
// extern int32_t badInstanceContext(int32_t x);
// extern int32_t badInput(void *context, char x);
// extern char badOutput(void *context);
// extern void logMessage(void *context, int32_t pointer, int32_t length);
// extern void logMessageWithContextData(void *context, int32_t pointer, int32_t length);
import "C"
import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
	"unsafe"
)

func getImportedFunctionBytes(wasmFile ...string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", path.Join(wasmFile...))

	bytes, _ := wasm.ReadBytes(modulePath)

	return bytes
}

//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

func testInstanceImport(t *testing.T) {
	imports, err := wasm.NewImports().Namespace("env").Append("sum", sum, C.sum)
	assert.NoError(t, err)

	instance, err := wasm.NewInstanceWithImports(getImportedFunctionBytes("examples", "imported_function.wasm"), imports)
	defer instance.Close()

	assert.NoError(t, err)

	add1, exists := instance.Exports["add1"]
	assert.Equal(t, true, exists)

	result, err := add1(1, 2)

	assert.Equal(t, wasm.I32(4), result)
	assert.NoError(t, err)
}

func testModuleImport(t *testing.T) {
	imports, err := wasm.NewImports().Namespace("env").Append("sum", sum, C.sum)
	assert.NoError(t, err)

	module, err := wasm.Compile(getImportedFunctionBytes("examples", "imported_function.wasm"))
	defer module.Close()

	assert.NoError(t, err)

	instance, err := module.InstantiateWithImports(imports)
	defer instance.Close()

	assert.NoError(t, err)

	add1, exists := instance.Exports["add1"]
	assert.Equal(t, true, exists)

	result, err := add1(1, 2)

	assert.Equal(t, wasm.I32(4), result)
	assert.NoError(t, err)
}

func testInstanceImportMissingImports(t *testing.T) {
	_, err := wasm.NewInstance(getImportedFunctionBytes("examples", "imported_function.wasm"))

	assert.EqualError(t, err, "Failed to instantiate the module:\n    link error: Import not found, namespace: env, name: sum")
}

func testModuleImportMissingImports(t *testing.T) {
	module, _ := wasm.Compile(getImportedFunctionBytes("examples", "imported_function.wasm"))
	defer module.Close()

	_, err := module.Instantiate()

	assert.EqualError(t, err, "Failed to instantiate the module:\n    error instantiating from module")
}

//export missingContext
func missingContext() int32 {
	return 7
}

func testImportMissingInstanceContext(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", missingContext, C.missingContext)
	assert.EqualError(t, err, "Imported function `foo` must at least have one argument for the instance context.")
}

//export badInstanceContext
func badInstanceContext(x int32) int32 {
	return x + 7
}

func testImportBadTypeForInstanceContext(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", badInstanceContext, C.badInstanceContext)
	assert.EqualError(t, err, "The instance context of the `foo` imported function must be of kind `unsafe.Pointer`; given `int32`; is it missing?")
}

//export badInput
func badInput(context unsafe.Pointer, x C.char) int32 {
	return 7
}

func testImportBadInput(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", badInput, C.badInput)
	assert.EqualError(t, err, "Invalid input type for the `foo` imported function; given `int8`; only accept `int32`, `int64`, `float32`, and `float64`.")
}

//export badOutput
func badOutput(context unsafe.Pointer) C.char {
	return 'a'
}

func testImportBadOutput(t *testing.T) {
	_, err := wasm.NewImports().Append("foo", badOutput, C.badOutput)
	assert.EqualError(t, err, "Invalid output type for the `foo` imported function; given `int8`; only accept `int32`, `int64`, `float32`, and `float64`.")
}

var loggedMessage = ""

//export logMessage
func logMessage(context unsafe.Pointer, pointer int32, length int32) {
	var instanceContext = wasm.IntoInstanceContext(context)
	var memory = instanceContext.Memory().Data()

	loggedMessage = string(memory[pointer : pointer+length])
}

func testImportInstanceContext(t *testing.T) {
	imports, err := wasm.NewImports().Append("log_message", logMessage, C.logMessage)
	assert.NoError(t, err)

	instance, err := wasm.NewInstanceWithImports(getImportedFunctionBytes("log.wasm"), imports)
	assert.NoError(t, err)

	defer instance.Close()

	doSomething, exists := instance.Exports["do_something"]
	assert.Equal(t, true, exists)

	result, err := doSomething()

	assert.Equal(t, wasm.TypeVoid, result.GetType())
	assert.NoError(t, err)
	assert.Equal(t, "hello", loggedMessage)
	loggedMessage = ""
}

//export logMessageWithContextData
func logMessageWithContextData(context unsafe.Pointer, pointer int32, length int32) {
	var instanceContext = wasm.IntoInstanceContext(context)
	var memory = instanceContext.Memory().Data()
	var logMessage = (*logMessageContext)(instanceContext.Data())

	logMessage.message = string(memory[pointer : pointer+length])
}

type logMessageContext struct {
	message string
}

func testImportInstanceContextData(t *testing.T) {
	imports, err := wasm.NewImports().Append("log_message", logMessageWithContextData, C.logMessageWithContextData)
	assert.NoError(t, err)

	instance, err := wasm.NewInstanceWithImports(getImportedFunctionBytes("log.wasm"), imports)
	assert.NoError(t, err)

	defer instance.Close()

	contextData := logMessageContext{message: "first"}
	instance.SetContextData(unsafe.Pointer(&contextData))

	doSomething := instance.Exports["do_something"]

	result, err := doSomething()

	assert.NoError(t, err)
	assert.Equal(t, wasm.TypeVoid, result.GetType())
	assert.Equal(t, "hello", contextData.message)
}
