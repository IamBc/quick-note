package main

type Storager interface {
	getNote(readOnlyHash *string, editHash *string) (n note, err error)
	setNote(newNote note) (n note, err error)
}
