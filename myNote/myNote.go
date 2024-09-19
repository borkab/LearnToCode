package main

import "time"

// myNote is a note that holds data in his content field. It has a title, some info
// about when it was created and updated, and has a unique ID to be able to search for it.
type myNote struct {
	Title   string
	Content string
	Info
	ID string
}

// info is a struct that holds two dates: when was a note created and when it was updated.
type Info struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Container is a map of a pointer of a note that holds all the created notes with the value(the note) and key(ID)
var Container = make(map[string]*myNote)

// create is function that takes all the data to a note and creates a new one. It generates
// the Created time and the ID fields for this new note. it returns an error message if
// the creation has failed. It also puts the new note in our Container.
func Create(*myNote) error {}

/*

// find is a funcion that takes an ID of a note and searches through the Note database.
// it returns a boolean if it was found or not, and an error message if the ID was not found.
func Find(ID string) (bool, error) {}

// update is a function that takes an ID and if it was found updates it.
// it generates an upgrade time and returns an error message if it was not found or
// if the update failed
func Update(ID string) error {}

// delete is a function that takes an ID of a note as argument and if it find() it deletes
// this note from the database. it returns an error message if it was not found or if
// there was a problem by the deletion.
func Delete(ID string) error {}


*/
