<p align="center">
  <a href="https://wasmer.io" target="_blank" rel="noopener noreferrer">
    <img width="400" src="https://raw.githubusercontent.com/wasmerio/wasmer/master/logo.png" alt="Wasmer logo">
  </a>
</p>

<p align="center">
  <a href="https://spectrum.chat/wasmer">
    <img src="https://withspectrum.github.io/badge/badge.svg" alt="Join the Wasmer Community" valign="middle"></a>
  <a href="https://github.com/wasmerio/wasmer/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/wasmerio/wasmer.svg" alt="License" valign="middle"></a>
</p>

`go-ext-wasm` is a Go library for executing WebAssembly binaries.

# Install

So far, to install the `go-ext-wasm` library, just run this command in your shell:

```sh
$ go get https://github.com/wasmerio/go-ext-wasm
```

# Example

There is a toy program in `examples/simple.go`, written in Rust (or
any other language that compiles to WebAssembly):

```rust
#[no_mangle]
pub extern fn sum(x: i32, y: i32) -> i32 {
    x + y
}
```

After compilation to WebAssembly, the
[`examples/simple.wasm`](https://github.com/wasmerio/go-ext-wasm/blob/master/examples/simple.wasm)
binary file is generated. ([Download
it](https://github.com/wasmerio/go-ext-wasm/raw/master/examples/simple.wasm)).

Then, we can execute it in Go:

```go
package main

import (
	"fmt"
	wasm "https://github.com/wasmerio/go-ext-wasm"
)

func main(){
	bytes, _ := wasm.ReadBytes("simple.wasm")
	instance, _ := wasm.NewInstance(bytes)
	result, _ := instance.Call(
		"sum",
        wasm.I32(5),
        wasm.I32(37),
	)

	fmt.Println(result) // 42!
}
```

And then, finally, enjoy by running:

```sh
$ go run examples/simple.wasm
```

# Development

The Go library is written in Go and Rust.

To build both parts, run the following commands:

```sh
$ just rust
$ just go
```

To build the Go part, run:

(Yes, you need [`just`]).

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


[`just`]: https://github.com/casey/just/
[license]: https://github.com/wasmerio/wasmer/blob/master/LICENSE
