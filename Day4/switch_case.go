package main

import (
	"fmt"
)

func main() {
	// switch
	// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，
	// 直到匹配为止。 switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break。
	// p.s. python没有switch，java的switch需要在每个case后面加break。
	// Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，
	// 而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。

	var score int
	var grade string = "None"

	fmt.Println("Please enter your score in this test: ")
	fmt.Scanln(&score)

	// check if each case is equal to score
	switch score {
	case 100:
		grade = "You fucking aced this test! 100!"
	case 0, 1, 2:
		grade = "How can you be so trash? 0!" // case can have multiple values
	}

	if grade == "None" {
		// check each condition in each case
		switch {
		case score >= 92:
			grade = "A"
		case score >= 82:
			grade = "B"
		case score >= 72:
			grade = "C"
		case score >= 62:
			grade = "D"
		case score < 62:
			grade = "F"
		}
	}

	fmt.Printf("Your score is %d, and your grade is %s", score, grade)
}
