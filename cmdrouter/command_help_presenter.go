package cmdrouter

// CommandHelpPresenter is the interface for displaying help messages for commands
type CommandHelpPresenter interface {
	Help(c *Command)
}
