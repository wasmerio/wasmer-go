// A WASM module can import and export entities, like functions, memories, globals
// and tables.
//
// This example illustrates the basics of using these entities.
//
// You can run the example directly by executing in Wasmer root:
//
// ```shell
// go test examples/example_imports_exports_test.go
// ```
//
// Ready?

package wasmer

import (
	"fmt"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func ExampleExtern() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
		(module
		  (func $host_function (import "" "host_function") (result i32))
		  (global $host_global (import "env" "host_global") i32)

		  (func $function (export "guest_function") (result i32) (global.get $global))
		  (global $global (export "guest_global") i32 (i32.const 42))
		  (table $table (export "guest_table") 1 1 funcref)
		  (memory $memory (export "guest_memory") 1))
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
	hostFunction := wasmer.NewFunction(
		store,
		wasmer.NewFunctionType(wasmer.NewValueTypes(), wasmer.NewValueTypes(wasmer.I32)),
		func(args []wasmer.Value) ([]wasmer.Value, error) {
			return []wasmer.Value{wasmer.NewI32(42)}, nil
		},
	)

	fmt.Println("Creating the imported global...")
	hostGlobal := wasmer.NewGlobal(
		store,
		wasmer.NewGlobalType(wasmer.NewValueType(wasmer.I32), wasmer.IMMUTABLE),
		wasmer.NewValue(42, wasmer.I32),
	)

	// Create an import object.
	//
	// Imports are stored in namespaces. We'll need to register each of the
	// namespaces with a name and add the imported entities there.
	//
	// Note that the namespace can also have an empty name.
	//
	// Our module requires us to import:
	//   * A function `host_function` in a namespace with an empty name;
	//   * A global `host_global` in the `env` namespace.
	//
	// Let's do this!
	importObject := wasmer.NewImportObject()
	importObject.Register(
		"",
		map[string]wasmer.IntoExtern{
			"host_function": hostFunction,
		},
	)
	importObject.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"host_global": hostGlobal,
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
	// The WASM module exports some entities:
	//   * A function: `guest_function`
	//   * A global: `guest_global`
	//   * A memory: `guest_memory`
	//   * A table: `guest_table`
	//
	// Let's get them.
	fmt.Println("Getting the exported function...")
	function, err := instance.Exports.GetFunction("guest_function")

	if err != nil {
		panic(fmt.Sprintln("Failed to get the exported function:", err))
	}

	fmt.Printf("Got the exported function: %T\n", function)

	fmt.Println("Getting the exported global...")
	global, err := instance.Exports.GetGlobal("guest_global")

	if err != nil {
		panic(fmt.Sprintln("Failed to get the exported global:", err))
	}

	fmt.Printf("Got the exported global: %T\n", global)

	fmt.Println("Getting the exported memory...")
	memory, err := instance.Exports.GetMemory("guest_memory")

	if err != nil {
		panic(fmt.Sprintln("Failed to get the exported memory:", err))
	}

	fmt.Printf("Got the exported memory: %T\n", memory)

	fmt.Println("Getting the exported table...")
	table, err := instance.Exports.GetTable("guest_table")

	if err != nil {
		panic(fmt.Sprintln("Failed to get the exported table:", err))
	}

	fmt.Printf("Got the exported table: %T\n", table)

	// Output:
	// Compiling module...
	// Creating the imported function...
	// Creating the imported global...
	// Instantiating module...
	// Getting the exported function...
	// Got the exported function: func(...interface {}) (interface {}, error)
	// Getting the exported global...
	// Got the exported global: *wasmer.Global
	// Getting the exported memory...
	// Got the exported memory: *wasmer.Memory
	// Getting the exported table...
	// Got the exported table: *wasmer.Table
}
