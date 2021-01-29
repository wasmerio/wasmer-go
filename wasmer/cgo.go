package wasmer

// #cgo CFLAGS: -I${SRCDIR}/packaged/include
// #cgo LDFLAGS: -lwasmer
//
// #cgo linux,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/packaged/lib/linux-amd64 -L${SRCDIR}/packaged/lib/linux-amd64
// #cgo darwin,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/packaged/lib/darwin-amd64 -L${SRCDIR}/packaged/lib/darwin-amd64
// #cgo windows,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/packaged/lib/windows-amd64 -L${SRCDIR}/packaged/lib/windows-amd64
//
// #include <wasmer_wasm.h>
import "C"
