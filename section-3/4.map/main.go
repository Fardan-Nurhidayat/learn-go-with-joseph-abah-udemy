package main

import "fmt"

func main() {
	// Map adalah kumpulan pasangan key-value
	// Map adalah unordered collection
	// Map adalah mutable
	// Map adalah reference type
	// Map adalah unordered collection
	// Map adalah unordered collection

	// Cara membuat map
	// dengan opsi make
	mapExample := make(map[string]string)
	mapExample["name"] = "John"
	mapExample["age"] = "30"
	mapExample["city"] = "New York"
	fmt.Println(mapExample)

	// dengan opsi literal
	mapExample2 := map[string]string{
		"name": "John",
		"age":  "30",
		"city": "New York",
	}
	fmt.Println(mapExample2)

	// Cara mengakses map
	fmt.Println(mapExample["name"])
	fmt.Println(mapExample["age"])
	fmt.Println(mapExample["city"])

	// Cara menghapus map
	delete(mapExample, "name")
	fmt.Println(mapExample)

	// Cara mengecek apakah map kosong
	fmt.Println(len(mapExample))

	// Cara mengecek apakah key ada di map
	_, ok := mapExample["name"]
	fmt.Println(ok)

	// Cara iterasi map
	for key, value := range mapExample {
		fmt.Println(key, value)
	}

}
