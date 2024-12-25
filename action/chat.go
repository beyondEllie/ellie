package actions

import (
	"bufio"
	"fmt"
	"os"
)


type OpenAIRequest struct {
	Model    string  `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func Chat(openaiApikey string){
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Talk to me: ")
		msg,readerErr := reader.ReadString('\n')
		if readerErr != nil{
			fmt.Println("Something went wrong",readerErr)
		}
		if msg == "exit" {
			break
		}
		
		output := chatWithOpenAI(msg)
		fmt.Println(output)
	}




}
func chatWithOpenAI(message string) string {
	url := "https://api.openai.com/v1/chat/completions"
	return "I am a chatbot"

}
