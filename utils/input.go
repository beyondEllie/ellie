package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/mattn/go-isatty"
	"github.com/tacheraSasi/ellie/styles"
)

// GetInput prompts the user for input and returns the trimmed string.
func GetInput(promptText string) (string, error) {
	// Try to use go-prompt for interactive input
	if isTerminalInteractive() {
		result := prompt.Input(styles.InputPrompt.Sprintf("%s ", promptText), completer)
		return strings.TrimSpace(result), nil
	}
	// Fallback to bufio for non-interactive environments
	reader := bufio.NewReader(os.Stdin)
	styles.InputPrompt.Print(promptText, " ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// isTerminalInteractive checks if the terminal is interactive
func isTerminalInteractive() bool {
	return isatty.IsTerminal(os.Stdin.Fd())
}

// completer is a no-op for now, but can be extended for suggestions
func completer(d prompt.Document) []prompt.Suggest {
	return nil
}
