package usecases

import "github.com/ac2393921/todo-go/backend/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
