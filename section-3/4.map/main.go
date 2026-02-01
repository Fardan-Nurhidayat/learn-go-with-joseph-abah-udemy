package main

import "fmt"

func main() {
	words := []string{"go", "php", "go", "java", "go", "php"}
	wordCount := make(map[string]uint8)

	for _, word := range words {
		wordCount[word]++
	}
	for word, count := range wordCount {
		fmt.Println(word, count)
	}

	studentGrades := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Printf("%+v\n", studentGrades)

	var configs = make(map[string]string)
	if configs == nil {
		fmt.Println("config is nil, initializing map")
	}

}
