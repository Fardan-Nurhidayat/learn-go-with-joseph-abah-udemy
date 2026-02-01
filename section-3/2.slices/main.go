package main

import (
	"fmt"
)

func main() {
	names := []string{"Alice", "Bob", ""}
	kelas := make([]string, 2, 5)
	kelas = append(kelas, "Charlie")
	kelas = append(kelas, "David")

	names = append(names, "Lutfi", "Ivan", "Arif")
	fmt.Printf("%+v\n", names)

	fmt.Println(len(names))
	fmt.Println(cap(names))

	bil1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bil2 := []int{11, 12, 13, 14, 15}

	bil1 = append(bil1, bil2...)
	fmt.Printf("%+v\n", bil1)

	fmt.Print("========================== \n")
	fmt.Println(kelas)

	fmt.Println("======================================")
	s := make([]int, 0, 5)

	s = append(s, 10)
	s = append(s, 20)
	s = append(s, 30)
	s = append(s, 40)
	s = append(s, 50)
	s = append(s, 60)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	fmt.Println("=============================")
	groups := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(groups)
	groups = append(groups, "G", "H", "I", "J", "K")
	fmt.Println(groups)

	fmt.Println("==================")
	result := make([]int, 0, 1000) // Preallocate capacity
	for i := 0; i < 1000; i++ {
		result = append(result, i) // Minimal reallocation
	}
}
