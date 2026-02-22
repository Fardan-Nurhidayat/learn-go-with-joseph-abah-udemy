package main

import (
	"fmt"
)

func main() {
	// Contoh slice dengan mendefinisikan secara langsung
	sliceExample := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(sliceExample)

	// Contoh slice dengan mendefiniskannya dengan make
	// Parameter pertama adalah tipe data, parameter kedua adalah panjang slice, parameter ketiga adalah kapasitas
	sliceExample2 := make([]string, 3)
	sliceExample2[0] = "Alice"
	sliceExample2[1] = "Bob"
	sliceExample2[2] = "Charlie"
	fmt.Println(sliceExample2)

	// Contoh slice jika mendefinisikan panjang dan kapasitas
	// Kapasitas sendiri akan berfungsi untuk mengalokasikan memori sebelumnya
	// Jadi jika panjang slice lebih dari kapasitas, maka akan terjadi realokasi memori
	// Namun, jika panjang slice kurang dari kapasitas, maka tidak akan terjadi realokasi memori
	// Hal ini akan mempengaruhi performa program , maka sebaiknya mendefinisikan kapasitas sesuai dengan kebutuhan
	sliceExample3 := make([]string, 3, 5)
	sliceExample3[0] = "Alice"
	sliceExample3[1] = "Bob"
	sliceExample3[2] = "Charlie"
	fmt.Println(sliceExample3)
	fmt.Printf("Length: %d, Capacity: %d\n", len(sliceExample3), cap(sliceExample3))

	// Contoh slice dengan append
	sliceExample4 := []string{"Alice", "Bob", "Charlie"}
	sliceExample4 = append(sliceExample4, "David")
	fmt.Println(sliceExample4)
	fmt.Println("Contoh jika tanpa mendefinisikan kapasitas")
	fmt.Printf("Length: %d, Capacity: %d\n", len(sliceExample4), cap(sliceExample4))

	// Contoh slice dengan append dan kapasitas
	sliceExample5 := make([]string, 3, 5)
	sliceExample5[0] = "Alice"
	sliceExample5[1] = "Bob"
	sliceExample5[2] = "Charlie"
	sliceExample5 = append(sliceExample5, "David")
	fmt.Println(sliceExample5)
	fmt.Println("==== Contoh dengan kapasitas ====")
	fmt.Printf("Length: %d, Capacity: %d\n", len(sliceExample5), cap(sliceExample5))

	// Contoh jika slice diprint dengan for loop
	for i, v := range sliceExample5 {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	// Contoh jika slice diprint dengan for loop tanpa index
	for _, v := range sliceExample5 {
		fmt.Printf("Value: %s\n", v)
	}

	// Contoh jika slice diprint dengan for loop tanpa value
	for i := range sliceExample5 {
		fmt.Printf("Index: %d\n", i)
	}

}
