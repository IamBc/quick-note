package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

var m *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder

func AddQuestionRoutes(r *mux.Router) {
	w := NewWriterMemory()
	this.w = &w
	handler := NewAPIHandlerREST(this.w)
	//s := r.PathPrefix("/g/").Subrouter()
	r.HandleFunc("/g/{NoteID}", handler.getNote)
	r.HandleFunc("/g/", handler.getNote)
	r.HandleFunc("/save/", handler.setNote)
	r.HandleFunc("/ddddddd/", handler.getNote)

}

func setup() {
	//mux router with added question routes
	m = mux.NewRouter()
	AddQuestionRoutes(m)

	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
}

func TestGetNoteNilHash(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("GET", "/g/", nil)
	if err != nil {
		t.Fatal("Creating 'GET' request failed!")
	}
	m.ServeHTTP(respRec, req)
	glog.Info(respRec.Code)
	if respRec.Code != http.StatusBadRequest {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}
}

func TestGetNote(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("GET", "/g/", nil)
	if err != nil {
		t.Fatal("Creating 'GET' request failed!")
	}
	req.Header.Add("xauthhash", `non-existing-hash"`)
	m.ServeHTTP(respRec, req)
	glog.Info(respRec.Code)
	if respRec.Code != http.StatusBadRequest {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}
}

func TestSetNoteNilHash(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("POST", "/save/", nil)
	if err != nil {
		t.Fatal("Creating 'POST' request failed!")
	}
	m.ServeHTTP(respRec, req)
	glog.Info(respRec.Code)
	if respRec.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}
}

func TestSetNoteHash(t *testing.T) {
	setup()

	//Testing get of non existent question type
	req, err = http.NewRequest("POST", "/save/", bytes.NewBuffer([]byte(`asd`)))
	if err != nil {
		t.Fatal("Creating 'POST' request failed!")
	}

	req.Header.Add("xauthhash", `ad`)
	m.ServeHTTP(respRec, req)
	glog.Info(respRec.Code)
	if respRec.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}
}
