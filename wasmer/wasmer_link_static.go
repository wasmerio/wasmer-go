// +build static_build

package wasmer

// #cgo LDFLAGS: -lm -ldl -lrt -L${SRCDIR} -lwasmer_runtime_c_api
import "C"
