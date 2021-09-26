package main

import (
	"fmt"
)

func main() {
	// var名称类型是声明单个变量的语法。
	// first: designate the type
	var variable_A int
	variable_A = 10
	println("This is variable_A:", variable_A)

	// second: give the value so the type is recognized
	var variable_B = 10.0
	fmt.Printf("This is variable_B: %.2f\n", variable_B)

	// third: ignore 'var'
	variable_C := "10"
	println("This is variable_C:", variable_C)
	// := means this is a new variable
	// = means that we are assigning a value to an existing variable

	// multiple variables
	// first
	var v1, v2, v3 int
	v1, v2, v3 = 1, 2, 3
	println("This is v1v2v3: ", v1, v2, v3)

	// second
	var v4, v5, v6 = 4, 5, 6
	println("This is v4v5v6: ", v4, v5, v6)

	// third
	var (
		v7 int = 7
		v8 int = 8
	)
	println("This is v7v8: ", v7, v8)

	// constant
	const const1 string = "abc" // 显式类型定义
	const const2 = "abc"        // 隐式类型定义
	const (
		Unknown = 0
		Male    = 1
		Female  = 2
	)
	// 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		x uint16 = 16
		y
		s = "abc"
		z
	)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	// output
	// uint16,16
	// string,abc

	// iota: 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
