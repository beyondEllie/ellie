package actions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/styles"
)

// SmartSuggest provides intelligent command suggestions based on context
func SmartSuggest() {
	styles.GetInfoStyle().Println("\n💡 Smart Command Suggestions")
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	
	// Analyze current directory
	analyzeCurrentDirectory()
	
	// Check for common tasks
	suggestCommonTasks()
	
	// Check for Git repository
	suggestGitOperations()
	
	// Check for project files
	suggestProjectOperations()
	
	// Suggest maintenance tasks
	suggestMaintenance()
	
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
}

func analyzeCurrentDirectory() {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}
	
	styles.GetHighlightStyle().Printf("\n📂 Current Directory: %s\n", filepath.Base(cwd))
	
	// Count files
	files, err := os.ReadDir(cwd)
	if err != nil {
		return
	}
	
	var dirCount, fileCount int
	for _, file := range files {
		if file.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
	}
	
	fmt.Printf("   Files: %d | Directories: %d\n", fileCount, dirCount)
}

func suggestCommonTasks() {
	styles.GetHighlightStyle().Println("\n⚡ Quick Actions:")
	
	// Check time of day
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("   • ellie start-day - Start your development day")
	}
	
	// Common commands
	fmt.Println("   • ellie health - Check system health")
	fmt.Println("   • ellie sysinfo - View system information")
	fmt.Println("   • ellie history - View command history")
}

func suggestGitOperations() {
	// Check if in a Git repository
	if !isGitRepo() {
		return
	}
	
	styles.GetHighlightStyle().Println("\n🔀 Git Suggestions:")
	
	// Check for uncommitted changes
	if hasUncommittedChanges() {
		fmt.Println("   ⚠️  You have uncommitted changes")
		fmt.Println("   • ellie git status - View changes")
		fmt.Println("   • ellie git commit - Create a commit")
	}
	
	// Check for unpushed commits
	if hasUnpushedCommits() {
		fmt.Println("   📤 You have unpushed commits")
		fmt.Println("   • ellie git push - Push your changes")
	}
	
	// Suggest common Git operations
	if !hasUncommittedChanges() && !hasUnpushedCommits() {
		fmt.Println("   ✅ Repository is clean and synced")
		fmt.Println("   • ellie git pull - Check for updates")
	}
}

func suggestProjectOperations() {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}
	
	// Check for common project files
	projectType := detectProjectType(cwd)
	if projectType != "" {
		styles.GetHighlightStyle().Printf("\n🚀 Detected %s Project:\n", projectType)
		suggestProjectCommands(projectType)
	}
}

func detectProjectType(dir string) string {
	// Check for various project indicators
	indicators := map[string]string{
		"package.json": "Node.js",
		"go.mod":       "Go",
		"Cargo.toml":   "Rust",
		"pom.xml":      "Maven/Java",
		"build.gradle": "Gradle/Java",
		"requirements.txt": "Python",
		"Pipfile":      "Python",
		"composer.json": "PHP",
		"Gemfile":      "Ruby",
		"Makefile":     "C/C++",
	}
	
	for file, projectType := range indicators {
		if _, err := os.Stat(filepath.Join(dir, file)); err == nil {
			return projectType
		}
	}
	
	return ""
}

func suggestProjectCommands(projectType string) {
	commands := map[string][]string{
		"Node.js": {
			"npm install - Install dependencies",
			"npm run dev - Start development server",
			"npm test - Run tests",
			"npm run build - Build for production",
		},
		"Go": {
			"go build - Build the project",
			"go test ./... - Run tests",
			"go mod tidy - Clean dependencies",
			"go run . - Run the application",
		},
		"Python": {
			"pip install -r requirements.txt - Install dependencies",
			"python -m pytest - Run tests",
			"python main.py - Run the application",
		},
		"Rust": {
			"cargo build - Build the project",
			"cargo test - Run tests",
			"cargo run - Run the application",
			"cargo clippy - Lint code",
		},
	}
	
	if suggestions, ok := commands[projectType]; ok {
		for _, suggestion := range suggestions {
			fmt.Printf("   • %s\n", suggestion)
		}
	}
}

func suggestMaintenance() {
	styles.GetHighlightStyle().Println("\n🔧 Maintenance Suggestions:")
	
	// Check system health score
	score := calculateHealthScore()
	if score < 70 {
		fmt.Println("   ⚠️  System health is below optimal")
		fmt.Println("   • ellie health - View detailed health report")
	}
	
	// Check for updates
	if shouldCheckUpdates() {
		fmt.Println("   📦 Consider checking for system updates")
		fmt.Println("   • ellie update - Update system packages")
	}
	
	// Suggest disk cleanup if needed
	if disks, err := getDiskUsage(); err == nil {
		for _, disk := range disks {
			if disk.UsagePercent > 80 {
				fmt.Printf("   💾 Disk %s is %.0f%% full\n", disk.Mount, disk.UsagePercent)
				fmt.Println("   • Consider cleaning up unused files")
			}
		}
	}
}

func isGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Stderr = nil
	err := cmd.Run()
	return err == nil
}

func hasUncommittedChanges() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return len(strings.TrimSpace(string(out))) > 0
}

func hasUnpushedCommits() bool {
	cmd := exec.Command("git", "log", "@{u}..", "--oneline")
	cmd.Stderr = nil
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return len(strings.TrimSpace(string(out))) > 0
}

func shouldCheckUpdates() bool {
	// Simple heuristic: suggest updates once per day
	// In a real implementation, you'd track the last update check
	return time.Now().Hour() == 10 // Suggest at 10 AM
}

// ContextHelp provides context-aware help based on current situation
func ContextHelp() {
	styles.GetInfoStyle().Println("\n🤖 Ellie's Context-Aware Assistant")
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	
	// Get current context
	cwd, _ := os.Getwd()
	projectType := detectProjectType(cwd)
	isGit := isGitRepo()
	
	styles.GetHighlightStyle().Println("\n📍 Current Context:")
	fmt.Printf("   Location: %s\n", cwd)
	
	if projectType != "" {
		fmt.Printf("   Project Type: %s\n", projectType)
	}
	
	if isGit {
		fmt.Println("   Git Repository: Yes")
		
		// Get branch name
		cmd := exec.Command("git", "branch", "--show-current")
		if out, err := cmd.Output(); err == nil {
			fmt.Printf("   Current Branch: %s\n", strings.TrimSpace(string(out)))
		}
	}
	
	// Provide relevant help
	styles.GetHighlightStyle().Println("\n💡 What can I help you with?")
	
	if isGit {
		fmt.Println("\n   Git Operations:")
		fmt.Println("   • ellie git status - Check repository status")
		fmt.Println("   • ellie git commit - Create a conventional commit")
		fmt.Println("   • ellie git push - Push changes")
		fmt.Println("   • ellie git pull - Pull latest changes")
	}
	
	if projectType != "" {
		fmt.Printf("\n   %s Operations:\n", projectType)
		suggestProjectCommands(projectType)
	}
	
	fmt.Println("\n   System Operations:")
	fmt.Println("   • ellie health - Check system health")
	fmt.Println("   • ellie sysinfo - View system info")
	fmt.Println("   • ellie start-day - Start your dev day")
	
	fmt.Println("\n   Productivity:")
	fmt.Println("   • ellie todo add \"task\" - Add a todo")
	fmt.Println("   • ellie project add <name> <path> - Add project")
	fmt.Println("   • ellie focus - Activate focus mode")
	
	styles.GetInfoStyle().Println("\n═══════════════════════════════════════════════════════")
	styles.GetInfoStyle().Println("💬 Type 'ellie chat' to talk to me or 'ellie --help' for all commands")
}

// WorkflowAnalysis analyzes your command history and provides insights
func WorkflowAnalysis() {
	styles.GetInfoStyle().Println("\n📊 Workflow Analysis & Insights")
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	
	// Get command history
	history := getRecentCommandHistory(50)
	
	if len(history) == 0 {
		styles.GetErrorStyle().Println("No command history found")
		return
	}
	
	// Analyze patterns
	analyzeCommandPatterns(history)
	
	// Suggest optimizations
	suggestWorkflowOptimizations(history)
	
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
}

func getRecentCommandHistory(limit int) []string {
	var commands []string
	
	historyFile := getHistoryFilePath()
	if historyFile == "" {
		return commands
	}
	
	content, err := os.ReadFile(historyFile)
	if err != nil {
		return commands
	}
	
	lines := strings.Split(string(content), "\n")
	start := len(lines) - limit
	if start < 0 {
		start = 0
	}
	
	for i := start; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			commands = append(commands, line)
		}
	}
	
	return commands
}

func getHistoryFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	
	// Try different shell history files
	historyFiles := []string{
		filepath.Join(home, ".bash_history"),
		filepath.Join(home, ".zsh_history"),
		filepath.Join(home, ".local/share/fish/fish_history"),
	}
	
	for _, file := range historyFiles {
		if _, err := os.Stat(file); err == nil {
			return file
		}
	}
	
	return ""
}

