// A Wasm module can export entities, like functions, memories,
// globals and tables.
//
// This example illustrates how to use exported functions. They come
// in 2 flavors:
//
//   1. Dynamic functions,
//   2. Native function.
//
// You can run the example directly by executing in Wasmer root:
//
// ```shell
// go test examples/example_exports_function_test.go
// ```
//
// Ready?

package wasmer

import (
	"fmt"

	"github.com/iden3/wasmer-go/wasmer"
)

func ExampleFunction() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
		(module
		  (type $sum_t (func (param i32 i32) (result i32)))
		  (func $sum_f (type $sum_t) (param $x i32) (param $y i32) (result i32)
		    local.get $x
		    local.get $y
		    i32.add)
		  (export "sum" (func $sum_f)))
	`)

	// Create an Engine
	engine := wasmer.NewEngine()

	// Create a Store
	store := wasmer.NewStore(engine)

	fmt.Println("Compiling module...")
	module, err := wasmer.NewModule(store, wasmBytes)

	if err != nil {
		fmt.Println("Failed to compile module:", err)
	}

	// Create an empty import object.
	importObject := wasmer.NewImportObject()

	fmt.Println("Instantiating module...")
	// Let's instantiate the Wasm module.
	instance, err := wasmer.NewInstance(module, importObject)

	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}

	// Here we go.
	//
	// The Wasm module exports a function called `sum`. Let's get
	// it.
	sum, err := instance.Exports.GetRawFunction("sum")

	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `sum` function:", err))
	}

	fmt.Println("Calling `sum` function...")
	// Let's call the `sum` exported function.
	result, err := sum.Call(1, 2)

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `sum` function:", err))
	}

	fmt.Println("Result of the `sum` function:", result)

	// That was fun. But what if we can get rid of the `Call` call? Well,
	// that's possible with the native flavor. The function will seem like
	// it's a standard Go function.
	sumNative := sum.Native()

	fmt.Println("Calling `sum` function (natively)...")
	// Let's call the `sum` exported function. The parameters are
	// statically typed Rust values of type `i32` and `i32`. The
	// result, in this case particular case, in a unit of type `i32`.
	result, err = sumNative(3, 4)

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `sum` function natively:", err))
	}

	fmt.Println("Result of the `sum` function:", result)

	// Output:
	// Compiling module...
	// Instantiating module...
	// Calling `sum` function...
	// Result of the `sum` function: 3
	// Calling `sum` function (natively)...
	// Result of the `sum` function: 7
}
