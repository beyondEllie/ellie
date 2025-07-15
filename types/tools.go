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

// TODO:Will change this to more accurate struct with frameworks and stuff...
type ServerTool struct{
	Name        string
	Description string
	CheckCmd    string
	Install     map[string]string // OS: command
	Configure   map[string]string // OS: post-install configuration
	DefaultInstall bool
}