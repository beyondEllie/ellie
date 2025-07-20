package actions

import (
	"fmt"
	"strings"
)

// SmartRun uses LLM to turn users request to commands and runs them
func SmartRun(args []string) {
	userPrompt := strings.Join(args[1:], " ")
	fmt.Println(userPrompt)
}
