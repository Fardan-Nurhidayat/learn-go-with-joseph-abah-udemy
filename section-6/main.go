package main

import "fmt"

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name string
	Age  int
	Address
}

// Emebeding struct adalah cara untuk menyisipkan satu struct ke dalam struct lainnya tanpa harus mendeklarasikan field secara eksplisit.

func main() {
	p := Person{
		Name: "Luna",
		Age:  28,
		Address: Address{
			Street:  "Jalan Merpati No 10",
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}
	fmt.Println("Name :", p.Name)
	fmt.Println("Age :", p.Age)
	fmt.Println("Street :", p.Street)
	fmt.Println("City :", p.City)
	fmt.Println("Country :", p.Country)

	fmt.Println("Full Address :", p.FullAddress())
	fmt.Println("Full Person Info :", p.FullPersonInfo())
	p.UpdateCity("Tasikmalaya")
	fmt.Println("Updated City :", p.City)
	fmt.Println("Full Address after update :", p.FullAddress())
}

// Inheritance-like Behavior
func (a *Address) FullAddress() string {
	return a.Street + ", " + a.City + ", " + a.Country
}

func (p Person) FullPersonInfo() string {
	return fmt.Sprintf("%s, %d years old, lives at %s", p.Name, p.Age, p.FullAddress())
}

func (a *Address) UpdateCity(newCity string) {
	a.City = newCity
}
