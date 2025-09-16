# Ellie Core Rust-Go FFI Integration

This package provides Go bindings for the Rust `elliecore` library, enabling high-performance system operations with Rust's memory safety guarantees.

## Architecture

The integration works through CGO bindings that call into a Rust static library (`libelliecore.a`). The Rust library provides C-compatible FFI functions that the Go code calls directly.

### FFI Functions Available

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

## Memory Management

The Go code properly manages C string memory by:
1. Converting Go strings to C strings with `C.CString()`
2. Freeing C strings with `defer C.free(unsafe.Pointer(cstr))`
3. Using helper function `cStringToGoString()` to safely convert and free returned C strings

## Building

The build process requires:
1. Rust toolchain for building the static library
2. CGO enabled for Go compilation
3. Proper linking flags in the CGO LDFLAGS directive

Use `make all` to build both Rust and Go components.