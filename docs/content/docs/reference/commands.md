# Command Reference

This is the complete reference for all Ellie CLI commands. Commands are organized by category for easy navigation.

## Core Commands

### Configuration

#### `ellie config`
Configure Ellie CLI settings.

```bash
ellie config
```

**What it does:**
- Prompts for username, OpenAI API key, and email
- Creates configuration file at `~/.ellie/.ellie.env`
- Sets secure file permissions

#### `ellie reset-config`
Reset Ellie CLI configuration.

```bash
ellie reset-config
```

**What it does:**
- Removes existing configuration
- Prompts for new configuration values
- Creates fresh configuration file

#### `ellie whoami`
Display current user information.

```bash
ellie whoami
```

**Output:**
```
Your majesty, Tach
```

### System Information

#### `ellie sysinfo`
Display comprehensive system information.

```bash
ellie sysinfo
```

**What it shows:**
- Operating system details
- Hardware specifications
- Memory usage
- CPU information
- Disk space
- Network interfaces

#### `ellie network-status`
Display detailed network information.

```bash
ellie network-status
```

**What it shows:**
- Active network interfaces
- IP addresses
- Connection status
- Network services
- WiFi information (platform-specific)

#### `ellie pwd`
Display current working directory.

```bash
ellie pwd
```

**Output:**
```
/Users/tach/projects/ellie
```

## File Operations

### `ellie list <directory>`
List files and directories.

```bash
# List current directory
ellie list .

# List specific directory
ellie list ~/Documents

# List home directory
ellie list ~
```

### `ellie create-file <path>`
Create a new file.

```bash
# Create empty file
ellie create-file my-notes.txt

# Create file with content
ellie create-file config.json '{"name": "project"}'
```

### `ellie open-explorer`
Open current directory in file explorer.

```bash
ellie open-explorer
```

### `ellie open <path>`
Open specific path in file explorer.

```bash
ellie open ~/Documents
ellie open /Applications
```

## Service Management

### `ellie start <service>`
Start system services.

```bash
# Start individual services
ellie start apache
ellie start mysql
ellie start postgres

# Start all services
ellie start all

# List available services
ellie start list
```

### `ellie stop <service>`
Stop system services.

```bash
# Stop individual services
ellie stop apache
ellie stop mysql

# Stop all services
ellie stop all
```

### `ellie restart <service>`
Restart system services.

```bash
# Restart individual services
ellie restart apache
ellie restart mysql

# Restart all services
ellie restart all
```

## Git Workflows

### `ellie git status`
Show Git repository status.

```bash
ellie git status
```

### `ellie git commit`
Create a conventional commit.

```bash
ellie git commit
```

**Interactive prompts:**
- Commit type (feat, fix, docs, etc.)
- Scope (optional)
- Description
- Body (optional)
- Breaking changes (optional)

### `ellie git push`
Push commits to remote repository.

```bash
ellie git push
```

### `ellie git pull`
Pull latest changes from remote repository.

```bash
ellie git pull
```

### `ellie setup-git`
Configure Git credentials.

```bash
ellie setup-git
```

## AI Integration

### `ellie chat`
Start interactive AI chat session.

```bash
ellie chat
```

**Features:**
- Context-aware conversations
- Code assistance
- System troubleshooting
- Best practices advice

### `ellie review <file>`
Review code using AI.

```bash
ellie review main.go
ellie review package.json
```

**What it analyzes:**
- Code quality
- Security issues
- Best practices
- Performance improvements

## Productivity Tools

### Todo Management

#### `ellie todo add "<task>" [category] [priority]`
Add a new todo item.

```bash
# Basic todo
ellie todo add "Fix login bug"

# With category and priority
ellie todo add "Implement OAuth2" auth high
ellie todo add "Update docs" documentation medium
```

**Priorities:**
- `high` - 🔴 High priority
- `medium` - 🟡 Medium priority
- `low` - 🟢 Low priority

#### `ellie todo list`
List all todo items.

```bash
ellie todo list
```

**Output format:**
```
Your todos:
📁 api:
  ❌ #1: Fix login bug 🔴 High
  ✅ #2: Update docs 🟡 Medium
```

#### `ellie todo complete <id>`
Mark todo as complete.

```bash
ellie todo complete 1
```

#### `ellie todo delete <id>`
Delete a todo item.

```bash
ellie todo delete 1
```

#### `ellie todo edit <id> <field> <value>`
Edit todo fields.

```bash
# Edit priority
ellie todo edit 1 priority high

# Edit task
ellie todo edit 1 task "Fix authentication bug"

# Edit category
ellie todo edit 1 category backend
```

### Project Management

#### `ellie project add <name> <path> [description] [tags...]`
Add a new project.

```bash
# Basic project
ellie project add api ~/projects/api

# With description and tags
ellie project add frontend ~/projects/frontend "React App" frontend,react
ellie project add backend ~/projects/backend "API Server" backend,nodejs
```

