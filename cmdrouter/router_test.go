package cmdrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestRouter() (
	r *Router,
	handlerA *TestHandler,
	handlerB *TestHandler,
) {
	handlerA = &TestHandler{false}
	handlerB = &TestHandler{false}

	r = NewRouter(
		"!",
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
	)

	return
}

func TestCommandRouteNotCommand(t *testing.T) {
	r, a, b := createTestRouter()

	r.route(&Context{
		"abc",
		map[string]interface{}{},
	})

	assert.False(t, a.wasCalled)
	assert.False(t, b.wasCalled)
}

func TestCommandRouteCommand(t *testing.T) {
	r, a, b := createTestRouter()

	r.route(&Context{
		"!a",
		map[string]interface{}{},
	})

	assert.True(t, a.wasCalled)
	assert.False(t, b.wasCalled)
}
