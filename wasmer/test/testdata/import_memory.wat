(module
  (import "env" "memory" (memory (;0;) 1))
  (func (export "read_memory") (result i32)
    i32.const 0
    i32.load)
)
