package dictionary

import "time"

type Entry struct {
	Word      string
	Meaning   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
