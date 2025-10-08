# Git Remote Management

Ellie provides comprehensive commands for managing Git remotes.

## List Remotes

Lists all configured remotes with their URLs.

**Usage:**

```bash
ellie git remote-list
```

**Output:**
Shows all remotes in verbose format with both fetch and push URLs.

## Add Remote

Adds a new remote repository.

**Usage:**

```bash
ellie git remote-add
```

You will be prompted for:

- Remote name (e.g., "origin", "upstream")
- Remote URL (e.g., `https://github.com/user/repo.git`)

**Example:**

```text
ellie git remote-add
Remote name (origin) ➜ upstream
Remote URL (https://github.com/user/repo.git) ➜ https://github.com/original/repo.git
Remote 'upstream' added
```

## Remove Remote

Removes an existing remote repository.

**Usage:**

```bash
ellie git remote-remove
```

You will be prompted for the remote name to remove.

**Example:**

```text
ellie git remote-remove
Remote name (origin) ➜ old-remote
Remote 'old-remote' removed
```

## Fetch Changes

Fetches changes from all remotes without merging.

**Usage:**

```bash
ellie git fetch
```

This command fetches from all configured remotes using `git fetch --all`.