use std::ffi::{CStr, CString};
use std::os::raw::{c_char, c_int};
use std::process::{Command, Stdio};
use std::{env, fs, path::Path};
use std::io::Write;
use std::fs::OpenOptions;
use std::time::UNIX_EPOCH;
use std::collections::HashMap;

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

// ============ ADDITIONAL SYSTEM OPERATIONS ============

#[unsafe(no_mangle)]
pub extern "C" fn create_dir(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    match fs::create_dir_all(path_str) {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn copy_file(src: *const c_char, dst: *const c_char) -> *mut c_char {
    let src_c = unsafe { CStr::from_ptr(src) };
    let dst_c = unsafe { CStr::from_ptr(dst) };
    let src_str = match src_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid source path").unwrap().into_raw(),
    };
    let dst_str = match dst_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid destination path").unwrap().into_raw(),
    };
    match fs::copy(src_str, dst_str) {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn move_file(src: *const c_char, dst: *const c_char) -> *mut c_char {
    let src_c = unsafe { CStr::from_ptr(src) };
    let dst_c = unsafe { CStr::from_ptr(dst) };
    let src_str = match src_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid source path").unwrap().into_raw(),
    };
    let dst_str = match dst_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid destination path").unwrap().into_raw(),
    };
    match fs::rename(src_str, dst_str) {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn get_file_hash(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    
    match fs::read(path_str) {
        Ok(contents) => {
            use std::collections::hash_map::DefaultHasher;
            use std::hash::{Hash, Hasher};
            let mut hasher = DefaultHasher::new();
            contents.hash(&mut hasher);
            let hash = hasher.finish();
            CString::new(format!("{:x}", hash)).unwrap().into_raw()
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn get_system_info() -> *mut c_char {
    let mut info = HashMap::new();
    
    // Get OS info
    info.insert("os", env::consts::OS);
    info.insert("arch", env::consts::ARCH);
    info.insert("family", env::consts::FAMILY);
    
    // Get CPU count
    let cpu_count = std::thread::available_parallelism()
        .map(|n| n.get().to_string())
        .unwrap_or_else(|_| "unknown".to_string());
    info.insert("cpu_count", &cpu_count);
    
    // Format as JSON-like string
    let json = format!(
        "{{\"os\":\"{}\",\"arch\":\"{}\",\"family\":\"{}\",\"cpu_count\":{}}}",
        info["os"], info["arch"], info["family"], cpu_count
    );
    
    CString::new(json).unwrap().into_raw()
}

#[unsafe(no_mangle)]
pub extern "C" fn get_disk_usage(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    
    // Use platform-specific commands for disk usage
    #[cfg(target_os = "windows")]
    let cmd = format!("dir /-c \"{}\"", path_str);
    #[cfg(not(target_os = "windows"))]
    let cmd = format!("df -h \"{}\"", path_str);
    
    let output = Command::new(SHELL)
        .arg(SHELL_ARG)
        .arg(&cmd)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output();
    
    match output {
        Ok(out) => {
            let stdout = String::from_utf8_lossy(&out.stdout);
            if !out.status.success() {
                let stderr = String::from_utf8_lossy(&out.stderr);
                CString::new(format!("Error: {}", stderr.trim())).unwrap().into_raw()
            } else {
                CString::new(stdout.trim()).unwrap().into_raw()
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn get_network_interfaces() -> *mut c_char {
    // Use platform-specific commands for network interfaces
    #[cfg(target_os = "windows")]
    let cmd = "ipconfig";
    #[cfg(target_os = "macos")]
    let cmd = "ifconfig | grep -E '^[a-zA-Z]|inet '";
    #[cfg(target_os = "linux")]
    let cmd = "ip addr show";
    #[cfg(not(any(target_os = "windows", target_os = "macos", target_os = "linux")))]
    let cmd = "ifconfig";
    
    let output = Command::new(SHELL)
        .arg(SHELL_ARG)
        .arg(cmd)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output();
    
    match output {
        Ok(out) => {
            let stdout = String::from_utf8_lossy(&out.stdout);
            if !out.status.success() {
                let stderr = String::from_utf8_lossy(&out.stderr);
                CString::new(format!("Error: {}", stderr.trim())).unwrap().into_raw()
            } else {
                CString::new(stdout.trim()).unwrap().into_raw()
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn ping_host(host: *const c_char) -> *mut c_char {
    let host_c = unsafe { CStr::from_ptr(host) };
    let host_str = match host_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid host").unwrap().into_raw(),
    };
    
    // Use platform-specific ping commands
    #[cfg(target_os = "windows")]
    let cmd = format!("ping -n 1 {}", host_str);
    #[cfg(not(target_os = "windows"))]
    let cmd = format!("ping -c 1 {}", host_str);
    
    let output = Command::new(SHELL)
        .arg(SHELL_ARG)
        .arg(&cmd)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output();
    
    match output {
        Ok(out) => {
            let stdout = String::from_utf8_lossy(&out.stdout);
            let stderr = String::from_utf8_lossy(&out.stderr);
            if !out.status.success() {
                CString::new(format!("Ping failed: {}", stderr.trim())).unwrap().into_raw()
            } else {
                CString::new(stdout.trim()).unwrap().into_raw()
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn get_process_list() -> *mut c_char {
    // Use platform-specific commands for process listing
    #[cfg(target_os = "windows")]
    let cmd = "tasklist";
    #[cfg(target_os = "macos")]
    let cmd = "ps aux";
    #[cfg(target_os = "linux")]
    let cmd = "ps aux";
    #[cfg(not(any(target_os = "windows", target_os = "macos", target_os = "linux")))]
    let cmd = "ps aux";
    
    let output = Command::new(SHELL)
        .arg(SHELL_ARG)
        .arg(cmd)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output();
    
    match output {
        Ok(out) => {
            let stdout = String::from_utf8_lossy(&out.stdout);
            if !out.status.success() {
                let stderr = String::from_utf8_lossy(&out.stderr);
                CString::new(format!("Error: {}", stderr.trim())).unwrap().into_raw()
            } else {
                CString::new(stdout.trim()).unwrap().into_raw()
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn kill_process(pid: *const c_char) -> *mut c_char {
    let pid_c = unsafe { CStr::from_ptr(pid) };
    let pid_str = match pid_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid PID").unwrap().into_raw(),
    };
    
    // Use platform-specific commands for killing processes
    #[cfg(target_os = "windows")]
    let cmd = format!("taskkill /PID {} /F", pid_str);
    #[cfg(not(target_os = "windows"))]
    let cmd = format!("kill {}", pid_str);
    
    let output = Command::new(SHELL)
        .arg(SHELL_ARG)
        .arg(&cmd)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output();
    
    match output {
        Ok(out) => {
            if out.status.success() {
                CString::new("OK").unwrap().into_raw()
            } else {
                let stderr = String::from_utf8_lossy(&out.stderr);
                CString::new(format!("Error: {}", stderr.trim())).unwrap().into_raw()
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn path_join(path1: *const c_char, path2: *const c_char) -> *mut c_char {
    let path1_c = unsafe { CStr::from_ptr(path1) };
    let path2_c = unsafe { CStr::from_ptr(path2) };
    let path1_str = match path1_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path1").unwrap().into_raw(),
    };
    let path2_str = match path2_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path2").unwrap().into_raw(),
    };
    
    let joined = Path::new(path1_str).join(path2_str);
    match joined.to_str() {
        Some(path) => CString::new(path).unwrap().into_raw(),
        None => CString::new("Error: Invalid path encoding").unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn path_absolute(path: *const c_char) -> *mut c_char {
    let path_c = unsafe { CStr::from_ptr(path) };
    let path_str = match path_c.to_str() {
        Ok(s) => s,
        Err(_) => return CString::new("Invalid path").unwrap().into_raw(),
    };
    
    match fs::canonicalize(path_str) {
        Ok(abs_path) => {
            match abs_path.to_str() {
                Some(path) => CString::new(path).unwrap().into_raw(),
                None => CString::new("Error: Invalid path encoding").unwrap().into_raw(),
            }
        }
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}
