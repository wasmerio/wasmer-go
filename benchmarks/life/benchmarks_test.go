package wasmer_test

import (
	life "github.com/perlin-network/life/exec"
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

func Benchmark_Life_Nbody(b *testing.B) {
	benchmarkLife(b, "nbody.wasm", "main", N)
}

func Benchmark_Life_Fibonacci_Recursive(b *testing.B) {
	benchmarkLife(b, "fib_recursive.wasm", "app_main")
}

func Benchmark_Life_Pollard_Rho_128(b *testing.B) {
	benchmarkLife(b, "pollard_rho_128.wasm", "app_main")
}

func Benchmark_Life_Snappy_Compress(b *testing.B) {
	benchmarkLife(b, "snappy_compress.wasm", "app_main")
}

func benchmarkLife(b *testing.B, wasmFile string, exportName string, exportValues ...int64) {
	wasmBytes := GetBytes(wasmFile)

	for i := 0; i < b.N; i++ {
		vm, _ := life.NewVirtualMachine(wasmBytes, life.VMConfig{}, &life.NopResolver{}, nil)
		entryID, _ := vm.GetFunctionExport(exportName)
		result, _ := vm.Run(entryID, exportValues...)

		_ = result
	}
}
