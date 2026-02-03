package main

import "fmt"

func main() {
	i := 1
Next: // label is declared.
	fmt.Println(i * 2)
	i++
	if i <= 10 {
		goto Next // execution jumps
	}
}
