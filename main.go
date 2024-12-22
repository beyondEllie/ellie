package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "start":
			start(args)
			break
		case "stop":
			fmt.Println("starting mysql")
			break
		default:
			fmt.Println("Hello Tach, what can i do for you today?")
		}
	} else {
		fmt.Println("Hello Tach, what can i do for you today?")
	}

}

func start(args []string) {
	if len(args) > 2 {
		switch args[1] {
		case "apache":
			fmt.Println("starting apache")
			break
		case "mysql":
			fmt.Println("starting mysql")
			break
		case "all":
			fmt.Println("starting mysql & apache")
			break
		default:
			fmt.Println("Hello Tach, what can i do for you today?")
		}
	} else {
		fmt.Println("Hello Tach, what can i do for you today?")
	}
}

func stop() {

}
