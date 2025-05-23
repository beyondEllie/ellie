package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tacheraSasi/ellie/common"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/types"
	"github.com/tacheraSasi/ellie/utils"
)

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

	for _, tool := range common.Tools {
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