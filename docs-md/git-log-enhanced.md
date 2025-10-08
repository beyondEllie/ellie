# Enhanced Git Log Operations

Ellie provides powerful commands for exploring Git history and finding specific commits.

## Standard Log

Displays a pretty formatted git log with graph visualization.

**Usage:**

```bash
ellie git log
```

Shows commit history with `--oneline`, `--graph`, `--decorate`, and `--all` flags for a comprehensive view.

## Search Commit History

Searches for commits containing specific text in their commit messages.

**Usage:**

```bash
ellie git log-search
```

You will be prompted for a search term to find in commit messages.

**Example:**

```text
ellie git log-search
Search term (bug fix) ➜ authentication
```

This searches for commits mentioning "authentication" in their messages.

## Filter by Author

Shows commits made by a specific author.

**Usage:**

```bash
ellie git log-author
```

You will be prompted for an author name or email address.

**Example:**

```text
ellie git log-author
Author name or email (john@example.com) ➜ jane.smith@company.com
```

## Filter by Date

Shows commits made since a specific date.

**Usage:**

```bash
ellie git log-since
```

You will be prompted for a date in YYYY-MM-DD format.

**Example:**

```text
ellie git log-since
Date (YYYY-MM-DD) (2023-01-01) ➜ 2024-01-15
```

Shows all commits made since January 15th, 2024.

## Use Cases

- **log-search**: Find commits related to specific features or bug fixes
- **log-author**: Review contributions by team members
- **log-since**: See recent changes or changes since a milestone