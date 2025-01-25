// A Wasm module can export entities, like functions, memories,
// globals and tables.
//
// This example illustrates how to use exported globals. They come
// in 2 flavors:
//
//  1. Immutable globals (const),
//  2. Mutable globals.
//
// You can run the example directly by executing in Wasmer root:
//
// ```shell
// go test examples/example_exports_global_test.go
// ```
//
// Ready?
package wasmer

import (
	"fmt"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func ExampleGlobal_Set() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
		(module
		  (global $one (export "one") f32 (f32.const 1))
		  (global $some (export "some") (mut f32) (f32.const 0))

		  (func (export "get_one") (result f32) (global.get $one))
		  (func (export "get_some") (result f32) (global.get $some))

		  (func (export "set_some") (param f32) (global.set $some (local.get 0))))
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
	// The Wasm module exports some globals. Let's get them.
	one, err := instance.Exports.GetGlobal("one")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `one` global:", err))
	}

	some, err := instance.Exports.GetGlobal("some")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `some` global:", err))
	}

	fmt.Println("Getting globals types information...")
	// Let's get the globals types. The results are `GlobalType`s.
	oneType := one.Type()
	someType := some.Type()

	fmt.Printf("`one` type: %s %s\n", oneType.Mutability(), oneType.ValueType().Kind().String())
	fmt.Printf("`some` type: %s %s\n", someType.Mutability(), someType.ValueType().Kind().String())

	fmt.Println("Getting global values...")
	// Getting the values of globals can be done in two ways:
	//   1. Through an exported function,
	//   2. Using the Global API directly.
	//
	// We will use an exported function for the `one` global
	// and the Global API for `some`.
	getOne, err := instance.Exports.GetFunction("get_one")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `get_one` function:", err))
	}

	oneValue, err := getOne()
	if err != nil {
		panic(fmt.Sprintln("Failed to call the `get_one` function:", err))
	}

	someValue, err := some.Get()
	if err != nil {
		panic(fmt.Sprintln("Failed to get the `some` global value:", err))
	}

	fmt.Printf("`one` value: %.1f\n", oneValue)
	fmt.Printf("`some` value: %.1f\n", someValue)

	fmt.Println("Setting global values...")
	// Trying to set the value of a immutable global (`const`)
	// will result in a `RuntimeError`.
	err = one.Set(float32(42.0), wasmer.F32)

	if err == nil {
		panic(fmt.Sprintln("Setting value to `one` did not error"))
	}

	oneValue, err = getOne()

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `get_one` function:", err))
	}

	fmt.Printf("`one` value: %.1f\n", oneValue)

	// Setting the values of globals can be done in two ways:
	//   1. Through an exported function,
	//   2. Using the Global API directly.
	//
	// We will use both for the `some` global.
	setSome, err := instance.Exports.GetFunction("set_some")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `set_some` function:", err))
	}

	_, err = setSome(float32(21.0))

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `set_some` function:", err))
	}

	someValue, err = some.Get()

	if err != nil {
		panic(fmt.Sprintln("Failed to get the `some` global value:", err))
	}

	fmt.Printf("`some` value after `set_some`: %.1f\n", someValue)

	err = some.Set(float32(42.0), wasmer.F32)

	if err != nil {
		panic(fmt.Sprintln("Failed to set the `some` global value:", err))
	}

	someValue, err = some.Get()

	if err != nil {
		panic(fmt.Sprintln("Failed to get the `some` global value:", err))
	}

	fmt.Printf("`some` value after `set`: %.1f\n", someValue)

	// Output:
	// Compiling module...
	// Instantiating module...
	// Getting globals types information...
	// `one` type: const f32
	// `some` type: var f32
	// Getting global values...
	// `one` value: 1.0
	// `some` value: 0.0
	// Setting global values...
	// `one` value: 1.0
	// `some` value after `set_some`: 21.0
	// `some` value after `set`: 42.0
}
