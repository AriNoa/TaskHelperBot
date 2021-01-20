package discorddb

import "encoding/json"

// Table holds the data marshalled to JSON and the corresponding Discord message
type Table struct {
	Value     []byte
	MessageID string
}

// NewTable is a constructor for Router
func NewTable(value interface{}, messageID string) (*Table, error) {
	json, err := json.Marshal((value))

	if err != nil {
		return nil, err
	}

	return &Table{
		json,
		messageID,
	}, nil
}
