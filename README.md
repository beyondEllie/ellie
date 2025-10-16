# Ellie - Your Personal Command-Line Companion 🤖✨

**Meet Ellie** - Your intelligent CLI companion that takes the hassle out of system management and automation. With real-time health monitoring, smart context-aware assistance, and powerful automation capabilities, Ellie transforms how you interact with your terminal. Built with ❤️ by **Tachera Sasi**

## ✨ What's New & Impressive

🏥 **Intelligent System Monitoring** - Real-time health dashboard with CPU, memory, disk tracking, and proactive alerts  
🤖 **Smart Context-Aware Assistant** - Get intelligent suggestions based on your location, project type, and Git status  
⚙️ **Powerful Automation Scheduler** - Automate routine tasks with daily, hourly, weekly, or custom schedules  
📊 **Workflow Analysis** - Understand your command patterns and get optimization tips  
🎯 **Quick Health Checks** - Instant system health verification for scripts and workflows

> 💡 **Try it now**: `ellie impress` - See all the impressive features in action!

## Quick Installation ⚡

## Brew
``` bash
# One-time setup 
brew tap beyondEllie/ellie

# Install CLI
brew install ellie
```

### macOS (Intel)
```bash
curl -O -L https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_mac_amd64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_mac_amd64.tar.gz
```

### macOS (Apple Silicon)
```bash
curl -O -L https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_mac_amd64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_mac_arm64.tar.gz
```

<!-- ### Linux
```bash
curl -O -L https://github.com/tacheraSasi/ellie/releases/download/v0.0.91/ellie_linux_amd64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_linux_amd64.tar.gz
``` -->

### Verify Installation
```bash
ellie --version
# Should output: Ellie CLI Version: 0.0.91
```

## Getting Started with Ellie 🌟

### 🛠️ Configuration Setup

When you first run **Ellie**, you'll need to configure it with your details. Here's how to do it across all operating systems:

1. **Run the Configuration Setup**: Upon first use, Ellie will automatically prompt you to enter the configuration details.

```bash
$ ellie
🔧 Setting up Ellie CLI configuration...
-> Enter your username: Tach
-> Enter your OpenAI API key: sk-123...
-> Enter your Email (optional): this@that.com
✅ Configuration saved successfully at /home/tach/ellie/.ellie.env
🔧 Want to edit it? Open: /home/tach/ellie/.ellie.env
```

- **Username**: Your preferred username.
- **OpenAI API Key**: The API key for ChatGPT integration.
- **Email**: Optional, used for personalizing interactions.

2. **Manual Configuration Editing**:
   - If you'd like to modify the configuration later, simply open the `.ellie.env` file located in your home directory (or equivalent) and adjust the values.
   
   **Linux/MacOS**:
   ```bash
   nano ~/.ellie/.ellie.env
   ```

   **Windows**:
   ```bash
   notepad %HOMEPATH%\.ellie\.ellie.env
   ```

### Where is Your Configuration File Located?

- **Linux/MacOS**:  
  `/home/username/ellie/.ellie.env`
  
- **Windows**:  
  `C:\Users\YourUsername\ellie\.ellie.env`

The configuration file is created automatically, and you can edit it anytime to update details like your OpenAI key or username.

---

## What's New in v0.0.9? 🎉

- **Git Superpowers** 🚀 - Full Conventional Commits workflow with interactive prompts
- **Smarter UI** 🎨 - Colorized output and emoji-driven interface
- **Enhanced Service Control** 🔧 - Manage Apache/MySQL with single commands
- **Network Wizardry** 🌐 - WiFi connection management made simple
- **AI Integration** 🧠 - Built-in ChatGPT functionality

```bash
# Just look how pretty it is! ✨
$ ellie git commit
📝 Conventional Commit Builder
─────────────────────────────
🔧 Type (feat, fix, docs, style, refactor, perf, test, chore, revert) ➜ feat
🎯 Scope (optional) ➜ auth
📌 Description ➜ Add OAuth2 support
💬 Body (optional):
◎ Press Enter twice to finish:
Implemented Google and GitHub providers
Updated configuration schema

💥 Breaking change? (Y/n) ➜ y
📣 Breaking change details ➜ Changed config format
🔗 Issue number (optional) ➜ 42

✨ Commit Preview:
──────────────────
feat(auth): Add OAuth2 support

Implemented Google and GitHub providers
Updated configuration schema

BREAKING CHANGE: Changed config format

Refs #42
──────────────────
✅ Successfully committed and pushed!
```

