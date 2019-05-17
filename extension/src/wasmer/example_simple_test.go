package wasmer

import (
	"fmt"
	"path"
	"runtime"
)

func Example() {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "/../../examples/simple.wasm")

	// Reads the WebAssembly module as bytes.
	bytes, _ := ReadBytes(module_path)

	// Instantiates the WebAssembly module.
	instance, _ := NewInstance(bytes)

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
