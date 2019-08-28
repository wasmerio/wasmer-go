(module
  (type $hello_wasm_type (func (result i32)))

  (func $hello (type $hello_wasm_type) (result i32)
    i32.const 42)

  (export "hello" (func $hello)))