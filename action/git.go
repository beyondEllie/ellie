package actions

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tacheraSasi/ellie/styles"
)

var (
	allowedTypes = []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"}
)

// GitConventionalCommit
func GitConventionalCommit() {
	reader := bufio.NewReader(os.Stdin)

	styles.Cyan.Println("\nConventional Commit Builder")
	styles.Cyan.Println("─────────────────────────────")

	commitType := getCommitType(reader)
	scope := getScope(reader)
	description := getRequiredInput(reader, "Description")
	body := getMultilineInput(reader, "Body (optional)")
	breakingDetail, isBreaking := getBreakingChange(reader)
	issueRef := getIssueReference(reader)
	trailers := getTrailers(reader)

	header := buildHeader(commitType, scope, description)
	commitMessage := buildCommitMessage(header, body, breakingDetail, isBreaking, issueRef, trailers)

	displayCommitPreview(commitMessage)
	if !confirmAction("Commit with this message?") {
		styles.ErrorStyle.Println("Commit canceled")
		os.Exit(0)
	}

	executeGitWorkflow(commitMessage)
	styles.SuccessStyle.Println("Successfully committed and pushed!")
}

func getCommitType(reader *bufio.Reader) string {
	for {
		input := promptInput(reader, "Type", "feat, fix, docs, style, refactor, perf, test, chore, revert")
		if isValidCommitType(input) {
			return input
		}
		styles.ErrorStyle.Printf("Invalid type: %s\n", input)
	}
}

func getScope(reader *bufio.Reader) string {
	return promptInput(reader, "Scope (optional)", "e.g., authentication")
}

func getRequiredInput(reader *bufio.Reader, label string) string {
	for {
		input := promptInput(reader, label, "")
		if input != "" {
			return input
		}
		styles.ErrorStyle.Println("This field is required")
	}
}

func getMultilineInput(reader *bufio.Reader, label string) string {
	styles.Cyan.Printf("\n%s\n", label)
	styles.Yellow.Println("Press Enter twice to finish")
	var lines []string
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" && len(lines) > 0 {
			break
		}
		if line != "" {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "\n")
}

func getBreakingChange(reader *bufio.Reader) (string, bool) {
	if !confirmAction("Breaking change?") {
		return "", false
	}
	return getRequiredInput(reader, "Breaking change details"), true
}

func getIssueReference(reader *bufio.Reader) string {
	input := promptInput(reader, "Issue number (optional)", "e.g., 123")
	if input == "" {
		return ""
	}
	return fmt.Sprintf("Refs #%s", input)
}

func getTrailers(reader *bufio.Reader) []string {
	var trailers []string
	styles.Cyan.Println("\nGit Trailers (e.g., Reviewed-by: Name)")
	styles.Yellow.Println("Leave empty to finish")
	for {
		input := promptInput(reader, "   Add trailer", "Key: Value")
		if input == "" {
			break
		}
		if isValidTrailer(input) {
			trailers = append(trailers, input)
		} else {
			styles.ErrorStyle.Println("Invalid format. Use 'Key: Value'")
		}
	}
	return trailers
}

func buildHeader(commitType, scope, description string) string {
	if scope != "" {
		return fmt.Sprintf("%s(%s): %s", commitType, scope, description)
	}
	return fmt.Sprintf("%s: %s", commitType, description)
}

func buildCommitMessage(header, body, breaking string, isBreaking bool, issue string, trailers []string) string {
	var msg strings.Builder
	msg.WriteString(header)

	if body != "" {
		msg.WriteString("\n\n" + body)
	}

	if isBreaking {
		msg.WriteString("\n\nBREAKING CHANGE: " + breaking)
	}

	if issue != "" {
		msg.WriteString("\n\n" + issue)
	}

	if len(trailers) > 0 {
		msg.WriteString("\n\n" + strings.Join(trailers, "\n"))
	}

	return msg.String()
}

func displayCommitPreview(message string) {
	styles.Magenta.Println("\nCommit Preview:")
	fmt.Println("──────────────────")
	fmt.Println(message)
	fmt.Println("──────────────────")
}

func executeGitWorkflow(message string) {
	runGitCommand("add", ".")
	runGitCommand("commit", "-m", message)
	runGitCommand("push")
}

func runGitCommand(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		styles.ErrorStyle.Printf("Git error: %v\n", err)
		os.Exit(1)
	}
}

