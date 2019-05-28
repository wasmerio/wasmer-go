package wasmer_test

import (
	life "github.com/perlin-network/life/exec"
	"github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
)

const N = 100000

func GetBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "benchmarks", "nbody.wasm")

	bytes, _ := wasmer.ReadBytes(modulePath)

	return bytes
}

func BenchmarkWasmer(b *testing.B) {
	wasmBytes := GetBytes()

	for i := 0; i < b.N; i++ {
		instance, _ := wasmer.NewInstance(wasmBytes)
		result, _ := instance.Exports["main"](N)

		_ = result
	}
}

func BenchmarkLife(b *testing.B) {
	wasmBytes := GetBytes()

	for i := 0; i < b.N; i++ {
		vm, _ := life.NewVirtualMachine(wasmBytes, life.VMConfig{}, &life.NopResolver{}, nil)
		entryID, _ := vm.GetFunctionExport("main")
		result, _ := vm.Run(entryID, N)

		_ = result
	}
}
