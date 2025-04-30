package actions

import (
	"fmt"
	"os"

	"github.com/tacheraSasi/ellie/utils"
)

func ServerInit() {
	// if utils.IsServerInitialized() {
	// 	fmt.Println("Server is already initialized.")
	// 	return
	// }

	serverName, err := utils.GetInput("Enter the server name: ")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	err = os.Mkdir(serverName, 0755)
	if err != nil {
		fmt.Println("Error creating server directory:", err)
		return
	}

	err = os.WriteFile(serverName+"/config.json", []byte("{}"), 0644)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		return
	}

	fmt.Printf("Server '%s' initialized successfully.\n", serverName)
}