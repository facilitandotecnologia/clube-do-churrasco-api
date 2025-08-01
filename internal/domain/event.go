package domain

import "time"

type Event struct {
	ID          string
	Date        time.Time
	Description string
	Attenders   []string
	Details     string
}
