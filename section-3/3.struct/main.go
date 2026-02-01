package main

import "fmt"

type Gender int

const ( 
	Male Gender = iota
	Female
)

type Person struct {
	Name   string
	Age    int
	Gender Gender
}


func main() {
	
	p := &Person{
		Name : "Arif",
		Age  : 21,
		Gender : Male,
	}


	fmt.Printf("%+v\n", p)
}