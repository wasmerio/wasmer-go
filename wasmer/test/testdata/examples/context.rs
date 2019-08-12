extern "C" {
    fn get_ctx_addr() -> i64;
}

#[no_mangle]
pub extern "C" fn run() -> i64 {
    unsafe { get_ctx_addr() }
}

fn main() {}
