package notice

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateNextNoticeTime(t *testing.T) {
	n := New(
		0,
		"test",
		time.Date(2020, time.December, 31, 22, 42, 24, 0, time.UTC),
		3*time.Hour+4*time.Minute,
		true,
	)

	err := n.UpdateNextNoticeTime()

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, n.Next, time.Date(2021, time.January, 1, 1, 46, 24, 0, time.UTC))
}

func TestUpdateNextNoticeTimeError(t *testing.T) {
	n := New(
		0,
		"test",
		time.Date(2020, time.December, 31, 22, 42, 24, 0, time.UTC),
		0,
		true,
	)

	err := n.UpdateNextNoticeTime()

	assert.Equal(t, err, errors.New("Notice's repeat has not value"))
}
