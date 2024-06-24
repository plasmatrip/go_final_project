package repository

import (
	"log"
	"todo/model"
)

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
