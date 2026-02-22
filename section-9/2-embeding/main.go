package main

import "fmt"

type Address struct {
	Street, City, State, ZipCode string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type ContactInfo struct {
	Email string
	Phone string
}

func (c ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name : %s\n", c.Name)
	fmt.Printf("Company Address : %s\n", c.Address.FullAddress())
	fmt.Printf("Street : %s\n", c.Street)
	fmt.Printf("City : %s\n", c.City)
	fmt.Printf("State : %s\n", c.State)
	fmt.Printf("ZipCode : %s\n", c.ZipCode)

}

type Employee struct {
	ID   int
	Name string
	ContactInfo
	Company
}

func main() {
	contact := ContactInfo{
		Email: "[EMAIL_ADDRESS]",
		Phone: "1234567890",
	}

	company := Company{
		Name: "Google",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
		ContactInfo: contact,
	}

	employee := Employee{
		ID:          1,
		Name:        "Joe",
		ContactInfo: contact,
		Company:     company,
	}

	employee.GetProfile()
	employee.DisplayContact()

}
