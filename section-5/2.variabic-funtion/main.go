package main

import "fmt"

func add(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	sum1 := add(1, 2, 3)
	sum2 := add(10, 20, 30, 40, 50)
	fmt.Println(sum1)
	fmt.Println(sum2)

}
