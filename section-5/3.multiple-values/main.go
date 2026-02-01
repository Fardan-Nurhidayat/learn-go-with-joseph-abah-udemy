package main

import (
	"errors"
	"fmt"
)

type Users struct {
	ID   int
	Name string
	Age  int
}

var users []Users
var userIndexByName map[string]int

func init() {
	users = make([]Users, 0)
	userIndexByName = make(map[string]int)
}

func calculateStats(numbers []int) (int, int, float64) {
	if len(numbers) == 0 {
		return 0, 0, 0.0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	avg := float64(sum) / float64(len(numbers))
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return sum, max, avg
}

func substract(a, b uint) (result uint) {
	result = a * b
	return
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func addUser(name string, age int) Users {
	newUser := Users{
		ID:   len(users) + 1,
		Name: name,
		Age:  age,
	}
	users = append(users, newUser)
	userIndexByName[name] = len(users) - 1
	return newUser
}

func findUser(name string) (*Users, error) {
	index, exists := userIndexByName[name]
	if exists {
		return &users[index], nil
	}
	return nil, fmt.Errorf("user not found: %s", name)
}

func main() {
	// numbers := []int{10, 20, 30, 40, 50}
	// total, maximum, average := calculateStats(numbers)
	// fmt.Printf("Total: %d, Maximum: %d, Average: %.2f\n", total, maximum, average)

	// fmt.Println("=======================")
	// result := substract(10, 5)
	// fmt.Printf("Result: %d\n", result)

	// fmt.Println("===========================")
	// divResult, err := divide(10, 2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Division Result: %d\n", divResult)
	// }

	fmt.Println("===========================")
	fmt.Println(users)

	fmt.Println("=======================")
	addUser("Alice", 30)
	addUser("Bob", 30)

	fmt.Println("======================")
	fmt.Println(users)

	user, err := findUser("LUtfi")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found User: %+v\n", *user)
	}
}
