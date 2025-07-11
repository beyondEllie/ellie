# Ellie CLI Demo Guide

This guide demonstrates all major features of Ellie CLI. Copy and paste the commands into your terminal to try them out!

---

## 🛠️ System Management

```bash
ellie sysinfo                # Show system information
ellie network-status         # Show network status
ellie pwd                    # Print working directory
ellie list ~                 # List files in your home directory
ellie create-file demo.txt   # Create a new file
ellie open-explorer          # Open file explorer in current directory
ellie open ~                 # Open file explorer at home directory
```

---

## 🔌 Service Management

```bash
ellie start apache           # Start Apache service
ellie stop mysql             # Stop MySQL service
ellie restart all            # Restart all services
ellie start list             # List available services
```

---

## 🚀 Git Workflows

```bash
ellie git status             # Show git status
ellie git commit             # Create a conventional commit
ellie git push               # Push commits
ellie git pull               # Pull latest changes
ellie setup-git              # Configure git credentials
```

---

## 📦 Package Management

```bash
ellie install neofetch       # Install a package
ellie update                 # Update all packages
```

---

## 📂 File & Network Operations

```bash
ellie connect-wifi "SSID" "password"   # Connect to WiFi
```

---

## 🤖 AI Integration

```bash
ellie chat                   # Start interactive AI chat session
```

---

## 📝 Todo Management

```bash
ellie todo add "Write docs" projectX high   # Add a todo with category and priority
ellie todo list                            # List all todos
ellie todo complete 1                      # Mark todo #1 as complete
ellie todo delete 1                        # Delete todo #1
ellie todo edit 2 priority medium          # Edit priority of todo #2
```

---

## 🚀 Project Management

```bash
ellie project add api ~/tach/api "API Service" backend,nodejs   # Add a project
ellie project list                                                  # List all projects
ellie project search nodejs                                         # Search projects
ellie project delete api                                            # Delete a project
ellie switch api                                                    # Switch to a project
```

---

## ⚡ Alias Management

```bash
ellie alias add gs="git status"     # Add alias
ellie alias list                    # List aliases
ellie alias delete gs               # Delete alias
```

---

## 🌅 Daily Setup

```bash
ellie start-day                     # Start your dev day
ellie day-start add apps "code"     # Add app to daily setup
ellie day-start add services "mysql"# Add service to daily setup
ellie day-start add git_repos "~/projects/api" # Add repo to daily setup
ellie day-start list                # List daily setup
```

---

## 🎨 Theme Management

```bash
ellie theme set dark                # Set theme to dark
ellie theme show                    # Show current theme
```

---

## 📜 Command History

```bash
ellie history                       # Show recent commands
```

---

## 📧 Email & Fun

```bash
ellie send-mail                     # Send an email
ellie weather                       # Show weather info
ellie joke                          # Tell a joke
ellie remind                        # Set a reminder
```

---

## 🧑 User & About

```bash
ellie whoami                        # Show current user
ellie about                         # Show about info
```

---

## 🛠️ Configuration

```bash
ellie config                        # Configure Ellie CLI
ellie reset-config                  # Reset configuration
```

---

## 🧪 Review & Focus

```bash
ellie review main.go                # Review a file with LLM
ellie focus                         # Activate focus mode
```

---

## 🏗️ Development Environment

```bash
ellie dev-init                      # Initialize dev environment
ellie dev-init --all                # Install all recommended tools
```

---

## 🏃 Run System Commands

```bash
ellie run ls -la                    # Run a system command
```

---

For more details, use:
```bash
ellie <command> --help
``` 