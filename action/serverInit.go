package actions

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/common"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/types"
	"github.com/tacheraSasi/ellie/utils"
)

// ServerInitSession represents a server environment setup session
type ServerInitSession struct {
	OS             string
	ServerName     string
	SuccessCount   int
	SkippedCount   int
	FailedCount    int
	StartTime      time.Time
	EndTime        time.Time
	InstalledTools []string
	FailedTools    []string
	Framework      string
}

// Framework represents a server framework with required tools and setup commands
type Framework struct {
	Name          string
	Description   string
	RequiredTools []string
	SetupCommands []string
}

// frameworks is a map of available server frameworks
var frameworks = map[string]Framework{
	"general": {
		Name:          "General",
		Description:   "General server setup with common tools",
		RequiredTools: []string{"Git", "Node.js", "Python", "Docker", "NGINX"},
		SetupCommands: []string{},
	},
	"laravel": {
		Name:          "Laravel",
		Description:   "PHP-based Laravel framework server",
		RequiredTools: []string{"PHP", "Composer", "Node.js", "MySQL", "NGINX"},
		SetupCommands: []string{
			"composer global require laravel/installer",
			"laravel new project --no-interaction",
		},
	},
	"nodejs": {
		Name:          "Node.js",
		Description:   "Node.js server with Express",
		RequiredTools: []string{"Node.js", "Yarn"},
		SetupCommands: []string{
			"mkdir project && cd project",
			"yarn init -y",
			"yarn add express",
		},
	},
	"django": {
		Name:          "Django",
		Description:   "Python-based Django framework server",
		RequiredTools: []string{"Python", "PostgreSQL"},
		SetupCommands: []string{
			"pip install django psycopg2-binary",
			"django-admin startproject project",
		},
	},
}

// isInstalled checks if a command is available and working
func isInstalled(cmd string) bool {
	parts := strings.Split(cmd, " ")
	if len(parts) == 0 {
		return false
	}
	_, err := exec.LookPath(parts[0])
	if err != nil {
		return false
	}
	out, err := exec.Command(parts[0], parts[1:]...).CombinedOutput()
	return err == nil && len(out) > 0
}

