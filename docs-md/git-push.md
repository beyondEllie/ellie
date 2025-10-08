# Git Push Operations

Ellie provides enhanced push commands for various Git workflows.

## Push Tags

Pushes all local tags to the remote repository.

**Usage:**

```bash
ellie git push-tags
```

This command pushes all tags using `git push --tags`, making them available to other developers.

## Force Push

Performs a force push with safety measures.

**Usage:**

```bash
ellie git push-force
```

**Safety Features:**

- Uses `--force-with-lease` instead of `--force` for safety
- Confirms with user before executing
- Shows warning about overwriting remote history

**Example:**

```text
ellie git push-force
Force Push
──────────
WARNING: Force push can overwrite remote history!
Are you sure you want to force push? (Y/n) ➜ y
Force push completed
```

## Push with Upstream

Pushes current branch and sets upstream tracking.

**Usage:**

```bash
ellie git push-upstream
```

You will be prompted for the branch name, or leave empty to use the current branch.

**Example:**

```text
ellie git push-upstream
Branch name (current if empty) ➜ 
Push with upstream completed
```

This is equivalent to `git push -u origin HEAD` for new branches.

## When to Use Each Command

- **push-tags**: When releasing versions or sharing tags with the team
- **push-force**: When you need to rewrite history (use with extreme caution)
- **push-upstream**: When pushing a new branch for the first time