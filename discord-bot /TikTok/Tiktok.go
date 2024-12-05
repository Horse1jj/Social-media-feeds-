package TikTok

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MonitorTikTok(s *discordgo.Session, config map[string]interface{}) {
	tiktokUsers := config["TikTokUsers"].([]string)
	discordChannelID := config["DiscordChannelID"].(string)

	for {
		for _, user := range tiktokUsers {
			// Fetch TikTok data (scraping or third-party APIs)
			log.Printf("Checking TikTok updates for %s...", user)
			// Notify in Discord if there's a new video
		}

		time.Sleep(10 * time.Minute)
	}
}

