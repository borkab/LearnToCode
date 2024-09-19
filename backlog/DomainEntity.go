package main

import "time"

/*
Note is a struct, that represent a sticky note. It contains a title, a content,
a date when it was created and a date when it was updated.
*/
type Note struct {
	ID      string
	Title   string
	Content string
	Info
}

type Info struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create creates a new Note entity. It takes an id, a title and a content as strings
// and returns a pointer of a new note
func Create(id, title, content string) *Note {
	now := time.Now()
	return &Note{
		ID:      id,
		Title:   title,
		Content: content,
		Info: Info{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}
