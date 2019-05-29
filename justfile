# Compile the library.
build go-build-args='-v':
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
	if ! test -f libwasmer_runtime_c_api.${dylib_extension}; then
		cargo build --release
		ln -s ../target/release/deps/libwasmer_runtime_c_api-*.${dylib_extension} libwasmer_runtime_c_api.${dylib_extension}
	fi
	go build {{go-build-args}} .

# Compile the Rust part for this specific system.
rust:
	cargo build --release

# Generate cgo debug objects.
debug-cgo:
	cd wasmer/ && \
		go tool cgo bridge.go && \
		echo "Browse wasmer/_obj/"

# Run all the tests.
test:
	#!/usr/bin/env bash
	#export DYLD_PRINT_LIBRARIES=y
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
