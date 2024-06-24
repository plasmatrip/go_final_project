package api

import (
	"todo/repository"
)

type TodoHandlers struct {
	Repo *repository.Todo
}

func NewTodoHandlers(todo *repository.Todo) *TodoHandlers {
	return &TodoHandlers{Repo: todo}
}
