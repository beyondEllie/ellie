package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

type DevTool struct {
	Name          string
	Description   string
	CheckCmd      string
	Install       map[string]string // OS: command
	DefaultInstall bool
}

var Tools []DevTool = []DevTool{
	{
		Name:          "Git",
		Description:   "Version control system",
		CheckCmd:      "git --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":    "brew install git",
			"linux":  "sudo apt-get install git -y",
			"windows": "choco install git -y",
		},
	},
	{
		Name:          "Node.js",
		Description:   "JavaScript runtime",
		CheckCmd:      "node --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":    "brew install node",
			"linux":  "curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash - && sudo apt-get install -y nodejs",
			"windows": "choco install nodejs-lts",
		},
	},
	{
		Name:          "Docker",
		Description:   "Containerization platform",
		CheckCmd:      "docker --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install --cask docker",
			"linux":  "curl -fsSL https://get.docker.com | sh",
			"windows": "choco install docker-desktop",
		},
	},
	{
		Name:          "Go",
		Description:   "Go programming language",
		CheckCmd:      "go version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install go",
			"linux":  "sudo apt install golang -y",
			"windows": "choco install golang",
		},
	},
	{
		Name:          "Python",
		Description:   "Python programming language",
		CheckCmd:      "python3 --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":    "brew install python",
			"linux":  "sudo apt install python3 -y",
			"windows": "choco install python",
		},
	},
	{
		Name:          "VS Code",
		Description:   "Code editor",
		CheckCmd:      "code --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install --cask visual-studio-code",
			"linux":  "sudo snap install --classic code",
			"windows": "choco install vscode",
		},
	},
}

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

func runInstallCommand(tool DevTool, currentOS string) bool {
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

	// Verify installation
	if !isInstalled(tool.CheckCmd) {
		styles.ErrorStyle.Printf("❌ Installation verification failed for %s\n", tool.Name)
		return false
	}

	styles.SuccessStyle.Printf("✅ Successfully installed %s\n", tool.Name)
	return true
}


func DevInit(installAll bool) {
	currentOS := utils.GetOS()
	if currentOS == "unknown" {
		styles.ErrorStyle.Println("❌ Unsupported operating system")
		return
	}



	styles.HeaderStyle.Println("🚀 Development Environment Setup")
	styles.InfoStyle.Printf("Detected OS: %s\n\n", strings.ToUpper(currentOS))

	successCount := 0
	skippedCount := 0
	failedCount := 0

	for _, tool := range Tools {
		styles.Highlight.Printf("\n%s - %s\n", tool.Name, tool.Description)
		
		if isInstalled(tool.CheckCmd) {
			styles.SuccessStyle.Printf("✅ %s is already installed\n", tool.Name)
			skippedCount++
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
			skippedCount++
			continue
		}

		_, exists := tool.Install[currentOS]
		if !exists {
			styles.ErrorStyle.Printf("❌ No installation command for %s on %s\n", tool.Name, currentOS)
			failedCount++
			continue
		}

		if runInstallCommand(tool, utils.GetOS()) {
			successCount++
		} else {
			failedCount++
		}
	}

	styles.HeaderStyle.Println("\n📊 Installation Summary:")
	styles.SuccessStyle.Printf("✅ Success: %d\n", successCount)
	styles.InfoStyle.Printf("⏩ Skipped: %d\n", skippedCount)
	styles.ErrorStyle.Printf("❌ Failed: %d\n", failedCount)

	if failedCount > 0 {
		styles.WarningStyle.Println("\nℹ️  Some installations failed. You may need to:")
		styles.InfoStyle.Println("  - Check internet connection")
		styles.InfoStyle.Println("  - Verify package manager is installed")
		styles.InfoStyle.Println("  - Run with administrator privileges")
	}
}