package ui

import (
	"github.com/ac2393921/todo-go/backend/entities"
)

type Service interface {
	GetAllTodos() ([]entities.Todo, error)
}
