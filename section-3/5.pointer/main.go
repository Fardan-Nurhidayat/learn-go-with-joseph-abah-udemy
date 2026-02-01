package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 // dereferencing
	fmt.Printf("modifyPointer: %+v\n", val)
}

func main() {
	// var a int = 42
	// var b *int = &a
	// fmt.Println("Value of a:", a)
	// fmt.Println("Address of a:", &a)
	// fmt.Println("Value of b (address of a):", b)
	// fmt.Println("Value pointed to by b:", *b)

	// *b = 100
	// fmt.Println("New value of a after modifying through b:", a)
	num := 10
	modifyValue(num)
	fmt.Println(num)

	modifyPointer(&num)
	fmt.Println(num)

	grade := 50
	gradePtr := &grade
	fmt.Printf("gradePtr grade: %+v\n", gradePtr)
	fmt.Printf("gradePtr: %+v\n", *(&gradePtr))
}
