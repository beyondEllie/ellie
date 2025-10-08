# Git Maintenance Operations

Ellie provides commands for maintaining and optimizing your Git repository.

## Reflog

Shows the reference log, which tracks when branch tips and other references were updated.

**Usage:**

```bash
ellie git reflog
```

**Use Cases:**

- Recover lost commits after a reset
- See history of branch switches
- Find commits that aren't in any branch

## Clean Untracked Files

Removes untracked files and directories from the working tree.

**Usage:**

```bash
ellie git clean
```

**Safety Features:**

- Shows confirmation prompt before proceeding
- Warning message about permanent deletion
- Uses `-fd` flags to remove files and directories

**Example:**

```text
ellie git clean
Clean Untracked Files
─────────────────────
This will remove untracked files. Use with caution!
Continue with clean? (Y/n) ➜ y
Clean completed
```

## Garbage Collection

Optimizes the repository by cleaning up unnecessary files and optimizing storage.

**Usage:**

```bash
ellie git gc
```

This runs `git gc --aggressive --prune=now` to:

- Remove unreachable objects
- Pack loose objects
- Optimize repository structure
- Free up disk space

## Repository Integrity Check

Checks the integrity and connectivity of objects in the repository.

**Usage:**

```bash
ellie git fsck
```

This runs `git fsck --full` to verify:

- Object integrity
- Reference validity  
- Repository consistency
- Detect corruption

## When to Use These Commands

- **reflog**: When you need to recover lost work or understand recent changes
- **clean**: Before important operations to ensure a clean working directory
- **gc**: Periodically to optimize repository performance and size
- **fsck**: When suspicious of repository corruption or after system crashes

## Safety Notes

- **clean** permanently removes files - ensure you don't need any untracked files
- **gc** is generally safe but may take time on large repositories
- **fsck** is read-only and safe to run anytime
