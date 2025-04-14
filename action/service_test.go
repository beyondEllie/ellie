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

func TestStartService(t *testing.T) {
	tests := []struct {
		name        string
		service     Service
		output      []byte
		err         error
		shouldError bool
	}{
		{
			name:        "Successful start",
			service:     services["apache"],
			output:      []byte("Service started successfully"),
			err:         nil,
			shouldError: false,
		},
		{
			name:        "Failed start",
			service:     services["apache"],
			output:      []byte("Failed to start service"),
			err:         errors.New("start failed"),
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalRunner := cmdRunner
			cmdRunner = &MockCommandRunner{
				mockOutput: tt.output,
				mockError:  tt.err,
			}
			defer func() { cmdRunner = originalRunner }()

			startService(tt.service)
			// Since startService doesn't return errors, we can't directly test for them
			// Instead, we can verify the expected behavior through the mock
		})
	}
}

func TestStopService(t *testing.T) {
	tests := []struct {
		name        string
		service     Service
		output      []byte
		err         error
		shouldError bool
	}{
		{
			name:        "Successful stop",
			service:     services["apache"],
			output:      []byte("Service stopped successfully"),
			err:         nil,
			shouldError: false,
		},
		{
			name:        "Failed stop",
			service:     services["apache"],
			output:      []byte("Failed to stop service"),
			err:         errors.New("stop failed"),
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalRunner := cmdRunner
			cmdRunner = &MockCommandRunner{
				mockOutput: tt.output,
				mockError:  tt.err,
			}
			defer func() { cmdRunner = originalRunner }()

			stopService(tt.service)
			// Since stopService doesn't return errors, we can't directly test for them
			// Instead, we can verify the expected behavior through the mock
		})
	}
}

func TestHandleSingleService(t *testing.T) {
	tests := []struct {
		name        string
		action      string
		serviceName string
		status      string
		output      []byte
		err         error
	}{
		{
			name:        "Start already running service",
			action:      "start",
			serviceName: "apache",
			status:      "running",
			output:      []byte("Service is running"),
			err:         nil,
		},
		{
			name:        "Start stopped service",
			action:      "start",
			serviceName: "apache",
			status:      "stopped",
			output:      []byte("Service started"),
			err:         nil,
		},
		{
			name:        "Stop running service",
			action:      "stop",
			serviceName: "apache",
			status:      "running",
			output:      []byte("Service stopped"),
			err:         nil,
		},
		{
			name:        "Stop already stopped service",
			action:      "stop",
			serviceName: "apache",
			status:      "stopped",
			output:      []byte("Service is stopped"),
			err:         nil,
		},
		{
			name:        "Restart running service",
			action:      "restart",
			serviceName: "apache",
			status:      "running",
			output:      []byte("Service restarted"),
			err:         nil,
		},
		{
			name:        "Restart stopped service",
			action:      "restart",
			serviceName: "apache",
			status:      "stopped",
			output:      []byte("Service restarted"),
			err:         nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalRunner := cmdRunner
			cmdRunner = &MockCommandRunner{
				mockOutput: tt.output,
				mockError:  tt.err,
			}
			defer func() { cmdRunner = originalRunner }()

			// Mock the service status check
			service := services[tt.serviceName]
			originalStatus := service.StatusCmd
			service.StatusCmd = "echo " + tt.status
			defer func() { service.StatusCmd = originalStatus }()

			// This should not panic
			assert.NotPanics(t, func() {
				handleSingleService(tt.action, tt.serviceName)
			})
		})
	}
}

func TestServiceStruct(t *testing.T) {
	tests := []struct {
		name        string
		service     Service
		expected    Service
		description string
	}{
		{
			name:    "Apache service",
			service: services["apache"],
			expected: Service{
				Name:        "apache",
				DisplayName: "Apache Web Server",
				Windows:     "httpd",
				Linux:       "apache2",
				MacOS:       "httpd",
				CheckCmd:    "apache2 -v",
				StatusCmd:   "apache2 status",
			},
			description: "Apache service should have correct configuration",
		},
		{
			name:    "MySQL service",
			service: services["mysql"],
			expected: Service{
				Name:        "mysql",
				DisplayName: "MySQL Database",
				Windows:     "mysql",
				Linux:       "mysql",
				MacOS:       "mysql",
				CheckCmd:    "mysql --version",
				StatusCmd:   "mysqladmin status",
			},
			description: "MySQL service should have correct configuration",
		},
		{
			name:    "PostgreSQL service",
			service: services["postgres"],
			expected: Service{
				Name:        "postgres",
				DisplayName: "PostgreSQL Database",
				Windows:     "postgres",
				Linux:       "postgresql",
				MacOS:       "postgresql",
				CheckCmd:    "postgres --version",
				StatusCmd:   "pg_isready",
			},
			description: "PostgreSQL service should have correct configuration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected.Name, tt.service.Name, "Service name should match")
			assert.Equal(t, tt.expected.DisplayName, tt.service.DisplayName, "Display name should match")
			assert.Equal(t, tt.expected.Windows, tt.service.Windows, "Windows command should match")
			assert.Equal(t, tt.expected.Linux, tt.service.Linux, "Linux command should match")
			assert.Equal(t, tt.expected.MacOS, tt.service.MacOS, "MacOS command should match")
			assert.Equal(t, tt.expected.CheckCmd, tt.service.CheckCmd, "Check command should match")
			assert.Equal(t, tt.expected.StatusCmd, tt.service.StatusCmd, "Status command should match")
		})
	}
}
