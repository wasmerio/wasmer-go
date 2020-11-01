package utils

import (
	"fmt"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func GetGlobal(instance *wasmer.Instance, name string) *wasmer.Global {
	global, err := instance.Exports.GetGlobal(name)

	if err != nil {
		panic(fmt.Sprintf("Failed to get the `%s` global: %s\n", name, err))
	}

	return global
}

func GetFunction(instance *wasmer.Instance, name string) func(...interface{}) (interface{}, error) {
	fn, err := instance.Exports.GetFunction(name)

	if err != nil {
		panic(fmt.Sprintf("Failed to get the `%s` function: %s\n", name, err))
	}

	return fn
}

func GetRawFunction(instance *wasmer.Instance, name string) *wasmer.Function {
	fn, err := instance.Exports.GetRawFunction(name)

	if err != nil {
		panic(fmt.Sprintf("Failed to get the `%s` function: %s\n", name, err))
	}

	return fn
}
