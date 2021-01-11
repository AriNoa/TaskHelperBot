package cmdrouter

import "github.com/bwmarrin/discordgo"

// Context is the Context passed to the command handler
type Context struct {
	Session  discordgo.Session
	Event    discordgo.MessageCreate
	Argument string
	Props    map[string]interface{}
}
