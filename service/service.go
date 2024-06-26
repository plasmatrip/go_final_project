package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"

	"strconv"
	"strings"
	"time"
	"todo/configs"
	"todo/model"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	var nextDate time.Time

	startDate, err := time.Parse(configs.DateLayout, date)
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
		if len([]rune(repeat)) < 3 {
			log.Printf("неправильный формат правила повторения задачи [repeat=%s]", repeat)
			return "", fmt.Errorf("неправильный формат правила повторения задачи")
		}

		rule := strings.Split(repeat, " ")
		weekdays := strings.Split(rule[1], ",")
		targetDays := []time.Time{}

		var add int
		for _, weekday := range weekdays {
			day, err := strconv.Atoi(weekday)
			if err != nil {
				log.Printf("недопустимый символ в интервале повторения задачи [repeat=%s]", repeat)
				return "", fmt.Errorf("недопустимый символ в интервале повторения задачи")
			}

			if day > 7 {
				log.Printf("недопустимое значени в интервале повторения задачи [repeat=%s]", repeat)
				return "", fmt.Errorf("недопустимое значени в интервале повторения задачи")
			}

			if day <= int(now.Weekday()) {
				add = 7 - int(now.Weekday()) + day
			} else {
				add = day - int(now.Weekday())
			}
			targetDays = append(targetDays, now.AddDate(0, 0, add))
		}
	outW:
		for {
			for _, weekday := range targetDays {
				if nextDate.Compare(weekday) == 0 {
					break outW
				}
			}
			nextDate = nextDate.AddDate(0, 0, 1)
		}
	case 'm':
		ruleD := []string{}
		ruleM := []string{}
		days := []int{}
		months := []int{}

		rule := strings.Split(repeat, " ")
		if len(rule) < 2 {
			log.Printf("неправильный формат правила повторения задачи [repeat=%s]", repeat)
			return "", fmt.Errorf("неправильный формат правила повторения задачи")
		}

		for i, value := range rule {
			switch i {
			case 1:
				//if len([]rune(value)) > 1 {
				ruleD = append(ruleD, strings.Split(value, ",")...)
				// } else {
				// 	ruleD = append(ruleD, value)
				// }
			case 2:
				//if len([]rune(value)) > 1 {
				ruleM = append(ruleM, strings.Split(value, ",")...)
				// } else {
				// 	ruleM = append(ruleM, value)
				// }
			}
		}

		for _, d := range ruleD {
			day, err := strconv.Atoi(d)
			if err != nil {
				log.Printf("недопустимый символ в интервале повторения задачи [repeat=%s]", repeat)
				return "", fmt.Errorf("недопустимый символ в интервале повторения задачи")
			}

			if day > 31 || day == 0 || day < -31 {
				log.Printf("недопустимое значени в интервале повторения задачи [repeat=%s]", repeat)
				return "", fmt.Errorf("недопустимое значени в интервале повторения задачи")
			}

			days = append(days, day)
		}
		sort.Slice(days, func(i, j int) bool { return days[i] < days[j] })

		for _, m := range ruleM {
			month, err := strconv.Atoi(m)
			if err != nil {
				log.Printf("недопустимый символ в интервале повторения задачи [repeat=%s]", repeat)
				return "", fmt.Errorf("недопустимый символ в интервале повторения задачи")
			}

			if month < 1 || month > 31 {
				log.Printf("недопустимое значени в интервале повторения задачи [repeat=%s]", repeat)
				return "", fmt.Errorf("недопустимое значени в интервале повторения задачи")
			}

			months = append(months, month)
		}

		targetDays := []time.Time{}
		if len(months) == 0 {
			for _, day := range days {
				if day < 0 {
					if now.AddDate(0, int(now.Month()), day).After(now) {
						targetDays = append(targetDays, now.AddDate(0, int(now.Month()), day))
					} else {
						targetDays = append(targetDays, now.AddDate(0, int(now.Month())+1, day))
					}
				} else if day <= now.Day() {
					targetDays = append(targetDays, time.Date(now.Year(), now.Month(), day, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0))
				} else if day > now.Day() {
					targetDays = append(targetDays, time.Date(now.Year(), now.Month(), day, 0, 0, 0, 0, time.UTC))
				}
			}
		}

	outM:
		for {
			for _, targetDay := range targetDays {
				if nextDate.Compare(targetDay) == 0 {
					break outM
				}
			}
			nextDate = nextDate.AddDate(0, 0, 1)
		}
	default:
		log.Printf("неподдерживаемый формат [repeat=%s]", repeat)
		return "", fmt.Errorf("неподдерживаемый формат")
	}

	return nextDate.Format(configs.DateLayout), nil
}

func ErrorResponse(w http.ResponseWriter, message string, err error) {
	error := model.Error{Message: fmt.Errorf("%s [%s]", message, err.Error()).Error()}
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

func CheckTask(task *model.Task) error {
	if len(task.Id) > 0 {
		if _, err := strconv.Atoi(task.Id); err != nil {
			return errors.New("идентификатор не может быть преобразован в число")
		}
	}

	if len(task.Title) == 0 {
		return errors.New("заголовок задачи не может быть пустым")
	}

	now := time.Now()

	if len(task.Date) == 0 {
		task.Date = now.Format(configs.DateLayout)
	} else {
		date, err := time.Parse(configs.DateLayout, task.Date)
		if err != nil {
			return errors.New("дата представлена в формате, отличном от 20060102")
		}

		if date.Before(now) {
			task.Date = now.Format(configs.DateLayout)
		}
	}

	if len(task.Repeat) > 0 {
		_, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}
	return nil
}
