package actions

import (
	"fmt"
	"os"
	"strings"

	"github.com/tacheraSasi/ellie/chat"
	"github.com/tacheraSasi/ellie/llm"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// Review reads a file and asks the LLM to review the code.
func Review(file string) {
	openaiApikey := os.Getenv("OPENAI_API_KEY")
	if openaiApikey == "" {
		styles.ErrorStyle.Println("Error: OpenAI API key is required. Please set your OPENAI_API_KEY environment variable.")
		return
	}

	code, err := os.ReadFile(file)
	if err != nil {
		styles.ErrorStyle.Printf("Error reading file: %v\n", err)
		return
	}

	config := llm.Config{
		APIKey:  openaiApikey,
		Model:   "gpt-4",
		Timeout: 60,
	}

	provider, err := llm.NewProvider("openai", config)
	if err != nil {
		styles.ErrorStyle.Printf("Error creating provider: %v\n", err)
		return
	}
	// Create a new chat session
	session := chat.NewChatSession(provider)

	prompt := fmt.Sprintf(
		`You are an expert software engineer. Please review the following code for bugs, security issues, code quality, and best practices. Provide actionable suggestions and a summary. 
File: %s

Code:
%s
`, file, string(code))

	styles.InfoStyle.Println("Reviewing code with Ellie...")

	done := make(chan bool)
	responseChan := make(chan string)
	errorChan := make(chan error)

	go utils.ShowLoadingSpinner("Thinking...", done)

	go func() {
		response, err := session.SendMessage(prompt)
		if err != nil {
			errorChan <- err
			return
		}
		responseChan <- response
	}()

	select {
	case err := <-errorChan:
		done <- true
		styles.ErrorStyle.Printf("\nError: %v\n", err)
	case response := <-responseChan:
		done <- true
		if strings.TrimSpace(response) == "" {
			styles.WarningStyle.Println("\nNo response received from AI.")
			return
		}
		renderedOutput, err := utils.RenderMarkdown(response)
		if err != nil {
			styles.ErrorStyle.Println("\nRaw response:", response)
			styles.ErrorStyle.Println("Error rendering Markdown:", err)
			return
		}
		fmt.Println("\n" + renderedOutput)
	}
}
