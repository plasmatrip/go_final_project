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

	database.Open()
	defer database.Close()

	r := chi.NewRouter()

	r.Mount("/", http.FileServer(http.Dir(config.WebDir)))
	r.Get("/api/nextdate", controller.HandleNextDate)
	r.Post("/api/task", controller.HandleAddTask)
	r.Get("/api/task", controller.HandleGetTasks)

	err := http.ListenAndServe(":"+config.Port, r)
	if err != nil {
		log.Panic(err)
	}
}
