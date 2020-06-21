// +build static_build

package wasmer

// #cgo !darwin LDFLAGS: -L${SRCDIR}/static -lwasmer -lm -ldl -lrt
// #cgo darwin LDFLAGS: -L${SRCDIR}/static -lwasmer -lm -ldl
import "C"