// Helper functions
func promptInput(reader *bufio.Reader, label string, placeholder string) string {
	styles.Cyan.Printf("%s ", label)
	if placeholder != "" {
		styles.Yellow.Printf("(%s) ", placeholder)
	}
	fmt.Print("➜ ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func confirmAction(question string) bool {
	input := promptInput(bufio.NewReader(os.Stdin), question, "Y/n")
	return strings.ToLower(input) != "n"
}

func isValidCommitType(t string) bool {
	for _, allowed := range allowedTypes {
		if t == allowed {
			return true
		}
	}
	return false
}

func isValidTrailer(trailer string) bool {
	return strings.Contains(trailer, ":") && len(strings.Split(trailer, ":")) >= 2
}

// GitPush handles standard push workflow
func GitPush() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nQuick Push")
	styles.Cyan.Println("─────────────")

	message := promptInput(reader, "Message", "")
	if message == "" {
		styles.ErrorStyle.Println("Commit message required")
		os.Exit(1)
	}

	executeGitWorkflow("Ellie: " + message)
}

// GitPull executes git pull with feedback
func GitPull() {
	styles.Cyan.Println("\nPulling Changes")
	styles.Cyan.Println("─────────────────")
	runGitCommand("pull")
	styles.SuccessStyle.Println("Pull completed")
}

// GitStatus shows enhanced status output
func GitStatus() {
	styles.Cyan.Println("\nRepository Status")
	styles.Cyan.Println("───────────────────")
	runGitCommand("status", "-sb")
}

// GitBranchCreate creates a new branch
func GitBranchCreate() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nCreate New Branch")
	styles.Cyan.Println("────────────────────")
	branch := promptInput(reader, "Branch name", "feature/my-feature")
	if branch == "" {
		styles.ErrorStyle.Println("Branch name required")
		return
	}
	runGitCommand("checkout", "-b", branch)
	styles.SuccessStyle.Printf("Created and switched to branch '%s'\n", branch)
}

// GitBranchSwitch switches to an existing branch
func GitBranchSwitch() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nSwitch Branch")
	styles.Cyan.Println("────────────────")
	branch := promptInput(reader, "Branch name", "main")
	if branch == "" {
		styles.ErrorStyle.Println("Branch name required")
		return
	}
	runGitCommand("checkout", branch)
	styles.SuccessStyle.Printf("Switched to branch '%s'\n", branch)
}

// GitBranchDelete deletes a branch
func GitBranchDelete() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nDelete Branch")
	styles.Cyan.Println("────────────────")
	branch := promptInput(reader, "Branch name", "feature/my-feature")
	if branch == "" {
		styles.ErrorStyle.Println("Branch name required")
		return
	}
	runGitCommand("branch", "-d", branch)
	styles.SuccessStyle.Printf("Deleted branch '%s'\n", branch)
}

// GitStashSave saves changes to a new stash
func GitStashSave() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nStash Changes")
	styles.Cyan.Println("────────────────")
	msg := promptInput(reader, "Stash message (optional)", "WIP")
	if msg == "" {
		runGitCommand("stash", "save")
	} else {
		runGitCommand("stash", "save", msg)
	}
	styles.SuccessStyle.Println("Changes stashed")
}

// GitStashPop pops the latest stash
func GitStashPop() {
	styles.Cyan.Println("\nPop Stash")
	styles.Cyan.Println("────────────")
	runGitCommand("stash", "pop")
	styles.SuccessStyle.Println("Stash applied")
}

// GitStashList lists all stashes
func GitStashList() {
	styles.Cyan.Println("\nStash List")
	styles.Cyan.Println("─────────────")
	runGitCommand("stash", "list")
}

// GitTagCreate creates a new tag
func GitTagCreate() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nCreate Tag")
	styles.Cyan.Println("─────────────")
	tag := promptInput(reader, "Tag name", "v1.0.0")
	if tag == "" {
		styles.ErrorStyle.Println("Tag name required")
		return
	}
	runGitCommand("tag", tag)
	styles.SuccessStyle.Printf("Tag '%s' created\n", tag)
}

// GitTagList lists all tags
func GitTagList() {
	styles.Cyan.Println("\nTag List")
	styles.Cyan.Println("───────────")
	runGitCommand("tag", "--list")
}

// GitTagDelete deletes a tag
func GitTagDelete() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nDelete Tag")
	styles.Cyan.Println("────────────")
	tag := promptInput(reader, "Tag name", "v1.0.0")
	if tag == "" {
		styles.ErrorStyle.Println("Tag name required")
		return
	}
	runGitCommand("tag", "-d", tag)
	styles.SuccessStyle.Printf("Tag '%s' deleted\n", tag)
}

