package Commands

import (
	"github.com/bwmarrin/discordgo"
)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "Here are the available commands:\n" +
		"/help - Show this help message\n" +
		"/channelinfo - Get info about a channel (e.g., subscriber count, etc.)"
	s.ChannelMessageSend(m.ChannelID, helpMessage)
}
