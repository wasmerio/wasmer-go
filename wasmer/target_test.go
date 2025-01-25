package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testTarget(t *testing.T) {
	triple, err := NewTriple("x86_64-apple-darwin")
	assert.NoError(t, err)

	cpuFeatures := NewCpuFeatures()
	cpuFeatures.Add("sse2")

	target := NewTarget(triple, cpuFeatures)

	_ = target
}
