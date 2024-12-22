package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		switch args[1] {
		case "apache":
			fmt.Println("starting apache")
			break
		case "mysql":
			fmt.Println("starting mysql")
			break

		}
	}else{
		fmt.Println("Hello Tach, what can i do for you today?")
	}
}
