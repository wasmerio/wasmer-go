package wasmer

import (
	"io/ioutil"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetBytes(moduleFileName string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", moduleFileName)
	bytes, _ := ioutil.ReadFile(modulePath)

	return bytes
}

func testGetInstance(t *testing.T) (*Instance, func()) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, testGetBytes("tests.wasm"))
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	return instance, func() {
		runtime.KeepAlive(store)
		runtime.KeepAlive(module)
		runtime.KeepAlive(instance)
	}
}
