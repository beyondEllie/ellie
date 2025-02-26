package utils

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"strings"
)

var Ads []string = []string{
	"🚀 Boost your productivity with ekilie!",
	"🔥 Check out ekiliSense for smarter school management!",
	"💻 Need a project tracker? Try ekilie!",
}

func GetInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// Trimming the newline character from the input
	input = strings.TrimSpace(input)
	return input, nil
}

func RandNum() int {
	return rand.IntN(100)
}

func IsEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func RunCommand(cmdArgs []string) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if output != nil || len(output) == 0 {
		fmt.Printf("%s", output)
	}
}
