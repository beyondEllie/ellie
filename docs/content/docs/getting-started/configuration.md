# Configuration Guide

Ellie CLI is highly configurable and can be tailored to your preferences and workflow. This guide covers all configuration options and customization features.

## Initial Configuration

When you first run Ellie, it will prompt you for basic configuration:

```bash
ellie
```

You'll be asked to provide:
- **Username**: Your preferred username for personalized interactions
- **OpenAI API Key**: For AI chat features (optional but recommended)
- **Email**: For email features and personalization (optional)

## Configuration File

Ellie stores its configuration in a `.env` file:

- **Linux/macOS**: `~/.ellie/.ellie.env`
- **Windows**: `%HOMEPATH%\.ellie\.ellie.env`

### Manual Configuration

You can edit the configuration file directly:

```bash
# Linux/macOS
nano ~/.ellie/.ellie.env

# Windows
notepad %HOMEPATH%\.ellie\.ellie.env
```

### Configuration Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `USERNAME` | Your username for personalized interactions | `Tach` |
| `EMAIL` | Your email address | `user@example.com` |
| `OPENAI_API_KEY` | OpenAI API key for AI features | `sk-...` |
| `RELAY_API_KEY` | EkiliRelay API key for email features | `relay_key_...` |

## Theme Configuration

Ellie supports multiple themes for a personalized experience:

### Available Themes

- **light**: Light theme with dark text
- **dark**: Dark theme with light text  
- **auto**: Automatically matches your system theme

### Setting Themes

```bash
# Set theme
ellie theme set dark

# Show current theme
ellie theme show
```

### Theme Persistence

Themes are automatically saved and persist across sessions.

## AI Configuration

### OpenAI Integration

To use Ellie's AI features, you need an OpenAI API key:

1. Get an API key from [OpenAI](https://platform.openai.com/api-keys)
2. Add it to your configuration:
   ```bash
   ellie config
   ```
   Or edit the config file directly:
   ```
   OPENAI_API_KEY=sk-your-api-key-here
   ```

### AI Models

Ellie supports multiple OpenAI models:

- **gpt-3.5-turbo** (default) - Fast and cost-effective
- **gpt-4** - More capable but slower
- **gpt-4o** - Latest model with enhanced capabilities
- **gpt-4o-mini** - Optimized for speed and cost

You can change the model in the code or configuration.

## Email Configuration

### EkiliRelay Setup

For email features, configure your EkiliRelay API key:

1. Get an API key from [EkiliRelay Console](https://relay.ekilie.com/console)
2. Add it to your configuration:
   ```
   RELAY_API_KEY=your-relay-key-here
   ```

## Shell Integration

### Alias Management

Ellie can manage shell aliases automatically:

```bash
# Add an alias
ellie alias add gs="git status"

# List aliases
ellie alias list

# Remove an alias
ellie alias delete gs
```

Aliases are automatically added to your shell configuration file (`.zshrc`, `.bashrc`, etc.).

### Shell Detection

Ellie automatically detects your shell:
- **zsh**: Uses `~/.zshrc`
- **bash**: Uses `~/.bashrc`
- **fallback**: Defaults to `~/.zshrc`

## Daily Setup Configuration

Configure your daily development routine:

```bash
# Add apps to open daily
ellie day-start add apps "code"

# Add services to start daily
ellie day-start add services "mysql"

# Add Git repos to check daily
ellie day-start add git_repos "~/projects/api"

# List your daily setup
ellie day-start list
```

## Project Management

### Project Configuration

Configure project-specific settings:

```bash
# Add a project
ellie project add api ~/projects/api "API Service" backend,nodejs

# Switch to a project
ellie switch api
```

## Environment Variables

### System Environment

Ellie respects standard environment variables:

- `HOME` - User home directory
- `SHELL` - Current shell
- `PATH` - Executable search path
- `EDITOR` - Default text editor

### Ellie-Specific Variables

You can set Ellie-specific environment variables:

```bash
export ELLIE_CONFIG_DIR=~/custom/config
export ELLIE_LOG_LEVEL=debug
```

## Logging and Debugging

### Log Levels

Configure logging verbosity:

- **info** (default) - Standard information
- **debug** - Detailed debugging information
- **error** - Only error messages

### Debug Mode

Enable debug mode for troubleshooting:

```bash
export ELLIE_DEBUG=true
ellie --help
```

## Security Considerations

### API Key Security

- Store API keys securely in the configuration file
- The config file has restricted permissions (600)
- Never commit API keys to version control
- Use environment variables for CI/CD environments

### File Permissions

Ellie automatically sets secure permissions on configuration files:
- Configuration files: `600` (user read/write only)
- Log files: `644` (user read/write, group/other read)

## Backup and Migration

### Backup Configuration

```bash
# Backup your configuration
cp ~/.ellie/.ellie.env ~/.ellie/.ellie.env.backup
```

### Migration

To migrate Ellie to a new system:

1. Copy the configuration file
2. Copy the `.ellie` directory
3. Reinstall Ellie on the new system
4. Restore your configuration

## Advanced Configuration

### Custom Commands

You can extend Ellie with custom commands by editing the source code or creating plugins.

### Integration with Other Tools

Ellie integrates with:
- **Git**: Enhanced Git workflows
- **Docker**: Container management
- **Kubernetes**: Cluster management
- **Cloud providers**: AWS, GCP, Azure

## Troubleshooting Configuration

### Common Issues

**Configuration not loading**
```bash
# Check file permissions
ls -la ~/.ellie/.ellie.env

# Reset configuration
ellie reset-config
```

**API key issues**
```bash
# Verify API key format
echo $OPENAI_API_KEY | head -c 10

# Test API connection
ellie chat
```

**Theme not applying**
```bash
# Force theme refresh
ellie theme set light
ellie theme set dark
```

### Getting Help

For configuration issues:
- Check the [Troubleshooting Guide](/docs/reference/troubleshooting)
- Review the [Command Reference](/docs/reference/commands)
- Open an issue on [GitHub](https://github.com/tacheraSasi/ellie/issues) 