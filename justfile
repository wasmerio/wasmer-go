# Compile the Rust part.
rust:
        cargo build --release

# Compile the Go part.
go:
	#!/usr/bin/env bash
	set -euo pipefail
	cd extension
	case "{{os()}}" in
		"macos")
			dylib_extension="dylib"
			;;
		"windows")
			dylib_extension="dll"
			;;
		*)
			dylib_extension="so"
	esac
	test -f libwasmer_runtime_c_api.${dylib_extension} && rm libwasmer_runtime_c_api.${dylib_extension}
	ln -s ../target/release/deps/libwasmer_runtime_c_api-*.${dylib_extension} libwasmer_runtime_c_api.${dylib_extension}
	go build -o ../target/go/wasm -ldflags="-r $(pwd)" src/wasmer/*.go

# Run a Go program, like `just go-run examples/simple.go`.
go-run FILE:
	GOPATH=$(pwd)/extension go run {{FILE}}

# Generate cgo debug objects.
debug-cgo:
	cd extension/src/wasmer/ && \
		go tool cgo wasmer.go && \
		echo "Browse extension/src/wasmer/_obj/"

# Run the tests.
test:
	GOPATH=$(pwd)/extension go test -ldflags="-r $(pwd)" wasmer -test.v

# Local Variables:
# mode: makefile
# End:
