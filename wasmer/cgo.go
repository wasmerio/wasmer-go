package wasmer

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/include/ -L${SRCDIR}/include/ -lwasmer_c_api
// #include <wasmer_wasm.h>
import "C"
