# Compile the Rust part.
rust:
        cargo build --release

# Compile the Go part.
go:
	#!/usr/bin/env bash
	set -euo pipefail
	cd extension
	test -f libwasmer_runtime_c_api.dylib && rm libwasmer_runtime_c_api.dylib
	ln -s ../target/release/deps/libwasmer_runtime_c_api-*.dylib libwasmer_runtime_c_api.dylib
	go build -o ../target/go/wasm -ldflags="-r $(pwd)" src/wasmer/wasmer.go
	echo "Generated in ./target/go/wasm"

# Run a Go program, like `just go-run examples/simple.go`.
go-run FILE:
	GOPATH=$(pwd)/extension go run {{FILE}}

# Generate cgo debug objects.
debug-cgo:
	cd extension/src/wasmer/ && \
		go tool cgo wasmer.go && \
		echo "Browse extension/src/wasmer/_obj/"

# Local Variables:
# mode: makefile
# End:
