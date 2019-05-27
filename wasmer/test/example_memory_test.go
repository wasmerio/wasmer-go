package wasmertest

import (
	"fmt"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
)

func memoryWasmFile() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "testdata", "examples", "memory.wasm")
}

func Example_memory() {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes(memoryWasmFile())

	// Instantiates the WebAssembly mdule.
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// Calls the `return_hello` exported function. This function
	// returns a pointer to a string.
	result, _ := instance.Exports["return_hello"]()

	// Gets the pointer value as an integer.
	pointer := result.ToI32()

	// Reads the memory.
	memory := instance.Memory.Data()

	fmt.Println(string(memory[pointer : pointer+13]))

	// Output:
	// Hello, World!
}
