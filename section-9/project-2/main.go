package main

import "fmt"

type Tenantable interface {
	fmt.Stringer
	CalculateValue(hours float64) float64
}

type Car struct {
	Name        string
	HourlyPrice float64
}

func (c Car) CalculateValue(hours float64) float64 {
	return c.HourlyPrice * hours
}

func (c Car) String() string {
	return fmt.Sprintf("%s (Hourly price: $%.2f)", c.Name, c.HourlyPrice)
}

type Motor struct {
	Name        string
	HourlyPrice float64
}

func (m Motor) CalculateValue(hours float64) float64 {
	return m.HourlyPrice * hours
}

func (m Motor) String() string {
	return fmt.Sprintf("%s (Hourly price: $%.2f)", m.Name, m.HourlyPrice)
}

func PrintRentSummary[P fmt.Stringer](tenant P) {
	fmt.Printf(" - Processing: %s\n", tenant)
}

func ProcessPaying(tenans []Tenantable, hours float64) {
	fmt.Printf("\n--- Processing Payment for %.1f Hours\n", hours)
	totalPaying := 0.0
	for _, ten := range tenans {
		PrintRentSummary(ten)
		pay := ten.CalculateValue(hours)
		fmt.Printf("    Rent Fee: $%.2f\n", pay)
		totalPaying += pay
	}
	fmt.Printf("\nTotal Rental Paying: $%.2f\n", totalPaying)
	fmt.Println("--------------------------")
}

func main() {
	fmt.Println("Welcome to the Rent Processor!")

	car1 := Car{Name: "Toyota", HourlyPrice: 10.0}
	motor1 := Motor{Name: "Honda", HourlyPrice: 5.0}

	rentList := []Tenantable{
		car1,
		motor1,
		Car{Name: "Civic", HourlyPrice: 12.0},
		Motor{Name: "Yamaha", HourlyPrice: 6.0},
	}

	ProcessPaying(rentList, 2.5) // Example: renting for 2.5 hours
}
