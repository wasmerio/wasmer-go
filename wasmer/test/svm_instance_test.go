package wasmertest

import (
	"testing"
)

func TestNewSvmInstanceNoImports(t *testing.T) {
	testNewSvmInstanceNoImports(t)
}

func TestNewSvmInstanceWithImports(t *testing.T) {
	testNewSvmInstanceWithImports(t)
}
