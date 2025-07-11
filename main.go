package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	actions "github.com/tacheraSasi/ellie/action"
	"github.com/tacheraSasi/ellie/command"
	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

const (
	VERSION = configs.VERSION
)

var ICON any = static.Icon()

// User name from the saved files during initialization
var CurrentUser string = configs.GetEnv("USERNAME")


func main() {
	// Setup global flags
	showHelp := flag.Bool("help", false, "Show help information")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	// Handle global flags
	if *showVersion {
		styles.GetSuccessStyle().Printf("Ellie CLI v%s\n", VERSION)
		return
	}
	if *showHelp {
		showHelpFunc()
		return
	}

	// Initialize configuration
	configs.Init()

	// Get remaining arguments after flags
	args := flag.Args()

	// Interactive mode if no commands
	if len(args) == 0 {
		actions.Chat(configs.GetEnv("OPENAI_API_KEY"))
		return
	}

	handleCommand(args)
}

func handleCommand(args []string) {
	if len(args) == 0 {
		actions.Chat(configs.GetEnv("OPENAI_API_KEY"))
		return
	}

	cmdName := args[0]

	// Check if the command is an alias
	if actions.ExecuteAlias(cmdName) {
		return
	}

	cmd, exists := command.Registry[cmdName]
	if !exists {
		matches := getClosestMatchingCmd(command.Registry, cmdName)
		// fmt.Println(matches)
		if len(matches) > 0 {
			styles.GetErrorStyle().Printf("Unknown command: %s\n", cmdName)
			styles.GetInfoStyle().Println("Did you mean:")
			for _, m := range matches {
				styles.GetInfoStyle().Printf("  %s\n", m)
			}
		} else {
			styles.GetErrorStyle().Printf("Unknown command: %s\n", cmdName)
			showHelpFunc()
		}
		os.Exit(1)
	}

	if cmd.PreHook != nil {
		cmd.PreHook()
	}

	if len(cmd.SubCommands) > 0 && len(args) > 1 {
		handleSubCommand(cmd, args[1:])
		return
	}

	if len(args)-1 < cmd.MinArgs {
		styles.GetErrorStyle().Printf("Invalid usage for %s\n", cmdName)
		styles.GetInfoStyle().Println("Usage:", cmd.Usage)
		os.Exit(1)
	}

	cmd.Handler(args)
}

// Returns a list of command names that closely match the input
func getClosestMatchingCmd(cmdMap map[string]command.Command, cmdArg string) []string {
	var list []string
	for cmd := range cmdMap {
		distance := levenshtein.DistanceForStrings([]rune(cmdArg), []rune(cmd), levenshtein.DefaultOptions)
		maxLen := len(cmdArg)
		if len(cmd) > maxLen {
			maxLen = len(cmd)
		}
		similarity := 1.0 - (float64(distance) / float64(maxLen))
		if similarity > 0.4 {
			list = append(list, cmd)
		}
	}
	return list
}

func handleSubCommand(parentCmd command.Command, args []string) {
	subCmdName := args[0]
	subCmd, exists := parentCmd.SubCommands[subCmdName]
	if !exists {
		styles.GetErrorStyle().Println("Unknown subcommand:", subCmdName)
		os.Exit(1)
	}

	if len(args)-1 < subCmd.MinArgs {
		styles.GetErrorStyle().Printf("Invalid usage for %s\n", subCmdName)
		styles.GetInfoStyle().Println("Usage:", subCmd.Usage)
		os.Exit(1)
	}

	if subCmd.PreHook != nil {
		subCmd.PreHook()
	}

	subCmd.Handler(args)
}

