package database

import (
	"database/sql"
)

// var db *sql.DB

var schema = `
	CREATE TABLE scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
		title VARCHAR(256) NOT NULL DEFAULT "",
		comment TEXT NOT NULL DEFAULT "",
		repeat VARCHAR(128) NOT NULL DEFAULT ""
	);
	CREATE INDEX schedule_date ON scheduler (date);
`

type Todo struct {
	db *sql.DB
}

func NewToDo() *Todo {

	return &Todo{db: open()}
}
