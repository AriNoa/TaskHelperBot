package cmdrouter

// RouteHandler is a struct that implements CommandHandler
type RouteHandler struct {
	Command       *Command
	HelpPresenter CommandHelpPresenter
}

// Handle of RouteHandler calls the subcommand handle according to Argument
func (rh *RouteHandler) Handle(ctx *Context) {
	cmd, arg := DetachCommandFrom(ctx.Argument)

	if c, ok := rh.Command.SubCommands[cmd]; ok {

		ctx.Argument = arg

		c.Handler.Handle(ctx)

		return
	}

	rh.HelpPresenter.Help(rh.Command)
}
