package database

import (
	"database/sql"
	"log"
	"os"
	"todo/config"
)

func open() *sql.DB {
	var err error
	db, err := sql.Open("sqlite", config.DBFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if _, err := os.Stat(config.DBFile); err != nil {
		if _, err := os.Stat(config.DBDir); err != nil {
			if err := os.Mkdir(config.DBDir, 0755); err != nil {
				log.Fatal(err)
			}
		}
		_, err = db.Exec(schema)
		if err != nil {
			log.Panic(err)
		}
	}
	return db
}
