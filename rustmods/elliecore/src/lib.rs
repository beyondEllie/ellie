use std::ffi::{CStr, CString};
use std::os::raw::{c_char, c_int};
use std::process::{Command, Stdio};
use std::{env, fs, path::Path};
use std::io::Write;
use std::fs::OpenOptions;
use std::time::UNIX_EPOCH;

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

#[unsafe(no_mangle)]
pub extern "C" fn write_file(path: *const c_char, content: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let content_c = unsafe { CStr::from_ptr(content) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    let content_str = match content_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid content").unwrap().into_raw(),
    };
    match fs::write(path_str, content_str) {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn append_file(path: *const c_char, content: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let content_c = unsafe { CStr::from_ptr(content) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    let content_str = match content_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid content").unwrap().into_raw(),
    };
    let mut file = match OpenOptions::new().append(true).create(true).open(path_str) {
        Ok(f) => f,
        Err(e) => return CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    };
    match file.write_all(content_str.as_bytes()) {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn delete_file(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    match fs::remove_file(path_str) {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn list_dir(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    let mut entries = Vec::new();
    match fs::read_dir(path_str) {
        Ok(read_dir) => {
            for entry in read_dir {
                if let Ok(e) = entry {
                    if let Ok(name) = e.file_name().into_string() {
                        entries.push(name);
                    }
                }
            }
            CString::new(entries.join("\n")).unwrap().into_raw()
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn get_env(key: *const c_char) -> *mut c_char {
    let key_c = unsafe { CStr::from_ptr(key) };
    let key_str = match key_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid key").unwrap().into_raw(),
    };
    match env::var(key_str) {
        Ok(val) => CString::new(val).unwrap().into_raw(),
        Err(_) => CString::new("").unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn set_env(key: *const c_char, value: *const c_char) -> *mut c_char {
    let key_c = unsafe { CStr::from_ptr(key) };
    let value_c = unsafe { CStr::from_ptr(value) };
    let key_str = match key_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid key").unwrap().into_raw(),
    };
    let value_str = match value_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid value").unwrap().into_raw(),
    };
    unsafe {
        env::set_var(key_str, value_str);
    }
    CString::new("OK").unwrap().into_raw()
}

#[unsafe(no_mangle)]
pub extern "C" fn get_cwd() -> *mut c_char {
    match env::current_dir() {
        Ok(path) => CString::new(path.display().to_string()).unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn change_dir(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    match env::set_current_dir(path_str) {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn file_exists(path: *const c_char) -> c_int {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return 0,
    };
    if Path::new(path_str).exists() {
        1
    } else {
        0
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn file_metadata(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    match fs::metadata(path_str) {
        Ok(meta) => {
            let size = meta.len();
            let readonly = meta.permissions().readonly();
            let modified = meta.modified().ok().and_then(|m| m.duration_since(UNIX_EPOCH).ok()).map(|d| d.as_secs()).unwrap_or(0);
            let json = format!("{{\"size\":{},\"readonly\":{},\"modified\":{}}}", size, readonly, modified);
            CString::new(json).unwrap().into_raw()
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}
