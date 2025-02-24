package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/tacheraSasi/ellie/utils"
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
		msg, err := utils.GetInput("Talk to me: ")
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if strings.EqualFold(msg, "exit") {
			fmt.Println("Goodbye!")
			break
		}
		// fmt.Println(utils.Ads[1])

		output := chatWithOpenAI(msg, openaiApikey)
		if output == "" {
			fmt.Println("No response received.")
			continue
		}

		// Rendering markdown with glamour
		renderedOutput, err := renderMarkdown(output)
		if err != nil {
			fmt.Println("Error rendering Markdown:", err)
			continue
		}
		if utils.IsEven(utils.RandNum()){
			fmt.Println(utils.Ads[1])
		}
		fmt.Println(renderedOutput)
	}
}

func chatWithOpenAI(message, openaiApikey string) string {
	url := "https://api.openai.com/v1/chat/completions"
	instructions := fmt.Sprintf("You are Ellie, a local Linux AI assistant and friend. Everything about you: %s, You were created by Tachera sasi he is so  brilliant and handsome", getReadmeContent())

	reqBody := OpenAIRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "system", Content: instructions},
			{Role: "user", Content: message},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return ""
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiApikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: received status code %d", resp.StatusCode)
		return ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return ""
	}

	var openAIResponse OpenAIResponse
	if err := json.Unmarshal(respBody, &openAIResponse); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return ""
	}

	if len(openAIResponse.Choices) > 0 {
		return openAIResponse.Choices[0].Message.Content
	}

	return "No response received from OpenAI."
}

func ChatWithGemini(){}

func getReadmeContent() string {
	content, err := os.ReadFile("./README.md")
	if err != nil {
		log.Printf("Error reading README.md: %v", err)
		return "README.md file not found or unreadable."
	}
	return string(content)
}

func renderMarkdown(input string) (string, error) {
	// Rendering Markdown with glamour
	rendered, err := glamour.Render(input, "dark") 
	if err != nil {
		return "", err
	}
	return rendered, nil
}
