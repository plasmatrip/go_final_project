package main

import (
	"log"
	"net/http"
	"todo/api"
	"todo/api/middleware"
	"todo/configs"
	"todo/repository"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi/v5"
)

func main() {
	configs.LoadEnv()

	configs.StartLog()
	defer configs.StopLog()

	db := repository.NewToDo()
	defer db.Close()

	r := chi.NewRouter()

	todoHandlers := api.NewTodoHandlers(db)

	r.Mount("/", http.FileServer(http.Dir(configs.WebDir)))
	r.Get("/api/nextdate", todoHandlers.NextDate)
	r.Post("/api/task", middleware.Auth(todoHandlers.AddTask))
	r.Get("/api/task", middleware.Auth(todoHandlers.GetTask))
	r.Put("/api/task", middleware.Auth(todoHandlers.UpdateTask))
	r.Get("/api/tasks", middleware.Auth(todoHandlers.GetTasks))
	r.Post("/api/task/done", middleware.Auth(todoHandlers.TaskDone))
	r.Delete("/api/task", middleware.Auth(todoHandlers.DeleteTask))
	r.Post("/api/signin", todoHandlers.Login)

	err := http.ListenAndServe(":"+configs.Port, r)
	if err != nil {
		log.Panic(err)
	}

}
