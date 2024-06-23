package model

import "fmt"

type Error struct {
	Message string `json:"error"`
}

type SavedTask struct {
	Id string `json:"id"`
}

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (t Task) String() string {
	return fmt.Sprintf("id=%s date=%s title=%s comment=%s repeat=%s", t.Id, t.Date, t.Title, t.Comment, t.Repeat)
}