// GitLogPretty prints a pretty git log
func GitLogPretty() {
	styles.Cyan.Println("\nGit Log")
	styles.Cyan.Println("──────────")
	runGitCommand("log", "--oneline", "--graph", "--decorate", "--all")
}

// GitDiff shows the diff
func GitDiff() {
	styles.Cyan.Println("\nGit Diff")
	styles.Cyan.Println("───────────")
	runGitCommand("diff")
}

// GitMerge merges a branch into the current branch
func GitMerge() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nMerge Branch")
	styles.Cyan.Println("───────────────")
	branch := promptInput(reader, "Branch to merge", "feature/my-feature")
	if branch == "" {
		styles.ErrorStyle.Println("Branch name required")
		return
	}
	runGitCommand("merge", branch)
	styles.SuccessStyle.Printf("Merged branch '%s'\n", branch)
}

// GitRebase rebases the current branch onto another
func GitRebase() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nRebase Branch")
	styles.Cyan.Println("────────────────")
	branch := promptInput(reader, "Branch to rebase onto", "main")
	if branch == "" {
		styles.ErrorStyle.Println("Branch name required")
		return
	}
	runGitCommand("rebase", branch)
	styles.SuccessStyle.Printf("Rebased onto '%s'\n", branch)
}

// GitCherryPick cherry-picks a commit
func GitCherryPick() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nCherry-pick Commit")
	styles.Cyan.Println("─────────────────────")
	commit := promptInput(reader, "Commit hash", "abc1234")
	if commit == "" {
		styles.ErrorStyle.Println("Commit hash required")
		return
	}
	runGitCommand("cherry-pick", commit)
	styles.SuccessStyle.Printf("Cherry-picked commit '%s'\n", commit)
}

// GitReset resets the current branch to a commit
func GitReset() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nReset Branch")
	styles.Cyan.Println("───────────────")
	commit := promptInput(reader, "Commit hash or ref", "HEAD~1")
	if commit == "" {
		styles.ErrorStyle.Println("Commit hash or ref required")
		return
	}
	runGitCommand("reset", "--hard", commit)
	styles.SuccessStyle.Printf("Reset to '%s'\n", commit)
}

// GitBisect starts a bisect session
func GitBisect() {
	styles.Cyan.Println("\nGit Bisect")
	styles.Cyan.Println("─────────────")
	styles.Yellow.Println("This will help you find the commit that introduced a bug.")
	styles.Yellow.Println("Use 'git bisect good' and 'git bisect bad' as prompted.")
	runGitCommand("bisect", "start")
}

// GitBisectGood marks the current commit as good
func GitBisectGood() {
	styles.Cyan.Println("\nMark Good Commit")
	styles.Cyan.Println("──────────────────")
	runGitCommand("bisect", "good")
}

// GitBisectBad marks the current commit as bad
func GitBisectBad() {
	styles.Cyan.Println("\nMark Bad Commit")
	styles.Cyan.Println("─────────────────")
	runGitCommand("bisect", "bad")
}

// GitBisectReset ends the bisect session
func GitBisectReset() {
	styles.Cyan.Println("\nReset Bisect")
	styles.Cyan.Println("────────────────")
	runGitCommand("bisect", "reset")
}

// GitRemoteList lists all remotes
func GitRemoteList() {
	styles.Cyan.Println("\nRemote List")
	styles.Cyan.Println("───────────")
	runGitCommand("remote", "-v")
}

// GitRemoteAdd adds a new remote
func GitRemoteAdd() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nAdd Remote")
	styles.Cyan.Println("──────────")
	name := promptInput(reader, "Remote name", "origin")
	if name == "" {
		styles.ErrorStyle.Println("Remote name required")
		return
	}
	url := promptInput(reader, "Remote URL", "https://github.com/user/repo.git")
	if url == "" {
		styles.ErrorStyle.Println("Remote URL required")
		return
	}
	runGitCommand("remote", "add", name, url)
	styles.SuccessStyle.Printf("Remote '%s' added\n", name)
}

// GitRemoteRemove removes a remote
func GitRemoteRemove() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nRemove Remote")
	styles.Cyan.Println("─────────────")
	name := promptInput(reader, "Remote name", "origin")
	if name == "" {
		styles.ErrorStyle.Println("Remote name required")
		return
	}
	runGitCommand("remote", "remove", name)
	styles.SuccessStyle.Printf("Remote '%s' removed\n", name)
}