---

## What's New in v0.0.9? 🎉

- **Todo Management** 📝 - Lightweight task tracking with completion status
- **Project Switcher** 🚀 - Quick navigation between projects
- **Enhanced Alias System** ⚡ - Create custom command shortcuts
- **Cross-Platform Support** 🖥️ - Works seamlessly on Windows, macOS, and Linux
- **Persistent Storage** 💾 - All data saved in Ellie's config directory

```bash
# Manage your tasks like a pro! ✨
$ ellie todo add "Implement OAuth2"
✅ Added todo #1: Implement OAuth2

$ ellie todo list
Your todos:
❌ #1: Implement OAuth2
✅ #2: Fix login bug

# Switch projects in a flash! ⚡
$ ellie project add api ~/projects/api
✅ Added project 'api'

$ ellie switch api
✅ Switched to project 'api'
```

---

## 🚀 Quick Start Guide

New to Ellie? Here's how to get the most out of your new companion:

### First Steps
```bash
# See what Ellie can do for you
ellie impress

# Get context-aware help
ellie assist

# Check your system health
ellie health

# Get smart suggestions
ellie suggest
```

### System Monitoring
```bash
# Full health dashboard
ellie health

# Quick health check
ellie quickcheck

# Real-time monitoring (Ctrl+C to stop)
ellie monitor

# Check for issues
ellie alerts
```

### Automation
```bash
# Quick setup for common tasks
ellie automate quick

# Add custom automation
ellie automate add "Morning Check" @09:00 "ellie health"

# List all automations
ellie automate list

# Run due tasks
ellie automate run
```

### Smart Assistant
```bash
# Get suggestions based on context
ellie suggest

# Analyze your workflow
ellie workflow

# Time-based suggestions
ellie time-suggest
```

📖 **For detailed documentation**: See [FEATURES.md](./FEATURES.md)

---

## Installation ⚡

```bash
# 1. Clone the repository
git clone https://github.com/tacheraSasi/ellie.git
cd ellie

# 2. Install dependencies
go get github.com/fatih/color

# 3. Build (choose your method)
make build  # or
go build -o ellie

# 4. Run with personality!
./ellie greet
```

---

## Core Features 🌟

### 🛠️ System Management
```bash
# Service Control
ellie start apache    # Start Apache
ellie restart mysql   # Restart MySQL
ellie stop all        # Stop all services

# System Insights
ellie sysinfo         # Show hardware/software specs
ellie network-status  # Detailed network analysis

# Command History
ellie history         # View recent commands (cross-platform)

# Daily Setup
ellie start-day       # Start your dev day (opens apps, services, checks Git)
ellie day-start add apps "code"  # Add apps to open
ellie day-start add services "mysql"  # Add services to start
ellie day-start add git_repos "~/projects/api"  # Add Git repos to check
ellie day-start list  # View your daily setup configuration

# Command Aliases
ellie alias add gs="git status"  # Create custom shortcuts
ellie alias list                 # View all aliases
ellie alias delete gs            # Remove an alias

# Todo Management
ellie todo add "Fix login bug" api high  # Add task with category and priority
ellie todo list                          # View categorized tasks
ellie todo complete 1                    # Mark task as done
ellie todo delete 1                      # Remove a task
ellie todo edit 1 priority high          # Update task priority

# Project Management
ellie project add api ~/projects/api "API Service" backend,nodejs  # Add project with description and tags
ellie project list                                                    # View all projects with details
ellie project search nodejs                                           # Search projects by name/tag/description
ellie switch api                                                      # Quick switch to project
```

### 📂 File Operations
```bash
ellie list ~/projects    # Visual directory listing
ellie create-file draft.md  # Create files with safety checks
ellie open-explorer     # Launch GUI file manager
```

