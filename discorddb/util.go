package discorddb

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// DeleteMessageWithLog deletes the message after displaying the content of the message
func DeleteMessageWithLog(session *discordgo.Session, channelID string, messageID string) error {
	message, err := session.ChannelMessage(channelID, messageID)

	if err != nil {
		return err
	}

	log.Printf("Delete message \"[%s]%s\"", message.Author.Username, message.Content)

	return session.ChannelMessageDelete(channelID, messageID)
}
