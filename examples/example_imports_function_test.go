// A Wasm module can import entities, like functions, memories,
// globals and tables.
//
// This example illustrates how to use imported functions. They come
// in 2 flavors:
//
//   1. Dynamic functions, where parameters and results are of a
//      slice of `Value`,
//   2. Native function, where parameters and results are statically
//      typed Rust values.
//
// You can run the example directly by executing in Wasmer root:
//
// ```shell
// go test examples/example_imports_function_test.go
// ```
//
// Ready?

package wasmer

import (
	"fmt"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func ExampleImportsFunction() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
	(module
		(func $multiply_dynamic (import "env" "multiply_dynamic") (param i32) (result i32))
		(func $multiply_typed (import "env" "multiply_typed") (param i32) (result i32))
		(type $sum_t (func (param i32) (param i32) (result i32)))
		(func $sum_f (type $sum_t) (param $x i32) (param $y i32) (result i32)
		  (call $multiply_dynamic (local.get $x))
		  (call $multiply_typed (local.get $y))
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
		panic(fmt.Sprintln("Failed to compile module:", err))
	}

	// Here we go.
	//
	// Before we can instantiate our module, we need to define
	// the entities we will import.
	//
	// We won't go into details here as creating entities will be
	// covered in more detail in other examples.
	fmt.Println("Creating the imported function...")
	multiply_dynamic := wasmer.NewFunction(
		store,
		wasmer.NewFunctionType(wasmer.NewValueTypes(wasmer.I32), wasmer.NewValueTypes(wasmer.I32)),
		func(args []wasmer.Value) ([]wasmer.Value, error) {
			fmt.Println("Calling `multiply_dynamic`...")

			result := args[0].I32() * 2
			fmt.Println("Result of `multiply_dynamic`: ", result)

			return []wasmer.Value{wasmer.NewI32(result)}, nil
		},
	)

	type MyEnvironment struct{}

	environment := MyEnvironment{}
	multiply := func(environment interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
		fmt.Println("Calling `multiply_typed`...")
		result := args[0].I32() * 3

		fmt.Println("Result of `multiply_typed`: ", result)
		return []wasmer.Value{wasmer.NewI32(result)}, nil
	}
	multiply_typed := wasmer.NewFunctionWithEnvironment(
		store,
		wasmer.NewFunctionType(wasmer.NewValueTypes(wasmer.I32), wasmer.NewValueTypes(wasmer.I32)),
		environment,
		multiply,
	)

	// Create an import object.
	importObject := wasmer.NewImportObject()
	importObject.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"multiply_dynamic": multiply_dynamic,
			"multiply_typed":   multiply_typed,
		},
	)

	fmt.Println("Instantiating module...")
	// Let's instantiate the Wasm module.
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}

	// Here we go.
	//
	// The Wasm module exports a function called `sum`. Let's get it.
	run, err := instance.Exports.GetFunction("sum")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `run` function:", err))
	}

	result, err := run(1, 2)
	if err != nil {
		panic(fmt.Sprintln("`run` met error"))
	}

	fmt.Println("Results of `sum`:", result)

	if result.(int32) != 8 {
		panic("result is wrong")
	}

	// Output:
	// Compiling module...
	// Creating the imported function...
	// Instantiating module...
	// Calling `multiply_dynamic`...
	// Result of `multiply_dynamic`:  2
	// Calling `multiply_typed`...
	// Result of `multiply_typed`:  6
	// Results of `sum`: 8
}
