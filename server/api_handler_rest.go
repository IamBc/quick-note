package main

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func NewAPIHandlerREST(w Storager) APIHandlerREST {
	var handler APIHandlerREST
	handler.w = w
	return handler
}

type APIHandlerREST struct {
	w Storager
}

func (handler *APIHandlerREST) getNote(w http.ResponseWriter, r *http.Request) {

	glog.Info(`GetNote xauthhash: `, r.Header.Get("xauthhash"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, xauthhash")

	if r.Header.Get("xauthhash") == `` {
		glog.Info(`getNote: xauthash is empty!`)
		http.Error(w, `Non existant note or wrong authentication.`, http.StatusBadRequest)
		return
	}

	if r.Header.Get("xnoteid") == `` {
		glog.Info(`getNote: xnoteid is empty!`)
		http.Error(w, `Non existant note or wrong authentication.`, http.StatusBadRequest)
		return
	}

	xauthhash := r.Header.Get("xauthhash")
	note, err := handler.w.getNote(&xauthhash, r.Header.Get("xnoteid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Error(`Cloud not fetch the note: `, err.Error())
		return
	}
	w.Write([]byte(note.Payload))
}

func (handler *APIHandlerREST) setNote(w http.ResponseWriter, r *http.Request) {

	glog.Info(r.Header.Get("xauthhash"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, xauthhash")
	glog.Info(`setNote: xauthhash`, r.Header.Get("xauthhash"), `request body: `, r.Body)
	if r.Header.Get("xauthhash") == `` {
		http.Error(w, `Bad request parameters!`, http.StatusBadRequest)
		return
	}

	if r.Header.Get("xnoteid") == `` {
		http.Error(w, `Bad request parameters!`, http.StatusBadRequest)
		return
	}

	xauthhash := r.Header.Get("xauthhash")
	_, err := handler.w.getNote(&xauthhash, r.Header.Get("xnoteid"))
	if err != nil && err.Error() != "This note is free. Save it to make it yours." {
		glog.Info("Note is already taken")
		http.Error(w, `This note is already taken.`, http.StatusBadRequest)
	}

	var newNote note
	newNote.IsEncrypted = true
	newNote.EditHash = r.Header.Get("xauthhash")
	newNote.NoteID = r.Header.Get("xnoteid")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Info(`Cloud not parse request payload: `, err)
		http.Error(w, `Internal error. Try again later.`, http.StatusInternalServerError)
		return
	}
	newNote.Payload = string(body)
	// Currently setNote doesn't return any errors. Only possible error is not
	//enough memory which would cause the OS to kill the process anyways...
	handler.w.setNote(newNote)
	w.Write([]byte(`OK`))
}
