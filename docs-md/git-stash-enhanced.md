# Enhanced Git Stash Operations

Git stash allows you to temporarily save uncommitted changes and restore them later.

## Save Stash

Saves current uncommitted changes to a new stash.

**Usage:**

```bash
ellie git stash-save
```

You will be prompted for an optional stash message.

**Example:**

```text
ellie git stash-save
Stash message (optional) (WIP) ➜ work in progress on authentication
Changes stashed
```

## Pop Stash

Applies the most recent stash and removes it from the stash list.

**Usage:**

```bash
ellie git stash-pop
```

This is equivalent to applying the stash and then dropping it.

## List Stashes

Shows all saved stashes with their messages and references.

**Usage:**

```bash
ellie git stash-list
```

**Output Format:**

```text
stash@{0}: WIP on main: 1234567 Latest commit message
stash@{1}: On feature: work in progress on authentication
```

## Show Stash Contents

Displays the changes contained in a specific stash.

**Usage:**

```bash
ellie git stash-show
```

You will be prompted for a stash reference (defaults to `stash@{0}`).

**Example:**

```text
ellie git stash-show
Stash reference (optional) (stash@{0}) ➜ stash@{1}
```

## Apply Stash

Applies a stash without removing it from the stash list.

**Usage:**

```bash
ellie git stash-apply
```

You will be prompted for a stash reference (defaults to `stash@{0}`).

## Drop Stash

Removes a stash from the stash list without applying it.

**Usage:**

```bash
ellie git stash-drop
```

You will be prompted for a stash reference (defaults to `stash@{0}`).

## Stash Workflow

1. **Before switching branches**: Use `stash-save` to save uncommitted work
2. **Review stashes**: Use `stash-list` and `stash-show` to see what's saved
3. **Restore work**: Use `stash-pop` to restore and remove, or `stash-apply` to keep the stash
4. **Clean up**: Use `stash-drop` to remove stashes you no longer need
