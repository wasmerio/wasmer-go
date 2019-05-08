package wasmer

import (
	"path"
	"runtime"
	"testing"
)

func GetBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "..", "..", "tests", "tests.wasm")

	bytes, _ := ReadBytes(module_path)

	return bytes
}

func TestValidate(t *testing.T) {
	output := Validate(GetBytes())

	if !output {
		t.Error("Failed to validate the module.")
	}
}
