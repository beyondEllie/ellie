package actions

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockCommandRunner implements CommandRunner for testing
type MockCommandRunner struct {
	mockOutput []byte
	mockError  error
}

func (m *MockCommandRunner) Run(name string, args ...string) error {
	return m.mockError
}

func (m *MockCommandRunner) Output(name string, args ...string) ([]byte, error) {
	return m.mockOutput, m.mockError
}

func (m *MockCommandRunner) CombinedOutput(name string, args ...string) ([]byte, error) {
	return m.mockOutput, m.mockError
}

func TestIsServiceInstalled(t *testing.T) {
	tests := []struct {
		name     string
		service  Service
		output   []byte
		err      error
		expected bool
	}{
		{
			name:     "Service installed",
			service:  services["apache"],
			output:   []byte("/usr/sbin/apache2"),
			err:      nil,
			expected: true,
		},
		{
			name:     "Service not installed",
			service:  services["apache"],
			output:   nil,
			err:      errors.New("command not found"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace command runner with mock
			originalRunner := cmdRunner
			cmdRunner = &MockCommandRunner{
				mockOutput: tt.output,
				mockError:  tt.err,
			}
			defer func() { cmdRunner = originalRunner }()

			result := isServiceInstalled(tt.service)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetServiceStatus(t *testing.T) {
	tests := []struct {
		name     string
		service  Service
		output   []byte
		err      error
		expected string
	}{
		{
			name:     "Service running",
			service:  services["apache"],
			output:   []byte("apache2 is running"),
			err:      nil,
			expected: "running",
		},
		{
			name:     "Service stopped",
			service:  services["apache"],
			output:   []byte("apache2 is stopped"),
			err:      nil,
			expected: "stopped",
		},
		{
			name:     "Error getting status",
			service:  services["apache"],
			output:   nil,
			err:      errors.New("command failed"),
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace command runner with mock
			originalRunner := cmdRunner
			cmdRunner = &MockCommandRunner{
				mockOutput: tt.output,
				mockError:  tt.err,
			}
			defer func() { cmdRunner = originalRunner }()

			result := getServiceStatus(tt.service)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHandleService(t *testing.T) {
	tests := []struct {
		name        string
		action      string
		serviceName string
		output      []byte
		err         error
	}{
		{
			name:        "Start existing service",
			action:      "start",
			serviceName: "apache",
			output:      []byte("Service started"),
			err:         nil,
		},
		{
			name:        "Stop existing service",
			action:      "stop",
			serviceName: "apache",
			output:      []byte("Service stopped"),
			err:         nil,
		},
		{
			name:        "Restart existing service",
			action:      "restart",
			serviceName: "apache",
			output:      []byte("Service restarted"),
			err:         nil,
		},
		{
			name:        "Unknown service",
			action:      "start",
			serviceName: "unknown",
			output:      nil,
			err:         nil,
		},
		{
			name:        "All services",
			action:      "start",
			serviceName: "all",
			output:      []byte("All services processed"),
			err:         nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace command runner with mock
			originalRunner := cmdRunner
			cmdRunner = &MockCommandRunner{
				mockOutput: tt.output,
				mockError:  tt.err,
			}
			defer func() { cmdRunner = originalRunner }()

			// This should not panic
			assert.NotPanics(t, func() {
				HandleService(tt.action, tt.serviceName)
			})
		})
	}
}

func TestListServices(t *testing.T) {
	// Replace command runner with mock
	originalRunner := cmdRunner
	cmdRunner = &MockCommandRunner{
		mockOutput: []byte("Service is running"),
		mockError:  nil,
	}
	defer func() { cmdRunner = originalRunner }()

	// Verify the function doesn't panic
	assert.NotPanics(t, func() {
		ListServices()
	})
}

// Helper function to mock exec.Command
var execCommand = exec.Command
