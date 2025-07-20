package actions

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/common"
	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/types"
	"github.com/tacheraSasi/ellie/utils"
)

// isInstalled checks if a command is available and working
func IsInstalled(cmd string) bool {
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
func runServerInstallCommand(tool types.DevTool, currentOS string) bool {
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
	if !IsInstalled(tool.CheckCmd) {
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
func CheckPackageManager(os string) bool {
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
	if IsInstalled(pmCommand) {
		styles.SuccessStyle.Printf("✅ %s is already installed\n", pmName)
		return true
	}
	styles.ErrorStyle.Printf("❌ %s is required but not installed\n", pmName)
	styles.InfoStyle.Printf("📦 Please install %s to proceed\n", pmName)
	return false
}

// ServerInit initializes a server environment for a chosen framework
func ServerInit() {
	session := &common.ServerInitSession{
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
	if !CheckPackageManager(session.OS) {
		styles.ErrorStyle.Println("❌ Please install a package manager and try again")
		return
	}

	// Get server name
	serverName, err := utils.GetInput("Enter the server name: ")
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error reading server name: %v\n", err)
		return
	}
	serverDir := filepath.Join(configs.GetEllieDir(), "servers", serverName)
	session.ServerName = serverDir

	// Create server directory and config file
	err = os.MkdirAll(serverDir, 0755)
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error creating server directory: %v\n", err)
		return
	}
	err = os.WriteFile(filepath.Join(serverDir, "config.json"), []byte("{}"), 0644)
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error creating config file: %v\n", err)
		return
	}

	// Display framework options
	styles.HeaderStyle.Println("\n📋 Choose Your Server Framework")
	frameworkList := []common.Framework{
		common.Frameworks["general"],
		common.Frameworks["laravel"],
		common.Frameworks["nodejs"],
		common.Frameworks["django"],
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
	os.Chdir(serverDir)
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
		if IsInstalled(tool.CheckCmd) {
			styles.SuccessStyle.Printf("✅ %s already installed\n", tool.Name)
			session.SkippedCount++
			continue
		}
		if runServerInstallCommand(*tool, session.OS) {
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
	styles.InfoStyle.Printf("📂 Server directory: %s\n", serverDir)
	styles.SuccessStyle.Println("\n🎉 Server environment ready!")
}
