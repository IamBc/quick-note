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
	str := r.Header.Get("xauthhash")
	note, err := handler.w.getNote(nil, &str)
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
	w.Write([]byte("Hello world!"))
	if r.Header.Get("xauthhash") == `` {
		http.Error(w, `Internal error. Try again later.`, http.StatusBadRequest)
		return
	}

	var newNote note
	newNote.IsEncrypted = true
	newNote.EditHash = r.Header.Get("xauthhash")
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
