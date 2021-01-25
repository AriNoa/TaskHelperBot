package discordbot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"

	cmdr "github.com/AriNoa/CommandRouterGo"
)

// LoopInterval is the interval at which the loopEvent function is called
const LoopInterval = time.Minute

var router cmdr.Router

// Run is a function to run DiscordBot
func Run() {
	session, err := discordgo.New()
	if err != nil {
		fmt.Println("Error in create session")
		panic(err)
	}

	discordToken := loadToken()
	session.Token = discordToken

	session.AddHandler(onReady)
	session.AddHandler(onMessageCreate)

	router = *newRouter()

	if err = session.Open(); err != nil {
		panic(err)
	}
	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	go func() {
		for {
			time.Sleep(LoopInterval)
			loopEvent()
		}
	}()

	<-sc
	return
}
