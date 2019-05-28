# Compile the Rust part for this specific system.
rust:
	cargo build --release

# Compile the Go part.
go go-build-args='':
	#!/usr/bin/env bash
	set -euo pipefail
	cd wasmer
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
	test -f libwasmer_runtime_c_api.${dylib_extension} || \
		ln -s ../target/release/deps/libwasmer_runtime_c_api-*.${dylib_extension} libwasmer_runtime_c_api.${dylib_extension}
	go build -ldflags="-r $(pwd)" -v {{go-build-args}} .

# Generate cgo debug objects.
debug-cgo:
	cd wasmer/ && \
		go tool cgo bridge.go && \
		echo "Browse wasmer/_obj/"

# Run all the tests.
test:
	#!/usr/bin/env bash
	export LD_LIBRARY_PATH=$(pwd)/wasmer
	cd wasmer/test/
	# Run the tests.
	go test -test.v $(find . -type f \( -name "*_test.go" \! -name "example_*.go" \! -name "benchmark*.go" \) ) imports.go
	# Run the short examples.
	go test -test.v example_test.go
	# Run the long examples.
	go test -test.v $(find . -type f \( -name "example_*_test.go" \! -name "example_import_test.go" \) )

# Run benchmarks. Subjects can be `wasmer`, `wagon` or `life`. Filter is a regex to select the benchmarks.
bench subject='wagon' filter='.*':
	cd benchmarks/{{subject}} && go test -bench '{{filter}}' benchmarks_test.go

# Local Variables:
# mode: makefile
# End:
