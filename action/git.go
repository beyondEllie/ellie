package actions

import (
	"fmt"
	"os/exec"

	"github.com/tacheraSasi/ellie/utils"
)

func GitPush() {
	commitMsg, err := utils.GetInput("Enter commit message: ")
	if err != nil {
		fmt.Println("Error reading commit message:", err)
		return
	}

	// Run 'git add .'
	cmd := exec.Command("git", "add", ".")
	output, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error running git add: %s\n", cmdErr)
		fmt.Printf("Output: %s\n", output)
		return
	}

	// Run 'git commit -m <commitMsg>'
	cmd = exec.Command("git", "commit", "-m", commitMsg)
	output, cmdErr = cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error running git commit: %s\n", cmdErr)
		fmt.Printf("Output: %s\n", output)
		return
	}

	// Run 'git push'
	cmd = exec.Command("git", "push")
	output, cmdErr = cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error running git push: %s\n", cmdErr)
		fmt.Printf("Output: %s\n", output)
		return
	}

	fmt.Printf("Output: %s\n", output)        
}

func GitStatus(){
	cmd := exec.Command("git", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if output != nil || len(output) == 0 {
		fmt.Printf("%s", output)
	}
}