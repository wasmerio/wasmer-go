package wasmertest

// #include <stdlib.h>
//
// extern int64_t get_ctx_addr(void *context);
import "C"

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/spacemeshos/go-ext-wasm/wasmer"
	"testing"
	"unsafe"
)

//export get_ctx_addr
func get_ctx_addr(context unsafe.Pointer) int64 {
	instanceContext := (wasm.IntoInstanceContext(context))

	return int64(uintptr((unsafe.Pointer)(instanceContext.Context())))
}

func testInstanceContextAddress(t *testing.T) {
	imports, err := wasm.NewImports().Namespace("env").Append("get_ctx_addr", get_ctx_addr, C.get_ctx_addr)
	assert.NoError(t, err)

	instance, err := wasm.NewInstanceWithImports(getImportedFunctionBytes("examples", "context.wasm"), imports)
	assert.NoError(t, err)
	defer instance.Close()

	ctxAddress := int64(uintptr((unsafe.Pointer)(instance.Context())))

	run := instance.Exports["run"]
	result, err := run()

	assert.NoError(t, err)
	assert.Equal(t, wasm.I64(ctxAddress), result)
}
