package llm

import (
	"encoding/json"
	"fmt"
)

// OpenAIProvider implements the Provider interface for OpenAI
type OpenAIProvider struct {
	config Config
}

// NewOpenAIProvider creates a new OpenAI provider
func NewOpenAIProvider(config Config) (Provider, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required for OpenAI provider")
	}
	if config.Model == "" {
		config.Model = "gpt-3.5-turbo"
	}
	return &OpenAIProvider{config: config}, nil
}

// GetModel returns the model being used
func (p *OpenAIProvider) GetModel() string {
	return p.config.Model
}

// Chat sends a chat request to OpenAI
func (p *OpenAIProvider) Chat(messages []Message) (*Response, error) {
	url := "https://api.openai.com/v1/chat/completions"
	if p.config.BaseURL != "" {
		url = p.config.BaseURL
	}

	headers := map[string]string{
		"Authorization": "Bearer " + p.config.APIKey,
		"Content-Type":  "application/json",
	}

	requestBody := map[string]interface{}{
		"model":    p.config.Model,
		"messages": messages,
	}

	responseBody, err := makeRequest(url, headers, requestBody)
	if err != nil {
		return nil, fmt.Errorf("OpenAI API request failed: %w", err)
	}

	var response struct {
		Choices []struct {
			Message Message `json:"message"`
		} `json:"choices"`
		Usage Usage `json:"usage"`
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing OpenAI response: %w", err)
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no choices in OpenAI response")
	}

	return &Response{
		Content: response.Choices[0].Message.Content,
		Usage:   response.Usage,
	}, nil
}
