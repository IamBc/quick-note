package main

func GetWriterMemory() WriterMemory {
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

	return n, err
}

func (w *WriterMemory) setNote(newNote note) (n note, err error) {
	w.notesEdit[newNote.EditHash] = newNote
	if newNote.ReadOnlyHash != nil {
		w.notesReadOnly[*newNote.ReadOnlyHash] = newNote
	}

	return n, err
}
