package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Payload struct {
	Content    string `json:"content,omitempty"`
}

func check(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func SendWebhook(url string, payload Payload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)


	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to send webhook, status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	cfg, err := LoadConfig()
	check(err, "Failed to load config")

	discord := "https://discord.com/api/webhooks"
    webhookID := cfg.DiscordWebhookID
	webhookToken := cfg.DiscordWebhookToken

	if webhookID == "" || webhookToken == "" {
		log.Fatal("Missing Discord webhook credentials in environment variables")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <message>")
	}

	content := os.Args[1]

	webhookURL := fmt.Sprintf("%s/%s/%s", discord, webhookID, webhookToken)

    payload := Payload{
		Content: content,
	}
		
	err = SendWebhook(webhookURL, payload)

	check(err, "Failed to send webhook")

	if err != nil {
		log.Fatalf("Failed to send webhook: %v", err)
	}

	fmt.Println("Webhook sent successfully!")
}
