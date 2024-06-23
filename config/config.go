package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBFile string
var DBDir string
var WebDir string
var Port string
var DateLayout string

var logFile *os.File

func LoadEnv() {
	var exists bool

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("не найден .env файл")
	}

	WebDir, exists = os.LookupEnv("WEB_DIR")
	if !exists {
		log.Fatal("не найдена переменная окружения WEB_DIR")
	}

	Port, exists = os.LookupEnv("TODO_PORT")
	if !exists {
		log.Fatal("не найдена переменная окружения TODO_PORT")
	}

	DBFile, exists = os.LookupEnv("TODO_DBFILE")
	if !exists {
		log.Fatal("не найдена переменная окружения TODO_DBFILE")
	}

	DBDir, exists = os.LookupEnv("TODO_DB_DIR")
	if !exists {
		log.Fatal("не найдена переменная окружения TODO_DB_DIR")
	}

	DateLayout, exists = os.LookupEnv("TODO_DATE_LAYOUT")
	if !exists {
		log.Fatal("не найдена переменная окружения TODO_DATE_LAYOUT")
	}
}

func StartLog() {
	logFile, err := os.OpenFile("../log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("не удалось открыть файл ", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("логирование начато")
}

func StopLog() {
	log.Println("логирование окончено")
	logFile.Close()
}
