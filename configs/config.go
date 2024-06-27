package configs

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
var SearchLayout string

var logFile *os.File

func LoadEnv() {
	var exists bool

	if err := godotenv.Load("./.env"); err != nil {
		log.Println("не найден .env файл")
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

	SearchLayout, exists = os.LookupEnv("TODO_SEARCH_LAYOUT")
	if !exists {
		log.Fatal("не найдена переменная окружения TODO_SEARCH_LAYOUT")
	}

	_, exists = os.LookupEnv("TODO_PASSWORD")
	if !exists {
		log.Fatal("не найдена переменная окружения TODO_PASSWORD")
	}

	_, exists = os.LookupEnv("APP_LOG_DIR")
	if !exists {
		log.Fatal("не найдена переменная окружения APP_LOG_DIR")
	}

	_, exists = os.LookupEnv("APP_LOG_FILE")
	if !exists {
		log.Fatal("не найдена переменная окружения APP_LOG_FILE")
	}
}

func StartLog() {
	logDir, _ := os.LookupEnv("APP_LOG_DIR")
	logFile, _ := os.LookupEnv("APP_LOG_FILE")

	if _, err := os.Stat(logDir); err != nil {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Println("не удалось создать каталог для log-файла", err)
		}
	}
	l, err := os.OpenFile(logDir+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("не удалось открыть файл ", err)
	}
	log.SetOutput(l)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("логирование начато")
}

func StopLog() {
	log.Println("логирование окончено")
	logFile.Close()
}
