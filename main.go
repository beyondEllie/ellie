package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	actions "github.com/tacheraSasi/ellie/action"
)

const (
	VERSION    = "0.0.5"
	configPath = "/home/tach/tach/go/ellie/.env" //TODO: Will make this configurable.
)

// Command holds CLI command details.
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
		PreHook: func() { fmt.Println("Using focus") },
		Handler: actions.Focus,
	},
	"pwd": {
		Handler: func(_ []string) {actions.Pwd()},
	},
	"open-explorer": {
		Handler: func(_ []string) {actions.OpenExplorer()},
	},
	"play": {
		MinArgs: 1,
		Usage:   "play <media>",
		PreHook: func() { fmt.Println("Using the mpv player") },
		Handler: actions.Play,
	},
	"setup-git": {
		Handler: func(args []string) {
			actions.GitSetup(getEnv("PAT"), getEnv("USERNAME"))
		},
	},
	"sysinfo": {
		Handler: func(_ []string) {
			actions.SysInfo()
		},
	},
	"install": {
		MinArgs: 1,
		Usage:   "install <package>",
		Handler: func(args []string) { actions.InstallPackage(args[1]) },
	},
	"update": {
		Handler: func(_ []string) {
			actions.UpdatePackages()
		},
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
		Handler: func(_ []string) {actions.NetworkStatus()},
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
			"status": {Handler: func(_ []string) {actions.GitStatus()}},
			"push":   {Handler: func(_ []string) {actions.GitPush()}},
			// Using production-ready Conventional Commit implementation.
			"commit": {Handler: func(args []string) { actions.GitConventionalCommit() }},
			"pull":   {Handler: func(_ []string) {actions.GitPull()}},
		},
	},
	"start":   createServiceCommand("start"),
	"stop":    createServiceCommand("stop"),
	"restart": createServiceCommand("restart"),
	"--help":  {Handler: showHelp},
	"--version": {
		Handler: func(args []string) {
			fmt.Println("Ellie CLI Version:", VERSION)
		},
	},
	"--v": {
		Handler: func(args []string) {
			fmt.Println("Ellie CLI Version:", VERSION)
		},
	},
}

func main() {
	// Load environment variables.
	if err := godotenv.Load(configPath); err != nil {
		fmt.Println("Warning: .env file could not be loaded.")
	}

	// If no command is provided, fallback to a default interactive mode.
	if len(os.Args) < 2 {
		actions.Chat(getEnv("OPENAI_API_KEY"))
		return
	}

	handleCommand(os.Args[1:])
}

func handleCommand(args []string) {
	if len(args) == 0 {
		actions.Chat(getEnv("OPENAI_API_KEY"))
		return
	}

	cmdName := args[0]
	cmd, exists := commandRegistry[cmdName]
	if !exists {
		fmt.Println("Unknown command:", cmdName)
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
		fmt.Printf("Invalid usage: %s\n%s\n", cmdName, cmd.Usage)
		os.Exit(1)
	}

	cmd.Handler(args)
}

func handleSubCommand(parentCmd Command, args []string) {
	subCmdName := args[0]
	subCmd, exists := parentCmd.SubCommands[subCmdName]
	if !exists {
		fmt.Println("Unknown subcommand:", subCmdName)
		os.Exit(1)
	}

	if len(args)-1 < subCmd.MinArgs {
		fmt.Printf("Invalid usage: %s\n%s\n", subCmdName, subCmd.Usage)
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

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}

func greetUser(args []string) {
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 18 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}
}

func showHelp(args []string) {
	fmt.Println(`
Ellie CLI - AI-Powered System Management Tool

Usage: ellie <command> [arguments]

Core Commands:
  run <command>         Execute system commands
  pwd                   Print working directory
  open-explorer         Open file explorer
  play <media>          Play media files
  setup-git             Configure Git credentials
  sysinfo               Show system information
  install <package>     Install software packages
  update                Update system packages
  list <dir>            List directory contents
  create-file <path>    Create new files
  network-status        Show network information
  connect-wifi <creds>  Manage WiFi connections
  git status            Show Git status
  git push              Push commits
  git commit            Create a Conventional Commit and push
  git pull              Pull latest changes
  start, stop, restart  Manage services (apache, mysql, all)
`)
}
