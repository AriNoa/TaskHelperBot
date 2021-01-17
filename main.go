package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"

	cmdr "TaskHelperBot/cmdrouter"
	cmds "TaskHelperBot/commands"
)

var router cmdr.Router

func main() {
	session, err := discordgo.New()
	if err != nil {
		fmt.Println("Error in create session")
		panic(err)
	}

	discordToken := loadToken()
	session.Token = discordToken

	session.AddHandler(onReady)
	session.AddHandler(onMessageCreate)

	router = *cmdr.NewRouter(
		"!",
		[]*cmdr.Command{
			cmdr.NewCommand(
				"PingPong",
				"ping",
				"",
				"",
				"",
				[]*cmdr.Command{},
				&cmds.PingPongHandler{},
			),
		},
	)

	if err = session.Open(); err != nil {
		panic(err)
	}
	defer session.Close()

	loopch := make(chan bool)
	go loop(1*time.Minute, loopch)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-loopch
	<-sc
	return
}

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

func loop(d time.Duration, ch chan bool) {
	for {
		time.Sleep(d)
		fmt.Println("hello")
	}
}
