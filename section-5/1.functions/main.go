package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	result := add(3, 5)
	fmt.Println("Result of addition:", result)
	// fact := factorial(5)
	// fmt.Println("Factorial of 5:", fact)
	fact := factorial(4)
	fmt.Println("Factorial of 2:", fact)
}
