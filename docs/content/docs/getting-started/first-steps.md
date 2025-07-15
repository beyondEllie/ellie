# First Steps with Ellie

Welcome to Ellie CLI! This guide will help you get started with your first commands and explore the essential features that make Ellie your perfect terminal companion.

## Quick Start

### 1. Say Hello to Ellie

Start with a simple greeting:

```bash
ellie greet
```

Ellie will greet you based on the time of day and your configured username.

### 2. Check Your System

Get an overview of your system:

```bash
ellie sysinfo
```

This shows hardware specs, OS information, and system resources.

### 3. Explore Your Environment

See what Ellie knows about your setup:

```bash
ellie whoami
ellie pwd
ellie list ~
```

## Essential Commands

### System Management

**Get system information:**
```bash
ellie sysinfo
```

**Check network status:**
```bash
ellie network-status
```

**Open file explorer:**
```bash
ellie open-explorer
```

### File Operations

**List directory contents:**
```bash
ellie list ~/Documents
```

**Create a new file:**
```bash
ellie create-file my-notes.txt
```

**Read file contents:**
```bash
ellie read-file my-notes.txt
```

### Git Workflows

**Check Git status:**
```bash
ellie git status
```

**Create a conventional commit:**
```bash
ellie git commit
```

**Push your changes:**
```bash
ellie git push
```

## AI Integration

### Chat with Ellie

Start an interactive AI chat session:

```bash
ellie chat
```

Ask Ellie anything - from coding questions to system management help!

### Code Review

Have Ellie review your code:

```bash
ellie review main.go
```

Ellie will analyze your code for bugs, security issues, and improvements.

## Productivity Features

### Todo Management

**Add a task:**
```bash
ellie todo add "Fix login bug" api high
```

**List your todos:**
```bash
ellie todo list
```

**Mark a task complete:**
```bash
ellie todo complete 1
```

### Project Management

**Add a project:**
```bash
ellie project add api ~/projects/api "API Service" backend,nodejs
```

**Switch to a project:**
```bash
ellie switch api
```

**Search projects:**
```bash
ellie project search nodejs
```

### Alias Management

**Create a useful alias:**
```bash
ellie alias add gs="git status"
```

**List your aliases:**
```bash
ellie alias list
```

## Daily Workflow

### Start Your Day

Configure and use Ellie's daily setup:

```bash
# Add apps to open daily
ellie day-start add apps "code"

# Add services to start daily
ellie day-start add services "mysql"

# Start your day
ellie start-day
```

### Service Management

**Start services:**
```bash
ellie start apache
ellie start mysql
ellie start all
```

**Check service status:**
```bash
ellie start list
```

## Fun Features

### Entertainment

**Get a joke:**
```bash
ellie joke
```

**Check the weather:**
```bash
ellie weather
```

**Set a reminder:**
```bash
ellie remind
```

## Configuration

### Personalize Ellie

**Set your theme:**
```bash
ellie theme set dark
```

**Configure Ellie:**
```bash
ellie config
```

**Reset configuration:**
```bash
ellie reset-config
```

## Getting Help

### Built-in Help

**Show all commands:**
```bash
ellie --help
```

**Get command help:**
```bash
ellie <command> --help
```

### About Ellie

**Learn more:**
```bash
ellie about
```

## Your First Workflow

Here's a complete example of using Ellie for a typical development task:

```bash
# 1. Start your day
ellie start-day

# 2. Switch to your project
ellie switch api

# 3. Check Git status
ellie git status

# 4. Add a todo
ellie todo add "Implement user authentication" auth high

# 5. Review your code
ellie review auth.go

# 6. Commit your changes
ellie git commit

# 7. Push to remote
ellie git push

# 8. Mark todo complete
ellie todo complete 1
```

## Tips for Success

### 1. Use Aliases
Create aliases for your most common commands:
```bash
ellie alias add dev="ellie start-day && ellie switch api"
ellie alias add commit="ellie git commit"
```

### 2. Leverage AI
Use Ellie's AI features for:
- Code reviews
- Debugging help
- Best practices advice
- System troubleshooting

### 3. Organize Projects
Keep your projects organized with Ellie's project management:
```bash
ellie project add frontend ~/projects/frontend "React App" frontend,react
ellie project add backend ~/projects/backend "API Server" backend,nodejs
```

### 4. Daily Routine
Set up a daily routine that works for you:
```bash
ellie day-start add apps "code,terminal,chrome"
ellie day-start add services "mysql,redis"
ellie day-start add git_repos "~/projects/api,~/projects/frontend"
```

## Next Steps

Now that you've mastered the basics, explore:

- [Core Features](/docs/getting-started/features) - Deep dive into Ellie's capabilities
- [Command Reference](/docs/reference/commands) - Complete command documentation
- [Examples & Use Cases](/docs/reference/examples) - Real-world usage examples
- [Advanced Usage](/docs/reference/advanced) - Power user features

## Need Help?

- Check the [Troubleshooting Guide](/docs/reference/troubleshooting)
- Join our [Community](/docs/community)
- Open an issue on [GitHub](https://github.com/tacheraSasi/ellie/issues)

Welcome to the Ellie family! 🚀 