// GitFetch fetches from a remote
func GitFetch() {
	styles.Cyan.Println("\nFetch Changes")
	styles.Cyan.Println("─────────────")
	runGitCommand("fetch", "--all")
	styles.SuccessStyle.Println("Fetch completed")
}

// GitReflog shows the reflog
func GitReflog() {
	styles.Cyan.Println("\nReflog")
	styles.Cyan.Println("──────")
	runGitCommand("reflog", "--oneline")
}

// GitClean removes untracked files
func GitClean() {
	styles.Cyan.Println("\nClean Untracked Files")
	styles.Cyan.Println("─────────────────────")
	styles.Yellow.Println("This will remove untracked files. Use with caution!")
	if !confirmAction("Continue with clean?") {
		styles.ErrorStyle.Println("Clean canceled")
		return
	}
	runGitCommand("clean", "-fd")
	styles.SuccessStyle.Println("Clean completed")
}

// GitArchive creates an archive of the repository
func GitArchive() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nCreate Archive")
	styles.Cyan.Println("──────────────")
	format := promptInput(reader, "Format (tar/zip)", "tar.gz")
	if format == "" {
		format = "tar.gz"
	}
	filename := promptInput(reader, "Output filename", "archive.tar.gz")
	if filename == "" {
		filename = "archive.tar.gz"
	}
	ref := promptInput(reader, "Reference (branch/tag/commit)", "HEAD")
	if ref == "" {
		ref = "HEAD"
	}
	runGitCommand("archive", "--format="+format, "--output="+filename, ref)
	styles.SuccessStyle.Printf("Archive created: %s\n", filename)
}

// GitBlame shows blame information for a file
func GitBlame() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nBlame File")
	styles.Cyan.Println("──────────")
	file := promptInput(reader, "File path", "main.go")
	if file == "" {
		styles.ErrorStyle.Println("File path required")
		return
	}
	runGitCommand("blame", file)
}

// GitSubmoduleAdd adds a submodule
func GitSubmoduleAdd() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nAdd Submodule")
	styles.Cyan.Println("─────────────")
	url := promptInput(reader, "Submodule URL", "https://github.com/user/repo.git")
	if url == "" {
		styles.ErrorStyle.Println("Submodule URL required")
		return
	}
	path := promptInput(reader, "Local path", "submodules/repo")
	if path == "" {
		styles.ErrorStyle.Println("Local path required")
		return
	}
	runGitCommand("submodule", "add", url, path)
	styles.SuccessStyle.Printf("Submodule added at %s\n", path)
}

// GitSubmoduleUpdate updates all submodules
func GitSubmoduleUpdate() {
	styles.Cyan.Println("\nUpdate Submodules")
	styles.Cyan.Println("─────────────────")
	runGitCommand("submodule", "update", "--init", "--recursive")
	styles.SuccessStyle.Println("Submodules updated")
}

// GitSubmoduleStatus shows submodule status
func GitSubmoduleStatus() {
	styles.Cyan.Println("\nSubmodule Status")
	styles.Cyan.Println("────────────────")
	runGitCommand("submodule", "status", "--recursive")
}

// GitConfigSetUser sets user name and email
func GitConfigSetUser() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nSet User Configuration")
	styles.Cyan.Println("──────────────────────")
	name := promptInput(reader, "User name", "John Doe")
	if name == "" {
		styles.ErrorStyle.Println("User name required")
		return
	}
	email := promptInput(reader, "User email", "john@example.com")
	if email == "" {
		styles.ErrorStyle.Println("User email required")
		return
	}
	runGitCommand("config", "user.name", name)
	runGitCommand("config", "user.email", email)
	styles.SuccessStyle.Println("User configuration updated")
}

// GitConfigList lists all configuration
func GitConfigList() {
	styles.Cyan.Println("\nGit Configuration")
	styles.Cyan.Println("─────────────────")
	runGitCommand("config", "--list")
}

// GitConfigSetAlias sets a Git alias
func GitConfigSetAlias() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nSet Git Alias")
	styles.Cyan.Println("─────────────")
	alias := promptInput(reader, "Alias name", "st")
	if alias == "" {
		styles.ErrorStyle.Println("Alias name required")
		return
	}
	command := promptInput(reader, "Git command", "status")
	if command == "" {
		styles.ErrorStyle.Println("Git command required")
		return
	}
	runGitCommand("config", "alias."+alias, command)
	styles.SuccessStyle.Printf("Alias '%s' set to '%s'\n", alias, command)
}

