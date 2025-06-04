package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tacheraSasi/ellie/chat"
	"github.com/tacheraSasi/ellie/llm"
)

func main() {
	// Create a new LLM provider (OpenAI in this case)
	config := llm.Config{
		APIKey:  os.Getenv("OPENAI_API_KEY"),
		Model:   "gpt-3.5-turbo",
		Timeout: 30,
	}

	provider, err := llm.NewProvider("openai", config)
	if err != nil {
		fmt.Printf("Error creating provider: %v\n", err)
		os.Exit(1)
	}

	// Create a new chat session
	session := chat.NewChatSession(provider)

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Ellie! Type 'exit' to quit.")
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			break
		}

		// Send the message and get the response
		response, err := session.SendMessage(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("Ellie: %s\n", response)
		fmt.Println("----------------------------------------")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}