func analyzeCommandPatterns(history []string) {
	styles.GetHighlightStyle().Println("\n🔍 Command Patterns:")
	
	// Count command frequency
	cmdFreq := make(map[string]int)
	for _, cmd := range history {
		// Extract base command
		parts := strings.Fields(cmd)
		if len(parts) > 0 {
			baseCmd := parts[0]
			cmdFreq[baseCmd]++
		}
	}
	
	// Find top 5 commands
	type cmdCount struct {
		cmd   string
		count int
	}
	
	var topCmds []cmdCount
	for cmd, count := range cmdFreq {
		topCmds = append(topCmds, cmdCount{cmd, count})
	}
	
	// Simple bubble sort to get top 5
	for i := 0; i < len(topCmds)-1; i++ {
		for j := 0; j < len(topCmds)-i-1; j++ {
			if topCmds[j].count < topCmds[j+1].count {
				topCmds[j], topCmds[j+1] = topCmds[j+1], topCmds[j]
			}
		}
	}
	
	fmt.Println("   Top 5 Most Used Commands:")
	for i := 0; i < 5 && i < len(topCmds); i++ {
		fmt.Printf("   %d. %s (%d times)\n", i+1, topCmds[i].cmd, topCmds[i].count)
	}
}

func suggestWorkflowOptimizations(history []string) {
	styles.GetHighlightStyle().Println("\n💡 Optimization Suggestions:")
	
	// Check for repeated command patterns
	gitCmds := 0
	for _, cmd := range history {
		if strings.HasPrefix(cmd, "git ") {
			gitCmds++
		}
	}
	
	if gitCmds > 10 {
		fmt.Println("   • You use Git frequently! Try:")
		fmt.Println("     - ellie alias add gs=\"git status\"")
		fmt.Println("     - ellie alias add gc=\"git commit\"")
		fmt.Println("     - ellie git commit for conventional commits")
	}
	
	// Check for cd patterns
	cdCmds := 0
	for _, cmd := range history {
		if strings.HasPrefix(cmd, "cd ") {
			cdCmds++
		}
	}
	
	if cdCmds > 15 {
		fmt.Println("   • You navigate directories often! Try:")
		fmt.Println("     - ellie project add to save common locations")
		fmt.Println("     - ellie switch <project> for quick navigation")
	}
	
	// Check for ls/dir patterns
	lsCmds := 0
	for _, cmd := range history {
		if strings.HasPrefix(cmd, "ls") || strings.HasPrefix(cmd, "dir") {
			lsCmds++
		}
	}
	
	if lsCmds > 10 {
		fmt.Println("   • Try 'ellie list <dir>' for better directory listings")
	}
}

// TimeBasedSuggestions provides suggestions based on time of day
func TimeBasedSuggestions() {
	hour := time.Now().Hour()
	
	styles.GetInfoStyle().Println("\n⏰ Time-Based Suggestions")
	styles.GetInfoStyle().Println("─────────────────────────")
	
	if hour >= 6 && hour < 10 {
		fmt.Println("🌅 Good morning! Here's what you might want to do:")
		fmt.Println("   • ellie start-day - Start your development day")
		fmt.Println("   • ellie health - Check system health")
		fmt.Println("   • ellie todo list - Review your tasks")
	} else if hour >= 10 && hour < 12 {
		fmt.Println("☕ Mid-morning productivity time:")
		fmt.Println("   • ellie focus - Activate focus mode for deep work")
		fmt.Println("   • ellie suggest - Get context-aware suggestions")
	} else if hour >= 12 && hour < 14 {
		fmt.Println("🍽️  Lunchtime! Consider:")
		fmt.Println("   • Take a break")
		fmt.Println("   • ellie git commit - Commit your morning work")
		fmt.Println("   • ellie health - Quick system check")
	} else if hour >= 14 && hour < 18 {
		fmt.Println("🚀 Afternoon productivity:")
		fmt.Println("   • ellie workflow - Analyze your workflow")
		fmt.Println("   • ellie todo complete <id> - Mark tasks done")
		fmt.Println("   • ellie git push - Push your changes")
	} else if hour >= 18 && hour < 22 {
		fmt.Println("🌆 Evening wrap-up:")
		fmt.Println("   • ellie git status - Check uncommitted work")
		fmt.Println("   • ellie todo list - Review remaining tasks")
		fmt.Println("   • ellie health - Final system check")
	} else {
		fmt.Println("🌙 Late night! Remember to:")
		fmt.Println("   • Commit your work: ellie git commit")
		fmt.Println("   • Save your progress")
		fmt.Println("   • Get some rest! 😴")
	}
}
