# dev-init

Set up a complete development environment with essential tools and configurations.

## Usage
```sh
ellie dev-init [--all]
```

## Description

The enhanced `dev-init` command provides a comprehensive development environment setup that goes beyond just installing tools. It includes Git configuration, SSH key generation, useful aliases, and project templates.

## Features

### 🛠️ **Tool Installation**
- **Essential Tools**: Git, Node.js, Python, Yarn, Prettier, ESLint
- **Development Tools**: Docker, Go, VS Code, PostgreSQL, Redis
- **Cloud Tools**: AWS CLI, Terraform, kubectl, Helm
- **Infrastructure**: NGINX, GitHub CLI, .NET SDK, Java, PHP
- **DevOps**: Ansible, Vagrant

### 🔧 **Git Configuration**
- Sets up Git username and email
- Configures default branch to `main`
- Sets appropriate line ending handling

### 🔑 **SSH Key Generation**
- Generates RSA SSH key for GitHub/GitLab
- Displays public key for easy copying
- Checks for existing keys to avoid duplicates

### ⚡ **Useful Aliases**
- **Git Aliases**: `gs`, `ga`, `gc`, `gp`, `gl`, `gd`, `gco`, `gcb`, `gpl`, `gst`, `gstp`
- **Navigation**: `..`, `...`, `....`, `.....`
- **System**: `ll`, `la`, `l`, `c`, `h`, `j`, `ports`, `myip`, `weather`

### 📁 **Project Templates**
- **React App**: TypeScript React application
- **Node API**: Express.js API with essential packages
- **Python Web**: Flask web application with virtual environment
- **Go API**: Gin framework API
- **Docker Project**: Docker-based development environment

## Options

- `--all`: Install all recommended tools without prompting

## Examples

### Basic Setup
```sh
ellie dev-init
```
Interactive setup with prompts for each tool and configuration option.

### Install All Tools
```sh
ellie dev-init --all
```
Installs all recommended tools automatically.

## Interactive Features

### Tool Selection
The command will prompt you for each tool:
```
Install Node.js? (default: true) [Y/n]:
```

### Git Configuration
```
Enter your Git username: John Doe
Enter your Git email: john@example.com
```

### SSH Key Setup
```
Generate SSH key for GitHub/GitLab? [y/N]:
Enter your email for SSH key: john@example.com
```

### Project Templates
```
📋 Available project templates:
  1. react-app - React application with TypeScript
  2. node-api - Node.js API with Express
  3. python-web - Python web application with Flask
  4. go-api - Go API with Gin framework
  5. docker-project - Docker-based development environment

Create a project template? (number or 'n' for skip): 1
Enter project name: my-react-app
```

## Supported Operating Systems

- **macOS**: Uses Homebrew for package management
- **Linux**: Uses apt-get for Ubuntu/Debian systems
- **Windows**: Uses Chocolatey for package management

## Installation Summary

After completion, you'll see a detailed summary:
```
📊 Development Environment Setup Summary
────────────────────────────────────────
⏱️  Total time: 5m 30s
✅ Successfully installed: 8 tools
⏩ Skipped: 3 tools
❌ Failed: 0 tools

🎉 Successfully installed tools:
  • Git
  • Node.js
  • Python
  • Yarn
  • Prettier
  • ESLint
  • Docker
  • VS Code
```

## Next Steps

The command provides helpful next steps:
1. Restart your terminal to activate aliases
2. Configure your IDE/editor preferences
3. Set up your preferred Git hosting service
4. Install language-specific extensions
5. Add your SSH key to GitHub/GitLab (if generated)

## Pro Tips

- Use `ellie focus` for productive coding sessions
- Try `ellie git commit` for conventional commits
- Use `ellie project add` to track your projects
- Run `ellie start-day` to begin your dev day

## Troubleshooting

### Common Issues

**Installation Failures**
- Check internet connection
- Verify package manager is installed
- Run with administrator privileges
- Check system requirements

**SSH Key Issues**
- Ensure SSH is installed on your system
- Check permissions on `.ssh` directory
- Verify email format

**Alias Issues**
- Restart terminal after setup
- Check shell configuration file permissions
- Verify shell type (bash/zsh)

### Package Manager Requirements

- **macOS**: Install Homebrew (`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`)
- **Linux**: Ensure apt-get is available
- **Windows**: Install Chocolatey (`Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))`)

## Configuration Files

The command modifies these files:
- `~/.gitconfig` - Git configuration
- `~/.ssh/id_rsa` - SSH private key
- `~/.ssh/id_rsa.pub` - SSH public key
- `~/.zshrc` or `~/.bashrc` - Shell aliases 