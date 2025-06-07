package chat

import (
	"fmt"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/llm"
)

// ChatSession represents an ongoing chat session
type ChatSession struct {
	provider llm.Provider
	messages []llm.Message
	history  []string
}

// NewChatSession creates a new chat session with the specified provider
func NewChatSession(provider llm.Provider) *ChatSession {
	return &ChatSession{
		provider: provider,
		messages: make([]llm.Message, 0),
		history:  make([]string, 0),
	}
}

// SendMessage sends a message to the LLM and returns the response
func (s *ChatSession) SendMessage(content string) (string, error) {
	// Add user message to history
	s.messages = append(s.messages, llm.Message{
		Role:    "user",
		Content: content,
	})
	s.history = append(s.history, fmt.Sprintf("User: %s", content))

	// Get response from LLM
	response, err := s.provider.Chat(s.messages)
	if err != nil {
		return "", fmt.Errorf("failed to get response from LLM: %w", err) //Note: If this occurs
		//Its probably because the user is offline
	}

	// Add assistant message to history
	s.messages = append(s.messages, llm.Message{
		Role:    "assistant",
		Content: response.Content,
	})
	s.history = append(s.history, fmt.Sprintf("Assistant: %s", response.Content))

	return response.Content, nil
}

// GetHistory returns the chat history
func (s *ChatSession) GetHistory() []string {
	return s.history
}

// ClearHistory clears the chat history
func (s *ChatSession) ClearHistory() {
	s.messages = make([]llm.Message, 0)
	s.history = make([]string, 0)
}

// FormatMessage formats a message with timestamp
func FormatMessage(role, content string) string {
	timestamp := time.Now().Format("15:04:05")
	return fmt.Sprintf("[%s] %s: %s", timestamp, role, content)
}

// ParseMessage parses a formatted message into its components
func ParseMessage(formatted string) (timestamp, role, content string, err error) {
	parts := strings.SplitN(formatted, " ", 3)
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid message format")
	}

	timestamp = strings.Trim(parts[0], "[]")
	roleContent := strings.SplitN(parts[1], ":", 2)
	if len(roleContent) != 2 {
		return "", "", "", fmt.Errorf("invalid role format")
	}

	return timestamp, roleContent[0], parts[2], nil
}
