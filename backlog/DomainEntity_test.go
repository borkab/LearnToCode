package main

import (
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	id := "newID123"
	title := "first Note"
	content := "this is my first sticletti"

	got := Create(id, title, content)
	want := &Note{
		ID:      "newID123",
		Title:   "first Note",
		Content: "this is my first sticletti",
		Info: Info{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
