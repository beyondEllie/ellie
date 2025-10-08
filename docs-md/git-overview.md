# Git Operations Overview

Ellie provides comprehensive Git functionality with user-friendly commands. All Git operations are accessed through the `ellie git` command followed by a subcommand.

## Quick Reference

### Basic Operations

- `ellie git status` - Show repository status
- `ellie git commit` - Create conventional commits with guided prompts
- `ellie git push` - Push changes with commit message
- `ellie git pull` - Pull latest changes
- `ellie git init` - Initialize new repository
- `ellie git clone` - Clone remote repository
- `ellie git fetch` - Fetch from all remotes

### Branch Management

- `ellie git branch-create` - Create and switch to new branch
- `ellie git branch-switch` - Switch to existing branch
- `ellie git branch-delete` - Delete branch
- `ellie git branch-list` - List all branches
- `ellie git branch-list-remote` - List remote branches
- `ellie git branch-rename` - Rename branch

### Stash Operations

- `ellie git stash-save` - Save uncommitted changes
- `ellie git stash-pop` - Apply and remove latest stash
- `ellie git stash-list` - List all stashes
- `ellie git stash-show` - Show stash contents
- `ellie git stash-apply` - Apply stash without removing
- `ellie git stash-drop` - Remove stash

### Remote Management

- `ellie git remote-list` - List configured remotes
- `ellie git remote-add` - Add new remote
- `ellie git remote-remove` - Remove remote

### Log and History

- `ellie git log` - Pretty formatted commit history
- `ellie git log-search` - Search commit messages
- `ellie git log-author` - Filter commits by author
- `ellie git log-since` - Show commits since date
- `ellie git reflog` - Show reference log

### Diff Operations

- `ellie git diff` - Show unstaged changes
- `ellie git diff-staged` - Show staged changes
- `ellie git diff-branch` - Compare two branches

### Advanced Operations

- `ellie git merge` - Merge branch
- `ellie git rebase` - Rebase current branch
- `ellie git cherry-pick` - Apply specific commit
- `ellie git reset` - Reset to specific commit
- `ellie git revert` - Revert commit

### Tag Management

- `ellie git tag-create` - Create new tag
- `ellie git tag-list` - List all tags
- `ellie git tag-delete` - Delete tag

### Push Variants

- `ellie git push-tags` - Push all tags
- `ellie git push-force` - Force push with safety
- `ellie git push-upstream` - Push and set upstream

### Submodules

- `ellie git submodule-add` - Add submodule
- `ellie git submodule-update` - Update submodules
- `ellie git submodule-status` - Show submodule status

### Configuration

- `ellie git config-set-user` - Set user name and email
- `ellie git config-list` - Show all configuration
- `ellie git config-set-alias` - Create Git aliases

### Worktrees

- `ellie git worktree-add` - Create new worktree
- `ellie git worktree-list` - List worktrees
- `ellie git worktree-remove` - Remove worktree
- `ellie git worktree-prune` - Clean worktree info

### Bisect (Debugging)

- `ellie git bisect` - Start bisect session
- `ellie git bisect-good` - Mark commit as good
- `ellie git bisect-bad` - Mark commit as bad
- `ellie git bisect-reset` - End bisect session

### Maintenance

- `ellie git clean` - Remove untracked files
- `ellie git gc` - Garbage collection
- `ellie git fsck` - Check repository integrity

### Information

- `ellie git show` - Show commit details
- `ellie git blame` - Show file authorship
- `ellie git archive` - Create repository archive

## Features

- **User-friendly prompts**: All commands provide helpful prompts with examples
- **Safety measures**: Dangerous operations include confirmation prompts
- **No emojis**: Clean, professional output suitable for all environments
- **Comprehensive coverage**: Supports virtually all common Git workflows
- **Consistent interface**: All commands follow the same interaction patterns

## Getting Started

1. **Basic workflow**: `status` → `commit` → `push`
2. **Branch workflow**: `branch-create` → work → `merge`
3. **Collaboration**: `fetch` → `pull` → work → `push`
4. **Emergency**: `stash-save` → fix → `stash-pop`

For detailed documentation on specific command categories, see the individual documentation files.
