package wasmer_test

import (
	"fmt"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"strings"
	"testing"
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
	allocateResult, _ := instance.Exports["allocate"](lengthOfSubject)
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
	deallocate(inputPointer, lengthOfSubject)
	deallocate(outputPointer, lengthOfOutput)

	// Output:
	// Hello, Wasmer 🐹!
}

// Same, but demonstrate an allocation error by using the same string copying logic twice in a row,
// and using a string long enough so we're near the block boundary.
func Test_greetLong(t *testing.T) {
	// Instantiate the module.
	bytes, _ := wasm.ReadBytes(greetWasmFile())
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// Set the subject to greet.
	subject := make([]byte, 1<<16-52)
	for i := 0; i < len(subject); i++ {
		subject[i] = 'x'
	}
	lengthOfSubject := len(subject)

	// Allocate memory for the subject, and get a pointer to it.
	allocateResult, _ := instance.Exports["allocate"](lengthOfSubject)
	inputPointer := allocateResult.ToI32()

	// Write the subject into the memory.
	memory := instance.Memory.Data()[inputPointer:]

	for nth := 0; nth < lengthOfSubject; nth++ {
		memory[nth] = subject[nth]
	}

	// C-string terminates by NULL.
	memory[lengthOfSubject] = 0

	// Allocate and copy the input string a second time.

	// Allocate memory for the subject, and get a pointer to it.
	allocateResult, _ = instance.Exports["allocate"](lengthOfSubject)
	inputPointer2 := allocateResult.ToI32()

	// Write the subject into the memory.
	memory = instance.Memory.Data()[inputPointer2:]

	for nth := 0; nth < lengthOfSubject; nth++ {
		memory[nth] = subject[nth]
	}

	// C-string terminates by NULL.
	memory[lengthOfSubject] = 0

	// Run the `greet` function. Given the pointer to the subject (the *first* input we copied).
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

	// Deallocate the subject, and the output.
	deallocate := instance.Exports["deallocate"]
	deallocate(inputPointer, lengthOfSubject)
	deallocate(inputPointer2, lengthOfSubject)
	deallocate(outputPointer, lengthOfOutput)

	expected := "Hello, " + string(subject) + "!"
	if output.String() != expected {
		t.Errorf("Bad output, got %q", output.String()[:10] + "..." + output.String()[lengthOfOutput-10:])
	}
}
