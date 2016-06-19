package main

import (
	"net/http"

	"github.com/golang/glog"
)

func NewAPIHandlerREST() APIHandlerREST {
	var handler APIHandlerREST
	return handler
}

type APIHandlerREST struct {
}

/*TODO Business logic should be in another layer*/

func (handler *APIHandlerREST) getNote(w http.ResponseWriter, r *http.Request) {
	glog.Info("aaaaaaaaaaaaaaaaaaaaaaaaaa")
	glog.Info(r.Header.Get("xauthhash"))
	glog.Info(r.Header.Get("xedithash"))
	glog.Info(r.Header.Get("xreadonlyhash"))
	w.Write([]byte("hello world!"))
}

func (handler *APIHandlerREST) setNote(w http.ResponseWriter, r *http.Request) {

}
