package cmdrouter

// Command holds data about the command
type Command struct {
	Name        string
	Alias       string
	Summary     string
	Usage       string
	Example     string
	SubCommands []*Command
	Handler     CommandHandler
}
