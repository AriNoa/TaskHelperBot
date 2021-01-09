package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	session, err := discordgo.New()
	if err != nil {
		fmt.Println("Error in create session")
		panic(err)
	}

	discordToken := loadToken()
	session.Token = discordToken

	session.AddHandler(onMessageCreate)

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

func onMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot {
		return
	}

	if event.Content != "ping" {
		return
	}

	_, Err := session.ChannelMessageSend(event.ChannelID, "pong")
	if Err != nil {
		log.Println("failed send message: ", Err)
	}
	return
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
