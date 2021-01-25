package consoleapp

import (
	"bufio"
	"fmt"
	"os"
)

// Run is a function to run DiscordBot
func Run() {
	fmt.Println("Welcome to ConsoleApp")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "exit" {
			break
		}

		// TODO

		fmt.Printf("[%s]\n", txt)
	}

	fmt.Println("See you next time")
}
