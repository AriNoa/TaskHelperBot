package discorddb

// Table holds the data marshalled to JSON and the corresponding Discord message
type Table struct {
	Value     interface{}
	MessageID string
}
