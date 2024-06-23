package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"strconv"
	"strings"
	"time"
	"todo/config"
	"todo/model"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	var nextDate time.Time

	startDate, err := time.Parse(config.DateLayout, date)
	if err != nil {
		log.Printf("%s [date=%s]", err.Error(), date)
		return "", fmt.Errorf("переданное значение не может быть преобразовано в дату")
	}

	nextDate = startDate

	if len(repeat) == 0 {
		log.Printf("правило повторения задачи пустое")
		return "", fmt.Errorf("правило повторения задачи пустое")
	}

	switch repeat[0] {
	case 'd':
		if len([]rune(repeat)) < 3 {
			log.Printf("неправильный формат правила повторения задачи [repeat=%s]", repeat)
			return "", fmt.Errorf("неправильный формат правила повторения задачи")
		}

		rule := strings.Split(repeat, " ")

		days, err := strconv.Atoi(rule[1])
		if err != nil {
			log.Printf("недопустимый символ в интервале повторения задачи [repeat=%s]", repeat)
			return "", fmt.Errorf("недопустимый символ в интервале повторения задачи")
		}

		if days > 400 {
			log.Printf("превышен максимально допустимый интервал в правиле повторения задачи [repeat=%s]", repeat)
			return "", fmt.Errorf("превышен максимально допустимый интервал в правиле повторения задачи")
		}

		for nextDate = nextDate.AddDate(0, 0, days); nextDate.Before(now); {
			nextDate = nextDate.AddDate(0, 0, days)
		}
	case 'y':
		for nextDate = nextDate.AddDate(1, 0, 0); nextDate.Before(now); {
			nextDate = nextDate.AddDate(1, 0, 0)
		}
	case 'w':
		log.Printf("неподдерживаемый формат правила повторения задачи [repeat=%s]", repeat)
		return "", fmt.Errorf("неподдерживаемый формат правила повторения задачи")
	case 'm':
		log.Printf("неподдерживаемый формат правила повторения задачи [repeat=%s]", repeat)
		return "", fmt.Errorf("неподдерживаемый формат правила повторения задачи")
	default:
		log.Printf("неподдерживаемый формат [repeat=%s]", repeat)
		return "", fmt.Errorf("неподдерживаемый формат")
	}

	return nextDate.Format(config.DateLayout), nil
}

func ErrorResponse(w http.ResponseWriter, message string, err error) {
	error := model.Error{Message: fmt.Errorf("%s, %s", message, err.Error()).Error()}
	resp, err := json.Marshal(error)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
