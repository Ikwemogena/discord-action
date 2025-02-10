package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DiscordWebhookID    string
	DiscordWebhookToken string
}

func LoadConfig() (*config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &config{
		DiscordWebhookID:    os.Getenv("DISCORD_WEBHOOK_ID"),
		DiscordWebhookToken: os.Getenv("DISCORD_WEBHOOK_TOKEN"),
	}, nil
}