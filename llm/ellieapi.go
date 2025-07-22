package llm

import (
	"encoding/json"
	"fmt"
)

// EllieAPIProvider implements the Provider interface for Ellie's API
type EllieAPIProvider struct {
	config Config
}

// NewEllieAPIProvider creates a new EllieAPI provider
func NewEllieAPIProvider(config Config) (Provider, error) {
	if config.BaseURL == "" {
		config.BaseURL = "http://localhost:8000" // Default to local development server
	}

	return &EllieAPIProvider{
		config: config,
	}, nil
}

// Chat sends a chat request to the Ellie API
func (p *EllieAPIProvider) Chat(messages []Message) (*Response, error) {
	// Prepare request body
	requestBody := map[string]interface{}{
		"messages": messages,
	}

	// Make request to Ellie API
	responseData, err := makeRequest(
		fmt.Sprintf("%s/api/chat", p.config.BaseURL),
		map[string]string{
			"Content-Type": "application/json",
		},
		requestBody,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to make request to Ellie API: %w", err)
	}

	// Parse response
	var apiResponse struct {
		Content string `json:"content"`
		Usage   Usage  `json:"usage"`
	}

	if err := json.Unmarshal(responseData, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse Ellie API response: %w", err)
	}

	return &Response{
		Content: apiResponse.Content,
		Usage:   apiResponse.Usage,
	}, nil
}

// GetModel returns the model name
func (p *EllieAPIProvider) GetModel() string {
	return "ellie-api"
}