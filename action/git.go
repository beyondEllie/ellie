package actions

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/tacheraSasi/ellie/styles"
)

var (
	allowedTypes = []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"}
	successStyle = color.New(color.FgHiGreen, color.Bold)
)

// GitConventionalCommit
func GitConventionalCommit() {
	reader := bufio.NewReader(os.Stdin)

	styles.Cyan.Println("\n📝 Conventional Commit Builder")
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
		styles.ErrorStyle.Println("🚫 Commit canceled")
		os.Exit(0)
	}

	executeGitWorkflow(commitMessage)
	styles.SuccessStyle.Println("✅ Successfully committed and pushed!")
}

func getCommitType(reader *bufio.Reader) string {
	for {
		input := promptInput(reader, "🔧 Type", "feat, fix, docs, style, refactor, perf, test, chore, revert")
		if isValidCommitType(input) {
			return input
		}
		styles.ErrorStyle.Printf("⚠️  Invalid type: %s\n", input)
	}
}

func getScope(reader *bufio.Reader) string {
	return promptInput(reader, "🎯 Scope (optional)", "e.g., authentication")
}

func getRequiredInput(reader *bufio.Reader, label string) string {
	for {
		input := promptInput(reader, label, "")
		if input != "" { 
			return input
		}
		styles.ErrorStyle.Println("⚠️  This field is required")
	}
}

func getMultilineInput(reader *bufio.Reader, label string) string {
	styles.Cyan.Printf("\n%s\n", label)
	styles.Yellow.Println("◎ Press Enter twice to finish")
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
	if !confirmAction("💥 Breaking change?") {
		return "", false
	}
	return getRequiredInput(reader, "📣 Breaking change details"), true
}

func getIssueReference(reader *bufio.Reader) string {
	input := promptInput(reader, "🔗 Issue number (optional)", "e.g., 123")
	if input == "" {
		return ""
	}
	return fmt.Sprintf("Refs #%s", input)
}

func getTrailers(reader *bufio.Reader) []string {
	var trailers []string
	styles.Cyan.Println("\n🏷  Git Trailers (e.g., Reviewed-by: Name)")
	styles.Yellow.Println("◎ Leave empty to finish")
	for {
		input := promptInput(reader, "   Add trailer", "Key: Value")
		if input == "" {
			break
		}
		if isValidTrailer(input) {
			trailers = append(trailers, input)
		} else {
			styles.ErrorStyle.Println("⚠️  Invalid format. Use 'Key: Value'")
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
	styles.Magenta.Println("\n✨ Commit Preview:")
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
		styles.ErrorStyle.Printf("🚨 Git error: %v\n", err)
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
	styles.Cyan.Println("\n🚀 Quick Push")
	styles.Cyan.Println("─────────────")

	message := promptInput(reader, "💌 Message", "")
	if message == "" {
		styles.ErrorStyle.Println("🚫 Commit message required")
		os.Exit(1)
	}

	executeGitWorkflow("Ellie: " + message)
}

// GitPull executes git pull with feedback
func GitPull() {
	styles.Cyan.Println("\n🔄 Pulling Changes")
	styles.Cyan.Println("─────────────────")
	runGitCommand("pull")
	styles.SuccessStyle.Println("✅ Pull completed")
}

// GitStatus shows enhanced status output
func GitStatus() {
	styles.Cyan.Println("\n🔍 Repository Status")
	styles.Cyan.Println("───────────────────")
	runGitCommand("status", "-sb")
}
