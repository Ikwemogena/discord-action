package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Payload struct {
	Content    string `json:"content,omitempty"`
}

func check(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func NotifyDiscord(url string, payload Payload) error {
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
	discord := "https://discord.com/api/webhooks"

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <webhook_id> <webhook_token> <message>")
		os.Exit(1)
	}
	
	webhookID := os.Args[1]
	webhookToken := os.Args[2]
	content := strings.Join(os.Args[3:], " ")
	
	webhookURL := fmt.Sprintf("%s/%s/%s", discord, webhookID, webhookToken)

    payload := Payload{
		Content: content,
	}
		
	err := NotifyDiscord(webhookURL, payload)

	check(err, "Failed to send webhook")

	fmt.Println("Webhook sent successfully!")
}
