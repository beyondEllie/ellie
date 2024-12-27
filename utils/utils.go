package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string) (string, error) {
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