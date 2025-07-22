package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Provider defines the interface for different LLM providers
type Provider interface {
	Chat(messages []Message) (*Response, error)
	GetModel() string
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response represents the LLM response
type Response struct {
	Content string
	Usage   Usage
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Config holds the configuration for LLM providers
type Config struct {
	APIKey     string
	Model      string
	BaseURL    string
	Timeout    time.Duration
	MaxRetries int
}

// NewProvider creates a new LLM provider based on the type
func NewProvider(providerType string, config Config) (Provider, error) {
	switch providerType {
	case "openai":
		return NewOpenAIProvider(config)
	case "gemini":
		return NewGeminiProvider(config)
	case "ellieapi":
		return NewEllieAPIProvider(config)
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", providerType)
	}
}

// makeRequest is a helper function to make HTTP requests
func makeRequest(url string, headers map[string]string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return responseBody, nil
}
