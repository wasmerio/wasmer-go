package main

import (
	"fmt"
	"path"
	"runtime"
	wasm "wasmer"
)

func main(){
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "simple.wasm")

	bytes, _ := wasm.ReadBytes(module_path)

	fmt.Println(fmt.Sprintf("Is module valid? %t", wasm.Validate(bytes)))

	instance, _ := wasm.NewInstance(bytes)
	result, _ := instance.Call(
		"sum",
		[]wasm.Value{
			wasm.ValueI32(1),
			wasm.ValueI32(2),
		},
	)

	fmt.Print("Result of `sum(1, 2)` is: ")
	fmt.Println(result)
}
