package actions

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		hasError bool
	}{
		{
			name:     "10 seconds",
			input:    "10s",
			expected: 10 * time.Second,
			hasError: false,
		},
		{
			name:     "5 minutes",
			input:    "5m",
			expected: 5 * time.Minute,
			hasError: false,
		},
		{
			name:     "2 hours",
			input:    "2h",
			expected: 2 * time.Hour,
			hasError: false,
		},
		{
			name:     "3 days",
			input:    "3d",
			expected: 3 * 24 * time.Hour,
			hasError: false,
		},
		{
			name:     "1 week",
			input:    "1w",
			expected: 7 * 24 * time.Hour,
			hasError: false,
		},
		{
			name:     "2 weeks",
			input:    "2w",
			expected: 14 * 24 * time.Hour,
			hasError: false,
		},
		{
			name:     "1.5 days",
			input:    "1.5d",
			expected: time.Duration(1.5 * 24 * float64(time.Hour)),
			hasError: false,
		},
		{
			name:     "Complex duration (hours and minutes)",
			input:    "2h30m",
			expected: 2*time.Hour + 30*time.Minute,
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "invalid",
			expected: 0,
			hasError: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
			hasError: true,
		},
		{
			name:     "Negative duration",
			input:    "-5m",
			expected: -5 * time.Minute,
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseDuration(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error for input %q, but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %q: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("For input %q, expected %v, got %v", tt.input, tt.expected, result)
				}
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "30 seconds",
			duration: 30 * time.Second,
			expected: "30 seconds",
		},
		{
			name:     "5 minutes",
			duration: 5 * time.Minute,
			expected: "5 minutes",
		},
		{
			name:     "2 hours",
			duration: 2 * time.Hour,
			expected: "2.0 hours",
		},
		{
			name:     "1 day",
			duration: 24 * time.Hour,
			expected: "1.0 days",
		},
		{
			name:     "1 week",
			duration: 7 * 24 * time.Hour,
			expected: "1.0 weeks",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDuration(tt.duration)
			if result != tt.expected {
				t.Errorf("For duration %v, expected %q, got %q", tt.duration, tt.expected, result)
			}
		})
	}
}
