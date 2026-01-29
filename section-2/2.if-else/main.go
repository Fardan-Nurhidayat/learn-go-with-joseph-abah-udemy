package main

import "fmt"

func main() {
	tmp := 40

	if tmp > 35 {
		fmt.Println("Tmp greater than 35")
	}

	score := 80

	if score > 85 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else if score >= 65 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("You failed")
	}

	userAccess := map[string]bool{
		"Jane": true,
		"John": false,
	}

	if hasAccess, ok := userAccess["John"]; ok && hasAccess {
		fmt.Println("Jane can access the system")
	} else {
		fmt.Println("Access not guaranted")
	}

	users := map[string]int{
		"Fardan": 90,
		"Budi":   70,
	}

	if value, ok := users["Fardan"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("User not found")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Interger : %d\n", v)
		case string:
			fmt.Printf("String : %s\n", v)
		case bool:
			fmt.Printf("Boolean : %t\n", v)
		default:
			fmt.Println("Unknown Type")
		}
	}

	checkType(10)
	checkType("Golang")
	checkType(true)
	checkType(3.14)
}
