package actions

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	actions "github.com/tacheraSasi/ellie/action"
	"google.golang.org/api/option"
)

func geminiChat(prompt string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	actions.RenderMarkdown(resp)


}
