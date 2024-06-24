package repository

import (
	"database/sql"
	"fmt"
	"log"
	"todo/model"
)

func (d *Todo) Update(t model.Task) error {
	res, err := d.db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
		sql.Named("id", t.Id),
	)
	if err != nil {
		log.Printf("ошибка обновления данных: %s", err.Error())
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		log.Printf("запись с id=%s в таблице не найдена", t.Id)
		return fmt.Errorf("запись с id=%s в таблице не найдена", t.Id)
	}

	return nil
}
