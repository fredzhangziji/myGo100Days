package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Please enter an integer: ")
	fmt.Scanln(&a)

	// if, else if, else
	if a < 10 {
		fmt.Printf("a's value is: %d, and it's smaller than 10.\n", a)
	} else if a == 10 {
		fmt.Printf("a's value is: %d, bingo!\n", a)
	} else {
		fmt.Printf("a's value is: %d, and it's not smaller than 10.\n", a)
	}

	// if variation
	fmt.Println()
	if num := 10; num%2 == 0 { // assign num and check if num is even
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
}
