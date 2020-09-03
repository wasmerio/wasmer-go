package wasmer_test

import (
	"fmt"
	wasm "github.com/wasmerio/wasmer-go/wasmer"
	"path"
	"runtime"
)

func GetBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "test", "testdata", "tests.wasm")

	bytes, _ := wasm.ReadBytes(modulePath)

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

func ExampleCompile() {
	// Compiles the bytes into a WebAssembly module.
	module, err := wasm.Compile(GetBytes())
	defer module.Close()

	_ = err
}

func ExampleModule_Instantiate() {
	// Compiles the bytes into a WebAssembly module.
	module, _ := wasm.Compile(GetBytes())
	defer module.Close()

	// Instantiates the WebAssembly module.
	instance, _ := module.Instantiate()
	defer instance.Close()

	// Gets an exported function.
	sum, functionExists := instance.Exports["sum"]

	fmt.Println(functionExists)

	// Calls the `sum` exported function with Go values.
	result, _ := sum(1, 2)

	fmt.Println(result)

	// Output:
	// true
	// 3
}

func ExampleModule_Serialize() {
	// Compiles the bytes into a WebAssembly module.
	module1, _ := wasm.Compile(GetBytes())
	defer module1.Close()

	// Serializes the module into a sequence of bytes.
	serialization, _ := module1.Serialize()

	// Do something with `serialization`.
	// Then later…

	// Deserializes the module.
	module2, _ := wasm.DeserializeModule(serialization)
	defer module2.Close()
	// And enjoy!

	// Instantiates the WebAssembly module.
	instance, _ := module2.Instantiate()
	defer instance.Close()

	// Gets an exported function.
	sum, functionExists := instance.Exports["sum"]

	fmt.Println(functionExists)

	// Calls the `sum` exported function with Go values.
	result, _ := sum(1, 2)

	fmt.Println(result)

	// Output:
	// true
	// 3
}

func ExampleInstance_basic() {
	// Instantiates a WebAssembly instance from bytes.
	instance, err := wasm.NewInstance(GetBytes())
	defer instance.Close()

	_ = err
}

func ExampleInstance_exportedFunctions() {
	// Instantiates a WebAssembly instance from bytes.
	instance, _ := wasm.NewInstance(GetBytes())
	defer instance.Close()

	// Gets an exported function.
	sum, functionExists := instance.Exports["sum"]

	fmt.Println(functionExists)

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
	// Instantiates a WebAssembly instance from bytes.
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
	// Instantiates a WebAssembly instance from bytes.
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
