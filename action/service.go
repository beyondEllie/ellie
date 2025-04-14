package actions

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/styles"
)

// CommandRunner interface for executing commands
type CommandRunner interface {
	Run(name string, args ...string) error
	Output(name string, args ...string) ([]byte, error)
	CombinedOutput(name string, args ...string) ([]byte, error)
}

// RealCommandRunner implements CommandRunner using actual exec.Command
type RealCommandRunner struct{}

func (r *RealCommandRunner) Run(name string, args ...string) error {
	return exec.Command(name, args...).Run()
}

func (r *RealCommandRunner) Output(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}

func (r *RealCommandRunner) CombinedOutput(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).CombinedOutput()
}

// Default command runner
var cmdRunner CommandRunner = &RealCommandRunner{}

type Service struct {
	Name        string
	DisplayName string
	Windows     string
	Linux       string
	MacOS       string
	CheckCmd    string
	StatusCmd   string
}

var services = map[string]Service{
	"apache": {
		Name:        "apache",
		DisplayName: "Apache Web Server",
		Windows:     "httpd",
		Linux:       "apache2",
		MacOS:       "httpd",
		CheckCmd:    "apache2 -v",
		StatusCmd:   "apache2 status",
	},
	"mysql": {
		Name:        "mysql",
		DisplayName: "MySQL Database",
		Windows:     "mysql",
		Linux:       "mysql",
		MacOS:       "mysql",
		CheckCmd:    "mysql --version",
		StatusCmd:   "mysqladmin status",
	},
	"postgres": {
		Name:        "postgres",
		DisplayName: "PostgreSQL Database",
		Windows:     "postgres",
		Linux:       "postgresql",
		MacOS:       "postgresql",
		CheckCmd:    "postgres --version",
		StatusCmd:   "pg_isready",
	},
}

func isServiceInstalled(service Service) bool {
	checkCmd := "which"
	if runtime.GOOS == "windows" {
		checkCmd = "where"
	}

	err := cmdRunner.Run(checkCmd, service.Name)
	return err == nil
}

func getServiceStatus(service Service) string {
	output, err := cmdRunner.Output(service.StatusCmd)
	if err != nil {
		return "unknown"
	}

	status := string(output)
	if strings.Contains(status, "running") || strings.Contains(status, "RUNNING") {
		return "running"
	}
	return "stopped"
}

func HandleService(action string, serviceName string) {
	if serviceName == "all" {
		for name := range services {
			handleSingleService(action, name)
		}
		return
	}

	if _, exists := services[serviceName]; !exists {
		styles.ErrorStyle.Printf("Unknown service: %s\n", serviceName)
		return
	}

	handleSingleService(action, serviceName)
}

func handleSingleService(action string, serviceName string) {
	service := services[serviceName]

	// Check if service is installed
	if !isServiceInstalled(service) {
		styles.ErrorStyle.Printf("%s is not installed on your system\n", service.DisplayName)
		return
	}

	// Get current status
	status := getServiceStatus(service)

	// Handle different actions
	switch action {
	case "start":
		if status == "running" {
			styles.InfoStyle.Printf("%s is already running\n", service.DisplayName)
			return
		}
		startService(service)
	case "stop":
		if status == "stopped" {
			styles.InfoStyle.Printf("%s is already stopped\n", service.DisplayName)
			return
		}
		stopService(service)
	case "restart":
		if status == "running" {
			stopService(service)
			time.Sleep(2 * time.Second) // Give service time to stop
		}
		startService(service)
	}
}

func startService(service Service) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("net", "start", service.Windows)
	case "linux":
		cmd = exec.Command("sudo", "service", service.Linux, "start")
	case "darwin":
		cmd = exec.Command("sudo", "brew", "services", "start", service.MacOS)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("Failed to start %s: %s\n", service.DisplayName, err)
		return
	}

	styles.SuccessStyle.Printf("%s started successfully\n", service.DisplayName)
	if len(output) > 0 {
		styles.InfoStyle.Println(string(output))
	}
}

func stopService(service Service) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("net", "stop", service.Windows)
	case "linux":
		cmd = exec.Command("sudo", "service", service.Linux, "stop")
	case "darwin":
		cmd = exec.Command("sudo", "brew", "services", "stop", service.MacOS)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("Failed to stop %s: %s\n", service.DisplayName, err)
		return
	}

	styles.SuccessStyle.Printf("%s stopped successfully\n", service.DisplayName)
	if len(output) > 0 {
		styles.InfoStyle.Println(string(output))
	}
}

func ListServices() {
	styles.InfoStyle.Println("Available services:")
	for name, service := range services {
		installed := isServiceInstalled(service)
		status := getServiceStatus(service)

		statusEmoji := "❌"
		if status == "running" {
			statusEmoji = "✅"
		}

		installedEmoji := "❌"
		if installed {
			installedEmoji = "✅"
		}

		fmt.Printf("  %s %s (%s)\n", statusEmoji, service.DisplayName, name)
		fmt.Printf("    Status: %s\n", status)
		fmt.Printf("    Installed: %s\n", installedEmoji)
	}
}
