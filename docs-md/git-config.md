# Git Configuration Management

Ellie provides user-friendly commands for managing Git configuration.

## Set User Information

Sets the user name and email for Git commits.

**Usage:**

```bash
ellie git config-set-user
```

You will be prompted for:

- User name (e.g., "John Doe")
- User email (e.g., "john@example.com")

**Example:**

```text
ellie git config-set-user
User name (John Doe) ➜ Jane Smith
User email (john@example.com) ➜ jane@company.com
User configuration updated
```

## List Configuration

Lists all Git configuration settings.

**Usage:**

```bash
ellie git config-list
```

**Output:**
Shows all Git configuration in `key=value` format from local, global, and system configs.

## Set Alias

Creates a Git alias for common commands.

**Usage:**

```bash
ellie git config-set-alias
```

You will be prompted for:

- Alias name (e.g., "st", "co", "br")
- Git command (e.g., "status", "checkout", "branch")

**Example:**

```text
ellie git config-set-alias
Alias name (st) ➜ co
Git command (status) ➜ checkout
Alias 'co' set to 'checkout'
```

**Common Aliases:**

- `st` → `status`
- `co` → `checkout`
- `br` → `branch`
- `ci` → `commit`
- `unstage` → `reset HEAD --`
