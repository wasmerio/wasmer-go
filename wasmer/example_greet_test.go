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
	return path.Join(path.Dir(filename), "test", "testdata", "examples", "greet.wasm")
}

func Example_greet() {
	// Instantiate the module.
	bytes, _ := wasm.ReadBytes(greetWasmFile())
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// Set the subject to greet.
	subject := "Wasmer 🐹"
	lengthOfSubject := len(subject)

	// Allocate memory for the subject, and get a pointer to it.
	allocateResult, _ := instance.Exports["allocate"](lengthOfSubject+1)
	inputPointer := allocateResult.ToI32()

	// Write the subject into the memory.
	memory := instance.Memory.Data()[inputPointer:]

	for nth := 0; nth < lengthOfSubject; nth++ {
		memory[nth] = subject[nth]
	}

	// C-string terminates by NULL.
	memory[lengthOfSubject] = 0

	// Run the `greet` function. Given the pointer to the subject.
	greetResult, _ := instance.Exports["greet"](inputPointer)
	outputPointer := greetResult.ToI32()

	// Read the result of the `greet` function.
	memory = instance.Memory.Data()[outputPointer:]
	nth := 0
	var output strings.Builder

	for {
		if memory[nth] == 0 {
			break
		}

		output.WriteByte(memory[nth])
		nth++
	}

	lengthOfOutput := nth

	fmt.Println(output.String())

	// Deallocate the subject, and the output.
	deallocate := instance.Exports["deallocate"]
	deallocate(inputPointer, lengthOfSubject+1)
	deallocate(outputPointer, lengthOfOutput+1)

	// Output:
	// Hello, Wasmer 🐹!
}
