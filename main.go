package main

import (
	"github.com/AriNoa/TaskHelperBot/consoleapp"
	"github.com/AriNoa/TaskHelperBot/discordbot"
)

func main() {
	go discordbot.Run()
	consoleapp.Run()
}