// GitWorktreeAdd adds a new worktree
func GitWorktreeAdd() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nAdd Worktree")
	styles.Cyan.Println("────────────")
	path := promptInput(reader, "Worktree path", "../feature-branch")
	if path == "" {
		styles.ErrorStyle.Println("Worktree path required")
		return
	}
	branch := promptInput(reader, "Branch name", "feature/new-feature")
	if branch == "" {
		styles.ErrorStyle.Println("Branch name required")
		return
	}
	runGitCommand("worktree", "add", path, branch)
	styles.SuccessStyle.Printf("Worktree added at %s\n", path)
}

// GitWorktreeList lists all worktrees
func GitWorktreeList() {
	styles.Cyan.Println("\nWorktree List")
	styles.Cyan.Println("─────────────")
	runGitCommand("worktree", "list")
}

// GitWorktreeRemove removes a worktree
func GitWorktreeRemove() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nRemove Worktree")
	styles.Cyan.Println("───────────────")
	path := promptInput(reader, "Worktree path", "../feature-branch")
	if path == "" {
		styles.ErrorStyle.Println("Worktree path required")
		return
	}
	runGitCommand("worktree", "remove", path)
	styles.SuccessStyle.Printf("Worktree removed: %s\n", path)
}

// GitWorktreePrune cleans up worktree information
func GitWorktreePrune() {
	styles.Cyan.Println("\nPrune Worktrees")
	styles.Cyan.Println("───────────────")
	runGitCommand("worktree", "prune")
	styles.SuccessStyle.Println("Worktrees pruned")
}

// GitInit initializes a new Git repository
func GitInit() {
	styles.Cyan.Println("\nInitialize Repository")
	styles.Cyan.Println("─────────────────────")
	runGitCommand("init")
	styles.SuccessStyle.Println("Git repository initialized")
}

// GitClone clones a repository
func GitClone() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nClone Repository")
	styles.Cyan.Println("────────────────")
	url := promptInput(reader, "Repository URL", "https://github.com/user/repo.git")
	if url == "" {
		styles.ErrorStyle.Println("Repository URL required")
		return
	}
	dir := promptInput(reader, "Directory name (optional)", "")
	if dir == "" {
		runGitCommand("clone", url)
	} else {
		runGitCommand("clone", url, dir)
	}
	styles.SuccessStyle.Println("Repository cloned")
}

// GitShow shows information about a commit
func GitShow() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nShow Commit")
	styles.Cyan.Println("───────────")
	commit := promptInput(reader, "Commit hash (optional)", "HEAD")
	if commit == "" {
		commit = "HEAD"
	}
	runGitCommand("show", commit)
}

// GitRevert reverts a commit
func GitRevert() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nRevert Commit")
	styles.Cyan.Println("─────────────")
	commit := promptInput(reader, "Commit hash", "abc1234")
	if commit == "" {
		styles.ErrorStyle.Println("Commit hash required")
		return
	}
	runGitCommand("revert", commit)
	styles.SuccessStyle.Printf("Reverted commit '%s'\n", commit)
}

// GitGC performs garbage collection
func GitGC() {
	styles.Cyan.Println("\nGarbage Collection")
	styles.Cyan.Println("──────────────────")
	runGitCommand("gc", "--aggressive", "--prune=now")
	styles.SuccessStyle.Println("Garbage collection completed")
}

// GitFsck checks repository integrity
func GitFsck() {
	styles.Cyan.Println("\nCheck Repository Integrity")
	styles.Cyan.Println("──────────────────────────")
	runGitCommand("fsck", "--full")
}

// GitPushTags pushes all tags to remote
func GitPushTags() {
	styles.Cyan.Println("\nPush All Tags")
	styles.Cyan.Println("─────────────")
	runGitCommand("push", "--tags")
	styles.SuccessStyle.Println("All tags pushed to remote")
}

// GitPushForce force pushes changes
func GitPushForce() {
	styles.Cyan.Println("\nForce Push")
	styles.Cyan.Println("──────────")
	styles.Yellow.Println("WARNING: Force push can overwrite remote history!")
	if !confirmAction("Are you sure you want to force push?") {
		styles.ErrorStyle.Println("Force push canceled")
		return
	}
	runGitCommand("push", "--force-with-lease")
	styles.SuccessStyle.Println("Force push completed")
}

