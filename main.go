package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "modernc.org/sqlite"

	"go_final_project/config"
	"go_final_project/database"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	config.LoadEnv()
}

func handleNextDay(res http.ResponseWriter, req *http.Request) {

}

func main() {
	db, err := sql.Open("sqlite", config.DBFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	database.Check(db)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(config.WebDir)))
	mux.HandleFunc("/api/nextdate", handleNextDay)

	err = http.ListenAndServe(":"+config.Port, mux)
	if err != nil {
		panic(err)
	}
}
