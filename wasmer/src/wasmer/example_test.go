package wasmer_test

import (
	"fmt"
	"path"
	"runtime"
	wasm "wasmer"
)

func GetBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "testdata", "tests.wasm")

	bytes, _ := wasm.ReadBytes(module_path)

	return bytes
}

func ExampleReadBytes() {
	bytes, err := wasm.ReadBytes("my_program.wasm")

	_ = bytes
	_ = err

	// `bytes` now contains the WebAssembly module. They can be
	// used by `NewInstance` or `Validate`.
}

func ExampleValidate() {
	fmt.Println(wasm.Validate(GetBytes()))

	// Output:
	// true
}

func ExampleInstance_basic() {
	// Instantiates the WebAssembly module.
	instance, err := wasm.NewInstance(GetBytes())
	defer instance.Close()

	_ = err
}

func ExampleInstance_exportedFunctions() {
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	// Gets an exported function.
	sum, function_exists := instance.Exports["sum"]

	fmt.Println(function_exists)

	// Calls the `sum` exported function with Go values.
	result, _ := sum(1, 2)

	fmt.Println(result)

	result, _ = sum(wasm.I32(3), wasm.I32(4))

	// Calls the `sum` exported function with WebAssembly values.
	fmt.Println(result)

	// Output:
	// true
	// 3
	// 7
}

func ExampleInstance_memory() {
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	// Calls an exported function.
	result, _ := instance.Exports["string"]()

	// Gets the pointer value as an integer.
	pointer := result.ToI32()

	// Reads 13 bytes as a string from the memory.
	memory := instance.Memory.Data()

	fmt.Println(
		string(memory[pointer : pointer+13]),
	)

	// Output:
	// Hello, World!
}

func ExampleInstance_Close() {
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(GetBytes())

	// Closes/frees the instance (usually used with `defer`).
	defer instance.Close()

}

func ExampleMemory_Data() {
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	// Gets memory data as a slice of bytes.
	memory := instance.Memory.Data()

	// Then, regular slice operations are available.
	_ = memory[1:3]
}

func ExampleMemory_Length() {
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	// Gets memory data length.
	length := instance.Memory.Length()

	fmt.Println(length)

	// Output:
	// 1114112
}
