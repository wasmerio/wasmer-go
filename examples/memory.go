package main

import (
	"fmt"
	"path"
	"runtime"
	wasm "wasmer"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "memory.wasm")

	bytes, _ := wasm.ReadBytes(module_path)
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	result, _ := instance.Exports["return_hello"]()
	pointer := result.ToI32()

	memory := instance.Memory.GetData()

	fmt.Println(string(memory[pointer : pointer+13]))

	memory[pointer] = 72

	fmt.Println(string(memory[pointer : pointer+13]))

}

/*
<?php

declare(strict_types = 1);

require_once dirname(__DIR__) . '/vendor/autoload.php';

$instance = new Wasm\Instance(__DIR__ . '/memory.wasm');
$pointer = $instance->return_hello();

$memory = new Wasm\Uint8Array($instance->getMemoryBuffer(), $pointer);

$nth = 0;

while (0 !== $memory[$nth]) {
    echo chr($memory[$nth]);
    ++$nth;
}

echo "\n";

*/
