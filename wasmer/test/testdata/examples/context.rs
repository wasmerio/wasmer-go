// rustc context.rs --target=wasm32-unknown-unknown -Clinker-plugin-lto -Copt-level=2
// wapm run wasm2wat context.wasm > context.wast
// rm context.wasm
// wapm run wat2wasm context.wast
// rm context.wast
extern "C" {
    fn get_ctx_addr() -> i64;
}

#[no_mangle]
pub extern "C" fn run() -> i64 {
    unsafe { get_ctx_addr() }
}

fn main() {}
