package main

import (
	// General
	"flag"
	"net/http"
	_ "os"
	"strconv"

	"github.com/golang/glog"
)

var this app

func main() {
	configFile := flag.String("config_file", "", "The configuration file for the project")
	flag.Parse()
	this.GetConfiguration(*configFile)
	this.InitializeApp()
	glog.Info(http.ListenAndServe(":"+strconv.Itoa(this.config.BackendPort), this.router))
}
