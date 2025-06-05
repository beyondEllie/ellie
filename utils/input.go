package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/tacheraSasi/ellie/styles"
)

// GetInput prompts the user for input and returns the trimmed string.
func GetInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	styles.InputPrompt.Print(prompt, " ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
