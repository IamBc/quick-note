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
		return
	}
	//glog.Info(r.Header.Get("xedithash"))
	//glog.Info(r.Header.Get("xreadonlyhash"))
	str := r.Header.Get("xauthhash")
	note, err := handler.w.getNote(nil, &str)
	if err != nil {
		glog.Error(err)
	}
	w.Write([]byte(note.Payload))
}

func (handler *APIHandlerREST) setNote(w http.ResponseWriter, r *http.Request) {

	glog.Info(r.Header.Get("xauthhash"))
	glog.Info(r.Header.Get("xedithash"))
	glog.Info(r.Header.Get("xreadonlyhash"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, xauthhash")
	//TODO check for empty note
	glog.Info(`!!!!!!!!!!1setNote: xauthhash`, r.Header.Get("xauthhash"), `request body: `, r.Body)
	w.Write([]byte("Hello world!"))
	if r.Header.Get("xauthhash") == `` {
		return
	}

	var newNote note
	newNote.IsEncrypted = true
	newNote.EditHash = r.Header.Get("xauthhash")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Info(`Cloud not parse request payload: `, err)
		return
	}
	newNote.Payload = string(body)
	glog.Info("GOING TO SET THE NEWNOTE111")
	handler.w.setNote(newNote)
}
