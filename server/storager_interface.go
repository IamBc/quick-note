package main

type Storager interface {
	getNote(editHash *string, noteID string) (n note, err error)
	setNote(newNote note) (n note, err error)
}
