package commands

import (
	"log"

	cmdr "github.com/AriNoa/CommandRouterGo"
	"github.com/bwmarrin/discordgo"

	echo "TaskHelperBot/domain/echo/usecase"
)

// EchoPresenter is a struct that implements echo.Presenter interface
type EchoPresenter struct {
	Session *discordgo.Session
	Event   *discordgo.MessageCreate
}

// Echo is a method that implements echo.Presenter's Echo method
func (p *EchoPresenter) Echo(contents string) {
	channelID := p.Event.ChannelID

	p.Session.ChannelMessageSend(channelID, contents)
}

// EchoHandler is a struct that implements CommandHandler interface
type EchoHandler struct {
	Command       *cmdr.Command
	UseCase       echo.UseCase
	Presenter     *EchoPresenter
	HelpPresenter cmdr.CommandHelpPresenter
}

// Handle is a method that implements CommandHandler's Handle method
func (eh *EchoHandler) Handle(ctx *cmdr.Context) {
	s, err := FindSessionFrom(ctx.Props)
	if err != nil {
		log.Fatal(err)
	}

	e, err := FindEventFrom(ctx.Props)
	if err != nil {
		log.Fatal(err)
	}

	if ctx.Argument == "" {
		eh.HelpPresenter.Help(eh.Command, ctx.Props)
	}

	eh.Presenter.Session = s
	eh.Presenter.Event = e

	eh.UseCase.Echo(ctx.Argument)
}
