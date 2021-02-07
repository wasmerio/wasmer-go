// +build customlib

package wasmer

// User is expected to provide the CGO_LDFLAGS and CGO_CFLAGS that point to appropriate wasmer build directories
// for example:
// CGO_CFLAGS="-I/wasmer/lib/c-api/"
// CGO_LDFLAGS="-Wl,-rpath,/wasmer/target/x86_64-unknown-linux-musl/release/ -L/wasmer/target/x86_64-unknown-linux-musl/release/ -pthread -lwasmer_c_api -lm -ldl -static"
//
//
// #include <wasmer_wasm.h>
import "C"
