package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo/config"
	"todo/model"
	"todo/service"
)

func (h *TodoHandlers) AddTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	var buf bytes.Buffer

	log.Printf("получен запрос [%s]", r.RequestURI)

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		service.ErrorResponse(w, "ошибка десериализации JSON", err)
		return
	}
	log.Printf("данные в запросе: task[%s]", task.String())

	if len(task.Title) == 0 {
		log.Println("получена задача с пустым заголовком")
		service.ErrorResponse(w, "не указан заголовок задачи", errors.New("заголовок задачи не может быть пустым"))
		return
	}

	now := time.Now()

	if len(task.Date) == 0 {
		task.Date = now.Format(config.DateLayout)
	} else {
		date, err := time.Parse(config.DateLayout, task.Date)
		if err != nil {
			log.Printf("%s [task.date=%s]", err.Error(), task.Date)
			service.ErrorResponse(w, "переданное значение не может быть преобразовано в дату", errors.New("дата представлена в формате, отличном от 20060102"))
			return
		}

		if date.Before(now) {
			task.Date = now.Format(config.DateLayout)
		}
	}

	if len(task.Repeat) > 0 {
		_, err = service.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			service.ErrorResponse(w, "ошибка в данных", err)
			return
		}
	}

	id, err := h.Todo.Insert(task)
	if err != nil {
		log.Println(err.Error())
		service.ErrorResponse(w, "внутренняя ощибка сервера", err)
		return
	}

	log.Printf("добавлена задача id=%s", strconv.Itoa(id))

	savedTask, err := json.Marshal(model.SavedTask{Id: strconv.Itoa(id)})
	if err != nil {
		log.Println(err.Error())
		service.ErrorResponse(w, "ошибка сериализации данных", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(savedTask)
	if err != nil {
		log.Println(err.Error())
		service.ErrorResponse(w, "внутренняя ошибка сервера", err)
	}
}
