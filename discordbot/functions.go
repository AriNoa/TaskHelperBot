package discordbot

import (
	"fmt"
	"os"

	cmdr "github.com/AriNoa/CommandRouterGo"
	"github.com/bwmarrin/discordgo"
)

func onReady(session *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Name\t: %s\n", event.User.Username)
	fmt.Printf("ID\t: %s\n", event.User.ID)
}

func onMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot {
		return
	}

	props := map[string]interface{}{
		"session": session,
		"event":   event,
	}

	ctx := cmdr.Context{
		Argument: event.Content,
		Props:    props,
	}

	router.Route(&ctx)
}

func loadToken() string {
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("no discord token exists.")
	}
	return "Bot " + token
}

func loopEvent() {
	// TODO
}
