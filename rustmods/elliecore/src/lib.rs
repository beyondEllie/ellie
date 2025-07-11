use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use std::process::{Command, Stdio};

#[unsafe(no_mangle)]
pub extern "C" fn run_cmd(cmd: *const c_char) -> *mut c_char {
    let c_str = unsafe { CStr::from_ptr(cmd) };
    let command_str = match c_str.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid command").unwrap().into_raw(),
    };

    let output = Command::new("sh")
        .arg("-c")
        .arg(command_str)
        .stdout(Stdio::piped())
        .output();

    match output {
        Ok(out) => {
            let result = String::from_utf8_lossy(&out.stdout).to_string();
            CString::new(result).unwrap().into_raw()
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}
