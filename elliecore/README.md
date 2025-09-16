# Ellie Core Rust-Go FFI Integration

This package provides Go bindings for the Rust `elliecore` library, enabling high-performance system operations with Rust's memory safety guarantees.

## Architecture

The integration works through CGO bindings that call into a Rust static library (`libelliecore.a`). The Rust library provides C-compatible FFI functions that the Go code calls directly.

### FFI Functions Available

#### **Basic Operations**
- **Command Execution**
  - `run_cmd(cmd)` - Execute shell commands
  - `run_cmd_with_env(cmd, envs)` - Execute commands with custom environment variables

- **File Operations**
  - `read_file(path)` - Read file contents
  - `write_file(path, content)` - Write/overwrite file
  - `append_file(path, content)` - Append to file
  - `delete_file(path)` - Delete file
  - `list_dir(path)` - List directory contents
  - `file_exists(path)` - Check if file exists
  - `file_metadata(path)` - Get file metadata as JSON

- **Environment & Directory Operations**
  - `get_env(key)` - Get environment variable
  - `set_env(key, value)` - Set environment variable
  - `get_cwd()` - Get current working directory
  - `change_dir(path)` - Change directory

#### **Advanced System Operations** 🆕
- **Enhanced File System**
  - `create_dir(path)` - Create directories recursively
  - `copy_file(src, dst)` - Copy files efficiently
  - `move_file(src, dst)` - Move/rename files
  - `get_file_hash(path)` - Calculate file hash
  - `path_join(path1, path2)` - Platform-safe path joining
  - `path_absolute(path)` - Get absolute path

- **System Information**
  - `get_system_info()` - Get OS, architecture, CPU count as JSON
  - `get_disk_usage(path)` - Get disk usage statistics

- **Network Operations**
  - `get_network_interfaces()` - List network interfaces
  - `ping_host(host)` - Ping network hosts

- **Process Management**
  - `get_process_list()` - List all running processes
  - `kill_process(pid)` - Terminate processes by PID

## Memory Management

The Go code properly manages C string memory by:
1. Converting Go strings to C strings with `C.CString()`
2. Freeing C strings with `defer C.free(unsafe.Pointer(cstr))`
3. Using helper function `cStringToGoString()` to safely convert and free returned C strings

## Performance Benefits

- **10 command executions**: ~6.8ms via Rust FFI
- **100 file operations**: ~16.5ms via Rust FFI  
- **Large file handling**: Successfully handles 57KB+ files
- **Error handling**: Robust error reporting for edge cases
- **Cross-platform**: Unified API across Windows, macOS, and Linux

## Building

The build process requires:
1. Rust toolchain for building the static library
2. CGO enabled for Go compilation
3. Proper linking flags in the CGO LDFLAGS directive

Use `make all` to build both Rust and Go components.

## Verification

The integration is verified through:
- Static linking verification (Rust symbols in binary)
- Comprehensive FFI function testing (24 functions total)
- Edge case and performance testing
- Memory safety validation

## Example Usage

```go
// System information
sysInfo := elliecore.GetSystemInfo()
// Returns: {"os":"linux","arch":"x86_64","family":"unix","cpu_count":4}

// Network operations
networkInfo := elliecore.GetNetworkInterfaces()
pingResult := elliecore.PingHost("google.com")

// Advanced file operations
elliecore.CreateDir("/tmp/myapp")
elliecore.CopyFile("config.yaml", "/tmp/myapp/config.yaml")
hash := elliecore.GetFileHash("/tmp/myapp/config.yaml")

// Process management
processes := elliecore.GetProcessList()
elliecore.KillProcess("1234")
```