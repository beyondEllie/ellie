package types

// This struct represents a development tool with its properties
// and installation commands for different operating systems.
type DevTool struct {
	Name           string
	Description    string
	CheckCmd       string
	Install        map[string]string // OS: command
	Configure      map[string]string // OS: post-install configuration
	DefaultInstall bool
}