func createServiceCommand(action string) command.Command {
	return command.Command{
		SubCommands: map[string]command.Command{
			"apache":   {Handler: func(args []string) { actions.HandleService(action, "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService(action, "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService(action, "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService(action, "all") }},
		},
	}
}

func greetUser(args []string) {
	hour := time.Now().Hour()
	greeting := styles.GetSuccessStyle().Println
	message := "Good evening!"

	switch {
	case hour < 12:
		message = "Good morning!"
		greeting = styles.GetHighlightStyle().Println
	case hour < 18:
		message = "Good afternoon!"
		greeting = styles.GetInfoStyle().Println
	}

	greeting(message+",", CurrentUser)
}

func showHelpFunc() {
	styles.GetHeaderStyle().Println("Ellie CLI - AI-Powered System Management Tool")
	styles.GetInfoStyle().Println("Usage: ellie [--help] [--version] <command> [arguments]")

	styles.GetHeaderStyle().Println("Global Flags:")
	styles.GetInfoStyle().Println("  --help\tShow this help message")
	styles.GetInfoStyle().Println("  --version\tShow version information")

	styles.GetHeaderStyle().Println("Core Commands:")
	fmt.Println("  config \t\tConfigure Ellie CLI")
	fmt.Println("  reset-config\t\tReset Ellie CLI configuration")
	fmt.Println("  whoami\t\tShow current user")
	fmt.Println("  dev-init [--all]\tInitialize development environment (optionally install all tools)")
	fmt.Println("  greet\t\t\tGreet the user")
	fmt.Println("  send-mail\t\tSend an email")
	fmt.Println("  focus\t\t\tActivate focus mode")
	fmt.Println("  review <filename/path>\tReview code or files using LLMs")
	fmt.Println("  run <command>\t\tExecute system commands")
	fmt.Println("  pwd\t\t\tPrint working directory")
	fmt.Println("  open-explorer\t\tOpen the current dir in file explorer")
	fmt.Println("  open <path>\t\tOpen file explorer at path")
	fmt.Println("  play <media>\t\tPlay media files")
	fmt.Println("  setup-git\t\tConfigure Git credentials")
	fmt.Println("  sysinfo\t\tShow system information")
	fmt.Println("  install <package>\tInstall software packages")
	fmt.Println("  update\t\tUpdate system packages")
	fmt.Println("  list <dir>\t\tList directory contents")
	fmt.Println("  create-file <path>\tCreate new files")
	fmt.Println("  network-status\tShow network information")
	fmt.Println("  connect-wifi <SSID> <password>\tConnect to WiFi network")
	fmt.Println("  chat\t\t\tStart interactive AI chat session")
	fmt.Println("  history\t\tShow recent commands")
	fmt.Println("  start-day\t\tStart your dev day (apps, services, git)")
	fmt.Println("  weather\t\tShow weather info")
	fmt.Println("  joke\t\t\tTell a joke")
	fmt.Println("  remind\t\t\tSet a reminder")
	fmt.Println("  about\t\t\tShow about info")

	styles.GetHeaderStyle().Println("Git Commands:")
	fmt.Println("  git status\t\tShow Git status")
	fmt.Println("  git push\t\tPush commits")
	fmt.Println("  git commit\t\tCreate Conventional Commit")
	fmt.Println("  git pull\t\tPull latest changes")

	styles.GetHeaderStyle().Println("Service Management:")
	fmt.Println("  start <apache|mysql|postgres|all>\tStart services")
	fmt.Println("  stop <apache|mysql|postgres|all>\tStop services")
	fmt.Println("  restart <apache|mysql|postgres|all>\tRestart services")
	fmt.Println("  start list\t\tList available services")
	fmt.Println("  stop list\t\tList available services")
	fmt.Println("  restart list\t\tList available services")

	styles.GetHeaderStyle().Println("Alias Management:")
	fmt.Println("  alias add <name>=\"<command>\"\tAdd a new alias")
	fmt.Println("  alias list\t\tList all aliases")
	fmt.Println("  alias delete <name>\tDelete an alias")

	styles.GetHeaderStyle().Println("Todo Management:")
	fmt.Println("  todo add \"<task>\" [category] [priority]\tAdd a new todo")
	fmt.Println("  todo list\t\tList all todos")
	fmt.Println("  todo complete <id>\tMark todo as complete")
	fmt.Println("  todo delete <id>\tDelete a todo")
	fmt.Println("  todo edit <id> <field> <value>\tEdit a todo field")

	styles.GetHeaderStyle().Println("Project Management:")
	fmt.Println("  project add <name> <path> [description] [tags...]\tAdd a new project")
	fmt.Println("  project list\t\tList all projects")
	fmt.Println("  project delete <name>\tDelete a project")
	fmt.Println("  project search <query>\tSearch projects")
	fmt.Println("  switch <project-name>\tSwitch to a project")

	styles.GetHeaderStyle().Println("Day Start Configuration:")
	fmt.Println("  day-start add <type> <value>\tAdd to daily setup (apps/services/git_repos)")
	fmt.Println("  day-start list\t\tList daily setup configuration")

	styles.GetHeaderStyle().Println("Theme Management:")
	fmt.Println("  theme set <light|dark|auto>\tSet the theme mode")
	fmt.Println("  theme show\t\tShow current theme")

	styles.DimText.Println("For detailed command help, use 'ellie <command> --help'")
}
