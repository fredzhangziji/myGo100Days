package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n') // the input in the ReadString is the delimiter
	fmt.Println("Scanned data: ", s1)
}
