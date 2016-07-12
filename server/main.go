package main

import (
	// General
	"flag"
	"net/http"
	_ "os"
	"runtime"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

var this app

type app struct {
	APIPort string
	w       Storager
	router  *mux.Router
	handler APIHandler
}

func main() {
	flag.Parse()
	InitializeApp()
	glog.Info(http.ListenAndServe(":"+this.APIPort, this.router))
}

func InitializeApp() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	glog.Info("Initializing the API")
	this.APIPort = "7000"
	w := NewWriterMemory()
	this.w = &w
	handler := NewAPIHandlerREST(this.w)
	this.handler = &handler

	/* setup the  */
	this.router = mux.NewRouter().StrictSlash(false)
	this.router.HandleFunc("/g/{NoteID}", this.handler.getNote)
	this.router.HandleFunc("/save/", this.handler.setNote)
	this.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/")))

	//Initialize the API
	glog.Info("Listening on port: ", this.APIPort)
}
