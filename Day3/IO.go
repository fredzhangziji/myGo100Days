package main

import (
	"fmt"
)

func main() {
	// 	格式化打印占位符：
	// 	%v, 原样输出
	// 	%T，打印类型
	// 	%t, bool类型
	// 	%s，字符串
	// 	%f，浮点
	// 	%d，10进制的整数
	// 	%b，2进制的整数
	// 	%o，8进制
	// 	%x，%X，16进制
	// 		%x：0-9，a-f
	// 		%X：0-9，A-F
	// 	%c，打印字符
	// 	%p，打印地址
	// 	......

	a := 100           //int
	b := 3.14          //float64
	c := true          // bool
	d := "Hello World" //string
	e := `Ruby`        //string
	f := 'A'
	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)
	fmt.Printf("%T,%d,%c\n", f, f, f)
	fmt.Println("-----------------------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

	// Scan Scanf Scanln
	var x int
	var y float64
	fmt.Println("Please enter an integer, and a floating number")
	fmt.Scanln(&x, &y) // 读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	fmt.Printf("x's value: %d, y's value: %f\n", x, y)

	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d, y:%f\n", x, y)

}
