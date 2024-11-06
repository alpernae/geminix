// internal/prompt/handler.go
package prompt

import (
	"context"
	"fmt"
	"log"

	"github.com/alpernae/geminix/pkg/gemini"
	"github.com/google/generative-ai-go/genai"
)

func SendPrompt(ctx context.Context, client *gemini.GeminiClient, userInput string) {
	session := client.Model.StartChat()
	session.History = []*genai.Content{}

	resp, err := session.SendMessage(ctx, genai.Text(userInput))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}
