package commands

import (
	"log"

	cmdr "TaskHelperBot/cmdrouter"

	"github.com/bwmarrin/discordgo"
)

// PingPongHandler is a struct that implements CommandHandler
type PingPongHandler struct{}

// Handle of RouteHandler responds with a "pong" when it receives a "ping"
func (h *PingPongHandler) Handle(ctx *cmdr.Context) {
	var session *discordgo.Session
	var event *discordgo.MessageCreate

	// TODO
	if s, ok := ctx.Props["session"]; ok {
		if session, ok = s.(*discordgo.Session); !ok {
			log.Print("session type is not discordgo.Session")
			return
		}
	} else {
		log.Print("there are no session in props")
		return
	}

	if e, ok := ctx.Props["event"]; ok {
		if event, ok = e.(*discordgo.MessageCreate); !ok {
			log.Print("event type is not discordgo.MessageCreate")
			return
		}
	} else {
		log.Print("there are no event in props")
		return
	}

	session.ChannelMessageSend(event.ChannelID, "pong")
}
