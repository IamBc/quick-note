package main

import "errors"

type note struct {
	Payload      string
	IsEncrypted  bool
	ReadOnlyHash string
	EditHash     string
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

func (w *WriterMemory) getNote(readOnlyHash *string, editHash *string) (n note, err error) {
	if readOnlyHash != nil && editHash != nil {
		return n, errors.New("Only one argument can be defined.")
	}

	if readOnlyHash != nil && *readOnlyHash != "" {
		val, exists := w.notesReadOnly[*readOnlyHash]
		if !exists {
			return val, errors.New(`Key with readOnlyHash: ` + *readOnlyHash + ` doesn't exist.`)
		}
		return val, err
	} else if editHash != nil && *editHash != "" {
		val, exists := w.notesEdit[*editHash]
		if !exists {
			return val, errors.New(`Key with editHash: ` + *editHash + ` doesn't exist.`)
		}
		return val, err
	}

	return n, errors.New("System error.")
}

func (w *WriterMemory) setNote(newNote note) (n note, err error) {
	w.notesEdit[newNote.EditHash] = newNote
	if newNote.ReadOnlyHash != `` {
		w.notesReadOnly[newNote.ReadOnlyHash] = newNote
	}

	return n, err
}
