package main

import (
	"log"
	"net/http"
	"todo/api"
	"todo/configs"
	"todo/repository"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi/v5"
)

func main() {
	configs.StartLog()
	defer configs.StopLog()

	configs.LoadEnv()

	db := repository.NewToDo()
	defer db.Close()

	r := chi.NewRouter()

	todoHandlers := api.NewTodoHandlers(db)

	r.Mount("/", http.FileServer(http.Dir(configs.WebDir)))
	r.Get("/api/nextdate", todoHandlers.NextDate)
	r.Post("/api/task", todoHandlers.AddTask)
	r.Get("/api/task", todoHandlers.GetTask)
	r.Put("/api/task", todoHandlers.UpdateTask)
	r.Get("/api/tasks", todoHandlers.GetTasks)
	r.Post("/api/task/done", todoHandlers.TaskDone)
	r.Delete("/api/task", todoHandlers.DeleteTask)

	err := http.ListenAndServe(":"+configs.Port, r)
	if err != nil {
		log.Panic(err)
	}
}