// runInstallCommand installs a server tool
func runInstallCommand(tool types.DevTool, currentOS string) bool {
	styles.InfoStyle.Printf("🚀 Installing %s...\n", tool.Name)
	cmd := tool.Install[currentOS]
	if cmd == "" {
		cmd = tool.Install["common"]
	}
	if cmd == "" {
		styles.ErrorStyle.Printf("❌ No installation command for %s on %s\n", tool.Name, currentOS)
		return false
	}
	styles.DimText.Println("Running:", cmd)
	parts := strings.Split(cmd, " ")
	c := exec.Command(parts[0], parts[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		styles.ErrorStyle.Printf("❌ Failed to install %s: %v\n", tool.Name, err)
		return false
	}
	if !isInstalled(tool.CheckCmd) {
		styles.ErrorStyle.Printf("❌ Installation verification failed for %s\n", tool.Name)
		return false
	}
	styles.SuccessStyle.Printf("✅ Successfully installed %s\n", tool.Name)
	if configCmd, exists := tool.Configure[currentOS]; exists {
		styles.InfoStyle.Printf("⚙️ Configuring %s...\n", tool.Name)
		configParts := strings.Split(configCmd, " ")
		config := exec.Command(configParts[0], configParts[1:]...)
		config.Stdout = os.Stdout
		config.Stderr = os.Stderr
		if err := config.Run(); err != nil {
			styles.ErrorStyle.Printf("❌ Configuration failed for %s: %v\n", tool.Name, err)
			return true // Installation succeeded, configuration failed
		}
		styles.SuccessStyle.Printf("✅ Successfully configured %s\n", tool.Name)
	}
	return true
}

// checkPackageManager verifies if the required package manager is installed
func checkPackageManager(os string) bool {
	var pmName, pmCommand string
	switch os {
	case "mac":
		pmName, pmCommand = "Homebrew", "brew --version"
	case "linux":
		pmName, pmCommand = "apt", "apt --version"
	case "windows":
		pmName, pmCommand = "Chocolatey", "choco --version"
	default:
		return false
	}
	if isInstalled(pmCommand) {
		styles.SuccessStyle.Printf("✅ %s is already installed\n", pmName)
		return true
	}
	styles.ErrorStyle.Printf("❌ %s is required but not installed\n", pmName)
	styles.InfoStyle.Printf("📦 Please install %s to proceed\n", pmName)
	return false
}

// ServerInit initializes a server environment for a chosen framework
func ServerInit() {
	session := &ServerInitSession{
		OS:        utils.GetOS(),
		StartTime: time.Now(),
	}

	if session.OS == "unknown" {
		styles.ErrorStyle.Println("❌ Unsupported operating system")
		return
	}

	styles.HeaderStyle.Println("🚀 Rapid Server Environment Setup")
	styles.InfoStyle.Printf("Detected OS: %s\n\n", strings.ToUpper(session.OS))

	// Check package manager
	if !checkPackageManager(session.OS) {
		styles.ErrorStyle.Println("❌ Please install a package manager and try again")
		return
	}

	// Get server name
	serverName, err := utils.GetInput("Enter the server name: ")
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error reading server name: %v\n", err)
		return
	}
	session.ServerName = serverName

	// Create server directory and config file
	err = os.Mkdir(serverName, 0755)
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error creating server directory: %v\n", err)
		return
	}
	err = os.WriteFile(filepath.Join(serverName, "config.json"), []byte("{}"), 0644)
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error creating config file: %v\n", err)
		return
	}

	// Display framework options
	styles.HeaderStyle.Println("\n📋 Choose Your Server Framework")
	frameworkList := []Framework{
		frameworks["general"],
		frameworks["laravel"],
		frameworks["nodejs"],
		frameworks["django"],
	}
	for i, fw := range frameworkList {
		styles.InfoStyle.Printf("  %d. %s - %s\n", i+1, fw.Name, fw.Description)
	}
	choice, _ := utils.GetInput("Select a framework (number): ")
	chosenIndex := utils.StringToInt(choice) - 1
	if chosenIndex < 0 || chosenIndex >= len(frameworkList) {
		styles.ErrorStyle.Println("❌ Invalid framework choice")
		return
	}
	chosenFramework := frameworkList[chosenIndex]
	session.Framework = chosenFramework.Name

	// Change to server directory for setup
	originalDir, _ := os.Getwd()
	os.Chdir(serverName)
	defer os.Chdir(originalDir)

	// Install framework tools
	styles.HeaderStyle.Printf("\n🛠️ Setting Up %s Server Environment\n", chosenFramework.Name)
	for _, toolName := range chosenFramework.RequiredTools {
		var tool *types.DevTool
		for _, t := range common.Tools {
			if t.Name == toolName {
				tool = &t
				break
			}
		}
		if tool == nil {
			styles.ErrorStyle.Printf("❌ Tool %s not found in common.Tools\n", toolName)
			session.FailedCount++
			session.FailedTools = append(session.FailedTools, toolName)
			continue
		}
		if isInstalled(tool.CheckCmd) {
			styles.SuccessStyle.Printf("✅ %s already installed\n", tool.Name)
			session.SkippedCount++
			continue
		}
		if runInstallCommand(*tool, session.OS) {
			session.SuccessCount++
			session.InstalledTools = append(session.InstalledTools, tool.Name)
		} else {
			session.FailedCount++
			session.FailedTools = append(session.FailedTools, tool.Name)
		}
	}

	// Run framework setup commands
	if len(chosenFramework.SetupCommands) > 0 {
		styles.HeaderStyle.Println("\n⚙️ Running Framework Setup Commands")
		for _, cmd := range chosenFramework.SetupCommands {
			parts := strings.Split(cmd, " ")
			c := exec.Command(parts[0], parts[1:]...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			styles.DimText.Printf("Running: %s\n", cmd)
			if err := c.Run(); err != nil {
				styles.ErrorStyle.Printf("❌ Failed: %s\n", cmd)
			} else {
				styles.SuccessStyle.Printf("✅ Completed: %s\n", cmd)
			}
		}
	}

	// Summary
	session.EndTime = time.Now()
	styles.HeaderStyle.Println("\n📊 Server Setup Summary")
	styles.InfoStyle.Printf("Server: %s\n", session.ServerName)
	styles.InfoStyle.Printf("Framework: %s\n", session.Framework)
	styles.InfoStyle.Printf("Time: %s\n", session.EndTime.Sub(session.StartTime))
	styles.SuccessStyle.Printf("Installed: %d tools\n", session.SuccessCount)
	styles.InfoStyle.Printf("Skipped: %d tools\n", session.SkippedCount)
	styles.ErrorStyle.Printf("Failed: %d tools\n", session.FailedCount)
	if len(session.InstalledTools) > 0 {
		styles.SuccessStyle.Println("Installed tools:", strings.Join(session.InstalledTools, ", "))
	}
	if len(session.FailedTools) > 0 {
		styles.ErrorStyle.Println("Failed tools:", strings.Join(session.FailedTools, ", "))
	}
	styles.InfoStyle.Printf("📂 Server directory: %s\n", filepath.Join(originalDir, serverName))
	styles.SuccessStyle.Println("\n🎉 Server environment ready!")
}
