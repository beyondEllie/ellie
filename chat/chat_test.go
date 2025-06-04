package chat

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/tacheraSasi/ellie/llm"
)

// MockProvider is a mock implementation of llm.Provider for testing
type MockProvider struct {
	responses []string
	index     int
}

func (m *MockProvider) Chat(messages []llm.Message) (*llm.Response, error) {
	if m.index >= len(m.responses) {
		m.index = 0
	}
	response := m.responses[m.index]
	m.index++
	return &llm.Response{
		Content: response,
		Usage: llm.Usage{
			PromptTokens:     10,
			CompletionTokens: 20,
			TotalTokens:      30,
		},
	}, nil
}

func (m *MockProvider) GetModel() string {
	return "mock-model"
}

func TestNewChatSession(t *testing.T) {
	provider := &MockProvider{}
	session := NewChatSession(provider)

	if session.provider != provider {
		t.Errorf("NewChatSession() provider = %v, want %v", session.provider, provider)
	}
	if len(session.messages) != 0 {
		t.Errorf("NewChatSession() messages length = %d, want 0", len(session.messages))
	}
	if len(session.history) != 0 {
		t.Errorf("NewChatSession() history length = %d, want 0", len(session.history))
	}
}

func TestChatSession_SendMessage(t *testing.T) {
	provider := &MockProvider{
		responses: []string{"Hello!", "How can I help?"},
	}
	session := NewChatSession(provider)

	// Test first message
	response, err := session.SendMessage("Hi")
	if err != nil {
		t.Fatalf("SendMessage() error = %v", err)
	}
	if response != "Hello!" {
		t.Errorf("SendMessage() response = %v, want %v", response, "Hello!")
	}
	if len(session.messages) != 2 {
		t.Errorf("SendMessage() messages length = %d, want 2", len(session.messages))
	}
	if len(session.history) != 2 {
		t.Errorf("SendMessage() history length = %d, want 2", len(session.history))
	}

	// Test second message
	response, err = session.SendMessage("How are you?")
	if err != nil {
		t.Fatalf("SendMessage() error = %v", err)
	}
	if response != "How can I help?" {
		t.Errorf("SendMessage() response = %v, want %v", response, "How can I help?")
	}
	if len(session.messages) != 4 {
		t.Errorf("SendMessage() messages length = %d, want 4", len(session.messages))
	}
	if len(session.history) != 4 {
		t.Errorf("SendMessage() history length = %d, want 4", len(session.history))
	}
}

func TestChatSession_ClearHistory(t *testing.T) {
	provider := &MockProvider{
		responses: []string{"Hello!"},
	}
	session := NewChatSession(provider)

	// Send a message to populate history
	_, err := session.SendMessage("Hi")
	if err != nil {
		t.Fatalf("SendMessage() error = %v", err)
	}

	// Clear history
	session.ClearHistory()

	if len(session.messages) != 0 {
		t.Errorf("ClearHistory() messages length = %d, want 0", len(session.messages))
	}
	if len(session.history) != 0 {
		t.Errorf("ClearHistory() history length = %d, want 0", len(session.history))
	}
}

func TestFormatMessage(t *testing.T) {
	role := "User"
	content := "Hello, world!"
	formatted := FormatMessage(role, content)

	// Check if the formatted message contains all components
	if !strings.Contains(formatted, role) {
		t.Errorf("FormatMessage() missing role %s", role)
	}
	if !strings.Contains(formatted, content) {
		t.Errorf("FormatMessage() missing content %s", content)
	}
	if !strings.Contains(formatted, "[") || !strings.Contains(formatted, "]") {
		t.Errorf("FormatMessage() missing timestamp brackets")
	}
}

func TestParseMessage(t *testing.T) {
	timestamp := time.Now().Format("15:04:05")
	role := "User"
	content := "Hello, world!"
	formatted := fmt.Sprintf("[%s] %s: %s", timestamp, role, content)

	parsedTimestamp, parsedRole, parsedContent, err := ParseMessage(formatted)
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	if parsedTimestamp != timestamp {
		t.Errorf("ParseMessage() timestamp = %v, want %v", parsedTimestamp, timestamp)
	}
	if parsedRole != role {
		t.Errorf("ParseMessage() role = %v, want %v", parsedRole, role)
	}
	if parsedContent != content {
		t.Errorf("ParseMessage() content = %v, want %v", parsedContent, content)
	}

	// Test invalid format
	_, _, _, err = ParseMessage("invalid message")
	if err == nil {
		t.Error("ParseMessage() expected error for invalid format")
	}
}
