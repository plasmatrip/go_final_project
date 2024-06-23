package main

import (
	"log"
	"net/http"
	"todo/config"
	"todo/controller"
	"todo/database"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi/v5"
)

func main() {
	config.StartLog()
	defer config.StopLog()

	config.LoadEnv()

	db := database.NewToDo()
	defer db.Close()

	r := chi.NewRouter()

	todoHandlers := controller.NewTodoHandlers(db)

	r.Mount("/", http.FileServer(http.Dir(config.WebDir)))
	r.Get("/api/nextdate", todoHandlers.NextDate)
	r.Post("/api/task", todoHandlers.AddTask)
	r.Get("/api/tasks", todoHandlers.HandleGetTasks)

	err := http.ListenAndServe(":"+config.Port, r)
	if err != nil {
		log.Panic(err)
	}
}
