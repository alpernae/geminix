package main

import (
    "context"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"

    "github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)

var (
    apiKey    string
    configDir string
    prompt    string // -p ile alınacak prompt
)

func init() {
    // config dizini
    configDir = filepath.Join(os.Getenv("HOME"), ".config", "geminix") // kullanıcı dizini altında .config/geminix

    // config dizini yoksa oluşturulacak
    if err := os.MkdirAll(configDir, 0755); err != nil {
        log.Fatalf("Failed to create config directory: %v", err)
    }

    // config dosyasını kontrol et
    configFile := filepath.Join(configDir, "config")
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        // Eğer dosya yoksa -k parametresi ile API key alınacak
        flag.StringVar(&apiKey, "k", "", "API key for Gemini")
        flag.Parse()

        if apiKey == "" {
            log.Fatal("API key must be provided with the -k flag")
        }

        // API key'i .config/geminix/config dosyasına kaydedelim
        if err := ioutil.WriteFile(configFile, []byte(apiKey), 0600); err != nil {
            log.Fatalf("Failed to save API key to config file: %v", err)
        }

        // API key kaydedildiğinde herhangi bir çıktı vermek yerine sadece kaydedildiği bildirilir.
    } else {
        // Eğer config dosyası varsa, kaydedilen API key'i okuyalım
        loadedApiKey, err := ioutil.ReadFile(configFile)
        if err != nil {
            log.Fatalf("Error reading API key from config file: %v", err)
        }
        apiKey = string(loadedApiKey)
    }

    // -p parametresi ile prompt alınacak
    flag.StringVar(&prompt, "p", "", "Prompt to send to Gemini API")
    flag.Parse()
}

func main() {
    // API key ile client oluşturulacak
    client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
    if err != nil {
        log.Fatalf("Error creating Gemini client: %v", err)
    }
    defer client.Close()

    // Eğer prompt alınmamışsa, kullanıcının inputunu al
    if prompt == "" {
        fmt.Println("Please provide a prompt using -p.")
        return
    }

    // Chat başlatma ve cevap alma
    model := client.GenerativeModel("gemma-3-1b-it")
    model.SetTemperature(1)
    model.SetTopK(40)
    model.SetTopP(0.95)
    model.SetMaxOutputTokens(8192)
    model.ResponseMIMEType = "text/plain"

    session := model.StartChat()
    session.History = []*genai.Content{}

    // Gönderilen mesaj
    resp, err := session.SendMessage(context.Background(), genai.Text(prompt))
    if err != nil {
        log.Fatalf("Error sending message: %v", err)
    }

    // Yanıtı yazdırma
    for _, part := range resp.Candidates[0].Content.Parts {
        fmt.Printf("%v\n", part)
    }
}
