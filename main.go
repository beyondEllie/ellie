package main

import (
	"fmt"
	"os"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
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
			MinArgs: 2,
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
			MinArgs: 2,
			Usage:   "install <package>",
			Handler: func(a []string) { actions.InstallPackage(a[2]) },
		},
		"update": {
			Handler: func(_ []string) { actions.UpdatePackages() },
		},
		"list": {
			MinArgs: 2,
			Usage:   "list <directory>",
			Handler: func(a []string) { actions.ListFiles(a[2]) },
		},
		"create-file": {
			MinArgs: 2,
			Usage:   "create-file <path>",
			Handler: func(a []string) { actions.CreateFile(a[2]) },
		},
		"network-status": {
			Handler: func(_ []string) { actions.NetworkStatus() },
		},
		"connect-wifi": {
			MinArgs: 3,
			Usage:   "connect-wifi <SSID> <password>",
			Handler: func(a []string) { actions.ConnectWiFi(a[2], a[3]) },
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
		"start":    createServiceCommand("start"),
		"stop":     createServiceCommand("stop"),
		"restart":  createServiceCommand("restart"),
		"--help":   {Handler: showHelp},
		"--version": {Handler: func(_ []string) { fmt.Println("Ellie CLI Version:", VERSION) }},
	}
)

func main() {
	_ = godotenv.Load(configPath) 

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
		actions.Chat(getEnv("OPENAI_API_KEY"))
		return
	}

	if cmd.PreHook != nil {
		cmd.PreHook()
	}

	if len(cmd.SubCommands) > 0 && len(args) > 1 {
		handleSubCommand(cmd, args[1:])
		return
	}

	if len(args) <= cmd.MinArgs {
		fmt.Printf("Invalid usage: %s\n%s\n", cmdName, cmd.Usage)
		return
	}

	cmd.Handler(args)
}

func handleSubCommand(parentCmd Command, args []string) {
	subCmdName := args[0]
	subCmd, exists := parentCmd.SubCommands[subCmdName]
	if !exists {
		fmt.Printf("Unknown subcommand: %s\n", subCmdName)
		return
	}

	if len(args) <= subCmd.MinArgs {
		fmt.Printf("Invalid usage: %s\n%s\n", subCmdName, subCmd.Usage)
		return
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
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
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
	helpText := `
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

Service Management:
  start    <service>    Start system services
  stop     <service>    Stop system services
  restart  <service>    Restart services

Git Operations:
  git status            Check repository status
  git push              Push changes
  git commit            Commit changes
  git pull              Fetch and merge

Miscellaneous:
  --help                Show this help message
  --version             Display version information
  greet                 Get time-based greeting

Examples:
  ellie start apache
  ellie git status
  ellie play video.mp4
`
	fmt.Println(helpText)
}

func renderMd(content string) {
	result := markdown.Render(content, 80, 6)
	fmt.Println(string(result))
}