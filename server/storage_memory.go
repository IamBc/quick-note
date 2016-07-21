package main

import "errors"

type note struct {
	Payload      string
	IsEncrypted  bool
	ReadOnlyHash string
	EditHash     string
	NoteID       string
}

func NewWriterMemory() WriterMemory {
	var wm WriterMemory
	wm.notesEdit = make(map[string]note)
	wm.notesReadOnly = make(map[string]note)
	return wm
}

type WriterMemory struct {
	notesEdit     map[string]note
	notesReadOnly map[string]note
}

func (w *WriterMemory) getNote(editHash *string, noteID string) (n note, err error) {

	val, exists := w.notesEdit[noteID]
	if !exists {
		return val, errors.New("This note is free.")
	}

	if val.EditHash != *editHash {
		return val, errors.New(`Authentication error.`)
	}
	return val, err
}

func (w *WriterMemory) setNote(newNote note) (n note, err error) {
	w.notesEdit[newNote.NoteID] = newNote
	if newNote.ReadOnlyHash != `` {
		w.notesReadOnly[newNote.NoteID] = newNote
	}

	return n, err
}
