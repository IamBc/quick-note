package main

type note struct {
	Payload      string
	IsEncrypted  bool
	ReadOnlyHash *string
	EditHash     string
}
