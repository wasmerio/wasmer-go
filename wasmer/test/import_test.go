package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/wasmer-go/wasmer"
	"testing"
	"unsafe"
)

func TestInstanceImport(t *testing.T) {
	testInstanceImport(t)
}

func TestInstanceImportMultipleTypes(t *testing.T) {
	testInstanceImportMultipleTypes(t)
}

func TestModuleImport(t *testing.T) {
	testModuleImport(t)
}

func TestInstanceImportMissingImports(t *testing.T) {
	testInstanceImportMissingImports(t)
}

func TestModuleImportMissingImports(t *testing.T) {
	testModuleImportMissingImports(t)
}

func TestImportNoAFunction(t *testing.T) {
	_, err := wasm.NewImports().Append("sum", 42, unsafe.Pointer(nil))

	assert.EqualError(t, err, "Imported function `sum` must be a function; given `int`.")
}

func TestImportMissingInstanceContext(t *testing.T) {
	testImportMissingInstanceContext(t)
}

func TestImportBadTypeForInstanceContext(t *testing.T) {
	testImportBadTypeForInstanceContext(t)
}

func TestImportBadInput(t *testing.T) {
	testImportBadInput(t)
}

func TestImportBadOutput(t *testing.T) {
	testImportBadOutput(t)
}

func TestImportInstanceContext(t *testing.T) {
	testImportInstanceContext(t)
}

func TestImportInstanceContextData(t *testing.T) {
	testImportInstanceContextData(t)
}

func TestWasiImportObject(t *testing.T) {
	testWasiImportObject(t)
}

func TestImportMemory(t *testing.T) {
	testImportMemory(t)
}

func TestImportMemoryIsOwned(t *testing.T) {
	testImportMemoryIsOwned(t)
}
