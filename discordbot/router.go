package discordbot

import cmdr "github.com/AriNoa/CommandRouterGo"

func newRouter() *cmdr.Router {
	return cmdr.NewRouter(
		"!",
		[]*cmdr.Command{
			cmdr.NewCommand(
				"PingPong",
				"ping",
				"",
				"",
				"",
				[]*cmdr.Command{},
				nil, // TODO
			),
		},
	)
}