### 🌐 Network Management
```bash
ellie connect-wifi "Coffee Shop" "p4ssw0rd!"  # Secure WiFi connection
ellie network-status                         # Real-time connection stats
```

### 🤖 AI Integration
```bash
# Chat mode (when no command specified)
ellie 
Talk to me: How do I fix a 500 error in Apache?
```

### 🚀 Git Workflows
```bash
ellie git status       # Enhanced status display
ellie git commit       # Interactive conventional commit
ellie git push         # Smart push with pre-checks
ellie setup-git        # Configure credentials securely
```

---

## Conventional Commits Made Easy 📝

Ellie guides you through professional commit messages:
```bash
$ ellie git commit
📝 Conventional Commit Builder
─────────────────────────────
🔧 Choose from: feat|fix|docs|style|refactor|perf|test|chore|revert
🎯 Add scope (optional module/component)
📌 Write clear, concise description
💬 Detailed body (Markdown supported)
💥 Breaking changes detection
🔗 Automatic issue reference formatting
```

## Package Management 📦
```bash
ellie install neofetch    # Cross-platform installs
ellie update              # System-wide updates
```

## Service Management 🔌
Control services like a pro:
```bash
# Start/Restart/Stop services
ellie start apache
ellie restart mysql
ellie stop all

# Systemd integration
ellie check-service nginx  # Coming soon!
```

---

## Why Ellie? 🤔

1. **Human-Friendly** 😊 - Designed for actual humans
2. **Context-Aware** 🧠 - Remembers your workflow
3. **Safe & Secure** 🔒 - Validation on every operation
4. **Cross-Platform** 🖥️ - Works where you work
5. **Extensible** 🔌 - Add your own modules

---

## Real-World Magic ✨
```bash
# Full development workflow
ellie start all          # Fire up services
ellie git commit         # Create perfect commit
ellie connect-wifi Work_Network $PASSWORD  # Stay connected
ellie sysinfo            # Monitor resources
```

---

## Contribution Guide 🌱
Found a bug? Got an idea? Let's build together!
1. Fork the repo
2. Create your feature branch
3. Submit a PR with tests
4. Join our Discord (coming soon!)

```bash
# Happy coding! 🎉
ellie --version
Ellie CLI Version: 0.0.3
```

---

**Maintained with ❤️ by Tachera Sasi** - Because terminal shouldn't mean terminal boredom!

## What's New in v0.0.90? 🎉

- **Command History Viewer** 📜 - Cross-platform command history with pretty printing
- **Daily Setup Routine** 🌅 - One command to start your dev day
- **Enhanced Project Management** 🚀 - Quick project switching and Git status
- **Cross-Platform App Launcher** 🖥️ - Open apps on any OS
- **Smart Service Management** 🔧 - Start services and check Git status

```bash
# View your command history! 📜
$ ellie history
Recent Commands:
  1: git status
  2: cd ~/projects
  3: ellie todo add "Fix bug"

# Start your day with one command! 🌅
$ ellie start-day
Starting your development day...
Opening applications...
Starting services...
Checking Git repositories...
Pending tasks:
❌ #1: Fix login bug
✅ #2: Update docs
Your development environment is ready! 🚀
```

## What's New in v0.0.91? 🎉

- **Enhanced Todo System** 📝 - Categories, priorities, and editing
- **Smart Project Management** 🚀 - Tags, descriptions, and search
- **Usage Tracking** ⏰ - Last used timestamps for projects
- **Better Organization** 📁 - Categorized todos and tagged projects
- **Improved Search** 🔍 - Find projects by name, tag, or description

```bash
# Manage todos with categories and priorities! 📝
$ ellie todo add "Fix login bug" api high
✅ Added todo #1: Fix login bug [api] 🔴 High

$ ellie todo list
Your todos:
📁 api:
  ❌ #1: Fix login bug 🔴 High
  ✅ #2: Update docs 🟡 Medium

# Organize projects with tags and descriptions! 🚀
$ ellie project add api ~/projects/api "API Service" backend,nodejs
✅ Added project 'api'

$ ellie project search nodejs
Search results:
📁 api
   📝 API Service
   📂 /Users/me/projects/api
   🏷️  backend, nodejs
```
