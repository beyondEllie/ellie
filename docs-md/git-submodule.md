# Git Submodule Management

Git submodules allow you to include other Git repositories as subdirectories in your project.

## Add Submodule

Adds a new submodule to your repository.

**Usage:**

```bash
ellie git submodule-add
```

You will be prompted for:

- Submodule URL (e.g., `https://github.com/user/library.git`)
- Local path (e.g., "libs/library")

**Example:**

```text
ellie git submodule-add
Submodule URL (https://github.com/user/repo.git) ➜ https://github.com/vendor/awesome-lib.git
Local path (submodules/repo) ➜ libs/awesome-lib
Submodule added at libs/awesome-lib
```

## Update Submodules

Updates all submodules to their latest commits, initializing them if necessary.

**Usage:**

```bash
ellie git submodule-update
```

This command runs `git submodule update --init --recursive` to:

- Initialize any new submodules
- Update existing submodules to their recorded commits
- Handle nested submodules recursively

## Show Submodule Status

Displays the status of all submodules in the repository.

**Usage:**

```bash
ellie git submodule-status
```

**Output Format:**
- `-` indicates the submodule is not initialized
- `+` indicates the submodule has uncommitted changes
- `U` indicates the submodule has merge conflicts

## Best Practices

- Always commit submodule updates in the parent repository
- Use `submodule-update` after cloning a repository with submodules
- Be careful when updating submodules to avoid breaking dependencies
- Consider using specific commit hashes rather than branches for stability
