package wasmer_test

import (
	"fmt"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"strings"
)

func greetWasmFile() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "testdata", "examples", "greet.wasm")
}

func Example_greet() {
	// Instantiate the module.
	bytes, _ := wasm.ReadBytes(greetWasmFile())
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// Set the subject to greet.
	subject := "Wasmer 🐹"
	length_of_subject := len(subject)

	// Allocate memory for the subject, and get a pointer to it.
	allocate_result, _ := instance.Exports["allocate"](length_of_subject)
	input_pointer := allocate_result.ToI32()

	// Write the subject into the memory.
	memory := instance.Memory.Data()[input_pointer:]

	for nth := 0; nth < length_of_subject; nth++ {
		memory[nth] = subject[nth]
	}

	// C-string terminates by NULL.
	memory[length_of_subject] = 0

	// Run the `greet` function. Given the pointer to the subject.
	greet_result, _ := instance.Exports["greet"](input_pointer)
	output_pointer := greet_result.ToI32()

	// Read the result of the `greet` function.
	memory = instance.Memory.Data()[output_pointer:]
	nth := 0
	var output strings.Builder

	for {
		if memory[nth] == 0 {
			break
		}

		output.WriteByte(memory[nth])
		nth++
	}

	length_of_output := nth

	fmt.Println(output.String())

	// Deallocate the subject, and the output.
	deallocate := instance.Exports["deallocate"]
	deallocate(input_pointer, length_of_subject)
	deallocate(output_pointer, length_of_output)

	// Output:
	// Hello, Wasmer 🐹!
}
