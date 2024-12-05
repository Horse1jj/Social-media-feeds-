package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"discord-bot/YouTube"
	"discord-bot/TikTok"
	"discord-bot/Twitch"
	"discord-bot/Commands"
)

func main() {
	// Load configuration
	config := loadConfig("config.json")

	// Initialize Discord bot
	dg, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Add command handlers
	dg.AddMessageCreateHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		Commands.HandleCommands(s, m)
	})

	// Start social media monitors
	go YouTube.MonitorYouTube(dg, config)
	go TikTok.MonitorTikTok(dg, config)
	go Twitch.MonitorTwitch(dg, config)

	// Open WebSocket connection
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening WebSocket connection: %v", err)
	}
	defer dg.Close()

	fmt.Println("Bot is running. Press CTRL+C to exit.")
	select {}
}

