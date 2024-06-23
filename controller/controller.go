package controller

import (
	"todo/database"
)

type TodoHandlers struct {
	Todo *database.Todo
}

func NewTodoHandlers(todo *database.Todo) *TodoHandlers {
	return &TodoHandlers{Todo: todo}
}
