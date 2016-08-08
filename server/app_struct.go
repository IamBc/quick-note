package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

type app struct {
	w       Storager
	router  *mux.Router
	handler APIHandler
	Config  Config
}

type Config struct {
	MaxNoteCount  int
	MaxNoteLength int
	Endpoint      string

	BackendPort int
	SSLCertPath string
}

func (this *app) InitializeApp() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	glog.Info("Initializing the API")
	w := NewWriterMemory()
	this.w = &w
	handler := NewAPIHandlerREST(this.w, &this.Config)
	this.handler = &handler

	this.router = mux.NewRouter().StrictSlash(false)
	this.router.HandleFunc("/g/{NoteID}", this.handler.getNote)
	this.router.HandleFunc("/save/", this.handler.setNote)
	this.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/")))

	glog.Info("Listening on port: ", this.Config.BackendPort)
}

func (this *app) GetConfiguration(configFile string) error {
	dat, err := ioutil.ReadFile(configFile)
	if err != nil {
		glog.Error("Cloud not open file: ", err.Error())
		return err
	}

	err = json.Unmarshal(dat, &this.Config)
	if err != nil {
		glog.Error("Cloud not unmarshal file: ", err.Error())
		return err
	}
	return nil
}
