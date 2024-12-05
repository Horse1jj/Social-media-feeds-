package YouTube

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

var youtubeCache = make(map[string]bool)

func MonitorYouTube(s *discordgo.Session, config map[string]interface{}) {
	apiKey := config["YouTubeAPIKey"].(string)
	channels := config["YouTubeChannels"].([]string)
	discordChannelID := config["DiscordChannelID"].(string)

	for {
		for _, channelID := range channels {
			url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?channelId=%s&key=%s&part=snippet&order=date&maxResults=1", channelID, apiKey)
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Error fetching YouTube data: %v", err)
				continue
			}
			defer resp.Body.Close()

			// Parse response and send notification for new video
			// Add caching to prevent duplicates
		}

		time.Sleep(5 * time.Minute)
	}
}

