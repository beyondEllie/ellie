package llm

import (
	"encoding/json"
	"fmt"
)

// GeminiProvider implements the Provider interface for Google's Gemini
type GeminiProvider struct {
	config Config
}

// NewGeminiProvider creates a new Gemini provider
func NewGeminiProvider(config Config) (Provider, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required for Gemini provider")
	}
	if config.Model == "" {
		config.Model = "gemini-pro"
	}
	return &GeminiProvider{config: config}, nil
}

// GetModel returns the model being used
func (p *GeminiProvider) GetModel() string {
	return p.config.Model
}

// Chat sends a chat request to Gemini
func (p *GeminiProvider) Chat(messages []Message) (*Response, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent", p.config.Model)
	if p.config.BaseURL != "" {
		url = p.config.BaseURL
	}

	headers := map[string]string{
		"x-goog-api-key": p.config.APIKey,
		"Content-Type":   "application/json",
	}

	// Convert messages to Gemini format
	var contents []map[string]interface{}
	for _, msg := range messages {
		contents = append(contents, map[string]interface{}{
			"role": msg.Role,
			"parts": []map[string]string{
				{"text": msg.Content},
			},
		})
	}

	requestBody := map[string]interface{}{
		"contents": contents,
	}

	responseBody, err := makeRequest(url, headers, requestBody)
	if err != nil {
		return nil, fmt.Errorf("Gemini API request failed: %w", err)
	}

	var response struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
		PromptFeedback struct {
			TokenCount int `json:"tokenCount"`
		} `json:"promptFeedback"`
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing Gemini response: %w", err)
	}

	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no content in Gemini response")
	}

	return &Response{
		Content: response.Candidates[0].Content.Parts[0].Text,
		Usage: Usage{
			PromptTokens: response.PromptFeedback.TokenCount,
		},
	}, nil
}
