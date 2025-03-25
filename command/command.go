package command

type Command struct {
	MinArgs     int
	Usage       string
	Handler     func([]string)
	SubCommands map[string]Command
	PreHook     func()
}
