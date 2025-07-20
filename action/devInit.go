package actions

import (
	"fmt"
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

// DevInitSession represents a development environment setup session
type DevInitSession struct {
	OS              string
	InstallAll      bool
	SuccessCount    int
	SkippedCount    int
	FailedCount     int
	StartTime       time.Time
	EndTime         time.Time
	InstalledTools  []string
	FailedTools     []string
	ProjectTemplate string
	SetupGit        bool
	SetupSSH        bool
	SetupAliases    bool
}

// ProjectTemplate represents a project template configuration
type ProjectTemplate struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tools       []string `json:"tools"`
	Commands    []string `json:"commands"`
	Files       []string `json:"files"`
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

// runInstallCommand installs a development tool
func runInstallCommand(tool types.DevTool, currentOS string) bool {
	styles.InfoStyle.Printf("🚀 Installing %s...\n", tool.Name)

	cmd := tool.Install[currentOS]
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

	// Run configuration if available
	if configCmd, exists := tool.Configure[currentOS]; exists {
		styles.InfoStyle.Printf("⚙️  Configuring %s...\n", tool.Name)
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

// setupGitConfiguration configures Git with user information
func setupGitConfiguration() {
	styles.InfoStyle.Println("🔧 Setting up Git configuration...")

	// Get user information
	username, _ := utils.GetInput("Enter your Git username: ")
	email, _ := utils.GetInput("Enter your Git email: ")

	if username != "" {
		exec.Command("git", "config", "--global", "user.name", username).Run()
		styles.SuccessStyle.Printf("✅ Git username set to: %s\n", username)
	}

	if email != "" {
		exec.Command("git", "config", "--global", "user.email", email).Run()
		styles.SuccessStyle.Printf("✅ Git email set to: %s\n", email)
	}

	// Set default branch to main
	exec.Command("git", "config", "--global", "init.defaultBranch", "main").Run()
	exec.Command("git", "config", "--global", "core.autocrlf", "input").Run()

	styles.SuccessStyle.Println("✅ Git configuration completed")
}

// setupSSHKey generates SSH key for GitHub/GitLab
func setupSSHKey() {
	styles.InfoStyle.Println("🔑 Setting up SSH key...")

	// Check if SSH key already exists
	homeDir, _ := os.UserHomeDir()
	sshDir := filepath.Join(homeDir, ".ssh")
	idRsaPath := filepath.Join(sshDir, "id_rsa")

	if _, err := os.Stat(idRsaPath); err == nil {
		styles.InfoStyle.Println("✅ SSH key already exists")
		return
	}

	// Create SSH directory if it doesn't exist
	os.MkdirAll(sshDir, 0700)

	// Generate SSH key
	email, _ := utils.GetInput("Enter your email for SSH key: ")
	if email == "" {
		email = "user@example.com"
	}

	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-C", email, "-f", idRsaPath, "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		styles.ErrorStyle.Printf("❌ Failed to generate SSH key: %v\n", err)
		return
	}

	styles.SuccessStyle.Println("✅ SSH key generated successfully")
	styles.InfoStyle.Println("📋 Add this public key to your GitHub/GitLab account:")

	// Display public key
	publicKeyPath := idRsaPath + ".pub"
	if content, err := os.ReadFile(publicKeyPath); err == nil {
		styles.Highlight.Println(string(content))
	}
}

// setupUsefulAliases creates helpful development aliases
func setupUsefulAliases() {
	styles.InfoStyle.Println("⚡ Setting up useful aliases...")

	aliases := map[string]string{
		"gs":      "git status",
		"ga":      "git add .",
		"gc":      "git commit -m",
		"gp":      "git push",
		"gl":      "git log --oneline",
		"gd":      "git diff",
		"gco":     "git checkout",
		"gcb":     "git checkout -b",
		"gpl":     "git pull",
		"gst":     "git stash",
		"gstp":    "git stash pop",
		"ll":      "ls -la",
		"la":      "ls -A",
		"l":       "ls -CF",
		"..":      "cd ..",
		"...":     "cd ../..",
		"....":    "cd ../../..",
		".....":   "cd ../../../..",
		"c":       "clear",
		"h":       "history",
		"j":       "jobs -l",
		"ports":   "netstat -tulanp",
		"myip":    "curl http://ipecho.net/plain; echo",
		"weather": "curl wttr.in",
	}

	shell := os.Getenv("SHELL")
	var configFile string

	if strings.Contains(shell, "zsh") {
		configFile = filepath.Join(os.Getenv("HOME"), ".zshrc")
	} else if strings.Contains(shell, "bash") {
		configFile = filepath.Join(os.Getenv("HOME"), ".bashrc")
	} else {
		styles.WarningStyle.Println("⚠️  Unsupported shell, aliases not configured")
		return
	}

	// Read existing config
	content, _ := os.ReadFile(configFile)
	configContent := string(content)

	// Add aliases if they don't exist
	aliasSection := "\n# Ellie Development Aliases\n"
	for alias, command := range aliases {
		if !strings.Contains(configContent, fmt.Sprintf("alias %s=", alias)) {
			aliasSection += fmt.Sprintf("alias %s='%s'\n", alias, command)
		}
	}

	// Append aliases to config file
	if aliasSection != "\n# Ellie Development Aliases\n" {
		file, _ := os.OpenFile(configFile, os.O_APPEND|os.O_WRONLY, 0644)
		file.WriteString(aliasSection)
		file.Close()
		styles.SuccessStyle.Println("✅ Aliases added to shell configuration")
		styles.InfoStyle.Println("🔄 Restart your terminal or run 'source " + configFile + "' to activate aliases")
	} else {
		styles.InfoStyle.Println("✅ Aliases already configured")
	}
}

// getProjectTemplates returns available project templates
func getProjectTemplates() []ProjectTemplate {
	return []ProjectTemplate{
		{
			Name:        "react-app",
			Description: "React application with TypeScript",
			Tools:       []string{"Node.js", "Yarn"},
			Commands:    []string{"npx create-react-app my-app --template typescript", "cd my-app", "yarn start"},
		},
		{
			Name:        "node-api",
			Description: "Node.js API with Express",
			Tools:       []string{"Node.js", "Yarn"},
			Commands:    []string{"mkdir my-api", "cd my-api", "yarn init -y", "yarn add express cors dotenv"},
		},
		{
			Name:        "python-web",
			Description: "Python web application with Flask",
			Tools:       []string{"Python"},
			Commands:    []string{"mkdir my-python-app", "cd my-python-app", "python -m venv venv", "pip install flask"},
		},
		{
			Name:        "go-api",
			Description: "Go API with Gin framework",
			Tools:       []string{"Go"},
			Commands:    []string{"mkdir my-go-api", "cd my-go-api", "go mod init my-api", "go get github.com/gin-gonic/gin"},
		},
		{
			Name:        "docker-project",
			Description: "Docker-based development environment",
			Tools:       []string{"Docker"},
			Commands:    []string{"mkdir my-docker-project", "cd my-docker-project"},
		},
	}
}

// createProjectFromTemplate creates a new project from a template
func createProjectFromTemplate(template ProjectTemplate) {
	styles.InfoStyle.Printf("📁 Creating %s project...\n", template.Name)

	projectName, _ := utils.GetInput("Enter project name: ")
	if projectName == "" {
		projectName = "my-" + template.Name
	}

	// Create project directory
	projectPath := filepath.Join(".", projectName)
	os.MkdirAll(projectPath, 0755)

	// Change to project directory
	originalDir, _ := os.Getwd()
	os.Chdir(projectPath)
	defer os.Chdir(originalDir)

	// Execute template commands
	for _, cmd := range template.Commands {
		parts := strings.Split(cmd, " ")
		command := exec.Command(parts[0], parts[1:]...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		styles.DimText.Printf("Running: %s\n", cmd)
		if err := command.Run(); err != nil {
			styles.ErrorStyle.Printf("❌ Command failed: %s\n", cmd)
		}
	}

	styles.SuccessStyle.Printf("✅ Project '%s' created successfully!\n", projectName)
	styles.InfoStyle.Printf("📂 Project location: %s\n", projectPath)
}

// checkPackageManager verifies if the required package manager is installed
func checkPackageManager(os string) bool {
	var packageManagers []struct {
		name    string
		command string
		url     string
	}

	switch os {
	case "mac":
		packageManagers = []struct {
			name    string
			command string
			url     string
		}{
			{"Homebrew", "brew --version", "https://brew.sh"},
		}
	case "linux":
		packageManagers = []struct {
			name    string
			command string
			url     string
		}{
			{"apt", "apt --version", "https://help.ubuntu.com/community/AptGet/Howto"},
			{"snap", "snap --version", "https://snapcraft.io/docs/installing-snapd"},
		}
	case "windows":
		packageManagers = []struct {
			name    string
			command string
			url     string
		}{
			{"Chocolatey", "choco --version", "https://chocolatey.org/install"},
		}
	default:
		return false
	}

	// Check if any package manager is installed
	for _, pm := range packageManagers {
		if isInstalled(pm.command) {
			styles.SuccessStyle.Printf("✅ %s is already installed\n", pm.name)
			return true
		}
	}

	// No package manager found, prompt user to install
	styles.ErrorStyle.Printf("❌ No package manager found for %s\n", strings.ToUpper(os))
	styles.InfoStyle.Printf("📦 A package manager is required to install development tools on %s\n", strings.ToUpper(os))

	// Provide installation commands for each OS
	switch os {
	case "mac":
		styles.InfoStyle.Println("💻 To install Homebrew, run:")
		styles.Highlight.Println("/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\"")
		styles.InfoStyle.Printf("🔗 Installation guide: %s\n", packageManagers[0].url)
	case "linux":
		styles.InfoStyle.Println("💻 To install apt (usually pre-installed on Ubuntu/Debian):")
		styles.Highlight.Println("sudo apt update && sudo apt upgrade")
		styles.InfoStyle.Printf("🔗 Installation guide: %s\n", packageManagers[0].url)
		styles.InfoStyle.Println("💻 To install snap (alternative):")
		styles.Highlight.Println("sudo apt install snapd")
		styles.InfoStyle.Printf("🔗 Installation guide: %s\n", packageManagers[1].url)
	case "windows":
		styles.InfoStyle.Println("💻 To install Chocolatey, run in PowerShell as Administrator:")
		styles.Highlight.Println("Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))")
		styles.InfoStyle.Printf("🔗 Installation guide: %s\n", packageManagers[0].url)
	}

	// Ask user if they want to continue anyway
	styles.WarningStyle.Println("⚠️  Some tools may fail to install without a package manager")
	continueAnyway := utils.AskYesNo("Continue with dev init anyway?", false)

	if !continueAnyway {
		styles.InfoStyle.Println("🔄 Please install a package manager and run 'ellie dev init' again")
		return false
	}

	styles.WarningStyle.Println("⚠️  Proceeding without package manager - some installations may fail")
	return true
}

// DevInit is the enhanced development environment setup function
func DevInit(installAll bool) {
	session := &DevInitSession{
		OS:         utils.GetOS(),
		InstallAll: installAll,
		StartTime:  time.Now(),
	}

	if session.OS == "unknown" {
		styles.ErrorStyle.Println("❌ Unsupported operating system")
		return
	}

	styles.HeaderStyle.Println("🚀 Enhanced Development Environment Setup")
	styles.InfoStyle.Printf("Detected OS: %s\n\n", strings.ToUpper(session.OS))

	// Check for required package manager
	styles.HeaderStyle.Println("📦 Package Manager Check")
	if !checkPackageManager(session.OS) {
		return
	}
	styles.InfoStyle.Println("") // Add spacing

	// Welcome and options
	styles.InfoStyle.Println("Welcome to Ellie's Enhanced Dev Init!")
	styles.InfoStyle.Println("This will help you set up a complete development environment.")

	// Ask for setup preferences
	session.SetupGit = utils.AskYesNo("Configure Git with your credentials?", true)
	session.SetupSSH = utils.AskYesNo("Generate SSH key for GitHub/GitLab?", false)
	session.SetupAliases = utils.AskYesNo("Set up useful development aliases?", true)

	// Ask about project template
	styles.InfoStyle.Println("\n📋 Available project templates:")
	templates := getProjectTemplates()
	for i, template := range templates {
		styles.InfoStyle.Printf("  %d. %s - %s\n", i+1, template.Name, template.Description)
	}

	templateChoice, _ := utils.GetInput("Create a project template? (number or 'n' for skip): ")
	if templateChoice != "n" && templateChoice != "" {
		if choice := utils.StringToInt(templateChoice); choice > 0 && choice <= len(templates) {
			session.ProjectTemplate = templates[choice-1].Name
		}
	}

	// Install development tools
	styles.HeaderStyle.Println("\n🛠️  Installing Development Tools")
	styles.InfoStyle.Println("Installing essential development tools...")

	for _, tool := range common.Tools {
		styles.Highlight.Printf("\n%s - %s\n", tool.Name, tool.Description)

		if isInstalled(tool.CheckCmd) {
			styles.SuccessStyle.Printf("✅ %s is already installed\n", tool.Name)
			session.SkippedCount++
			continue
		}

		// Determine if we should install
		var shouldInstall bool
		if installAll {
			shouldInstall = tool.DefaultInstall
		} else {
			prompt := fmt.Sprintf("Install %s? (default: %t)", tool.Name, tool.DefaultInstall)
			shouldInstall = utils.AskYesNo(prompt, tool.DefaultInstall)
		}

		if !shouldInstall {
			styles.InfoStyle.Printf("⏩ Skipping %s installation\n", tool.Name)
			session.SkippedCount++
			continue
		}

		_, exists := tool.Install[session.OS]
		if !exists {
			styles.ErrorStyle.Printf("❌ No installation command for %s on %s\n", tool.Name, session.OS)
			session.FailedCount++
			session.FailedTools = append(session.FailedTools, tool.Name)
			continue
		}

		if runInstallCommand(tool, session.OS) {
			session.SuccessCount++
			session.InstalledTools = append(session.InstalledTools, tool.Name)
		} else {
			session.FailedCount++
			session.FailedTools = append(session.FailedTools, tool.Name)
		}
	}

	// Setup additional configurations
	styles.HeaderStyle.Println("\n⚙️  Additional Setup")

	if session.SetupGit {
		setupGitConfiguration()
	}

	if session.SetupSSH {
		setupSSHKey()
	}

	if session.SetupAliases {
		setupUsefulAliases()
	}

	// Create project template if requested
	if session.ProjectTemplate != "" {
		styles.HeaderStyle.Println("\n📁 Project Template Setup")
		for _, template := range templates {
			if template.Name == session.ProjectTemplate {
				createProjectFromTemplate(template)
				break
			}
		}
	}

	// Final summary
	session.EndTime = time.Now()
	showDevInitSummary(session)
}

// showDevInitSummary displays a comprehensive summary of the setup
func showDevInitSummary(session *DevInitSession) {
	duration := session.EndTime.Sub(session.StartTime)

	styles.HeaderStyle.Println("\n📊 Development Environment Setup Summary")
	styles.InfoStyle.Println("────────────────────────────────────────")
	styles.InfoStyle.Printf("⏱️  Total time: %s\n", formatDuration(duration))
	styles.SuccessStyle.Printf("✅ Successfully installed: %d tools\n", session.SuccessCount)
	styles.InfoStyle.Printf("⏩ Skipped: %d tools\n", session.SkippedCount)
	styles.ErrorStyle.Printf("❌ Failed: %d tools\n", session.FailedCount)

	if len(session.InstalledTools) > 0 {
		styles.SuccessStyle.Println("\n🎉 Successfully installed tools:")
		for _, tool := range session.InstalledTools {
			styles.SuccessStyle.Printf("  • %s\n", tool)
		}
	}

	if len(session.FailedTools) > 0 {
		styles.ErrorStyle.Println("\n❌ Failed installations:")
		for _, tool := range session.FailedTools {
			styles.ErrorStyle.Printf("  • %s\n", tool)
		}
	}

	// Next steps
	styles.HeaderStyle.Println("\n🚀 Next Steps:")
	styles.InfoStyle.Println("1. Restart your terminal to activate aliases")
	styles.InfoStyle.Println("2. Configure your IDE/editor preferences")
	styles.InfoStyle.Println("3. Set up your preferred Git hosting service")
	styles.InfoStyle.Println("4. Install language-specific extensions")

	if session.SetupSSH {
		styles.InfoStyle.Println("5. Add your SSH key to GitHub/GitLab")
	}

	// Tips
	styles.HeaderStyle.Println("\n💡 Pro Tips:")
	styles.InfoStyle.Println("• Use 'ellie focus' for productive coding sessions")
	styles.InfoStyle.Println("• Try 'ellie git commit' for conventional commits")
	styles.InfoStyle.Println("• Use 'ellie project add' to track your projects")
	styles.InfoStyle.Println("• Run 'ellie start-day' to begin your dev day")

	styles.SuccessStyle.Println("\n🎉 Your development environment is ready!")
}
