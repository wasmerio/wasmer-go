// +build static_build

package wasmer

// #cgo !darwin LDFLAGS: -L${SRCDIR} -l:libwasmer.a -lm -ldl -lrt
// #cgo darwin LDFLAGS: -L${SRCDIR} -l:libwasmer.a -lm -ldl
import "C"
