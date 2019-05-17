package wasmer_test

import (
	"fmt"
	"path"
	"runtime"
	wasm "wasmer"
)

func simpleWasmFile() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "/../../examples/simple.wasm")
}

func Example() {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes(simpleWasmFile())

	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(bytes)

	// Close the WebAssembly instance later.
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	sum := instance.Exports["sum"]

	// Calls that exported function with Go standard values. The
	// WebAssembly types are infered and values are casted
	// automatically.
	result, _ := sum(1, 2)

	fmt.Print("Result of `sum(1, 2)` is: ")
	fmt.Println(result)

	// Output:
	// Result of `sum(1, 2)` is: 3
}
