package main

/*
#cgo LDFLAGS: -L. -lwasmer_runtime_c_api
#include "./wasmer.h"
*/
import "C"
import (
	"fmt"
	"io/ioutil"
	"unsafe"
)

func main() {
	var bytes []byte
	var err error

	bytes, err = ioutil.ReadFile("simple.wasm")

	if err != nil {
		fmt.Print(err)
		return;
	}

	if (false == C.wasmer_validate((*C.uchar) (unsafe.Pointer(&bytes[0])), C.uint(len(bytes)))) {
		fmt.Println("Module is invalid.");
		return;
	}

	var imports []C.wasmer_import_t = []C.wasmer_import_t{}
	var instance *C.wasmer_instance_t = nil

	var compile_result C.wasmer_result_t = C.wasmer_instantiate(
		&instance,
		(*C.uchar) (unsafe.Pointer(&bytes[0])),
		C.uint(len(bytes)),
		(*C.wasmer_import_t) (unsafe.Pointer(&imports)),
		C.int(0),
	)

	if (C.WASMER_OK != compile_result) {
		fmt.Println("Failed to compile the module.");
		return;
	}

	var parameter_one C.wasmer_value_t;
	parameter_one.tag = C.WASM_I32;
	parameter_one.value[C.WASM_I32] = 1;

	var parameter_two C.wasmer_value_t;
	parameter_two.tag = C.WASM_I32;
	parameter_two.value[C.WASM_I32] = 2;

	var parameters []C.wasmer_value_t = []C.wasmer_value_t{parameter_one, parameter_two};

	var result_one C.wasmer_value_t;
	var results []C.wasmer_value_t = []C.wasmer_value_t{result_one};

	var call_result C.wasmer_result_t = C.wasmer_instance_call(
		instance,
		C.CString("sum"),
		(*C.wasmer_value_t) (unsafe.Pointer(&parameters[0])),
		C.int(len(parameters)),
		(*C.wasmer_value_t) (unsafe.Pointer(&results[0])),
		C.int(len(results)),
	);

	if (C.WASMER_OK != call_result) {
		fmt.Println("Failed to call the `sum` function.");
		return;
	}

	fmt.Print("Result is: ");
	fmt.Println(results[0].value[C.WASM_I32]);
}
