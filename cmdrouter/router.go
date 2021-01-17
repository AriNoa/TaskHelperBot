package cmdrouter

// Router represents the router for discord commands
type Router struct {
	Prefix   string
	Commands map[string]*Command
}

// NewRouter is a constructor for Router
func NewRouter(prefix string, commands []*Command) *Router {
	commandMap := MakeCommandMapFrom(commands)

	return &Router{
		prefix,
		commandMap,
	}
}

// Route is a router handler provided to discordgo session
func (r *Router) Route(ctx *Context) {
	//TODO
}
