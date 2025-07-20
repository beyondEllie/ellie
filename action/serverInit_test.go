package actions

import (
	"os/exec"
	"testing"
)

func TestCheckPackageManager_ServerInit(t *testing.T) {
	tests := []struct {
		name     string
		os       string
		expected bool
	}{
		{
			name:     "mac with brew",
			os:       "mac",
			expected: true, // Should be true if brew is installed
		},
		{
			name:     "linux with apt",
			os:       "linux",
			expected: true, // Should be true if apt is installed
		},
		{
			name:     "windows with choco",
			os:       "windows",
			expected: false, // Should be false if choco is not installed
		},
		{
			name:     "unknown os",
			os:       "unknown",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckPackageManager(tt.os)

			if tt.os == "mac" {
				_, err := exec.LookPath("brew")
				if err != nil {
					if result != false {
						t.Errorf("checkPackageManager() = %v, want %v (brew not found)", result, false)
					}
					return
				}
			}

			if tt.os == "linux" {
				_, err := exec.LookPath("apt")
				if err != nil {
					_, err = exec.LookPath("snap")
					if err != nil {
						if result != false {
							t.Errorf("checkPackageManager() = %v, want %v (neither apt nor snap found)", result, false)
						}
						return
					}
				}
			}

			if result != tt.expected {
				t.Errorf("checkPackageManager() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsInstalled_ServerInit(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected bool
	}{
		{
			name:     "existing command",
			command:  "ls",
			expected: true,
		},
		{
			name:     "non-existing command",
			command:  "nonexistentcommand12345",
			expected: false,
		},
		{
			name:     "command with args",
			command:  "ls -la",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInstalled(tt.command)
			if result != tt.expected {
				t.Errorf("isInstalled() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Note: Full integration test for ServerInit would require extensive mocking of user input and OS calls, so is omitted here.
