// Proyek Todo List sederhana dalam Go
package main

import (
	"errors"
	"fmt"
)

type Status string

const (
	planning Status = "Planning"
	ongoing  Status = "Ongoing"
	done            = "Done"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	Status      Status
}

var todos []Todo
var todoIndexByTitle map[string]int
var ID = 1
var (
	ErrorNotFound = errors.New("Resource not found")
)

func init() {
	todos = make([]Todo, 0)
	todoIndexByTitle = make(map[string]int)
}

func addTodo(title string, description string, status Status) (*Todo, error) {
	if len(title) < 3 {
		return nil, fmt.Errorf("A Title must be more than 3 characters")
	}
	if _, exists := todoIndexByTitle[title]; exists {
		return nil, fmt.Errorf("Todo Already Exists")
	}

	if !cekStatus(status) {
		return nil, fmt.Errorf("invalid status: %s", status)
	}

	newTodo := Todo{
		ID:          ID,
		Title:       title,
		Description: description,
		Status:      status,
	}
	ID++
	todos = append(todos, newTodo)
	todoIndexByTitle[title] = len(todos) - 1
	return &newTodo, nil
}

func cekStatus(s Status) bool {
	switch s {
	case planning, ongoing, done:
		return true
	default:
		return false
	}
}

func findTodoByTitle(title string) (*Todo, error) {
	index, exists := todoIndexByTitle[title]
	if exists {
		return &todos[index], nil
	}
	return nil, fmt.Errorf("Title not found: %s", title)
}

func main() {
	fmt.Println(todos)
	todo, err := addTodo("Re", "Reading book", planning)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Todo added:", todo)

	fmt.Println("======== List Todo ==========")
	fmt.Println(todos)
}
