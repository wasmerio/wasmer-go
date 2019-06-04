extern "C" {
    fn log_message(pointer: *const u8, length: u32);
}

#[no_mangle]
pub extern "C" fn do_something() {
    let message = "hello";

    unsafe { log_message(message.as_ptr(), message.len() as u32) }
}
