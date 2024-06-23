package database

import (
	"database/sql"
	"log"
	"os"
	"todo/config"
	"todo/model"
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

func NewToDo(db *sql.DB) *Todo {
	return &Todo{db: db}
}

func Open() *sql.DB {
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

func (d *Todo) Close() {
	d.Close()
}

func (d *Todo) Insert(t model.Task) (int, error) {
	res, err := d.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
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

func (d *Todo) GetTasks() ([]model.Task, error) {
	var res []model.Task

	rows, err := d.db.Query("SELECT * FROM scheduler LIMIT 25")
	if err != nil {
		log.Println(err)
		return []model.Task{}, err
	}
	defer rows.Close()

	for rows.Next() {
		t := model.Task{}
		err := rows.Scan(&t.Id, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			log.Printf("ошибка получения данных: %s", err.Error())
			return []model.Task{}, err
		}
		res = append(res, t)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ошибка итерации: %s", err.Error())
		return []model.Task{}, err
	}

	if res == nil {
		res = []model.Task{}
	}

	return res, nil
}
