package cmdrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHelpPresenter struct {
	wasCalled bool
}

func (p *TestHelpPresenter) Help(c *Command) {
	p.wasCalled = true
}

type TestHandler struct {
	wasCalled bool
}

func (th *TestHandler) Handle(ctx *Context) {
	th.wasCalled = true
}

func createTestCommand() (
	command *Command,
	presenter *TestHelpPresenter,
	handlerA *TestHandler,
	handlerB *TestHandler,
) {
	command = &Command{}
	presenter = &TestHelpPresenter{false}
	handlerA = &TestHandler{false}
	handlerB = &TestHandler{false}

	*command = *NewCommand(
		"Test",
		"test",
		"",
		"",
		"",
		[]*Command{
			NewCommand(
				"A",
				"a",
				"",
				"",
				"",
				[]*Command{},
				handlerA,
			),
			NewCommand(
				"B",
				"b",
				"",
				"",
				"",
				[]*Command{},
				handlerB,
			),
		},
		&RouteHandler{command, presenter},
	)

	return
}

func TestRouteHandlerHelp(t *testing.T) {
	c, p, a, b := createTestCommand()
	c.Handler.Handle(&Context{"", map[string]interface{}{}})

	assert.True(t, p.wasCalled)
	assert.False(t, a.wasCalled)
	assert.False(t, b.wasCalled)
}

func TestRouteHandlerSubCommandHandle(t *testing.T) {
	c, p, a, b := createTestCommand()
	c.Handler.Handle(&Context{"a", map[string]interface{}{}})

	assert.False(t, p.wasCalled)
	assert.True(t, a.wasCalled)
	assert.False(t, b.wasCalled)
}
