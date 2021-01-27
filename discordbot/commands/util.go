package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// FindSessionFrom is a function that finds Session from props
func FindSessionFrom(props map[string]interface{}) (*discordgo.Session, error) {
	var i interface{}
	var ok bool
	var s *discordgo.Session

	i, ok = props["session"]
	if !ok {
		return nil, errors.New("\"session\" is not found")
	}

	s, ok = i.(*discordgo.Session)
	if !ok {
		return nil, errors.New("\"session\"'s type is not *discordgo.Session")
	}

	return s, nil
}

// FindEventFrom is a function that finds Event from props
func FindEventFrom(props map[string]interface{}) (*discordgo.MessageCreate, error) {
	var i interface{}
	var ok bool
	var e *discordgo.MessageCreate

	i, ok = props["event"]
	if !ok {
		return nil, errors.New("\"event\" is not found")
	}

	e, ok = i.(*discordgo.MessageCreate)
	if !ok {
		return nil, errors.New("\"event\"'s type is not *discordgo.MessageCreate")
	}

	return e, nil
}
