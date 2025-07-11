use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use std::process::{Command, Stdio};
use std::env;
use std::fs;

#[cfg(target_os = "windows")]
const SHELL: &str = "cmd";
#[cfg(target_os = "windows")]
const SHELL_ARG: &str = "/C";

#[cfg(not(target_os = "windows"))]
const SHELL: &str = "sh";
#[cfg(not(target_os = "windows"))]
const SHELL_ARG: &str = "-c";

#[unsafe(no_mangle)]
pub extern "C" fn run_cmd(cmd: *const c_char) -> *mut c_char {
    let c_str = unsafe { CStr::from_ptr(cmd) };
    let command_str = match c_str.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid command").unwrap().into_raw(),
    };

    let output = Command::new(SHELL)
        .arg(SHELL_ARG)
        .arg(command_str)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output();

    match output {
        Ok(out) => {
            let stdout = String::from_utf8_lossy(&out.stdout);
            let stderr = String::from_utf8_lossy(&out.stderr);
            if !out.status.success() {
                let result = format!("Error: {}\nOutput:\n{}", stderr.trim(), stdout.trim());
                CString::new(result).unwrap().into_raw()
            } else if !stderr.trim().is_empty() {
                let result = format!("Output:\n{}\nError:\n{}", stdout.trim(), stderr.trim());
                CString::new(result).unwrap().into_raw()
            } else {
                CString::new(stdout.trim()).unwrap().into_raw()
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn run_cmd_with_env(cmd: *const c_char, envs: *const c_char) -> *mut c_char {
    let c_str = unsafe { CStr::from_ptr(cmd) };
    let command_str = match c_str.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid command").unwrap().into_raw(),
    };
    let env_str = unsafe { CStr::from_ptr(envs) };
    let env_pairs = match env_str.to_str() {
        Ok(s) => s,
        Err(_) => "",
    };
    let mut command = Command::new(SHELL);
    command.arg(SHELL_ARG).arg(command_str);
    for pair in env_pairs.split(';') {
        if let Some((k, v)) = pair.split_once('=') {
            command.env(k, v);
        }
    }
    let output = command.stdout(Stdio::piped()).stderr(Stdio::piped()).output();
    match output {
        Ok(out) => {
            let stdout = String::from_utf8_lossy(&out.stdout);
            let stderr = String::from_utf8_lossy(&out.stderr);
            if !out.status.success() {
                let result = format!("Error: {}\nOutput:\n{}", stderr.trim(), stdout.trim());
                CString::new(result).unwrap().into_raw()
            } else if !stderr.trim().is_empty() {
                let result = format!("Output:\n{}\nError:\n{}", stdout.trim(), stderr.trim());
                CString::new(result).unwrap().into_raw()
            } else {
                CString::new(stdout.trim()).unwrap().into_raw()
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn read_file(path: *const c_char) -> *mut c_char {
    let c_str = unsafe { CStr::from_ptr(path) };
    let file_path = match c_str.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    match fs::read_to_string(file_path) {
        Ok(content) => CString::new(content).unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}
