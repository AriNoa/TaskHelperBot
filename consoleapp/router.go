package consoleapp

import (
	cmdr "github.com/AriNoa/CommandRouterGo"

	"TaskHelperBot/consoleapp/commands"
	"TaskHelperBot/domain/usecase/echo"
)

func newRouter() *cmdr.Router {
	return cmdr.NewRouter(
		"!",
		[]*cmdr.Command{
			cmdr.NewCommand(
				"Echo",
				"echo",
				"任意の文字列を表示することができます",
				"echo [文字列]",
				"echo Hello world",
				[]*cmdr.Command{},
				&commands.EchoHandler{
					UseCase: &echo.Interactor{
						Presenter: &commands.EchoPresenter{},
					},
				},
			),
		},
	)
}
