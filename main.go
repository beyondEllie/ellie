package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	actions "github.com/tacheraSasi/ellie/action"
)

const (
	VERSION    = "0.0.3"
	configPath = ".env"
)

type Command struct {
	MinArgs     int
	Usage       string
	Handler     func([]string)
	SubCommands map[string]Command
	PreHook     func()
}

var (
	commandRegistry = map[string]Command{
		"run": {
			Handler: actions.Run,
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
			PreHook: func() { fmt.Println("Using the mpv player") },
			Handler: actions.Play,
		},
		"setup-git": {
			Handler: func(_ []string) {
				actions.GitSetup(getEnv("PAT"), getEnv("USERNAME"))
			},
		},
		"sysinfo": {
			Handler: func(_ []string) { actions.SysInfo() },
		},
		"install": {
			MinArgs: 1,
			Usage:   "install <package>",
			Handler: func(a []string) { actions.InstallPackage(a[1]) },
		},
		"update": {
			Handler: func(_ []string) { actions.UpdatePackages() },
		},
		"list": {
			MinArgs: 1,
			Usage:   "list <directory>",
			Handler: func(a []string) { actions.ListFiles(a[1]) },
		},
		"create-file": {
			MinArgs: 1,
			Usage:   "create-file <path>",
			Handler: func(a []string) { actions.CreateFile(a[1]) },
		},
		"network-status": {
			Handler: func(_ []string) { actions.NetworkStatus() },
		},
		"connect-wifi": {
			MinArgs: 2,
			Usage:   "connect-wifi <SSID> <password>",
			Handler: func(a []string) { actions.ConnectWiFi(a[1], a[2]) },
		},
		"greet": {
			Handler: func(_ []string) { greetUser() },
		},
		"git": {
			SubCommands: map[string]Command{
				"status": {Handler: func(_ []string) { actions.GitStatus() }},
				"push":   {Handler: func(_ []string) { actions.GitPush() }},
				"commit": {Handler: func(_ []string) { actions.GitCommitCmd() }},
				"pull":   {Handler: func(_ []string) { actions.GitPull() }},
			},
		},
		"start":     createServiceCommand("start"),
		"stop":      createServiceCommand("stop"),
		"restart":   createServiceCommand("restart"),
		"--help":    {Handler: showHelp},
		"--version": {Handler: func(_ []string) { fmt.Println("Ellie CLI Version:", VERSION) }},
	}
)

func main() {
	if err := godotenv.Load(configPath); err != nil {
		fmt.Println("Warning: .env file could not be loaded.")
	}

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
			"apache": {Handler: func(_ []string) { handleService(action, "apache") }},
			"mysql":  {Handler: func(_ []string) { handleService(action, "mysql") }},
			"all":    {Handler: func(_ []string) { handleService(action, "all") }},
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

func greetUser() {
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func showHelp(_ []string) {
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
`)
}
