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
	description := getRequiredInput(reader, "📌 Description")
	body := getMultilineInput(reader, "💬 Body (optional)")
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
	styles.Cyan.Println("\n🔍 Repository Status")
	styles.Cyan.Println("───────────────────")
	runGitCommand("status", "-sb")
}

// GitBranchCreate creates a new branch
func GitBranchCreate() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🌿 Create New Branch")
	styles.Cyan.Println("────────────────────")
	branch := promptInput(reader, "Branch name", "feature/my-feature")
	if branch == "" {
		styles.ErrorStyle.Println("🚫 Branch name required")
		return
	}
	runGitCommand("checkout", "-b", branch)
	styles.SuccessStyle.Printf("✅ Created and switched to branch '%s'\n", branch)
}

// GitBranchSwitch switches to an existing branch
func GitBranchSwitch() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🔀 Switch Branch")
	styles.Cyan.Println("────────────────")
	branch := promptInput(reader, "Branch name", "main")
	if branch == "" {
		styles.ErrorStyle.Println("🚫 Branch name required")
		return
	}
	runGitCommand("checkout", branch)
	styles.SuccessStyle.Printf("✅ Switched to branch '%s'\n", branch)
}

// GitBranchDelete deletes a branch
func GitBranchDelete() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🗑️ Delete Branch")
	styles.Cyan.Println("────────────────")
	branch := promptInput(reader, "Branch name", "feature/my-feature")
	if branch == "" {
		styles.ErrorStyle.Println("🚫 Branch name required")
		return
	}
	runGitCommand("branch", "-d", branch)
	styles.SuccessStyle.Printf("✅ Deleted branch '%s'\n", branch)
}

// GitStashSave saves changes to a new stash
func GitStashSave() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n📦 Stash Changes")
	styles.Cyan.Println("────────────────")
	msg := promptInput(reader, "Stash message (optional)", "WIP")
	if msg == "" {
		runGitCommand("stash", "save")
	} else {
		runGitCommand("stash", "save", msg)
	}
	styles.SuccessStyle.Println("✅ Changes stashed")
}

// GitStashPop pops the latest stash
func GitStashPop() {
	styles.Cyan.Println("\n📤 Pop Stash")
	styles.Cyan.Println("────────────")
	runGitCommand("stash", "pop")
	styles.SuccessStyle.Println("✅ Stash applied")
}

// GitStashList lists all stashes
func GitStashList() {
	styles.Cyan.Println("\n📚 Stash List")
	styles.Cyan.Println("─────────────")
	runGitCommand("stash", "list")
}

// GitTagCreate creates a new tag
func GitTagCreate() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🏷️ Create Tag")
	styles.Cyan.Println("─────────────")
	tag := promptInput(reader, "Tag name", "v1.0.0")
	if tag == "" {
		styles.ErrorStyle.Println("🚫 Tag name required")
		return
	}
	runGitCommand("tag", tag)
	styles.SuccessStyle.Printf("✅ Tag '%s' created\n", tag)
}

// GitTagList lists all tags
func GitTagList() {
	styles.Cyan.Println("\n🏷️ Tag List")
	styles.Cyan.Println("───────────")
	runGitCommand("tag", "--list")
}

// GitTagDelete deletes a tag
func GitTagDelete() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🗑️ Delete Tag")
	styles.Cyan.Println("────────────")
	tag := promptInput(reader, "Tag name", "v1.0.0")
	if tag == "" {
		styles.ErrorStyle.Println("🚫 Tag name required")
		return
	}
	runGitCommand("tag", "-d", tag)
	styles.SuccessStyle.Printf("✅ Tag '%s' deleted\n", tag)
}

// GitLogPretty prints a pretty git log
func GitLogPretty() {
	styles.Cyan.Println("\n📜 Git Log")
	styles.Cyan.Println("──────────")
	runGitCommand("log", "--oneline", "--graph", "--decorate", "--all")
}

// GitDiff shows the diff
func GitDiff() {
	styles.Cyan.Println("\n🔍 Git Diff")
	styles.Cyan.Println("───────────")
	runGitCommand("diff")
}

// GitMerge merges a branch into the current branch
func GitMerge() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🔗 Merge Branch")
	styles.Cyan.Println("───────────────")
	branch := promptInput(reader, "Branch to merge", "feature/my-feature")
	if branch == "" {
		styles.ErrorStyle.Println("🚫 Branch name required")
		return
	}
	runGitCommand("merge", branch)
	styles.SuccessStyle.Printf("✅ Merged branch '%s'\n", branch)
}

// GitRebase rebases the current branch onto another
func GitRebase() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🔄 Rebase Branch")
	styles.Cyan.Println("────────────────")
	branch := promptInput(reader, "Branch to rebase onto", "main")
	if branch == "" {
		styles.ErrorStyle.Println("🚫 Branch name required")
		return
	}
	runGitCommand("rebase", branch)
	styles.SuccessStyle.Printf("✅ Rebased onto '%s'\n", branch)
}

// GitCherryPick cherry-picks a commit
func GitCherryPick() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🍒 Cherry-pick Commit")
	styles.Cyan.Println("─────────────────────")
	commit := promptInput(reader, "Commit hash", "abc1234")
	if commit == "" {
		styles.ErrorStyle.Println("🚫 Commit hash required")
		return
	}
	runGitCommand("cherry-pick", commit)
	styles.SuccessStyle.Printf("✅ Cherry-picked commit '%s'\n", commit)
}

// GitReset resets the current branch to a commit
func GitReset() {
	reader := bufio.NewReader(os.Stdin)
	styles.Cyan.Println("\n🔄 Reset Branch")
	styles.Cyan.Println("───────────────")
	commit := promptInput(reader, "Commit hash or ref", "HEAD~1")
	if commit == "" {
		styles.ErrorStyle.Println("🚫 Commit hash or ref required")
		return
	}
	runGitCommand("reset", "--hard", commit)
	styles.SuccessStyle.Printf("✅ Reset to '%s'\n", commit)
}

// GitBisect starts a bisect session
func GitBisect() {
	styles.Cyan.Println("\n🕵️ Git Bisect")
	styles.Cyan.Println("─────────────")
	styles.Yellow.Println("◎ This will help you find the commit that introduced a bug.")
	styles.Yellow.Println("◎ Use 'git bisect good' and 'git bisect bad' as prompted.")
	runGitCommand("bisect", "start")
}

// GitBisectGood marks the current commit as good
func GitBisectGood() {
	styles.Cyan.Println("\n✅ Mark Good Commit")
	styles.Cyan.Println("──────────────────")
	runGitCommand("bisect", "good")
}

// GitBisectBad marks the current commit as bad
func GitBisectBad() {
	styles.Cyan.Println("\n❌ Mark Bad Commit")
	styles.Cyan.Println("─────────────────")
	runGitCommand("bisect", "bad")
}

// GitBisectReset ends the bisect session
func GitBisectReset() {
	styles.Cyan.Println("\n🔄 Reset Bisect")
	styles.Cyan.Println("────────────────")
	runGitCommand("bisect", "reset")
}
