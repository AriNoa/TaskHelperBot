package cmdrouter

// Context is the Context passed to the command handler
type Context struct {
	Argument string
	Props    map[string]interface{}
}
