// rustc counter.rs --target=wasm32-unknown-unknown -Clinker-plugin-lto -Copt-level=2
// wapm run wasm2wat counter.wasm > counter.wast
// rm counter.wasm
// wapm run wat2wasm counter.wast
// rm counter.wast
extern "C" {
    fn inc(x: i32);

    fn get() -> i32;

    fn copy_from_reg(reg_idx: i32);
}

#[no_mangle]
pub extern "C" fn inc_and_get(x: i32) -> i32 {
    unsafe {
        inc(x);
        get()
    }
}

#[no_mangle]
pub extern "C" fn copy_from_reg_and_get(reg_idx: i32) -> i32 {
    unsafe {
        copy_from_reg(reg_idx);
        get()
    }
}

fn main() {}