package cmdrouter

// Command holds data about the command
type Command struct {
	Name        string
	Alias       string
	Summary     string
	Usage       string
	Example     string
	SubCommands map[string]*Command
	Handler     CommandHandler
}

// NewCommand is a constructor for Command
func NewCommand(
	name string,
	alias string,
	summary string,
	usage string,
	example string,
	subCommands []*Command,
	handler CommandHandler,
) *Command {
	subCommandMap := MakeCommandMapFrom(subCommands)

	return &Command{
		name,
		alias,
		summary,
		usage,
		example,
		subCommandMap,
		handler,
	}
}
