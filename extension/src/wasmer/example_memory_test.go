package wasmer

import (
	"fmt"
	"path"
	"runtime"
)

func Example_memory() {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "/../../examples/memory.wasm")

	bytes, _ := ReadBytes(module_path)
	instance, _ := NewInstance(bytes)
	defer instance.Close()

	result, _ := instance.Exports["return_hello"]()
	pointer := result.ToI32()

	memory := instance.Memory.Data()

	fmt.Println(string(memory[pointer : pointer+13]))

	// Output:
	// Hello, World!
}
