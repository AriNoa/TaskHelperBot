package cmdrouter

import "strings"

// DetachCommandFrom function detaches the command from the string and returns it and its arguments
func DetachCommandFrom(str string) (command string, argument string) {
	slice := strings.Split(str, " ")

	command = slice[0]
	argument = strings.TrimLeft(strings.TrimLeft(str, command), " ")

	return
}

// MakeCommandMapFrom function make the map from the command slice with the command alias as the key.
func MakeCommandMapFrom(commands []*Command) map[string]*Command {
	return map[string]*Command{}
}
