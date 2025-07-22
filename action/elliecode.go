package actions

import (
	"fmt"
	"github.com/tacheraSasi/ellie/utils"
)

func StartEllieCode(){
	fmt.Println("Starting ellie code")
	utils.RunCommand([]string{"cd elliecode && go run main.go"},"failed to run")
}