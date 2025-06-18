package actions

import (
	"fmt"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

// ShowAbout displays the about information for Ellie
func ShowAbout(args []string) {
	styles.HeaderStyle.Println("Ellie - The AI-Powered CLI Companion")
	styles.InfoStyle.Println("Version:", configs.VERSION)
	fmt.Println()

	styles.Highlight.Println("Description:")
	styles.InfoStyle.Println("Your all-in-one terminal buddy for system management, Git workflows, and productivity hacks.")
	fmt.Println()

	styles.Highlight.Println("Core Features:")
	styles.InfoStyle.Println("• System Management")
	styles.InfoStyle.Println("• Git Workflows")
	styles.InfoStyle.Println("• Todo Management")
	styles.InfoStyle.Println("• Project Management")
	styles.InfoStyle.Println("• Network Management")
	styles.InfoStyle.Println("• AI Integration")
	fmt.Println()

	styles.SuccessStyle.Println("Built with ❤️ by Tachera Sasi")
}
