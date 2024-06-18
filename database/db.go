package database

import (
	"database/sql"
	"log"
	"os"
	"todo/config"
)

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

func Check(db *sql.DB) {
	if _, err := os.Stat(config.DBFile); err != nil {
		if _, err := os.Stat(config.DBDir); err != nil {
			if err := os.Mkdir(config.DBDir, 0755); err != nil {
				log.Fatal(err)
			}
		}
		_, err = db.Exec(schema)
		if err != nil {
			panic(err)
		}
	}
}
