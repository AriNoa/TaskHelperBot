package cmdrouter

// RouteHandler is a struct that implements CommandHandler
type RouteHandler struct {
	Command       *Command
	HelpPresenter CommandHelpPresenter
}

// Handle of RouteHandler calls the subcommand handle according to Argument
func (rh *RouteHandler) Handle(ctx *Context) {
	return
}
