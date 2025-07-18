package main

import (
	"fmt"
)

var (
	FuncVar = func(a, b int) int {
		return a + b
	}
)

func main() {
	// 使用匿名函数计算两个数的和
	res := func(a, b int) int {
		return a + b
	}(3, 5)
	fmt.Println("The sum is:", res)

	// 将匿名函数赋值给变量，可以打破在函数中定义函数
	sumFuncVar := func(a, b int) int {
		return a + b
	}
	fmt.Println("The sum using variable is:", sumFuncVar(10, 20))

	// 使用预定义的函数变量
	fmt.Println("The sum using FuncVar is:", FuncVar(15, 25))
}
