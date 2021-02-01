package commands

import (
	"fmt"

	cmdr "github.com/AriNoa/CommandRouterGo"

	echo "TaskHelperBot/domain/echo/usecase"
)

// EchoPresenter is a struct that implements echo.Presenter interface
type EchoPresenter struct{}

// Echo is a method that implements echo.Presenter's Echo method
func (p *EchoPresenter) Echo(contents string) {
	fmt.Printf("%s\n", contents)
}

// EchoHandler is a struct that implements CommandHandler interface
type EchoHandler struct {
	UseCase echo.UseCase
}

// Handle is a method that implements CommandHandler's Handle method
func (eh *EchoHandler) Handle(ctx *cmdr.Context) {
	eh.UseCase.Echo(ctx.Argument)
}
