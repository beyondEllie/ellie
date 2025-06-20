package command

type Command struct {
	MinArgs int    // Minimum number of arguments required
	Usage   string // Usage string for help output

	Handler     func([]string)     // Function to execute the command
	SubCommands map[string]Command // Subcommands for this command (if any)
	PreHook     func()             // Function to run before the command handler (optional)
}
