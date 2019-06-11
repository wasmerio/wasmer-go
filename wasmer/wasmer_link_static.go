// +build static_build

package wasmer

// #cgo LDFLAGS: -lm -ldl -lrt -L${SRCDIR} -l:libwasmer_runtime_c_api.a
import "C"
