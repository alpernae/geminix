// cmd/geminix/main.go
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alpernae/geminix/internal/prompt"
	"github.com/alpernae/geminix/pkg/gemini"
)

func main() {
	ctx := context.Background()

	// CLI parametreleri
	modelName := flag.String("m", "gemini-1.5-pro", "Model name to use")
	flag.Parse()

	// API anahtarını al
	apiKey := gemini.GetAPIKey()

	// API istemcisini oluştur
	client, err := gemini.NewGeminiClient(ctx, apiKey, *modelName)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}
	defer client.Client.Close()

	// Standart girdiden prompt alma
	var userInput string
	if len(flag.Args()) > 0 {
		userInput = strings.Join(flag.Args(), " ")
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter prompt: ")
		input, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(input)
	}

	// Prompt işleme ve API'ye gönderme
	prompt.SendPrompt(ctx, client, userInput)
}
