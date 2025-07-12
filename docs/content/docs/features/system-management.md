# System Management

Ellie provides powerful system management capabilities that work seamlessly across Windows, macOS, and Linux. From service management to file operations, Ellie makes system administration tasks simple and efficient.

## Service Management

Ellie can manage system services with simple commands, automatically detecting and using the appropriate service manager for your system.

### Starting Services

```bash
# Start individual services
ellie start apache
ellie start mysql
ellie start postgres

# Start all configured services
ellie start all

# List available services
ellie start list
```

### Stopping Services

```bash
# Stop individual services
ellie stop apache
ellie stop mysql

# Stop all services
ellie stop all
```

### Restarting Services

```bash
# Restart individual services
ellie restart apache
ellie restart mysql

# Restart all services
ellie restart all
```

### Supported Services

Ellie supports the following services out of the box:

- **Apache** - Web server
- **MySQL** - Database server
- **PostgreSQL** - Database server
- **Redis** - Cache server (coming soon)
- **Nginx** - Web server (coming soon)

### Cross-Platform Support

Ellie automatically detects your system and uses the appropriate service manager:

- **Linux**: systemd, systemctl
- **macOS**: brew services
- **Windows**: Windows Services

## System Information

Get comprehensive information about your system with a single command.

### Basic System Info

```bash
ellie sysinfo
```

This command provides:
- Operating system details
- Hardware specifications
- Memory usage
- CPU information
- Disk space
- Network interfaces

### Network Status

```bash
ellie network-status
```

Get detailed network information including:
- Active network interfaces
- IP addresses
- Connection status
- Network services
- WiFi information (macOS/Linux)

### Current Working Directory

```bash
ellie pwd
```

Display the current working directory with additional context.

## File Operations

Ellie provides cross-platform file operations that work consistently across all operating systems.

### Listing Files and Directories

```bash
# List current directory
ellie list .

# List specific directory
ellie list ~/Documents

# List with details
ellie list ~/Projects --details
```

### Creating Files

```bash
# Create a new file
ellie create-file my-notes.txt

# Create file with content
ellie create-file config.json '{"name": "project", "version": "1.0.0"}'
```

### File Explorer Integration

```bash
# Open current directory in file explorer
ellie open-explorer

# Open specific path
ellie open ~/Documents
```

### Cross-Platform File Operations

Ellie's file operations work seamlessly across platforms:

- **Path handling**: Automatic path conversion for different OS
- **Permissions**: Proper permission handling
- **Encoding**: UTF-8 support across platforms
- **Special characters**: Safe handling of special characters in filenames

## Network Management

### WiFi Connection

```bash
# Connect to WiFi network
ellie connect-wifi "Network Name" "password"

# List available networks (macOS/Linux)
ellie network-status
```

### Network Diagnostics

```bash
# Test network connectivity
ellie run ping google.com

# Check DNS resolution
ellie run nslookup google.com
```

## Process Management

### Running Commands

```bash
# Execute system commands
ellie run ls -la
ellie run ps aux
ellie run top
```

### Background Processes

```bash
# Start background process
ellie run "npm start" &

# Monitor processes
ellie run htop
```

## Package Management

### Installing Packages

```bash
# Install packages
ellie install neofetch
ellie install htop
ellie install tree
```

### Updating System

```bash
# Update all packages
ellie update
```

### Cross-Platform Package Management

Ellie automatically uses the appropriate package manager:

- **macOS**: Homebrew
- **Ubuntu/Debian**: apt
- **CentOS/RHEL**: yum/dnf
- **Windows**: Chocolatey (when available)

## System Monitoring

### Resource Usage

```bash
# Monitor CPU and memory
ellie run top

# Disk usage
ellie run df -h

# Memory usage
ellie run free -h
```

### Log Monitoring

```bash
# View system logs
ellie run journalctl -f

# View application logs
ellie run tail -f /var/log/apache2/access.log
```

## Security Features

### File Permissions

```bash
# Set file permissions
ellie run chmod 644 myfile.txt

# Set directory permissions
ellie run chmod 755 mydirectory
```

### User Management

```bash
# Check current user
ellie whoami

# Switch user (with proper permissions)
ellie run sudo -u username command
```

## Automation

### Scheduled Tasks

```bash
# Create cron job
ellie run "crontab -e"

# List scheduled tasks
ellie run crontab -l
```

### System Maintenance

```bash
# Clean temporary files
ellie run rm -rf /tmp/*

# Update system packages
ellie update
```

## Troubleshooting

### Service Issues

```bash
# Check service status
ellie start list

# View service logs
ellie run journalctl -u apache2

# Restart problematic service
ellie restart apache
```

### Network Issues

```bash
# Test connectivity
ellie run ping 8.8.8.8

# Check DNS
ellie run nslookup google.com

# Reset network (macOS)
ellie run sudo dscacheutil -flushcache
```

### File System Issues

```bash
# Check disk space
ellie run df -h

# Find large files
ellie run find / -size +100M

# Repair disk (macOS)
ellie run diskutil verifyDisk /
```

## Best Practices

### 1. Service Management
- Always check service status before starting/stopping
- Use `ellie start all` for development environments
- Monitor service logs for issues

### 2. File Operations
- Use relative paths when possible
- Check file permissions before operations
- Backup important files before bulk operations

### 3. Network Management
- Test network connectivity before operations
- Use secure passwords for WiFi connections
- Monitor network usage regularly

### 4. System Monitoring
- Set up regular system health checks
- Monitor resource usage trends
- Keep system packages updated

## Advanced Features

### Custom Service Definitions

You can extend Ellie's service management by adding custom service definitions to your configuration.

### System Integration

Ellie integrates with:
- **Docker**: Container management
- **Kubernetes**: Cluster management
- **Cloud providers**: AWS, GCP, Azure
- **Monitoring tools**: Prometheus, Grafana

### Automation Scripts

Create automation scripts that use Ellie's system management features:

```bash
#!/bin/bash
# Daily system maintenance script

# Start development services
ellie start all

# Update packages
ellie update

# Clean temporary files
ellie run rm -rf /tmp/*

# Check system health
ellie sysinfo
```

## Next Steps

Explore more system management features:
- [Git Workflows](/docs/features/git-workflows) - Version control integration
- [AI Integration](/docs/features/ai-integration) - AI-powered system assistance
- [Productivity Tools](/docs/features/productivity) - Task and project management 