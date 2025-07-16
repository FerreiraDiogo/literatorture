package dictionary

import "time"

type Entry struct {
	Word      string
	Meaning   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEntry(word, meaning string) *Entry {
	return &Entry{Word: word, Meaning: meaning, CreatedAt: time.Now(), UpdatedAt: time.Now()}
}
