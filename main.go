package main

import (
	"fmt"
	"os"

	"github.com/tacheraSasi/ellie/action"
	"github/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err !=nil{
		fmt.Println("Error loading .env file",err)
		return
	}
	pat := os.Getenv("PAT")
	username := os.Getenv("USERNAME")
	
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "pwd":
			action.Pwd()
		case "start":
			start(args)
		case "stop": 
			stop(args)
		case "restart":
			restart(args)
		case "setup-git":
			action.GitSetup(pat,username)
		default:
			fmt.Println("Hello Tach, what can I do for you today?")
		}
	} else {
		fmt.Println("Hello Tach, what can I do for you today?")
	}
}

func start(args []string) {
	if len(args) > 2 {
		switch args[2] {
		case "apache":
			action.StartApache()
		case "mysql":
			action.StartMysql()
		case "all":
			action.StartAll()
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
			action.StopApache()
		case "mysql":
			action.StopMysql()
		case "all":
			action.StopAll()
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
            action.StopApache()
            action.StartApache()
        case "mysql":
            action.StopMysql()
            action.StartMysql()
        case "all":
            action.StopAll()
            action.StartAll()
        default:
            fmt.Println("Service not recognized. Please choose 'apache', 'mysql', or 'all'.")
        }
    } else {
        fmt.Println("Please specify a service to restart: 'apache', 'mysql', or 'all'.")
    }
}
