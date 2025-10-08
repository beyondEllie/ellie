# Git Worktree Management

Git worktrees allow you to have multiple working directories attached to the same repository, enabling you to work on different branches simultaneously.

## Add Worktree

Creates a new worktree at a specified path with a new or existing branch.

**Usage:**

```bash
ellie git worktree-add
```

You will be prompted for:

- Worktree path (e.g., "../feature-branch")
- Branch name (e.g., "feature/new-feature")

**Example:**

```text
ellie git worktree-add
Worktree path (../feature-branch) ➜ ../hotfix
Branch name (feature/new-feature) ➜ hotfix/urgent-fix
Worktree added at ../hotfix
```

## List Worktrees

Lists all worktrees with their paths and associated branches.

**Usage:**

```bash
ellie git worktree-list
```

**Output:**
Shows all worktrees with their absolute paths, HEAD commits, and branch information.

## Remove Worktree

Removes a worktree and cleans up its references.

**Usage:**

```bash
ellie git worktree-remove
```

You will be prompted for the worktree path to remove.

**Example:**

```text
ellie git worktree-remove
Worktree path (../feature-branch) ➜ ../hotfix
Worktree removed: ../hotfix
```

## Prune Worktrees

Cleans up worktree administrative information for worktrees that no longer exist.

**Usage:**

```bash
ellie git worktree-prune
```

This command removes stale worktree references and is useful for cleanup after manually deleting worktree directories.

## Benefits of Worktrees

- Work on multiple branches simultaneously
- No need to stash changes when switching contexts
- Each worktree maintains its own working directory state
- Shared Git history and configuration
- Faster than cloning multiple repositories