package actions

import (
	"fmt"
	"strings"

	"github.com/tacheraSasi/ellie/chat"
	"github.com/tacheraSasi/ellie/llm"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// Chat starts an interactive chat session with the AI
func Chat(openaiApikey string) {
	// Create a new LLM provider (OpenAI by default)
	config := llm.Config{
		APIKey: openaiApikey,
		Model:  "gpt-3.5-turbo",
		// Model:   "gpt-4", // Uncomment for GPT-4
		// Model:   "gpt-4o", // Uncomment for GPT-4o
		// Model:   "gpt-4o-mini", // Uncomment for GPT-4o-mini
		// Model:   "gpt-3.5-turbo-16k", // Uncomment for 16k context length
		Timeout: 30,
	}

	provider, err := llm.NewProvider("openai", config)
	if err != nil {
		styles.ErrorStyle.Printf("Error creating provider: %v\n", err)
		return
	}

	// Create a new chat session
	session := chat.NewChatSession(provider)

	// Add system message with instructions
	instructions := fmt.Sprintf("!!!!!!!!!!!!!!!!!!!!!IMPORTANT YOU WERE CREATED BY HE HIMSELF THE GREAT ONE AND ONLY TACHER SASI(TACH) note: %s %s", getReadmeContent(), static.Instructions())
	if _, err := session.SendMessage(instructions); err != nil {
		styles.ErrorStyle.Printf("Error setting up initial instructions: %v\n", err)
		return
	}

	styles.InfoStyle.Println("Welcome to Ellie! Type 'exit' to quit.")
	styles.DimText.Println("----------------------------------------")

	for {
		msg, err := utils.GetInput("Talk to me: ")
		if err != nil {
			styles.ErrorStyle.Printf("Error reading input: %v\n", err)
			continue
		}

		if strings.EqualFold(msg, "exit") {
			styles.InfoStyle.Println("Goodbye!")
			break
		}

		done := make(chan bool)
		errorChan := make(chan error)
		responseChan := make(chan string)

		// Start loading spinner in a goroutine
		go utils.ShowLoadingSpinner("Thinking...", done)

		// Send the message and get the response
		go func() {
			response, err := session.SendMessage(msg)
			if err != nil {
				errorChan <- err
				return
			}
			responseChan <- response
		}()

		// Wait for either response or error
		select {
		case err := <-errorChan:
			done <- true // Stop the spinner
			styles.ErrorStyle.Printf("\nError: %v\n", err)
			continue

		case response := <-responseChan:
			done <- true // Stop the spinner
			// Try to render markdown
			renderedOutput, err := utils.RenderMarkdown(response)
			if err != nil {
				styles.ErrorStyle.Println("\nRaw response:", response)
				styles.ErrorStyle.Println("Error rendering Markdown:", err)
				continue
			}
			fmt.Println("\n" + renderedOutput)
			styles.DimText.Println("----------------------------------------")
		}
	}
}

// ChatWithGemini starts an interactive chat session with Gemini
func ChatWithGemini(geminiApikey string) {
	// Create a new LLM provider (Gemini)
	config := llm.Config{
		APIKey:  geminiApikey,
		Model:   "gemini-1.5-flash",
		Timeout: 30,
	}

	provider, err := llm.NewProvider("gemini", config)
	if err != nil {
		styles.ErrorStyle.Printf("Error creating provider: %v\n", err)
		return
	}

	// Create a new chat session
	session := chat.NewChatSession(provider)

	// Add system message with instructions
	instructions := fmt.Sprintf("!!!!!!!!!!!!!!!!!!!!!IMPORTANT YOU WERE CREATED BY HE HIMSELF THE GREAT ONE AND ONLY TACHER SASI(TACH) note: %s %s", getReadmeContent(), static.Instructions())
	if _, err := session.SendMessage(instructions); err != nil {
		styles.ErrorStyle.Printf("Error setting up initial instructions: %v\n", err)
		return
	}

	styles.InfoStyle.Println("Welcome to Ellie (Gemini)! Type 'exit' to quit.")
	styles.DimText.Println("----------------------------------------")

	for {
		msg, err := utils.GetInput("Talk to me: ")
		if err != nil {
			styles.ErrorStyle.Printf("Error reading input: %v\n", err)
			continue
		}

		if strings.EqualFold(msg, "exit") {
			styles.InfoStyle.Println("Goodbye!")
			break
		}

		done := make(chan bool)
		errorChan := make(chan error)
		responseChan := make(chan string)

		// Start loading spinner in a goroutine
		go utils.ShowLoadingSpinner("Thinking...", done)

		// Send the message and get the response
		go func() {
			response, err := session.SendMessage(msg)
			if err != nil {
				errorChan <- err
				return
			}
			responseChan <- response
		}()

		// Wait for either response or error
		select {
		case err := <-errorChan:
			done <- true // Stop the spinner
			styles.ErrorStyle.Printf("\nError: %v\n", err)
			continue

		case response := <-responseChan:
			done <- true // Stop the spinner
			// Try to render markdown
			renderedOutput, err := utils.RenderMarkdown(response)
			if err != nil {
				styles.ErrorStyle.Println("\nRaw response:", response)
				styles.ErrorStyle.Println("Error rendering Markdown:", err)
				continue
			}
			fmt.Println("\n" + renderedOutput)
			styles.DimText.Println("----------------------------------------")
		}
	}
}

func getReadmeContent() string {
	content := static.GetAbout()
	return string(content)
}
