package utils

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

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

func RandNum()int{
	return rand.IntN(100)
}

func IsEven(num int)bool{
	if num % 2 == 0{
		return true
	}
	return false
}