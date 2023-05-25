package main

import (
	"flag"
	"fmt"
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
	"io/ioutil"
	"time"
)

var count int

func init() {
	flag.IntVar(&count, "count", 1000000, "call host function count")
}

func main() {
	wasmBytes, _ := ioutil.ReadFile("./target/wasm32-unknown-unknown/release/simple.wasm")

	flag.Parse()

	t0 := time.Now().UnixNano()
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	call_total := 0
	host_func := wasmer.NewFunction(
		store,
		wasmer.NewFunctionType(wasmer.NewValueTypes(), wasmer.NewValueTypes()),
		func(args []wasmer.Value) ([]wasmer.Value, error) {
			call_total += 1
			return []wasmer.Value{}, nil
		},
	)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	importObject.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"host_func": host_func,
		},
	)

	instance, _ := wasmer.NewInstance(module, importObject)

	// Gets the `test_host_func` exported function from the WebAssembly instance.
	test_host_func, _ := instance.Exports.GetFunction("test_host_func")

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := test_host_func(count)

	t1 := time.Now().UnixNano()
	fmt.Println(fmt.Sprintf("result:%d, count:%d, call_total:%d, time:%d ms", result, count, call_total, (t1-t0)/1000000))
}
