package wasmertest

// // 1️⃣ Declare the `sum` function signature (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
import "C"

import (
	"fmt"
	wasm "github.com/wasmerio/wasmer-go/wasmer"
	"path"
	"runtime"
	"unsafe"
)

func importWasmFile() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "test", "testdata", "examples", "imported_function.wasm")
}

// 2️⃣ Write the implementation of the `sum` function, and export it (for cgo).
//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

func Example_import() {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes(importWasmFile())

	// 3️⃣ Declares the imported functions for WebAssembly.
	imports, _ := wasm.NewImports().Append("sum", sum, C.sum)

	// 4️⃣ Instantiates the WebAssembly module with imports.
	instance, _ := wasm.NewInstanceWithImports(bytes, imports)

	// Close the WebAssembly instance later.
	defer instance.Close()

	// Gets the `add1` exported function from the WebAssembly instance.
	add1 := instance.Exports["add1"]

	// Calls that exported function.
	// As a reminder: `add1` is defined as `fn add1(x: i32, y: i32) -> i32 { sum(x, y) + 1 }`
	result, _ := add1(1, 2)

	// Should output 4 (= 1 + 2 + 1).
	fmt.Println(result)

	// Output:
	// 4
}
