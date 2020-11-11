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
	GODEBUG=cgocheck=2 go test

# Local Variables:
# mode: makefile
# End:
