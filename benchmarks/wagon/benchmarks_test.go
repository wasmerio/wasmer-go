package wasmer_test

import (
	"bytes"
	"fmt"
	wagon_exec "github.com/go-interpreter/wagon/exec"
	wagon_wasm "github.com/go-interpreter/wagon/wasm"
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

func BenchmarkWagon(b *testing.B) {
	wasmBytes := GetBytes()

	for i := 0; i < b.N; i++ {
		module, _ := wagon_wasm.ReadModule(bytes.NewReader(wasmBytes), func(name string) (*wagon_wasm.Module, error) {
			return nil, fmt.Errorf("Module %q unknown.", name)
		})
		vm, _ := wagon_exec.NewVM(module)
		result, _ := vm.ExecCode(1, uint64(N))

		_ = result
	}
}
