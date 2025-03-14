package actions

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tacheraSasi/ellie/utils"
)

// GitConventionalCommit interactively builds a commit message per the Conventional Commits spec
// (see https://www.conventionalcommits.org/en/v1.0.0/).
func GitConventionalCommit() {
	reader := bufio.NewReader(os.Stdin)
	allowedTypes := []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "revert"}

	// Validate commit type.
	var commitType string
	for {
		fmt.Print("Enter commit type (feat, fix, docs, style, refactor, perf, test, chore, revert): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
		commitType = strings.TrimSpace(input)
		if isValidCommitType(commitType, allowedTypes) {
			break
		}
		fmt.Println("Invalid commit type. Please choose a valid type.")
	}
   
	// Get optional scope.
	fmt.Print("Enter commit scope (optional): ")
	scope, _ := reader.ReadString('\n')
	scope = strings.TrimSpace(scope)

	// Get commit description.
	fmt.Print("Enter commit description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading description:", err)
		os.Exit(1)
	}
	description = strings.TrimSpace(description)

	// Optional commit body.
	fmt.Print("Enter commit body (optional): ")
	body, _ := reader.ReadString('\n')
	body = strings.TrimSpace(body)

	// Ask for breaking change.
	fmt.Print("Is this a breaking change? (y/n): ")
	breakingInput, _ := reader.ReadString('\n')
	breakingInput = strings.TrimSpace(breakingInput)
	var breakingDetail string
	if strings.ToLower(breakingInput) == "y" {
		fmt.Print("Enter breaking change details: ")
		breakingDetail, _ = reader.ReadString('\n')
		breakingDetail = strings.TrimSpace(breakingDetail)
	}

	// Build commit header.
	header := ""
	if scope != "" {
		header = fmt.Sprintf("%s(%s): %s", commitType, scope, description)
	} else {
		header = fmt.Sprintf("%s: %s", commitType, description)
	}

	// Combine header, body, and breaking changes.
	commitMessage := header
	if body != "" {
		commitMessage += "\n\n" + body
	}
	if breakingDetail != "" {
		commitMessage += "\n\nBREAKING CHANGE: " + breakingDetail
	}

	fmt.Println("\nFinal commit message:")
	fmt.Println(commitMessage)
	fmt.Println()

	// Execute Git commands in sequence.
	if err := runGitCommand("add", "."); err != nil {
		fmt.Println("Error during git add:", err)
		os.Exit(1)
	}

	if err := runGitCommand("commit", "-m", commitMessage); err != nil {
		fmt.Println("Error during git commit:", err)
		os.Exit(1)
	}

	if err := runGitCommand("push"); err != nil {
		fmt.Println("Error during git push:", err)
		os.Exit(1)
	}

	fmt.Println("Git operations completed successfully.")
}

// runGitCommand executes a git command with the given arguments.
func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// isValidCommitType checks whether the provided commit type is in the list of allowed types.
func isValidCommitType(commitType string, allowed []string) bool {
	for _, t := range allowed {
		if commitType == t {
			return true
		}
	}
	return false
}

// GitPush executes a standard push command. This can be used for simpler commit flows.
func GitPush() {
	commitMsg, err := utils.GetInput("Enter commit message: ")
	if err != nil {
		fmt.Println("Error reading commit message:", err)
		os.Exit(1)
	}

	if err := runGitCommand("add", "."); err != nil {
		fmt.Printf("Error running git add: %v\n", err)
		os.Exit(1)
	}

	if err := runGitCommand("commit", "-m", "Ellie: "+commitMsg); err != nil {
		fmt.Printf("Error running git commit: %v\n", err)
		os.Exit(1)
	}

	if err := runGitCommand("push"); err != nil {
		fmt.Printf("Error running git push: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Push successful!")
}

// GitPull executes a git pull.
func GitPull() {
	if err := runGitCommand("pull"); err != nil {
		fmt.Printf("Error running git pull: %v\n", err)
		os.Exit(1)
	}
}

// GitStatus executes a git status.
func GitStatus() {
	if err := runGitCommand("status"); err != nil {
		fmt.Printf("Error running git status: %v\n", err)
		os.Exit(1)
	}
}
