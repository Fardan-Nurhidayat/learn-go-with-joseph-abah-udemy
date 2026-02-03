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

func updateStatus(title string, status Status) (*Todo, error) {
	index, exists := todoIndexByTitle[title]
	if !exists {
		return nil, ErrorNotFound
	}
	if !cekStatus(status) {
		return nil, fmt.Errorf("invalid status: %s", status)
	}
	todos[index].Status = status
	return &todos[index], nil
}

func deleteTodo(title string) (*Todo, error) {
	index, exists := todoIndexByTitle[title]
	if !exists {
		return nil, ErrorNotFound
	}
	last := len(todos) - 1
	todos[index] = todos[last]
	todoIndexByTitle[todos[index].Title] = index
	todos = todos[:last]
	delete(todoIndexByTitle, title)
	return &todos[index], nil

}

func main() {
	fmt.Println(todos)
	todo, err := addTodo("Read Book", "Reading book", planning)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	todo1, err := addTodo("Go Lang", "Learn Go Language", ongoing)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Todo added:", todo)
	fmt.Println("Todo Added", todo1)
	fmt.Println("======== List Todo ==========")
	fmt.Println(todos)

	fmt.Println("=========== Search Todo ============")
	foundTodo, err := findTodoByTitle("Read Book")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Found Todo:", foundTodo)

	fmt.Println("=========== Update Status =============")
	updatedTodo, err := updateStatus("Read Book", ongoing)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Updated Todo:", updatedTodo)
	fmt.Println("======== Final List Todo ==========")
	fmt.Println(todos)

	fmt.Println("=========== Delete Todo =============")
	deletedTodo, err := deleteTodo("Read Book")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deleted Todo:", deletedTodo)
	fmt.Println("======== Final List Todo ==========")
	fmt.Println(todos)
}
