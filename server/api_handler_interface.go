package main

import "net/http"

type APIHandler interface {
	getNote(w http.ResponseWriter, r *http.Request)
	setNote(w http.ResponseWriter, r *http.Request)
}
