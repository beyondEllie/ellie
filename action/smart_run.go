package actions

import (
	"fmt"
	"strings"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/llm"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/types"
	"github.com/tacheraSasi/ellie/utils"
)

// SmartRun uses LLM to process user requests and executes confirmed commands
func SmartRun(args []string) {
	userPrompt := strings.Join(args[1:], " ")
	openaiApiKey := configs.GetEnv("OPENAI_API_KEY")
	if openaiApiKey == "" {
		styles.ErrorStyle.Println("Error: OpenAI API key is required. Please set OPENAI_API_KEY.")
		return
	}

	config := llm.Config{
		Timeout: 60,
	}

	provider, err := llm.NewProvider("ellieapi", config)
	if err != nil {
		styles.ErrorStyle.Printf("Error creating provider: %v\n", err)
		return
	}
	userCtx := types.NewUserContext()
	instructions := fmt.Sprintf(`!!!!!!!!!!!!!!!!!!!!!IMPORTANT YOU ARE ELLIE note: %s %s`,
		getReadmeContent(),
		static.Instructions(*userCtx))
	
	prompt := `You are an expert terminal assistant. Respond to the user request following these rules:
1. If you need to execute a bash command to fulfill the request, wrap it EXACTLY like this: <cmd>COMMAND</cmd>.
2. Provide clear instructions/explanations OUTSIDE the tags.
3. NEVER include commands outside <cmd> tags.
User request: ` + instructions +userPrompt

	resp, err := provider.Chat([]llm.Message{{Role: "user", Content: prompt}})
	if err != nil {
		styles.ErrorStyle.Printf("LLM error: %v\n", err)
		return
	}

	// Process response
	responseContent := resp.Content
	instructions, commands := extractContent(responseContent)

	// Display instructions
	if instructions != "" {
		fmt.Println(instructions)
	}

	// Execute confirmed commands
	if len(commands) > 0 {
		executeCommands(commands)
	} else {
		styles.WarningStyle.Println("No actionable commands found in the response.")
	}
}

// extractContent separates instructions and commands from LLM response
func extractContent(response string) (string, []string) {
	var instructionsBuilder strings.Builder
	var commands []string
	startTag := "<cmd>"
	endTag := "</cmd>"

	// Process all <cmd> segments
	for {
		startIdx := strings.Index(response, startTag)
		if startIdx == -1 {
			break
		}

		// Capture content before command
		instructionsBuilder.WriteString(response[:startIdx])
		response = response[startIdx+len(startTag):]

		// Extract command
		endIdx := strings.Index(response, endTag)
		if endIdx == -1 {
			break // Unclosed tag, ignore
		}

		cmd := strings.TrimSpace(response[:endIdx])
		if cmd != "" {
			commands = append(commands, cmd)
		}
		response = response[endIdx+len(endTag):]
	}

	// Add remaining content
	instructionsBuilder.WriteString(response)
	return strings.TrimSpace(instructionsBuilder.String()), commands
}

// executeCommands prompts user and runs confirmed commands
func executeCommands(commands []string) {
	for _, cmd := range commands {
		styles.Cyan.Printf("\nCommand: %s\n", cmd)
		if utils.AskForConfirmation("Run this command?") {
			utils.RunCommand([]string{"bash", "-c", cmd}, "Command error:")
		} else {
			styles.InfoStyle.Println("Command skipped.")
		}
	}
}