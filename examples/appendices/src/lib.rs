
extern "C" {
    fn host_func();
}

#[export_name = "test_host_func"]
pub extern "C" fn test_host_func(count: i32) -> i32 {
    let mut i = 0;
    while i < count {
        i = i+1;
        unsafe { host_func() };
    }

    return i
}

