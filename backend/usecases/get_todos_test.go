package usecases_test

import (
	"fmt"
	"testing"

	"github.com/ac2393921/todo-go/backend/entities"
	"github.com/ac2393921/todo-go/backend/usecases"
	"github.com/gomagedon/expectate"
)

var dummyTodos = []entities.Todo{
	{
		Title:       "tpdp 1",
		Description: "description of todo 1",
		IsCompleted: true,
	},
	{
		Title:       "tpdp 2",
		Description: "description of todo 2",
		IsCompleted: false,
	},
	{
		Title:       "tpdp 3",
		Description: "description of todo 3",
		IsCompleted: true,
	},
}

type BadTodosRepo struct{}

func (BadTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("something went wrong")
}

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return dummyTodos, nil
}

func TestGetTodos(t *testing.T) {
	t.Run("Returns ErrInternal when TodoRepository returns error", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(BadTodosRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(usecases.ErrInternal)
		if todos != nil {
			t.Fatalf("Expected todos to be nil")
		}
	})

	// Test
	t.Run("Returns todos from TodosRepository", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(MockTodosRepo)

		todos, err := usecases.GetTodos(repo)
		expect(err).ToBe(nil)
		expect(todos).ToEqual((dummyTodos))
	})
}
