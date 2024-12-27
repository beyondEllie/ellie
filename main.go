package main

import (
	"fmt"
	"os"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/joho/godotenv"
	actions "github.com/tacheraSasi/ellie/action"
)

func main() {
	err := godotenv.Load("/home/tach/tach/go/ellie/.env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return
	}

	pat := os.Getenv("PAT")
	username := os.Getenv("USERNAME")
	openaiApikey := os.Getenv("OPENAI_API_KEY")

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "run":
			actions.Run(args)
		case "pwd":
			actions.Pwd()
		case "start":
			start(args)
		case "stop":
			stop(args)
		case "restart":
			restart(args)
		case "setup-git":
			actions.GitSetup(pat, username)
		case "git":
			git(args)
		case "sysinfo":
			actions.SysInfo()
		case "install":
			if len(args) > 2 {
				actions.InstallPackage(args[2])
			} else {
				fmt.Println("Please specify a package to install.")
			}
		case "update":
			actions.UpdatePackages()
		case "list":
			if len(args) > 2 {
				actions.ListFiles(args[2])
			} else {
				fmt.Println("Please specify a directory to list files.")
			}
		case "create-file":
			if len(args) > 2 {
				actions.CreateFile(args[2])
			} else {
				fmt.Println("Please specify a file path to create.")
			}
		case "network-status":
			actions.NetworkStatus()
		case "connect-wifi":
			if len(args) > 3 {
				actions.ConnectWiFi(args[2], args[3])
			} else {
				fmt.Println("Please provide SSID and password for Wi-Fi.")
			}
		case "greet":
			greetUser()
		default:
			actions.Chat(openaiApikey)
		}
	} else {
		actions.Chat(openaiApikey)
	}
}


func git(args []string) {
	if len(args) > 2 {
		switch args[2] {
		case "status":
			actions.GitStatus()
		case "push":
			actions.GitPush()
		// case "pull":
		// 	actions.GitPull()
		default:
			fmt.Println("Command not recognized. Please use 'status', 'push', or 'pull'.")
		}
	} else {
		fmt.Println("Please specify a git command: 'status','push', or 'pull'.")
	}
}

func start(args []string) {
	if len(args) > 2 {
		switch args[2] {
		case "apache":
			actions.StartApache()
		case "mysql":
			actions.StartMysql()
		case "all":
			actions.StartAll()
		default:
			fmt.Println("Service not recognized. Please choose 'apache', 'mysql', or 'all'.")
		}
	} else {
		fmt.Println("Please specify a service to start: 'apache', 'mysql', or 'all'.")
	}
}

func stop(args []string) {
	if len(args) > 2 {
		switch args[2] {
		case "apache":
			actions.StopApache()
		case "mysql":
			actions.StopMysql()
		case "all":
			actions.StopAll()
		default:
			fmt.Println("Service not recognized. Please choose 'apache', 'mysql', or 'all'.")
		}
	} else {
		fmt.Println("Please specify a service to stop: 'apache', 'mysql', or 'all'.")
	}
}

func restart(args []string) {
	if len(args) > 2 {
		switch args[2] {
		case "apache":
			actions.StopApache()
			actions.StartApache()
		case "mysql":
			actions.StopMysql()
			actions.StartMysql()
		case "all":
			actions.StopAll()
			actions.StartAll()
		default:
			fmt.Println("Service not recognized. Please choose 'apache', 'mysql', or 'all'.")
		}
	} else {
		fmt.Println("Please specify a service to restart: 'apache', 'mysql', or 'all'.")
	}
}

func greetUser() {
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 18 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}
}

func renderMd(content string) {
	fmt.Println(content)
	result := markdown.Render(string(content), 80, 6)

	fmt.Println(result)
}
