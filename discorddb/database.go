package discorddb

import (
	"encoding/json"
	"errors"

	"github.com/bwmarrin/discordgo"
)

// DataBase is a struct for using Discord as a database
type DataBase struct {
	Session   *discordgo.Session
	Tables    map[string]Table
	ChannelID string
}

// New is a constructor for DataBase
func New(s *discordgo.Session, channelID string) (*DataBase, error) {
	tables := map[string]Table{}

	msgs, err := s.ChannelMessages(channelID, 100, "", "", "")

	if err != nil {
		return nil, err
	}

	for _, m := range msgs {
		key, table, err := CreateTableFromMessage(s, m)

		if err != nil {
			return nil, err
		}

		if oldTable, ok := tables[key]; ok {
			DeleteMessageWithLog(s, channelID, oldTable.MessageID)
		}

		tables[key] = *table
	}

	return &DataBase{s, tables, channelID}, nil
}

// Create creates a new table
func (db *DataBase) Create(key string) error {
	if _, ok := db.Tables[key]; ok {
		return errors.New("the table has already been created")
	}

	data, err := json.Marshal(map[string]interface{}{key: nil})
	if err != nil {
		return err
	}

	m, err := db.Session.ChannelMessageSend(db.ChannelID, string(data))
	if err != nil {
		return err
	}

	db.Tables[key] = Table{data, m.ID}
	return nil
}
