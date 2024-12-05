package Twitch

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

var twitchLiveStatus = make(map[string]bool)

func MonitorTwitch(s *discordgo.Session, config map[string]interface{}) {
	clientID := config["TwitchClientID"].(string)
	accessToken := config["TwitchAccessToken"].(string)
	users := config["TwitchUsers"].([]string)
	discordChannelID := config["DiscordChannelID"].(string)

	for {
		for _, user := range users {
			url := fmt.Sprintf("https://api.twitch.tv/helix/streams?user_login=%s", user)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Client-ID", clientID)
			req.Header.Set("Authorization", "Bearer "+accessToken)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Error fetching Twitch data: %v", err)
				continue
			}
			defer resp.Body.Close()

			// Parse response and send notification if the user goes live
		}

		time.Sleep(1 * time.Minute)
	}
}

