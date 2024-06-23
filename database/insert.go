package database

import (
	"database/sql"
	"log"
	"todo/model"
)

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
