extern "C" {
    fn inc(x: i32);

    fn get() -> i32;
}

#[no_mangle]
pub extern "C" fn inc_and_get(x: i32) -> i32 {
    unsafe {
        inc(x);
        get()
    }
}

fn main() {}
