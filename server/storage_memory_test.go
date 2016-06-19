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

func TestGetNoteReadOnly(t *testing.T) {
	w := NewWriterMemory()
	_, err := w.setNote(note{Payload: "hello", ReadOnlyHash: "dadada"})
	str := "dadada"
	_, err = w.getNote(&str, nil)
	if err != nil {
		t.Fail()
	}
}

func TestGetNoteEdit(t *testing.T) {
	w := NewWriterMemory()
	_, err := w.setNote(note{Payload: "hello", ReadOnlyHash: "dadada", EditHash: "editHash"})
	str := "editHash"
	_, err = w.getNote(nil, &str)
	if err != nil {
		t.Fail()
	}
}

func TestGetNoteNotExistsReadOnly(t *testing.T) {
	w := NewWriterMemory()
	str := "dadada"
	_, err := w.getNote(&str, nil)
	if err == nil {
		t.Fail()
	}
}

func TestGetNoteNotExistsEdit(t *testing.T) {
	w := NewWriterMemory()
	str := "editHash"
	_, err := w.getNote(nil, &str)
	if err == nil {
		t.Fail()
	}
}

func TestGetNoteNotWrongInpArgs(t *testing.T) {
	w := NewWriterMemory()
	str := "editHash"
	_, err := w.getNote(&str, &str)
	if err == nil {
		t.Fail()
	}
}
