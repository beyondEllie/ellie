package llm

import (
	"encoding/json"
	"testing"
)

func TestNewProvider(t *testing.T) {
	tests := []struct {
		name         string
		providerType string
		config       Config
		wantErr      bool
	}{
		{
			name:         "valid OpenAI provider",
			providerType: "openai",
			config: Config{
				APIKey: "test-key",
				Model:  "gpt-3.5-turbo",
			},
			wantErr: false,
		},
		{
			name:         "valid Gemini provider",
			providerType: "gemini",
			config: Config{
				APIKey: "test-key",
				Model:  "gemini-pro",
			},
			wantErr: false,
		},
		{
			name:         "invalid provider type",
			providerType: "invalid",
			config: Config{
				APIKey: "test-key",
			},
			wantErr: true,
		},
		{
			name:         "missing API key for OpenAI",
			providerType: "openai",
			config: Config{
				Model: "gpt-3.5-turbo",
			},
			wantErr: true,
		},
		{
			name:         "missing API key for Gemini",
			providerType: "gemini",
			config: Config{
				Model: "gemini-pro",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := NewProvider(tt.providerType, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && provider == nil {
				t.Error("NewProvider() returned nil provider when no error was expected")
			}
		})
	}
}

func TestOpenAIProvider_GetModel(t *testing.T) {
	config := Config{
		APIKey: "test-key",
		Model:  "gpt-3.5-turbo",
	}
	provider, err := NewOpenAIProvider(config)
	if err != nil {
		t.Fatalf("Failed to create OpenAI provider: %v", err)
	}

	if got := provider.GetModel(); got != config.Model {
		t.Errorf("GetModel() = %v, want %v", got, config.Model)
	}
}

func TestGeminiProvider_GetModel(t *testing.T) {
	config := Config{
		APIKey: "test-key",
		Model:  "gemini-pro",
	}
	provider, err := NewGeminiProvider(config)
	if err != nil {
		t.Fatalf("Failed to create Gemini provider: %v", err)
	}

	if got := provider.GetModel(); got != config.Model {
		t.Errorf("GetModel() = %v, want %v", got, config.Model)
	}
}

func TestMessage_JSON(t *testing.T) {
	msg := Message{
		Role:    "user",
		Content: "Hello, AI!",
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	var unmarshaled Message
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal message: %v", err)
	}

	if unmarshaled.Role != msg.Role || unmarshaled.Content != msg.Content {
		t.Errorf("Unmarshaled message = %+v, want %+v", unmarshaled, msg)
	}
}
