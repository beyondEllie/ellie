package main

import (
	"fmt"
	"os"
	"yourmodule/action" // Replace with the correct module name
)

func main() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "start":
			start(args)
		case "stop":
			stop(args)
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
