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
	go build -o ../target/go/wasm -ldflags="-r $(pwd)" src/wasmer/

# Run a Go program, like `just go-run examples/simple.go`.
go-run FILE:
	GOPATH=$(pwd)/extension go run {{FILE}}

# Generate cgo debug objects.
debug-cgo:
	cd extension/src/wasmer/ && \
		go tool cgo wasmer.go && \
		echo "Browse extension/src/wasmer/_obj/"

# Run all the tests.
test:
	#!/usr/bin/env bash
	export LD_LIBRARY_PATH=$(pwd)/extension
	export GOPATH=$(pwd)/extension
	# Run the tests.
	go test -test.v $(find extension/src/wasmer -type f \( -name "*_test.go" \! -name "example_*.go" \) )
	# Run the short examples.
	go test -test.v extension/src/wasmer/example_test.go
	# Run the long examples.
	go test -test.v $(find extension/src/wasmer -type f \( -name "example_*_test.go" \) )

# Server the documentation.
doc:
	@echo 'Open http://localhost:6060/pkg/wasmer/'
	GOPATH=$(pwd)/extension godoc -http=:6060

# Local Variables:
# mode: makefile
# End:
