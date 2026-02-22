package main

import "fmt"

type User struct {
	Name string
}

// Tambahkan parameter 'newName' tipe string
func rename(u *User, newName string) {
	u.Name = newName
}

func updateValue(v *int, newValue int) {
	*v = newValue
}

func updateValueWithoutPointer(v int, newValue int) {
	v = newValue
}

func main() {
	var v int = 10
	fmt.Println(&v)

	updateValue(&v, 20)
	fmt.Println(v)

	updateValueWithoutPointer(v, 30)
	fmt.Println(v)

	ptr := new(string)
	fmt.Println(ptr)
	*ptr = "Hello"
	fmt.Println(*ptr)
	// user := User{
	// 	Name: "Bob",
	// }

	// fmt.Println("Sebelum:", user.Name)

	// // Panggil dengan menyertakan nama baru sebagai parameter
	// rename(&user, "Alice")
	// fmt.Println("Setelah Alice:", user.Name)

	// rename(&user, "Charlie")
	// fmt.Println("Setelah Charlie:", user.Name)
}
