package static

import "fmt"

func Instructions() string {
	prompt,err := GetStaticFile("./prompt.txt")
	if err != nil {
		return "Error loading instructions"
	}
	prompt = fmt.Sprintf(prompt) //  TODO:prompt.txt has placeholders for formatting will add them
	return prompt
}