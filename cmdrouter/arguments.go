package cmdrouter

import "strings"

// Arguments represent command arguments
type Arguments struct {
	raw   string
	slice []string
}

// New is constructor for Arguments
func New(raw string) *Arguments {
	return &Arguments{raw, strings.Split(raw, " ")}
}
