package notice

import "time"

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
