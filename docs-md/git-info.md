# Git Information Operations

These commands help you gather information about your repository and its contents.

## Show Commit Information

Displays detailed information about a specific commit.

**Usage:**

```bash
ellie git show
```

You will be prompted for a commit hash (defaults to HEAD).

**Example:**

```text
ellie git show
Commit hash (optional) (HEAD) ➜ abc123f
```

**Output includes:**

- Commit metadata (author, date, message)
- Full diff showing all changes
- File statistics

## Blame File

Shows line-by-line authorship information for a file.

**Usage:**

```bash
ellie git blame
```

You will be prompted for the file path to examine.

**Example:**

```text
ellie git blame
File path (main.go) ➜ src/auth.go
```

**Output Format:**

```text
abc123f (John Doe 2024-01-15 14:30:25 +0000  1) package auth
def456g (Jane Smith 2024-01-16 09:15:10 +0000  2) 
def456g (Jane Smith 2024-01-16 09:15:10 +0000  3) import "fmt"
```

## Create Archive

Creates an archive (tar.gz or zip) of the repository at a specific reference.

**Usage:**

```bash
ellie git archive
```

You will be prompted for:

- Format (tar.gz, zip, etc.)
- Output filename
- Reference (branch, tag, or commit)

**Example:**

```text
ellie git archive
Format (tar/zip) (tar.gz) ➜ zip
Output filename (archive.tar.gz) ➜ project-v1.0.zip
Reference (branch/tag/commit) (HEAD) ➜ v1.0.0
Archive created: project-v1.0.zip
```

## Repository Initialization

Creates a new Git repository in the current directory.

**Usage:**

```bash
ellie git init
```

This creates a new `.git` directory and initializes an empty repository.

## Clone Repository

Clones a remote repository to your local machine.

**Usage:**

```bash
ellie git clone
```

You will be prompted for:

- Repository URL
- Local directory name (optional)

**Example:**

```text
ellie git clone
Repository URL (https://github.com/user/repo.git) ➜ https://github.com/awesome/project.git
Directory name (optional) ➜ my-project
Repository cloned
```

## Use Cases

- **show**: Reviewing specific commits during code review
- **blame**: Finding who last modified specific lines for debugging
- **archive**: Creating releases or backups without Git history
- **init**: Starting a new project with Git version control
- **clone**: Getting a copy of an existing project
