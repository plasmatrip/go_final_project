package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/model"
	"todo/service"
)

func (h *TodoHandlers) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	// var tasks []model.Task
	tasks, err := h.Todo.GetTasks()
	if err != nil {
		log.Printf("ошибка получения данных: %s", err.Error())
		service.ErrorResponse(w, "ошибка получения данных", err)
	}

	foundTasks, err := json.Marshal(model.Tasks{Tasks: tasks})
	if err != nil {
		log.Println(err.Error())
		service.ErrorResponse(w, "ошибка сериализации данных", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(foundTasks)
	if err != nil {
		log.Println(err.Error())
		service.ErrorResponse(w, "внутренняя ошибка сервера", err)
	}
}
