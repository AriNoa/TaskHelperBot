package consoleapp

import (
	"bufio"
	"fmt"
	"os"

	cmdr "github.com/AriNoa/CommandRouterGo"
)

// Run is a function to run DiscordBot
func Run() {
	fmt.Println("Welcome to ConsoleApp")

	r := newRouter()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "exit" {
			break
		}

		ctx := cmdr.Context{
			Argument: txt,
			Props:    map[string]interface{}{},
		}

		r.Route(&ctx)
	}

	fmt.Println("See you next time")
}
