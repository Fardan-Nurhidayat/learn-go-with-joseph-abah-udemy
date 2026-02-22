package main

import "fmt"

type Address struct {
	Street, City, State, ZipCode string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	ID              int
	Name            string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetail() {
	fmt.Println("Customer ID:", c.ID)
	fmt.Println("Customer Name:", c.Name)
	fmt.Println("Billing Address:", c.BillingAddress.FullAddress())
	fmt.Println("Shipping Address:", c.ShippingAddress.FullAddress())
}

func main() {
	cus1 := Customer{
		ID:   1,
		Name: "Joe",
		BillingAddress: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
		ShippingAddress: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
	}

	cus1.PrintDetail()

	mainAddress := Address{
		Street:  "123 Main St",
		City:    "New York",
		State:   "NY",
		ZipCode: "10001",
	}

	cus2 := Customer{
		ID:              2,
		Name:            "Jane",
		BillingAddress:  mainAddress,
		ShippingAddress: mainAddress,
	}

	cus2.PrintDetail()
}
