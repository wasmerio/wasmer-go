// +build !static_build

package wasmer

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmer
import "C"