#### `ellie project list`
List all projects.

```bash
ellie project list
```

#### `ellie project search <query>`
Search projects by name, tag, or description.

```bash
ellie project search nodejs
ellie project search "API"
ellie project search backend
```

#### `ellie project delete <name>`
Delete a project.

```bash
ellie project delete api
```

#### `ellie switch <project-name>`
Switch to a project.

```bash
ellie switch api
```

### Alias Management

#### `ellie alias add <name>="<command>"`
Add a new alias.

```bash
ellie alias add gs="git status"
ellie alias add dev="ellie start-day && ellie switch api"
ellie alias add commit="ellie git commit"
```

#### `ellie alias list`
List all aliases.

```bash
ellie alias list
```

#### `ellie alias delete <name>`
Delete an alias.

```bash
ellie alias delete gs
```

## Daily Workflow

### `ellie start-day`
Start your development day.

```bash
ellie start-day
```

**What it does:**
- Opens configured applications
- Starts configured services
- Checks Git repositories
- Shows pending todos

### Daily Setup Configuration

#### `ellie day-start add <type> <value>`
Add items to daily setup.

```bash
# Add applications
ellie day-start add apps "code,terminal,chrome"

# Add services
ellie day-start add services "mysql,redis"

# Add Git repositories
ellie day-start add git_repos "~/projects/api,~/projects/frontend"
```

#### `ellie day-start list`
List daily setup configuration.

```bash
ellie day-start list
```

## System Commands

### `ellie run <command>`
Execute system commands.

```bash
# Basic commands
ellie run ls -la
ellie run ps aux
ellie run top

# Complex commands
ellie run "find . -name '*.go' -exec grep -l 'TODO' {} \;"
```

### `ellie install <package>`
Install software packages.

```bash
ellie install neofetch
ellie install htop
ellie install tree
```

### `ellie update`
Update system packages.

```bash
ellie update
```

## Network Management

### `ellie connect-wifi <SSID> <password>`
Connect to WiFi network.

```bash
ellie connect-wifi "Coffee Shop" "password123"
```

## Development Environment

### `ellie dev-init [--all]`
Initialize development environment.

```bash
# Basic setup
ellie dev-init

# Install all recommended tools
ellie dev-init --all
```

**What it installs:**
- Development tools
- Git configuration
- SSH keys
- Useful aliases
- Project templates

## Entertainment & Fun

### `ellie joke`
Tell a joke.

```bash
ellie joke
```

### `ellie weather`
Show weather information.

```bash
ellie weather
```

### `ellie remind`
Set a reminder.

```bash
ellie remind
```

## User Interface

### `ellie greet`
Greet the user based on time of day.

```bash
ellie greet
```

**Output examples:**
```
Good morning, Tach!
Good afternoon, Tach!
Good evening, Tach!
```

### `ellie focus`
Activate focus mode.

```bash
ellie focus
```

### Theme Management

#### `ellie theme set <theme>`
Set the theme.

```bash
ellie theme set light
ellie theme set dark
ellie theme set auto
```

#### `ellie theme show`
Show current theme.

```bash
ellie theme show
```

## Information & Help

### `ellie about`
Show about information.

```bash
ellie about
```

### `ellie history`
Show command history.

```bash
ellie history
```

### `ellie --help`
Show help information.

```bash
ellie --help
```

### `ellie --version`
Show version information.

```bash
ellie --version
```

## Email Features

### `ellie send-mail`
Send an email.

```bash
ellie send-mail
```

## Command Options

### Global Flags

- `--help` - Show help information
- `--version` - Show version information

### Common Patterns

#### Interactive Commands
Some commands are interactive and will prompt for input:
- `ellie git commit`
- `ellie config`
- `ellie chat`

#### File Paths
Commands that accept file paths support:
- Relative paths: `./file.txt`
- Absolute paths: `/path/to/file`
- Home directory: `~/file.txt`

#### Environment Variables
Commands respect environment variables:
- `HOME` - User home directory
- `SHELL` - Current shell
- `PATH` - Executable search path

## Error Handling

### Common Error Messages

**"Unknown command"**
- Check command spelling
- Use `ellie --help` to see available commands
- Ellie will suggest similar commands

**"Invalid usage"**
- Check command syntax
- Review required arguments
- Use `ellie <command> --help` for specific help

**"Permission denied"**
- Check file permissions
- Use `sudo` for system operations
- Verify user permissions

### Getting Help

For command-specific help:
```bash
ellie <command> --help
```

For general help:
```bash
ellie --help
```

For troubleshooting:
- Check the [Troubleshooting Guide](/docs/reference/troubleshooting)
- Open an issue on [GitHub](https://github.com/tacheraSasi/ellie/issues)

## Next Steps

- [Examples & Use Cases](/docs/reference/examples) - Real-world usage examples
- [Advanced Usage](/docs/reference/advanced) - Power user features
- [Troubleshooting](/docs/reference/troubleshooting) - Common issues and solutions 