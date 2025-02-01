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
// go test examples/example_early_exit_test.go
// ```
//
// Ready?

package wasmer

import (
	"fmt"
	"runtime"

	"github.com/wasmerio/wasmer-go/wasmer"
)

type exitCode struct {
	code int32
}

func (self *exitCode) Error() string {
	return fmt.Sprintf("exit code: %d", self.code)
}

func earlyExit(args []wasmer.Value) ([]wasmer.Value, error) {
	return nil, &exitCode{1}
}

func ExampleFunction_Call() {
	// Let's declare the Wasm module.
	//
	// We are using the text representation of the module here.
	wasmBytes := []byte(`
		(module
		  (type $run_t (func (param i32 i32) (result i32)))
		  (type $early_exit_t (func (param) (result)))
		  (import "env" "early_exit" (func $early_exit (type $early_exit_t)))
		  (func $run (type $run_t) (param $x i32) (param $y i32) (result i32)
		    (call $early_exit)
		    (i32.add
		        local.get $x
		        local.get $y))
  		(export "run" (func $run)))
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

	// Create an import object with the expected function.
	importObject := wasmer.NewImportObject()
	importObject.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"early_exit": wasmer.NewFunction(
				store,
				wasmer.NewFunctionType(wasmer.NewValueTypes(), wasmer.NewValueTypes()),
				earlyExit,
			),
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
	// Get the `run` function which we'll use as our entrypoint.
	fmt.Println("Calling `run` function...")
	run, err := instance.Exports.GetFunction("run")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `run` function:", err))
	}

	_, err = run(1, 7)

	if err == nil {
		panic(fmt.Sprintln("`run` did not error"))
	}

	fmt.Println("Exited early with:", err)

	// These lines are here to ensure that SetFinalizer works correctly
	runtime.GC()
	runtime.GC()

	// Output:
	// Compiling module...
	// Instantiating module...
	// Calling `run` function...
	// Exited early with: exit code: 1
}
