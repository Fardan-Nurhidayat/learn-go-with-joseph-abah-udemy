package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	k := 3
	for k < 6 {
		fmt.Println(k)
		k++
	}

	fmt.Println("------------  Skiping ------------")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\n------------ Arrays ------------")
	items := [3]string{"GO", "Jawa", "Python"}
	for index, value := range items {
		fmt.Println(index, value)
	}

	// for index := range items {
	// 	fmt.Println(index)
	// }

	fmt.Println("\n------------ Maps ------------")
	gradeStudent := map[string]float64{
		"Fardan":  90,
		"Budi":    80,
		"Joko":    70,
		"Santy":   85,
		"Michael": 95,
	}
	total := 0.0
	for _, grade := range gradeStudent {
		total += grade
	}
	fmt.Printf("Average Grade : %.2f\n", total/float64(len(gradeStudent)))

	users := map[string]int{
		"Fardan": 90,
		"Budi":   70,
		"Andi":   95,
		"Sari":   60,
	}

	maxValue := 0
	highestGrade := ""
	for name, grade := range users {
		if grade > 80 {
			fmt.Printf("Nama : %s , Grade %d \n", name, grade)

		}

		if grade > maxValue {
			maxValue = grade
			highestGrade = name
		}
	}

	fmt.Println("Nilai tertinggi: ", highestGrade)

}
