package main

import (
	"fmt"
	"path"
	"runtime"
	wasm "wasmer"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "memory.wasm")

	bytes, _ := wasm.ReadBytes(module_path)
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	result, _ := instance.Exports["return_hello"]()
	pointer := result.ToI32()

	memory := instance.Memory.Data()

	fmt.Println(string(memory[pointer : pointer+13])) // Hello, World!
}
