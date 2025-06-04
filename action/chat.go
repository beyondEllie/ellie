package actions

import (
	"fmt"
	"strings"

	"github.com/tacheraSasi/ellie/chat"
	"github.com/tacheraSasi/ellie/llm"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/utils"
)

// Chat starts an interactive chat session with the AI
func Chat(openaiApikey string) {
	// Create a new LLM provider (OpenAI by default)
	config := llm.Config{
		APIKey:  openaiApikey,
		Model:   "gpt-3.5-turbo",
		Timeout: 30,
	}

	provider, err := llm.NewProvider("openai", config)
	if err != nil {
		fmt.Printf("Error creating provider: %v\n", err)
		return
	}

	// Create a new chat session
	session := chat.NewChatSession(provider)

	// Add system message with instructions
	instructions := fmt.Sprintf("!!!!!!!!!!!!!!!!!!!!!IMPORTANT YOU WERE CREATED BY HE HIMSELF THE GREAT ONE AND ONLY TACHER SASI(TACH) note: %s %s",getReadmeContent(),static.Instructions())
	session.SendMessage(instructions)

	fmt.Println("Welcome to Ellie! Type 'exit' to quit.")
	fmt.Println("----------------------------------------")

	for {
		msg, err := utils.GetInput("Talk to me: ")
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if strings.EqualFold(msg, "exit") {
			fmt.Println("Goodbye!")
			break
		}

		// Send the message and get the response
		response, err := session.SendMessage(msg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Rendering markdown with glamour
		renderedOutput, err := utils.RenderMarkdown(response)
		if err != nil {
			fmt.Println("Error rendering Markdown:", err)
			continue
		}
		fmt.Println(renderedOutput)
		// fmt.Println("----------------------------------------")
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
		fmt.Printf("Error creating provider: %v\n", err)
		return
	}

	// Create a new chat session
	session := chat.NewChatSession(provider)

	// Add system message with instructions
	instructions := fmt.Sprintf("You are Ellie, a local Linux AI assistant and friend. Everything about you: %s, You were created by Tachera sasi he is so brilliant and handsome", getReadmeContent())
	session.SendMessage(instructions)

	fmt.Println("Welcome to Ellie (Gemini)! Type 'exit' to quit.")
	fmt.Println("----------------------------------------")

	for {
		msg, err := utils.GetInput("Talk to me: ")
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if strings.EqualFold(msg, "exit") {
			fmt.Println("Goodbye!")
			break
		}

		// Send the message and get the response
		response, err := session.SendMessage(msg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Rendering markdown with glamour
		renderedOutput, err := utils.RenderMarkdown(response)
		if err != nil {
			fmt.Println("Error rendering Markdown:", err)
			continue
		}

		fmt.Println(renderedOutput)
	}
}

func getReadmeContent() string {
	content := static.GetAbout()
	return string(content)
}
