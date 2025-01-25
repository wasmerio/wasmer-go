// With Wasmer you'll be able to interact with guest module memory.
//
// This example illustrates the basics of the basics of interacting with Wasm module memory:
//
//  1. How to load a Wasm modules as bytes
//  2. How to compile the module
//  3. How to create an instance of the module
//
// You can run the example directly by executing in Wasmer root:
//
// ```shell
// go test examples/example_memory_test.go
// ```
//
// Ready?
package wasmer

import (
	"fmt"
	"unsafe"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func ExampleMemory() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
		(module
		  (type $mem_size_t (func (result i32)))
		  (type $get_at_t (func (param i32) (result i32)))
		  (type $set_at_t (func (param i32) (param i32)))
		  (memory $mem 1)
		  (func $get_at (type $get_at_t) (param $idx i32) (result i32)
		    (i32.load (local.get $idx)))
		  (func $set_at (type $set_at_t) (param $idx i32) (param $val i32)
		    (i32.store (local.get $idx) (local.get $val)))
		  (func $mem_size (type $mem_size_t) (result i32)
		    (memory.size))
		  (export "get_at" (func $get_at))
		  (export "set_at" (func $set_at))
		  (export "mem_size" (func $mem_size))
		  (export "memory" (memory $mem)))
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

	// The module exports some utility functions, let's get them.
	//
	// These function will be used later in this example.
	memSize, err := instance.Exports.GetFunction("mem_size")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `mem_size` function:", err))
	}

	getAt, err := instance.Exports.GetFunction("get_at")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `get_at` function:", err))
	}

	setAt, err := instance.Exports.GetFunction("set_at")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `set_at` function:", err))
	}

	memory, err := instance.Exports.GetMemory("memory")
	if err != nil {
		panic(fmt.Sprintln("Failed to get the `memory` memory:", err))
	}

	// We now have an instance ready to be used.
	//
	// We will start by querying the most interesting information
	// about the memory: its size. There are mainly two ways of getting
	// this:
	// * the size as a number of `Page`s
	// * the size as a number of bytes
	//
	// The size in bytes can be found either by querying its pages or by
	// querying the memory directly.
	fmt.Println("Querying memory size...")
	size := memory.Size()
	fmt.Println("Memory size (pages):", size)
	fmt.Println("Memory size (pages as bytes):", size.ToBytes())
	fmt.Println("Memory size (bytes):", memory.DataSize())

	// Sometimes, the guest module may also export a function to let you
	// query the memory. Here we have a `mem_size` function, let's try it:
	result, err := memSize()
	if err != nil {
		panic(fmt.Sprintln("Failed to call the `mem_size` function:", err))
	}

	fmt.Println("Memory size (pages):", result)

	// Now that we know the size of our memory, it's time to see how we
	// can change this.
	//
	// A memory can be grown to allow storing more things into it. Let's
	// see how we can do that:
	fmt.Println("Growing memory...")
	memory.Grow(2)
	fmt.Println("New memory size (pages):", memory.Size())

	// Now that we know how to query and adjust the size of the memory,
	// let's see how we can write to it or read from it.
	//
	// We'll only focus on how to do this using exported function, the goal
	// is to show how to work with memory addresses. Here we'll use absolute
	// addresses to write and read a value.
	memAddr := 0x2220
	val := 0xFEFEFFE
	_, err = setAt(memAddr, val)

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `set_at` function:", err))
	}

	result, err = getAt(memAddr)

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `get_at` function:", err))
	}

	fmt.Printf("Value at 0x%x: %d\n", memAddr, result)

	// Now instead of using hard coded memory addresses, let's try to write
	// something at the end of the second memory page and read it.
	pageSize := 0x1_0000
	memAddr = (pageSize * 2) - int(unsafe.Sizeof(val))
	val = 0xFEA09
	_, err = setAt(memAddr, val)

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `set_at` function:", err))
	}

	result, err = getAt(memAddr)

	if err != nil {
		panic(fmt.Sprintln("Failed to call the `get_at` function:", err))
	}

	fmt.Printf("Value at 0x%x: %d\n", memAddr, result)

	// Output:
	// Compiling module...
	// Instantiating module...
	// Querying memory size...
	// Memory size (pages): 1
	// Memory size (pages as bytes): 65536
	// Memory size (bytes): 65536
	// Memory size (pages): 1
	// Growing memory...
	// New memory size (pages): 3
	// Value at 0x2220: 267382782
	// Value at 0x1fff8: 1042953
}
