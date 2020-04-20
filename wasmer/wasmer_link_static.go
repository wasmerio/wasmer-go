// +build static_build

package wasmer

// #cgo !darwin LDFLAGS: -lm -ldl -lrt -L${SRCDIR} -l:libwasmer_runtime_c_api.a
// #cgo darwin LDFLAGS: -lm -ldl -L${SRCDIR} -l:libwasmer_runtime_c_api.a
import "C"
