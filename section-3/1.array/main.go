package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)
	numbers[0] = 1
	numbers[1] = 2

	fmt.Printf("%+v\n", numbers)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("%+v\n", primes)

	for i, value := range primes {
		fmt.Printf("Index : %d , Value : %d \n", i, value)
	}

	// Matrix
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	fmt.Printf("%+v\n", matrix)

	names := [7]string{"Lutfi", "Ivan", "Arif", "Dani", "Said", "Budi", "Joko"}
	fmt.Println(names[1:5])
	fmt.Println(names[:3])
	fmt.Println(names[2:])
}
