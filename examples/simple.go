package main

import (
	"runtime"
	"path"
	"fmt"
	"wasmer"
)

func main(){
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "simple.wasm")

	bytes, _ := wasmer.ReadBytes(module_path)

	fmt.Println(fmt.Sprintf("Is module valid? %t", wasmer.Validate(bytes)))

	instance := wasmer.NewInstance(bytes)
	result := instance.Call("sum")

	fmt.Print("Result of `sum(1, 2)` is: ")
	fmt.Println(result)
}
