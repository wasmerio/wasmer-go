# Build the runtime shared library for this specific system.
build-runtime:
	#!/usr/bin/env bash
	set -euo pipefail

	# Build the shared library.
	cargo build --release

	# Find the shared library extension.
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

	# Link `wasmer/libwasmer_runtime_c_api.*`.
	rm -f wasmer/libwasmer_runtime_c_api.${dylib_extension}
	ln -s \
		'../'$( find target/release -name "libwasmer_runtime_c_api*.${dylib_extension}" -exec stat -n -f '%m ' {} ';' -print | sort -r | head -n 1 | cut -d ' ' -f 2 ) \
		wasmer/libwasmer_runtime_c_api.${dylib_extension}

	# Link `src/wasmer.h`.
	rm -f wasmer/wasmer.h
	ln -s \
		'../'$( find target/release/build -name 'wasmer.h' -exec stat -n -f '%m ' {} ';' -print | sort -r | head -n 1 | cut -d ' ' -f 2 ) \
		wasmer/wasmer.h

# Build the `wasmer` library.
build go-build-args='-v':
	cd wasmer && go build {{go-build-args}} .

# Build the `go-wasmer` bin.
build-bin go-build-args='-v':
	cd go-wasmer && go build {{go-build-args}} -o ../target/go/go-wasmer .

# Compile the Rust part for your specific system.
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
	cd wasmer
	# Run the tests.
	GODEBUG=cgocheck=2 go test -test.v $(find test -type f \( -name "*_test.go" \! -name "example_*.go" \! -name "benchmark*.go" \) ) test/imports.go
	# Run the short examples.
	go test -test.v example_test.go
	# Run the long examples.
	go test -test.v $(find . -type f \( -name "example_*_test.go" \! -name "_example_import_test.go" \) )

# Run benchmarks. Subjects can be `wasmer`, `wagon` or `life`. Filter is a regex to select the benchmarks.
bench subject='wasmer' filter='.*':
	cd benchmarks/{{subject}} && go test -bench '{{filter}}' benchmarks_test.go

# Local Variables:
# mode: makefile
# End:
