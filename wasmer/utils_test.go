package wasmer

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func testGetBytes() []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "tests.wasm")
	bytes, _ := ioutil.ReadFile(modulePath)

	return bytes
}

func testGetInstance(t *testing.T) *Instance {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(store, testGetBytes())

	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())

	assert.NoError(t, err)

	return instance
}
