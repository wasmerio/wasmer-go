// A Wasm module can sometimes not work, it will produce errors in some
// cases.
//
// In this example we'll see how to handle such errors in the most
// basic way. To do that we'll use a Wasm module that we know will
// produce an error.
//
// You can run the example directly by executing in Wasmer root:
//
// ```shell
// go test examples/example_errors_test.go
// ```
//
// Ready?

package wasmer

import (
	"fmt"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func ExampleError() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
		(module
		  (type $do_div_by_zero_t (func (result i32)))
		  (func $do_div_by_zero_f (type $do_div_by_zero_t) (result i32)
		    i32.const 4
		    i32.const 0
		    i32.div_s)
		  (type $div_by_zero_t (func (result i32)))
		  (func $div_by_zero_f (type $div_by_zero_t) (result i32)
		    call $do_div_by_zero_f)
		  (export "div_by_zero" (func $div_by_zero_f)))
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
	// The Wasm module exports a function called `div_by_zero`. As its name
	// implies, this function will try to do a division by zero and thus
	// produce an error.
	//
	// Let's get it.
	divByZero, err := instance.Exports.GetFunction("div_by_zero")

	if err != nil {
		panic(fmt.Sprintln("Failed to get the `div_by_zero` function:", err))
	}

	fmt.Println("Calling `div_by_zero` function...")
	_, err = divByZero()

	// When we call a function it can either succeed or fail. We expect it to fail.
	if err == nil {
		panic(fmt.Sprintln("`div_by_zero` did not error"))
	}

	fmt.Println("Error caught from `div_by_zero`:", err)

	// We'll try to match the error as a trap
	trap, ok := err.(*wasmer.TrapError)

	if !ok {
		panic(fmt.Sprintln("Error was not of the expected type"))
	}

	frames := trap.Trace()
	framesLen := len(frames)

	// Errors come with a trace we can inspect to get more
	// information on the execution flow.
	for index, frame := range frames {
		fmt.Printf("  Frame #%d: function index: %d\n", framesLen-index, frame.FunctionIndex())
	}

	// Output:
	// Compiling module...
	// Instantiating module...
	// Calling `div_by_zero` function...
	// Error caught from `div_by_zero`: integer divide by zero
	//   Frame #2: function index: 0
	//   Frame #1: function index: 50
}