// GitPushUpstream pushes and sets upstream branch
func GitPushUpstream() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nPush and Set Upstream")
	styles.Cyan.Println("─────────────────────")
	branch := promptInput(reader, "Branch name (current if empty)", "")
	if branch == "" {
		runGitCommand("push", "-u", "origin", "HEAD")
	} else {
		runGitCommand("push", "-u", "origin", branch)
	}
	styles.SuccessStyle.Println("Push with upstream completed")
}

// GitBranchList lists all branches
func GitBranchList() {
	styles.Cyan.Println("\nBranch List")
	styles.Cyan.Println("───────────")
	runGitCommand("branch", "-a")
}

// GitBranchListRemote lists remote branches
func GitBranchListRemote() {
	styles.Cyan.Println("\nRemote Branches")
	styles.Cyan.Println("───────────────")
	runGitCommand("branch", "-r")
}

// GitBranchRename renames a branch
func GitBranchRename() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nRename Branch")
	styles.Cyan.Println("─────────────")
	oldName := promptInput(reader, "Current branch name", "old-feature")
	if oldName == "" {
		styles.ErrorStyle.Println("Current branch name required")
		return
	}
	newName := promptInput(reader, "New branch name", "new-feature")
	if newName == "" {
		styles.ErrorStyle.Println("New branch name required")
		return
	}
	runGitCommand("branch", "-m", oldName, newName)
	styles.SuccessStyle.Printf("Branch renamed from '%s' to '%s'\n", oldName, newName)
}

// GitLogSearch searches commit history
func GitLogSearch() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nSearch Commit History")
	styles.Cyan.Println("─────────────────────")
	query := promptInput(reader, "Search term", "bug fix")
	if query == "" {
		styles.ErrorStyle.Println("Search term required")
		return
	}
	runGitCommand("log", "--grep="+query, "--oneline")
}

// GitLogAuthor shows commits by author
func GitLogAuthor() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nCommits by Author")
	styles.Cyan.Println("─────────────────")
	author := promptInput(reader, "Author name or email", "john@example.com")
	if author == "" {
		styles.ErrorStyle.Println("Author required")
		return
	}
	runGitCommand("log", "--author="+author, "--oneline")
}

// GitLogSince shows commits since a date
func GitLogSince() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nCommits Since Date")
	styles.Cyan.Println("──────────────────")
	date := promptInput(reader, "Date (YYYY-MM-DD)", "2023-01-01")
	if date == "" {
		styles.ErrorStyle.Println("Date required")
		return
	}
	runGitCommand("log", "--since="+date, "--oneline")
}

// GitDiffStaged shows staged changes
func GitDiffStaged() {
	styles.Cyan.Println("\nStaged Changes")
	styles.Cyan.Println("──────────────")
	runGitCommand("diff", "--staged")
}

// GitDiffBranch compares two branches
func GitDiffBranch() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nCompare Branches")
	styles.Cyan.Println("────────────────")
	branch1 := promptInput(reader, "First branch", "main")
	if branch1 == "" {
		styles.ErrorStyle.Println("First branch required")
		return
	}
	branch2 := promptInput(reader, "Second branch", "feature")
	if branch2 == "" {
		styles.ErrorStyle.Println("Second branch required")
		return
	}
	runGitCommand("diff", branch1+".."+branch2)
}

// GitStashShow shows stash contents
func GitStashShow() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nShow Stash Contents")
	styles.Cyan.Println("───────────────────")
	stash := promptInput(reader, "Stash reference (optional)", "stash@{0}")
	if stash == "" {
		stash = "stash@{0}"
	}
	runGitCommand("stash", "show", "-p", stash)
}

// GitStashDrop drops a stash
func GitStashDrop() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nDrop Stash")
	styles.Cyan.Println("──────────")
	stash := promptInput(reader, "Stash reference (optional)", "stash@{0}")
	if stash == "" {
		stash = "stash@{0}"
	}
	runGitCommand("stash", "drop", stash)
	styles.SuccessStyle.Printf("Stash %s dropped\n", stash)
}

// GitStashApply applies a stash without removing it
func GitStashApply() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\nApply Stash")
	styles.Cyan.Println("───────────")
	stash := promptInput(reader, "Stash reference (optional)", "stash@{0}")
	if stash == "" {
		stash = "stash@{0}"
	}
	runGitCommand("stash", "apply", stash)
	styles.SuccessStyle.Printf("Stash %s applied\n", stash)
}
