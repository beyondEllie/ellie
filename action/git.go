package actions

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/tacheraSasi/ellie/utils"
)

func GitPush() {
	commitMsg, err := utils.GetInput("Enter commit message: ")
	if err != nil {
		fmt.Println("Error reading commit message:", err)
		return
	}

	// 'git add .'
	cmd := exec.Command("git", "add", ".")
	output, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error running git add: %s\n", cmdErr)
		fmt.Printf("Output: %s\n", output)
		return
	}

	// 'git commit -m <commitMsg>'
	cmd = exec.Command("git", "commit", "-m", "Ellie: "+commitMsg)
	output, cmdErr = cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error running git commit: %s\n", cmdErr)
		fmt.Printf("Output: %s\n", output)
		return
	}

	//'git push'
	cmd = exec.Command("git", "push")
	output, cmdErr = cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error running git push: %s\n", cmdErr)
		fmt.Printf("Output: %s\n", output)
		return
	}

	fmt.Printf("Output: %s\n", output)
}

func GitPull(){	
	cmd := exec.Command("git","pull",".")
	output,err := cmd.CombinedOutput()
	if err != nil{
		log.Printf("Error: %s\n",err)
		return
	}
	fmt.Printf("OUTPUT: %s",output)
}

func GitStatus() {
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

func GitAdd(file string) {
	cmd := exec.Command("git", "add", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if output != nil || len(output) == 0 {
		fmt.Printf("%s", output)
	}
}

func GitCommit(commitMsg string) {
	// 'git commit -m <commitMsg>'
	cmd := exec.Command("git", "commit", "-m", "Ellie: "+commitMsg)
	output, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error running git commit: %s\n", cmdErr)
		fmt.Printf("Output: %s\n", output)
		return
	}
	if output != nil || len(output) == 0 {
		fmt.Printf("%s", output)
	}
}

func GitCommitCmd() {
	GitAdd(".")
	commitMsg, err := utils.GetInput("Enter the commit message: ")
	if err != nil {
		fmt.Printf("Error running git commit: %s\n", err)
	}
	GitCommit(commitMsg)
}
