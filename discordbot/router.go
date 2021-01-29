package discordbot

import (
	cmdr "github.com/AriNoa/CommandRouterGo"

	"TaskHelperBot/discordbot/commands"
	"TaskHelperBot/domain/usecase/echo"
)

func newEchoCommand() *cmdr.Command {
	echoPresenter := &commands.EchoPresenter{}
	var echoHandler commands.EchoHandler

	echoCommand := cmdr.NewCommand(
		"Echo",
		"echo",
		"任意の文字列を表示することができます",
		"echo [文字列]",
		"echo Hello world",
		[]*cmdr.Command{},
		&echoHandler,
	)

	echoHandler = commands.EchoHandler{
		Command: echoCommand,
		UseCase: &echo.Interactor{
			Presenter: echoPresenter,
		},
		Presenter:     echoPresenter,
		HelpPresenter: &commands.HelpPresenter{},
	}

	return echoCommand
}

func newRouter() *cmdr.Router {
	return cmdr.NewRouter(
		"!",
		[]*cmdr.Command{
			newEchoCommand(),
		},
	)
}
