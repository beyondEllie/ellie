package command

// Command holds CLI command details
type Command struct {
	MinArgs     int
	Usage       string
	Handler     func([]string)
	SubCommands map[string]Command
	PreHook     func()
}
