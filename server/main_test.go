package main

import "testing"

func TestMain(t *testing.T) {
	this.InitializeApp()
}

func TestConfigNoSuchFile(t *testing.T) {
	a := "asdasdadasd"
	this.GetConfiguration(a)
}

func TestConfigWrongJSON(t *testing.T) {
	a := "../example-config.json"
	this.GetConfiguration(a)
}

func TestConfigOK(t *testing.T) {
	a := "../example-wrong.json"
	this.GetConfiguration(a)
}
