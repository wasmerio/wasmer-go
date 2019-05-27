<p align="center">
  <a href="https://wasmer.io" target="_blank" rel="noopener noreferrer">
    <img width="400" src="https://raw.githubusercontent.com/wasmerio/wasmer/master/logo.png" alt="Wasmer logo">
  </a>
</p>

<p align="center">
  <a href="https://spectrum.chat/wasmer">
    <img src="https://withspectrum.github.io/badge/badge.svg" alt="Join the Wasmer Community" valign="middle"></a>
  <a href="https://godoc.org/github.com/wasmerio/go-ext-wasm/wasmer">
    <img src="https://godoc.org/github.com/wasmerio/go-ext-wasm?status.svg" alt="Read our API documentation" valign="middle"></a>
  <a href="https://goreportcard.com/report/github.com/wasmerio/go-ext-wasm">
    <img src="https://goreportcard.com/badge/github.com/wasmerio/go-ext-wasm" alt="Go Report Card" valign="middle" /></a>
  <a href="https://github.com/wasmerio/wasmer/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/wasmerio/wasmer.svg" alt="License" valign="middle"></a>
</p>

Wasm is a Go library for executing WebAssembly binaries.

# Install

To be defined.

# Documentation

[The documentation can be read online on godoc.org][documentation]. It
contains function descriptions, short examples, long examples
etc. Everything one need to start using Wasmer with Go!

[documentation]: https://godoc.org/github.com/wasmerio/go-ext-wasm/wasmer

# Examples

## Basic example: Exported function

There is a toy program in `wasmer/test/testdata/examples/simple.rs`,
written in Rust (or any other language that compiles to WebAssembly):

```rust
#[no_mangle]
pub extern fn sum(x: i32, y: i32) -> i32 {
    x + y
}
```

After compilation to WebAssembly, the
[`wasmer/test/testdata/examples/simple.wasm`](https://github.com/wasmerio/go-ext-wasm/blob/master/wasmer/test/testdata/examples/simple.wasm)
binary file is generated. ([Download
it](https://github.com/wasmerio/go-ext-wasm/raw/master/wasmer/test/testdata/examples/simple.wasm)).

Then, we can execute it in Go:

```go
package main

import (
	"fmt"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main(){
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("simple.wasm")
	
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	sum := instance.Exports["sum"]

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(5, 37)

	fmt.Println(result) // 42!
}
```

## Imported function

A WebAssembly module can export functions, this is how to run a
WebAssembly function, like we did in the previous example with
`instance.Exports["sum"](1, 2)`. Nonetheless, a WebAssembly module can
depend on “extern functions”, then called imported functions. For
instance, let's consider the basic following Rust program:

```rust
extern {
    fn sum(x: i32, y: i32) -> i32;
}

#[no_mangle]
pub extern fn add1(x: i32, y: i32) -> i32 {
    unsafe { sum(x, y) + 1 }
}
```

In this case, the `add1` function is a WebAssembly exported function,
whilst the `sum` function is a WebAssembly imported function (the
WebAssembly instance needs to _import_ it to complete the
program). Good news: We can write the implementation of the `sum`
function directly in Go!

First, we need to declare the `sum` function signature in C inside a
Go comment (with the help of [cgo]):

```go
package main

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
import "C"
```

Second, we declare the `sum` function implementation in Go. Notice the
`//export` which is the way cgo uses to map Go code to C code.

```go
//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}
```

Third, we use `NewImports` to create the WebAssembly imports. In this
code:

* `"sum"` is the imported function name,
* `sum` is the Go function pointer, and
* `C.sum` is the cgo function pointer.

```go
imports, _ := wasm.NewImports().Append("sum", sum, C.sum)
```

Finally, we use `NewInstanceWithImports` to inject the imports:

```go
bytes, _ := wasm.ReadBytes("imported_function.wasm")
instance, _ := wasm.NewInstanceWithImports(bytes, imports)
defer instance.Close()

// Gets and calls the `add1` exported function from the WebAssembly instance.
results, _ := instance.Exports["add1"](1, 2)

fmt.Println(result)
//   add1(1, 2)
// = sum(1 + 2) + 1
// = 1 + 2 + 1
// = 4
// QED
```

[cgo]: https://golang.org/cmd/cgo/

## Read the memory

A WebAssembly instance has a linear memory. Let's see how to read
it. Consider the following Rust program:

```rust
#[no_mangle]
pub extern fn return_hello() -> *const u8 {
    b"Hello, World!\0".as_ptr()
}
```

The `return_hello` function returns a pointer to a string. This string
is stored in the WebAssembly memory. Let's read it.

```go
bytes, _ := wasm.ReadBytes("memory.wasm")
instance, _ := wasm.NewInstance(bytes)
defer instance.Close()

// Calls the `return_hello` exported fucntion. This function returns a pointer to a string.
result, _ := instance.Exports["return_hello"]()

// Gets the pointer value as an integer.
pointer := result.ToI32()

// Reads the memory.
memory := instance.Memory.Data()

fmt.Println(string(memory[pointer : pointer+13])) // Hello, World!
```

In this example, we already know the string length, and we use a slice
to read a portion of the memory directly. Notice that the string
terminates by a null byte, which means we could iterate over the
memory starting from `pointer` until a null byte is met; that's a
similar approach.

For a more complete example, see the [Greet Example][greet-example].

[greet-example]: https://godoc.org/github.com/wasmerio/go-ext-wasm/wasmer#example-package--Greet

# Development

The Go library is written in Go and Rust.

To build both parts, run the following commands:

```sh
$ just rust
$ just go
```

To build the Go part, run:

(Yes, you need [`just`]).

[`just`]: https://github.com/casey/just/

# Testing

Once the library is build, run the following command:

```sh
$ just test
```

# What is WebAssembly?

Quoting [the WebAssembly site](https://webassembly.org/):

> WebAssembly (abbreviated Wasm) is a binary instruction format for a
> stack-based virtual machine. Wasm is designed as a portable target
> for compilation of high-level languages like C/C++/Rust, enabling
> deployment on the web for client and server applications.

About speed:

> WebAssembly aims to execute at native speed by taking advantage of
> [common hardware
> capabilities](https://webassembly.org/docs/portability/#assumptions-for-efficient-execution)
> available on a wide range of platforms.

About safety:

> WebAssembly describes a memory-safe, sandboxed [execution
> environment](https://webassembly.org/docs/semantics/#linear-memory) […].

# License

The entire project is under the BSD-3-Clause license. Please read [the
`LICENSE` file][license].


[license]: https://github.com/wasmerio/wasmer/blob/master/LICENSE
