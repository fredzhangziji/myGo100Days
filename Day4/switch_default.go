package main

import (
	"fmt"
)

func main() {
	// if none of the case is matched in the switch, then the default one will be entered (if there is any)
	// it's like the else statement in if-else condition

	var condition string
	var result string
	fmt.Println("Please enter the condition.")
	fmt.Scanln(&condition)

	switch condition {
	default:
		result = "good"
	case "bad":
		result = "bad"
	case "good":
		result = "good"
	}
	fmt.Println(result)
}
