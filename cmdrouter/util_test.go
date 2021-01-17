package cmdrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetachCommandFromEmpty(t *testing.T) {
	cmd, arg := DetachCommandFrom("")

	assert.Equal(t, cmd, "")
	assert.Equal(t, arg, "")
}

func TestDetachCommandFromCommandOnly(t *testing.T) {
	cmd, arg := DetachCommandFrom("command")

	assert.Equal(t, cmd, "command")
	assert.Equal(t, arg, "")
}

func TestDetachCommandFromCommands(t *testing.T) {
	cmd, arg := DetachCommandFrom("command arg")

	assert.Equal(t, cmd, "command")
	assert.Equal(t, arg, "arg")
}

func TestMakeCommandMapFromCommands(t *testing.T) {
	commandA := NewCommand(
		"A",
		"a",
		"",
		"",
		"",
		[]*Command{},
		nil,
	)
	commandB := NewCommand(
		"B",
		"b",
		"",
		"",
		"",
		[]*Command{},
		nil,
	)

	commandMap := MakeCommandMapFrom([]*Command{commandA, commandB})
	assert.Equal(
		t,
		commandMap,
		map[string]*Command{
			"a": commandA,
			"b": commandB,
		},
	)
}

func TestMakeCommandMapFromEmpty(t *testing.T) {
	commandMap := MakeCommandMapFrom([]*Command{})
	assert.Equal(
		t,
		commandMap,
		map[string]*Command{},
	)
}
