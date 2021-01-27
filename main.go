package main

import (
	"TaskHelperBot/consoleapp"
	"TaskHelperBot/discordbot"
)

func main() {
	go discordbot.Run()
	consoleapp.Run()
}
