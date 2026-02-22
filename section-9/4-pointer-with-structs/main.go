package main

import "fmt"

type Person struct {
	name string
	age  int8
}

func updateAge(p *Person, newAge int8) {
	p.age = newAge
}

func main() {
	PersonPointer := new(Person)
	PersonPointer.name = "John"
	PersonPointer.age = 30

	fmt.Printf("Person name : %s\n", PersonPointer.name)
	fmt.Printf("Person age : %d\n", PersonPointer.age)

	// Updated Person Age
	updateAge(PersonPointer, 35)
	fmt.Printf("Updated age : %d\n", PersonPointer.age)
}
