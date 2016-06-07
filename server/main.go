package main

import (
	// General
	"flag"
	"net/http"
	_ "os"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

var this app

func main() {
	flag.Parse()
	this.LoadSettings()

	//Initialize the API
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/g/{NoteID}", getNote)
	router.HandleFunc("/save/}", saveNote)
	glog.Info(http.ListenAndServe(":7000", router))

	/*glog.Info(`Hello world!`)
	a := GetWriterMemory()
	var testPayload string = "asd"
	var n note
	n.Payload = testPayload
	n.IsEncrypted = false

	glog.Info(`N: `, n)

	a.setNote(n)

	glog.Info(`NOTES EDIT: `, a.notesEdit)
	*/

}

/*
API Functions
They will do some maintenance work (logging/audit trail/statistics) and then pass the parameters to the Specific handler ie REST/JSONRPC and so on
Currently it's made simpler
*/
func getNote(w http.ResponseWriter, r *http.Request) {
	return
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	return
}
