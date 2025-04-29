package actions

import "github.com/tacheraSasi/ellie/styles"

type DevTool struct {
	Name     string
	CheckCmd string   // e.g. "git --version"
	Install  []string // OS specific install commands
}


func DevInit(){
	styles.InfoStyle.Println("Initializing development environment...")
	styles.InfoStyle.Println("Please wait while we set up your environment.")
	styles.InfoStyle.Println("This may take a few minutes.")
}