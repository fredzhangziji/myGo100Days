package main

import (
	"fmt"
)

func main() {
	// 如需贯通后续的case，就添加fallthrough
	// 如果某一个case后面加了fallthrough，那么不管后面这个case有没有match它的condition，它都会进入case。

	var height float32
	var choice string
	fmt.Println("Please enter your height in cm: ")
	fmt.Scanln(&height)
	fmt.Println("Do you want to grow your height? (Y/N)")
	fmt.Scanln(&choice)

	switch choice {
	case "Y":
		fmt.Printf("Your original height is: %.2fcm\n", height)
		height += 10
		fmt.Printf("Now your height is: %.2fcm\n", height)
		fallthrough // this will make case "Y" not break but continue to next case
	case "N":
		fmt.Printf("Your original height is: %.2fcm\n", height)
		height += 5
		fmt.Printf("We make you higher anyway, now you are %.2fcm\n", height)
	}

}
