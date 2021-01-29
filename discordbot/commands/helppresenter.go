package commands

import (
	"fmt"
	"log"

	cmdr "github.com/AriNoa/CommandRouterGo"
)

// HelpPresenter is a struct that implements cmdr.CommandHelpPresenter interface
type HelpPresenter struct{}

// Help is a method that implements cmdr.CommandHelpPresente's Help method
func (hp *HelpPresenter) Help(c *cmdr.Command, props map[string]interface{}) {
	s, err := FindSessionFrom(props)
	if err != nil {
		log.Fatal(err)
	}

	e, err := FindEventFrom(props)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("## Name\n%s", c.Name)
	summary := fmt.Sprintf("## Summary\n%s", c.Summary)
	usage := fmt.Sprintf("## Usage\n%s", c.Usage)
	example := fmt.Sprintf("## Example\n%s", c.Example)

	message := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n",
		name,
		summary,
		usage,
		example,
	)

	s.ChannelMessageSend(
		e.ChannelID,
		fmt.Sprintf("```markdown\n%s```", message),
	)
}
