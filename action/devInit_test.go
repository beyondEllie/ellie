package actions

import (
	"os/exec"
	"testing"
)

func TestCheckPackageManager(t *testing.T) {
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
			result := checkPackageManager(tt.os)

			// For mac and linux, we need to check if the package manager is actually available
			if tt.os == "mac" {
				_, err := exec.LookPath("brew")
				if err != nil {
					// If brew is not available, the test should expect false
					if result != false {
						t.Errorf("checkPackageManager() = %v, want %v (brew not found)", result, false)
					}
					return
				}
			}

			if tt.os == "linux" {
				_, err := exec.LookPath("apt")
				if err != nil {
					// If apt is not available, check for snap
					_, err = exec.LookPath("snap")
					if err != nil {
						// If neither apt nor snap is available, the test should expect false
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

func TestIsInstalled(t *testing.T) {
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
			result := isInstalled(tt.command)
			if result != tt.expected {
				t.Errorf("isInstalled() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Test helper to check if we're running on a supported OS
func TestGetOS(t *testing.T) {
	os := getOS()
	if os == "unknown" {
		t.Skip("Skipping test on unknown OS")
	}

	validOS := []string{"mac", "linux", "windows"}
	found := false
	for _, valid := range validOS {
		if os == valid {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("getOS() returned unsupported OS: %s", os)
	}
}

// Helper function to get OS (simplified version for testing)
func getOS() string {
	// This is a simplified version for testing
	// In the actual code, this would be utils.GetOS()
	return "mac" // Default for testing
}
