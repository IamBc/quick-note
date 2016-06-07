package main

import "github.com/golang/glog"

type note struct {
	Payload      string
	IsEncrypted  bool
	ReadOnlyHash *string
	EditHash     string
}

type app struct {
	APIPort string
	w       Storager
}

func (app *app) LoadSettings() {
	glog.Info(`LoadSettings`)
}
