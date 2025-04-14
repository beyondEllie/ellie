package actions

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tacheraSasi/ellie/styles"
)

func History(args []string) {
	// Get the appropriate history file based on OS and shell
	historyFile := getHistoryFile()
	if historyFile == "" {
		styles.ErrorStyle.Println("Could not determine shell history file")
		return
	}

	// Read the history file
	content, err := os.ReadFile(historyFile)
	if err != nil {
		styles.ErrorStyle.Printf("Error reading history file: %s\n", err)
		return
	}

	// Split into lines and get the last 50 commands
	lines := strings.Split(string(content), "\n")
	start := 0
	if len(lines) > 50 {
		start = len(lines) - 50
	}
	recentCommands := lines[start:]

	// Print the commands with numbering
	styles.InfoStyle.Println("Recent Commands:")
	for i, cmd := range recentCommands {
		if cmd != "" {
			fmt.Printf("%3d: %s\n", i+1, cmd)
		}
	}
}

func getHistoryFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	// Check shell type
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		return filepath.Join(homeDir, ".zsh_history")
	} else if strings.Contains(shell, "bash") {
		return filepath.Join(homeDir, ".bash_history")
	}

	// Fallback based on OS
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(homeDir, "AppData", "Roaming", "Microsoft", "Windows", "PowerShell", "PSReadLine", "ConsoleHost_history.txt")
	case "darwin", "linux":
		return filepath.Join(homeDir, ".bash_history")
	default:
		return ""
	}
}
