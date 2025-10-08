# Enhanced Git Diff Operations

Ellie provides specialized diff commands for comparing changes in various contexts.

## Standard Diff

Shows unstaged changes in the working directory.

**Usage:**

```bash
ellie git diff
```

Displays differences between the working directory and the index (staged area).

## Staged Changes Diff

Shows staged changes that will be included in the next commit.

**Usage:**

```bash
ellie git diff-staged
```

Equivalent to `git diff --staged`, showing what will be committed.

## Branch Comparison

Compares changes between two branches.

**Usage:**

```bash
ellie git diff-branch
```

You will be prompted for two branch names to compare.

**Example:**

```text
ellie git diff-branch
First branch (main) ➜ main
Second branch (feature) ➜ feature/user-auth
```

This shows differences between the `main` and `feature/user-auth` branches.

## Understanding Diff Output

- **Red lines** (starting with `-`): Removed content
- **Green lines** (starting with `+`): Added content
- **Context lines**: Unchanged lines shown for reference
- **File headers**: Show which files are being compared

## Common Workflows

1. **Before committing**: Use `diff-staged` to review what you're about to commit
2. **Code review**: Use `diff-branch` to see what changed in a feature branch
3. **Debugging**: Use standard `diff` to see what you've changed since last commit