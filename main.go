package main

import (
	"flag"
	"os"
	"time"

	actions "github.com/tacheraSasi/ellie/action"
	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

const (
	VERSION    = "0.0.7"
)

// Command holds CLI command details
type Command struct {
	MinArgs     int
	Usage       string
	Handler     func([]string)
	SubCommands map[string]Command
	PreHook     func()
}

var commandRegistry = map[string]Command{
	"run": {
		Handler: actions.Run,
	},
	"focus": {
		PreHook: func() { styles.InfoStyle.Println("Activating focus mode...") },
		Handler: actions.Focus,
	},
	"pwd": {
		Handler: func(_ []string) { actions.Pwd() },
	},
	"open-explorer": {
		Handler: func(_ []string) { actions.OpenExplorer() },
	},
	"play": {
		MinArgs: 1,
		Usage:   "play <media>",
		PreHook: func() { styles.InfoStyle.Println("Initializing media player...") },
		Handler: actions.Play,
	},
	"setup-git": {
		Handler: func(args []string) {
			actions.GitSetup(configs.GetEnv("PAT"), configs.GetEnv("USERNAME"))
		},
	},
	"sysinfo": {
		Handler: func(_ []string) { actions.SysInfo() },
	},
	"install": {
		MinArgs: 1,
		Usage:   "install <package>",
		Handler: func(args []string) { actions.InstallPackage(args[1]) },
	},
	"update": {
		Handler: func(_ []string) { actions.UpdatePackages() },
	},
	"list": {
		MinArgs: 1,
		Usage:   "list <directory>",
		Handler: func(args []string) { actions.ListFiles(args[1]) },
	},
	"create-file": {
		MinArgs: 1,
		Usage:   "create-file <path>",
		Handler: func(args []string) { actions.CreateFile(args[1]) },
	},
	"network-status": {
		Handler: func(_ []string) { actions.NetworkStatus() },
	},
	"connect-wifi": {
		MinArgs: 2,
		Usage:   "connect-wifi <SSID> <password>",
		Handler: func(args []string) { actions.ConnectWiFi(args[1], args[2]) },
	},
	"greet": {
		Handler: greetUser,
	},
	"git": {
		SubCommands: map[string]Command{
			"status": {Handler: func(_ []string) { actions.GitStatus() }},
			"push":   {Handler: func(_ []string) { actions.GitPush() }},
			"commit": {Handler: func(args []string) { actions.GitConventionalCommit() }},
			"pull":   {Handler: func(_ []string) { actions.GitPull() }},
		},
	},
	"start":   createServiceCommand("start"),
	"stop":    createServiceCommand("stop"),
	"restart": createServiceCommand("restart"),
	"whoami": {
		Handler: func(_ []string) {
			styles.Highlight.Println("Your majesty,", configs.GetEnv("USERNAME"))
		},
	},
}

func main() {
	// Setup global flags
	showHelp := flag.Bool("help", false, "Show help information")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	// Handle global flags
	if *showVersion {
		styles.SuccessStyle.Printf("Ellie CLI v%s\n", VERSION)
		return
	}
	if *showHelp {
		showHelpFunc(nil)
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
	cmd, exists := commandRegistry[cmdName]
	if !exists {
		styles.ErrorStyle.Println("Unknown command:", cmdName)
		showHelpFunc(nil)
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
		styles.ErrorStyle.Printf("Invalid usage for %s\n", cmdName)
		styles.InfoStyle.Println("Usage:", cmd.Usage)
		os.Exit(1)
	}

	cmd.Handler(args)
}

func handleSubCommand(parentCmd Command, args []string) {
	subCmdName := args[0]
	subCmd, exists := parentCmd.SubCommands[subCmdName]
	if !exists {
		styles.ErrorStyle.Println("Unknown subcommand:", subCmdName)
		os.Exit(1)
	}

	if len(args)-1 < subCmd.MinArgs {
		styles.ErrorStyle.Printf("Invalid usage for %s\n", subCmdName)
		styles.InfoStyle.Println("Usage:", subCmd.Usage)
		os.Exit(1)
	}

	if subCmd.PreHook != nil {
		subCmd.PreHook()
	}

	subCmd.Handler(args)
}

func createServiceCommand(action string) Command {
	return Command{
		SubCommands: map[string]Command{
			"apache": {Handler: func(args []string) { handleService(action, "apache") }},
			"mysql":  {Handler: func(args []string) { handleService(action, "mysql") }},
			"all":    {Handler: func(args []string) { handleService(action, "all") }},
		},
	}
}

func handleService(action, service string) {
	actionVerb := action + "ing"
	styles.InfoStyle.Printf("%s %s service...\n", actionVerb, service)
	
	switch action {
	case "start":
		switch service {
		case "apache":
			actions.StartApache()
		case "mysql":
			actions.StartMysql()
		case "all":
			actions.StartAll()
		}
	case "stop":
		switch service {
		case "apache":
			actions.StopApache()
		case "mysql":
			actions.StopMysql()
		case "all":
			actions.StopAll()
		}
	case "restart":
		switch service {
		case "apache":
			actions.StopApache()
			actions.StartApache()
		case "mysql":
			actions.StopMysql()
			actions.StartMysql()
		case "all":
			actions.StopAll()
			actions.StartAll()
		}
	}
}

func greetUser(args []string) {
	hour := time.Now().Hour()
	greeting := styles.SuccessStyle.Println
	message := "Good evening!"

	switch {
	case hour < 12:
		message = "Good morning!"
		greeting = styles.Highlight.Println
	case hour < 18:
		message = "Good afternoon!"
		greeting = styles.InfoStyle.Println
	}

	greeting(message)
}

func showHelpFunc(args []string) {
	styles.HeaderStyle.Println("Ellie CLI - AI-Powered System Management Tool")
	styles.InfoStyle.Println("Usage: ellie [--help] [--version] <command> [arguments]\n")

	styles.HeaderStyle.Println("Global Flags:")
	styles.InfoStyle.Println("  --help\tShow this help message")
	styles.InfoStyle.Println("  --version\tShow version information\n")

	styles.HeaderStyle.Println("Core Commands:")
	styles.InfoStyle.Println("  run <command>\t\tExecute system commands")
	styles.InfoStyle.Println("  pwd\t\t\tPrint working directory")
	styles.InfoStyle.Println("  open-explorer\t\tOpen file explorer")
	styles.InfoStyle.Println("  play <media>\t\tPlay media files")
	styles.InfoStyle.Println("  setup-git\t\tConfigure Git credentials")
	styles.InfoStyle.Println("  sysinfo\t\tShow system information")
	styles.InfoStyle.Println("  install <package>\tInstall software packages")
	styles.InfoStyle.Println("  update\t\tUpdate system packages")
	styles.InfoStyle.Println("  list <dir>\t\tList directory contents")
	styles.InfoStyle.Println("  create-file <path>\tCreate new files")
	styles.InfoStyle.Println("  network-status\tShow network information")
	styles.InfoStyle.Println("  connect-wifi <creds>\tManage WiFi connections")
	styles.InfoStyle.Println("  git status\t\tShow Git status")
	styles.InfoStyle.Println("  git push\t\tPush commits")
	styles.InfoStyle.Println("  git commit\t\tCreate Conventional Commit")
	styles.InfoStyle.Println("  git pull\t\tPull latest changes")
	styles.InfoStyle.Println("  start/stop/restart\tManage services (apache, mysql, all)\n")

	styles.DimText.Println("For detailed command help, use 'ellie <command> --help'")
}