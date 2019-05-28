package wasmer_test

import (
	"bytes"
	"fmt"
	wagon_exec "github.com/go-interpreter/wagon/exec"
	wagon_wasm "github.com/go-interpreter/wagon/wasm"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

const N = 100000

func GetBytes(wasmFile string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "benchmarks", wasmFile)

	bytes, _ := ioutil.ReadFile(modulePath)

	return bytes
}

func benchmarkWagon(b *testing.B, wasmFile string, exportName int64, exportValues ...uint64) {
	wasmBytes := GetBytes(wasmFile)

	for i := 0; i < b.N; i++ {
		module, _ := wagon_wasm.ReadModule(bytes.NewReader(wasmBytes), func(name string) (*wagon_wasm.Module, error) {
			return nil, fmt.Errorf("Module %q unknown.", name)
		})
		vm, _ := wagon_exec.NewVM(module)
		result, _ := vm.ExecCode(exportName, exportValues...)

		_ = result
	}
}

func Benchmark_Wagon_Nbody(b *testing.B) {
	benchmarkWagon(b, "nbody.wasm", 1, uint64(N))
}

func Benchmark_Wagon_Fibonacci_Recursive(b *testing.B) {
	benchmarkWagon(b, "fib_recursive.wasm", 1)
}

func Benchmark_Wagon_Pollard_Rho_128(b *testing.B) {
	benchmarkWagon(b, "pollard_rho_128.wasm", 0)
}

func Benchmark_Wagon_Snappy_Compress(b *testing.B) {
	benchmarkWagon(b, "snappy_compress.wasm", 0)
}
