package database

import (
	"database/sql"
	"log"
	"os"
	"todo/config"
	"todo/model"
)

var db *sql.DB

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

func Open() {
	var err error
	db, err = sql.Open("sqlite", config.DBFile)
	if err != nil {
		log.Fatal(err)
		return
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
}

func Close() {
	db.Close()
}

func Insert(t model.Task) (int, error) {
	res, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat))
	if err != nil {
		log.Printf("ошибка добавления в базу данных: %s", err.Error())
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("ошибка получения id записи: %s", err.Error())
		return 0, err
	}

	return int(id), nil
}

func GetTasks() ([]model.Task, error) {
	var res []model.Task

	rows, err := db.Query("SELECT * FROM scheduler LIMIT 25")
	if err != nil {
		log.Println(err)
		return res, err
	}
	defer rows.Close()

	return res, nil
}
