package main

import (
	"testing"

	"github.com/tacheraSasi/ellie/command"
)

func TestGetClosestMatchingCmd(t *testing.T) {
	// Create a test command registry
	testRegistry := map[string]command.Command{
		"chat":    {},
		"git":     {},
		"install": {},
		"list":    {},
		"open":    {},
		"push":    {},
	}

	tests := []struct {
		name        string
		input       string
		expected    []string
		shouldMatch bool
	}{
		{
			name:        "Exact match",
			input:       "chat",
			expected:    []string{"chat"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - chatt",
			input:       "chatt",
			expected:    []string{"chat"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - gitt",
			input:       "gitt",
			expected:    []string{"git"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - installx",
			input:       "installx",
			expected:    []string{"install"},
			shouldMatch: true,
		},
		{
			name:        "No match - completely different",
			input:       "xyz123",
			expected:    []string{},
			shouldMatch: false,
		},
		{
			name:        "Close typo - listt",
			input:       "listt",
			expected:    []string{"list"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - openn",
			input:       "openn",
			expected:    []string{"open"},
			shouldMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getClosestMatchingCmd(testRegistry, tt.input)

			if tt.shouldMatch {
				if len(result) == 0 {
					t.Errorf("Expected matches for input %q, but got none", tt.input)
					return
				}

				// Check if the expected commands are in the result
				for _, expected := range tt.expected {
					found := false
					for _, r := range result {
						if r == expected {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Expected %q in results for input %q, but got %v", expected, tt.input, result)
					}
				}
			} else {
				if len(result) != 0 {
					t.Errorf("Expected no matches for input %q, but got %v", tt.input, result)
				}
			}
		})
	}
}

func TestGetClosestMatchingSubCmd(t *testing.T) {
	// Create a test subcommand registry
	testSubCommands := map[string]command.Command{
		"status": {},
		"push":   {},
		"pull":   {},
		"commit": {},
		"branch": {},
		"add":    {},
		"list":   {},
		"delete": {},
	}

	tests := []struct {
		name        string
		input       string
		expected    []string
		shouldMatch bool
	}{
		{
			name:        "Close typo - statuss",
			input:       "statuss",
			expected:    []string{"status"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - pussh",
			input:       "pussh",
			expected:    []string{"push"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - pulll",
			input:       "pulll",
			expected:    []string{"pull"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - addd",
			input:       "addd",
			expected:    []string{"add"},
			shouldMatch: true,
		},
		{
			name:        "Close typo - listt",
			input:       "listt",
			expected:    []string{"list"},
			shouldMatch: true,
		},
		{
			name:        "No match - completely different",
			input:       "xyz123",
			expected:    []string{},
			shouldMatch: false,
		},
		{
			name:        "Close typo - branchh",
			input:       "branchh",
			expected:    []string{"branch"},
			shouldMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getClosestMatchingSubCmd(testSubCommands, tt.input)

			if tt.shouldMatch {
				if len(result) == 0 {
					t.Errorf("Expected matches for input %q, but got none", tt.input)
					return
				}

				// Check if the expected commands are in the result
				for _, expected := range tt.expected {
					found := false
					for _, r := range result {
						if r == expected {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Expected %q in results for input %q, but got %v", expected, tt.input, result)
					}
				}
			} else {
				if len(result) != 0 {
					t.Errorf("Expected no matches for input %q, but got %v", tt.input, result)
				}
			}
		})
	}
}
