package actions

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Model    string    `json:"model"`
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

func Chat(openaiApikey string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Talk to me: ")
		msg, readerErr := reader.ReadString('\n')
		if readerErr != nil {
			fmt.Println("Something went wrong", readerErr)
		}
		if msg == "exit" {
			break
		}

		output := chatWithOpenAI(msg,openaiApikey)
		fmt.Println(output)
	}

}
func chatWithOpenAI(message,openaiApikey string) string {
	url := "https://api.openai.com/v1/chat/completions"

	reqBody := OpenAIRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "system", Content: "You are Ellie, an AI therapist and friend."},
			{Role: "user", Content: message},
		},
	}

	body,err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshalling json", err)
		return ""
	}

	req, reqErr := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if reqErr != nil {
		fmt.Println("Error creating request", reqErr)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + openaiApikey)
	// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		defer resp.Body.Close()

		// Read the response
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response: %v", err)
		}

		// Parse the response
		var openAIResponse OpenAIResponse
		if err := json.Unmarshal(respBody, &openAIResponse); err != nil {
			log.Fatalf("Error unmarshalling response: %v", err)
		}

		// Return the first response from OpenAI
		if len(openAIResponse.Choices) > 0 {
			return openAIResponse.Choices[0].Message.Content
		}

		return "Sorry, I couldn't generate a response."

}
