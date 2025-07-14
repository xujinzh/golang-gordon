package main

import (
	"fmt"
)

func main() {

	var name string
	var age byte
	var salary float32
	var isPass bool

	fmt.Println("please input name, age, salary and isPass, use space split")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("name=%q, age=%d, salary=%f, isPass=%t\n", name, age, salary, isPass)
	
	// var name string
	// var age byte
	// var salary float32
	// var isPass bool
	// fmt.Printf("name: ")
	// fmt.Scanln(&name)

	// fmt.Printf("age: ")
	// fmt.Scanln(&age)
	// fmt.Printf("salary: ")
	// fmt.Scanln(&salary)
	// fmt.Printf("isPass: ")
	// fmt.Scanln(&isPass)
	// fmt.Printf("name=%q, age=%d, salary=%f, isPass=%t\n", name, age, salary, isPass)

	// var a int = 30
	// var b int = 40
	// var c int = 50
	// var max int
	// if a > b {
	// 	max = a
	// } else {
	// 	max = b
	// }

	// if max < c {
	// 	max = c
	// }
	// fmt.Printf("max(%v, %v, %v) = %v\n", a, b, c, max)

	// var a int = 30
	// var b int = 40
	// var max int
	// if a > b {
	// 	max = a
	// } else {
	// 	max = b
	// }

	// fmt.Printf("max = %v\n", max)

	// fmt.Println(10 / 4)
	// fmt.Println(10 / 4.0)

	// fmt.Println(-10 % 3) // a % b = a - a / b * b
	// fmt.Println(10 % -3)
	// fmt.Println(-10 % -3)

}
