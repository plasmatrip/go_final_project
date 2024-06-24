package repository

import (
	"database/sql"
	"fmt"
	"log"
)

func (d *Todo) Delete(id int) error {
	res, err := d.db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		log.Printf("ошибка удаления данных: %s", err.Error())
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		log.Printf("запись с id=%d в таблице не найдена", id)
		return fmt.Errorf("запись с id=%d в таблице не найдена", id)
	}

	return nil
}
