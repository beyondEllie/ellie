package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

type DayStartConfig struct {
	Apps     []string `json:"apps"`
	Services []string `json:"services"`
	GitRepos []string `json:"git_repos"`
}

var dayStartConfig DayStartConfig

func init() {
	loadDayStartConfig()
}

func loadDayStartConfig() {
	configFile := filepath.Join(configs.GetEllieDir(), "day-start.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		dayStartConfig = DayStartConfig{
			Apps:     []string{},
			Services: []string{},
			GitRepos: []string{},
		}
		return
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		styles.ErrorStyle.Println("Error reading day-start config:", err)
		return
	}

	if err := json.Unmarshal(data, &dayStartConfig); err != nil {
		styles.ErrorStyle.Println("Error parsing day-start config:", err)
		dayStartConfig = DayStartConfig{
			Apps:     []string{},
			Services: []string{},
			GitRepos: []string{},
		}
	}
}

func saveDayStartConfig() {
	configFile := filepath.Join(configs.GetEllieDir(), "day-start.json")
	data, err := json.MarshalIndent(dayStartConfig, "", "  ")
	if err != nil {
		styles.ErrorStyle.Println("Error saving day-start config:", err)
		return
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		styles.ErrorStyle.Println("Error writing day-start config:", err)
	}
}

func StartDay(args []string) {
	styles.InfoStyle.Println("Starting your development day...")

	// Open configured apps
	if len(dayStartConfig.Apps) > 0 {
		styles.InfoStyle.Println("Opening applications...")
		for _, app := range dayStartConfig.Apps {
			openApp(app)
		}
	}

	// Start configured services
	if len(dayStartConfig.Services) > 0 {
		styles.InfoStyle.Println("Starting services...")
		for _, service := range dayStartConfig.Services {
			HandleService("start", service)
		}
	}

	// Check Git status for configured repos
	if len(dayStartConfig.GitRepos) > 0 {
		styles.InfoStyle.Println("Checking Git repositories...")
		for _, repo := range dayStartConfig.GitRepos {
			checkGitRepo(repo)
		}
	}

	// Show pending todos
	styles.InfoStyle.Println("Pending tasks:")
	TodoList(args)

	styles.SuccessStyle.Println("Your development environment is ready! 🚀")
}

func openApp(app string) {
	var cmd string
	switch runtime.GOOS {
	case "windows":
		cmd = fmt.Sprintf("start %s", app)
	case "darwin":
		cmd = fmt.Sprintf("open %s", app)
	case "linux":
		cmd = fmt.Sprintf("xdg-open %s", app)
	default:
		styles.ErrorStyle.Printf("Unsupported OS for opening apps: %s\n", runtime.GOOS)
		return
	}

	Run([]string{"run", cmd})
}

func checkGitRepo(repo string) {
	// Save current directory
	originalDir, err := os.Getwd()
	if err != nil {
		styles.ErrorStyle.Printf("Failed to get current directory: %s\n", err)
		return
	}

	// Change to repo directory
	if err := os.Chdir(repo); err != nil {
		styles.ErrorStyle.Printf("Failed to change to repo %s: %s\n", repo, err)
		return
	}

	// Check Git status
	GitStatus()

	// Return to original directory
	if err := os.Chdir(originalDir); err != nil {
		styles.ErrorStyle.Printf("Failed to return to original directory: %s\n", err)
	}
}

func DayStartConfigAdd(args []string) {
	if len(args) < 3 {
		styles.ErrorStyle.Println("Usage: ellie day-start add <type> <value>")
		styles.InfoStyle.Println("Types: apps, services, git_repos")
		return
	}

	configType := args[1]
	value := args[2]

	switch configType {
	case "apps":
		dayStartConfig.Apps = append(dayStartConfig.Apps, value)
	case "services":
		dayStartConfig.Services = append(dayStartConfig.Services, value)
	case "git_repos":
		dayStartConfig.GitRepos = append(dayStartConfig.GitRepos, value)
	default:
		styles.ErrorStyle.Println("Invalid config type. Use: apps, services, git_repos")
		return
	}

	saveDayStartConfig()
	styles.SuccessStyle.Printf("Added %s to %s\n", value, configType)
}

func DayStartConfigList(args []string) {
	styles.InfoStyle.Println("Day Start Configuration:")
	fmt.Println("Apps:")
	for _, app := range dayStartConfig.Apps {
		fmt.Printf("  - %s\n", app)
	}
	fmt.Println("\nServices:")
	for _, service := range dayStartConfig.Services {
		fmt.Printf("  - %s\n", service)
	}
	fmt.Println("\nGit Repositories:")
	for _, repo := range dayStartConfig.GitRepos {
		fmt.Printf("  - %s\n", repo)
	}
}
