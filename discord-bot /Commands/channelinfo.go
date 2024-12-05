package Commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ChannelInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: /channelinfo <platform> <channel/user>")
		return
	}

	platform := args[0]
	channel := args[1]

	switch platform {
	case "youtube":
		// Fetch and display YouTube channel info
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Fetching info for YouTube channel: %s", channel))
	case "twitch":
		// Fetch and display Twitch user info
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Fetching info for Twitch user: %s", channel))
	case "tiktok":
		// Fetch and display TikTok user info
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Fetching info for TikTok user: %s", channel))
	default:
		s.ChannelMessageSend(m.ChannelID, "Invalid platform. Use youtube, twitch, or tiktok.")
	}
}

