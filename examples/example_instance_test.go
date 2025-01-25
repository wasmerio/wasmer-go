// Wasmer will let you easily run Wasm module in a Rust host.
//
// This example illustrates the basics of using Wasmer through a "Hello World"-like project:
//
//   1. How to load a Wasm modules as bytes
//   2. How to compile the module
//   3. How to create an instance of the module
//
// You can run the example directly by executing in Wasmer root:
//
// ```shell
// go test examples/example_instance_test.go
// ```
//
// Ready?

package wasmer

import (
	"fmt"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func ExampleInstance() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
		(module
		  (type $add_one_t (func (param i32) (result i32)))
		  (func $add_one_f (type $add_one_t) (param $value i32) (result i32)
		    local.get $value
		    i32.const 1
		    i32.add)
		  (export "add_one" (func $add_one_f)))
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

	// We now have an instance ready to be used.
	//
	// From an `Instance` we can fetch any exported entities from the Wasm module.
	// Each of these entities is covered in others examples.
	//
	// Here we are fetching an exported function. We won't go into details here
	// as the main focus of this example is to show how to create an instance out
	// of a Wasm module and have basic interactions with it.
	addOne, err := instance.Exports.GetFunction("add_one")
	if err != nil {
		panic(fmt.Sprintln("Failed to get the `add_one` function:", err))
	}

	fmt.Println("Calling `add_one` function...")
	result, err := addOne(1)
	if err != nil {
		panic(fmt.Sprintln("Failed to call the `add_one` function:", err))
	}

	fmt.Println("Results of `add_one`:", result)

	// Output:
	// Compiling module...
	// Instantiating module...
	// Calling `add_one` function...
	// Results of `add_one`: 2
}
