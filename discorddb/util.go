package discorddb

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// DeleteMessageWithLog deletes the message after displaying the content of the message
func DeleteMessageWithLog(s *discordgo.Session, channelID string, messageID string) error {
	m, err := s.ChannelMessage(channelID, messageID)

	if err != nil {
		return err
	}

	log.Printf("Delete message \"[%s]%s\"", m.Author.Username, m.Content)

	return s.ChannelMessageDelete(channelID, messageID)
}
