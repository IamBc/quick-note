package main

import "testing"

func TestNewWriterMemory(t *testing.T) {
	w := NewWriterMemory()
	if w.notesEdit == nil || w.notesReadOnly == nil {
		t.Fail()
	}
}

func TestSetNote(t *testing.T) {
	w := NewWriterMemory()
	_, err := w.setNote(note{Payload: "hello"})
	if err != nil {
		t.Fail()
	}
}

func TestSetNoteReadOnly(t *testing.T) {
	w := NewWriterMemory()
	_, err := w.setNote(note{Payload: "hello", ReadOnlyHash: "dadada"})
	if err != nil {
		t.Fail()
	}
}

func TestGetNoteEdit(t *testing.T) {
	w := NewWriterMemory()
	_, err := w.setNote(note{Payload: "hello", ReadOnlyHash: "dadada", EditHash: "editHash", NoteID: "aaa"})
	str := "editHash"
	_, err = w.getNote(&str, "aaa")
	if err != nil {
		t.Fail()
	}
}

func TestGetNoteNotExistsReadOnly(t *testing.T) {
	w := NewWriterMemory()
	_, err := w.getNote(nil, "aaa")
	if err == nil {
		t.Fail()
	}
}
