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

// UpdateNextNoticeTime updates the next notification time from the Repeat value
func (n *Notice) UpdateNextNoticeTime() error {
	if !n.Repeat.hasValue {
		return errors.New("Notice's repeat has not value")
	}

	n.Next.Add(n.Repeat.value)

	return nil
}
