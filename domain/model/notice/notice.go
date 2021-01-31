package notice

import (
	"errors"
	"time"
)

// Notice is a structure that expresses a notification
type Notice struct {
	ID       int
	Contents string
	Next     time.Time
	Repeat   struct {
		value    time.Duration
		hasValue bool
	}
	Enabled bool
}

// New is a constructor for Notice
func New(
	id int,
	contnts string,
	next time.Time,
	interval time.Duration,
	enabled bool,
) *Notice {
	var repeat struct {
		value    time.Duration
		hasValue bool
	}

	if interval != 0 {
		repeat.hasValue = true
		repeat.value = interval
	}

	return &Notice{
		id,
		contnts,
		next,
		repeat,
		enabled,
	}
}

// UpdateNextNoticeTime updates the next notification time from the Repeat value
func (n *Notice) UpdateNextNoticeTime() error {
	if !n.Repeat.hasValue {
		return errors.New("Notice's repeat has not value")
	}

	n.Next.Add(n.Repeat.value)

	return nil
}
