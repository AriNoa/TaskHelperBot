package discorddb

import (
	"encoding/json"
	"errors"
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

// CreateTableFromMessage creates a table from the message and returns the key and table
func CreateTableFromMessage(s *discordgo.Session, m *discordgo.Message) (string, *Table, error) {
	var hash map[string]interface{}

	err := json.Unmarshal([]byte(m.Content), &hash)

	if err != nil {
		return "", nil, err
	}

	_key, ok := hash["key"]
	if !ok {
		return "", nil, errors.New("message does not have \"key\"")
	}

	key, ok := _key.(string)
	if !ok {
		return "", nil, errors.New("key type is not string")
	}

	value, ok := hash["value"]
	if !ok {
		return "", nil, errors.New("message does not have \"value\"")
	}

	t, err := NewTable(value, m.ID)
	if err != nil {
		return "", nil, err
	}

	return key, t, nil
}
