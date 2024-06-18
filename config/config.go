package config

import "os"

var DBFile string
var DBDir string
var WebDir string
var Port string

func LoadEnv() {
	var exists bool

	WebDir, exists = os.LookupEnv("WEB_DIR")
	if !exists {
		panic("No environment variable WEB_DIR")
	}

	Port, exists = os.LookupEnv("TODO_PORT")
	if !exists {
		panic("No environment variable TODO_PORT")
	}

	DBFile, exists = os.LookupEnv("TODO_DBFILE")
	if !exists {
		panic("No environment variable TODO_DBFILE")
	}

	DBDir, exists = os.LookupEnv("TODO_DB_DIR")
	if !exists {
		panic("No environment variable TODO_DB_DIR")
	}
}
