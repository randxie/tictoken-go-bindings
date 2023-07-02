use std::ffi::CStr;
use tiktoken_rs::get_bpe_from_model as rs_get_bpe_from_model;
use tiktoken_rs::CoreBPE;

#[no_mangle]
pub extern "C" fn get_bpe_from_model(model: *const libc::c_char) -> *mut CoreBPE {
    let model_cstr = unsafe { CStr::from_ptr(model) };
    let model_name = model_cstr.to_str().unwrap();
    let bpe = rs_get_bpe_from_model(model_name).expect("failed to create bpe");
    return Box::into_raw(Box::new(bpe));
}

#[no_mangle]
pub extern "C" fn free_bpe(ptr: *mut ::libc::c_void) {
    if ptr.is_null() {
        return;
    }
    ptr.cast::<CoreBPE>();
}

#[no_mangle]
pub extern "C" fn encode(ptr: *mut libc::c_void, prompt: *const libc::c_char, len: *mut u32) -> *mut u32 {    
    let bpe: &CoreBPE;
    unsafe {
        bpe = ptr.cast::<CoreBPE>().as_ref().expect("failed to cast bpe");
    }
    let prompt_cstr = unsafe { CStr::from_ptr(prompt) };
    let prompt = prompt_cstr.to_str().unwrap();

    let vec = bpe.encode_with_special_tokens(prompt);
    unsafe {
        *len = vec.len() as u32;
    }

    // cast usize to u32 as return.
    let mut output: Vec<u32> = vec.iter().map(|&e| e as u32).collect();
    let output_ptr = output.as_mut_ptr();
    std::mem::forget(output);
    output_ptr
}