# Run all the tests.
test:
	#!/usr/bin/env bash
	#export DYLD_PRINT_LIBRARIES=y
	cd wasmer
	# Run the tests.
	GODEBUG=cgocheck=2 go test -test.v

examples:
	#!/usr/bin/env bash
	#export DYLD_PRINT_LIBRARIES=y
	cd examples
	# Run the examples.
	GODEBUG=cgocheck=2 go test -test.v

preview-doc ADDRESS="0.0.0.0" PORT="9090":
	@echo "Starting godoc preview..."
	@echo "Documentation preview will be available at: http://{{ADDRESS}}:{{PORT}}/pkg/github.com/wasmerio/wasmer-go/"
	godoc -http={{ADDRESS}}:{{PORT}} -timestamps

# Local Variables:
# mode: makefile
# End:
