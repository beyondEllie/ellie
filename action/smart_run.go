package actions

import (
	"fmt"
	"strings"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/llm"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// SmartRun uses LLM to turn users request to commands and runs them
func SmartRun(args []string) {
	userPrompt := strings.Join(args[1:], " ")
	openaiApiKey := configs.GetEnv("OPENAI_API_KEY")
	fmt.Println("openai key", openaiApiKey)
	if openaiApiKey == "" {
		styles.ErrorStyle.Println("Error: OpenAI API key is required. Please set your OPENAI_API_KEY environment variable.")
		return
	}

	llmConfig := llm.Config{
		APIKey:  openaiApiKey,
		Model:   "gpt-3.5-turbo",
		Timeout: 30,
	}
	provider, err := llm.NewProvider("openai", llmConfig)
	if err != nil {
		styles.ErrorStyle.Printf("Error creating LLM provider: %v\n", err)
		return
	}

	prompt := `You are an expert terminal assistant. Given a user request, output ONLY the most appropriate bash command to accomplish the task. Do not explain, just output the command. User request: ` + userPrompt
	resp, err := provider.Chat([]llm.Message{{Role: "user", Content: prompt}})
	if err != nil {
		styles.ErrorStyle.Printf("Error from LLM: %v\n", err)
		return
	}
	command := strings.TrimSpace(resp.Content)
	if command == "" {
		styles.ErrorStyle.Println("LLM did not return a command.")
		return
	}
	styles.Cyan.Printf("\n$ %s\n", command)
	utils.RunCommand([]string{"bash", "-c", command}, "Error running command:")
}
