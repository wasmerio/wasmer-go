extern "C" {
    fn sum_i32(x: i32, y: i32) -> i32;
    fn sum_i64(x: i64, y: i64) -> i64;
    fn sum_f32(x: f32, y: f32) -> f32;
    fn sum_f64(x: f64, y: f64) -> f64;
}

#[no_mangle]
pub extern "C" fn sum_i32_and_add_one(x: i32, y: i32) -> i32 {
    unsafe { sum_i32(x, y) + 1 }
}

#[no_mangle]
pub extern "C" fn sum_i64_and_add_one(x: i64, y: i64) -> i64 {
    unsafe { sum_i64(x, y) + 1 }
}

#[no_mangle]
pub extern "C" fn sum_f32_and_add_one(x: f32, y: f32) -> f32 {
    unsafe { sum_f32(x, y) + 1. }
}

#[no_mangle]
pub extern "C" fn sum_f64_and_add_one(x: f64, y: f64) -> f64 {
    unsafe { sum_f64(x, y) + 1. }
}
