# Installation Guide

Ellie CLI is available for macOS, Linux, and Windows. Choose the installation method that works best for your system.

## Quick Installation

### macOS

#### Using Homebrew (Recommended)
```bash
# One-time setup
brew tap beyondEllie/ellie

# Install Ellie CLI
brew install ellie
```

#### Direct Download
```bash
# For Intel Macs
curl -O -L https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_mac_amd64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_mac_amd64.tar.gz

# For Apple Silicon (M1/M2)
curl -O -L https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_mac_arm64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_mac_arm64.tar.gz
```

### Linux

#### Direct Download
```bash
curl -O -L https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_linux_amd64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_linux_amd64.tar.gz
```

#### Using Package Managers
```bash
# Ubuntu/Debian (when available)
sudo apt update
sudo apt install ellie

# CentOS/RHEL (when available)
sudo yum install ellie
```

### Windows

#### Direct Download
```powershell
# Download the Windows executable
Invoke-WebRequest -Uri "https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_windows_amd64.exe" -OutFile "ellie.exe"

# Move to a directory in your PATH (e.g., C:\Windows\System32)
Move-Item ellie.exe C:\Windows\System32\
```

#### Using Chocolatey (when available)
```powershell
choco install ellie
```

## Build from Source

If you prefer to build Ellie from source or want the latest development version:

### Prerequisites
- Go 1.21 or later
- Rust (for FFI components)
- Git

### Build Steps
```bash
# Clone the repository
git clone https://github.com/tacheraSasi/ellie.git
cd ellie

# Build everything (Rust + Go)
make all

# Or build manually
cd rustmods/elliecore && cargo build
cd ../..
go mod tidy
go build -o ellie

# Install globally (optional)
sudo mv ellie /usr/local/bin/
```

## Verify Installation

After installation, verify that Ellie is working correctly:

```bash
ellie --version
# Should output: Ellie CLI v0.0.91
```

## First Run

When you run Ellie for the first time, it will guide you through the initial configuration:

```bash
ellie
```

You'll be prompted to enter:
- **Username**: Your preferred username
- **OpenAI API Key**: For AI chat features (optional but recommended)
- **Email**: For personalized interactions (optional)

## Configuration File Location

Ellie stores its configuration in:
- **Linux/macOS**: `~/.ellie/.ellie.env`
- **Windows**: `%HOMEPATH%\.ellie\.ellie.env`

## Next Steps

Once installed, check out:
- [Configuration Guide](/docs/getting-started/configuration) - Learn how to customize Ellie
- [First Steps](/docs/getting-started/first-steps) - Get started with basic commands
- [Core Features](/docs/getting-started/features) - Explore what Ellie can do

## Troubleshooting

### Common Issues

**Permission Denied**
```bash
# Make sure the binary is executable
chmod +x /usr/local/bin/ellie
```

**Command Not Found**
```bash
# Ensure /usr/local/bin is in your PATH
echo $PATH | grep /usr/local/bin
```

**Rust FFI Issues**
```bash
# Rebuild the Rust components
cd rustmods/elliecore && cargo clean && cargo build
```

### Getting Help

If you encounter any issues:
- Check the [Troubleshooting Guide](/docs/reference/troubleshooting)
- Open an issue on [GitHub](https://github.com/tacheraSasi/ellie/issues)
- Join our [Community](/docs/community) for support 