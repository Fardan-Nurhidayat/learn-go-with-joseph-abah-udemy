package main

import "fmt"

func main() {
	var a uint8 = 10
	var b uint8 = 3

	result := func(x, y uint8) uint8 {
		return x + y
	}

	fmt.Printf("Result: %d\n", result(a, b))
}
