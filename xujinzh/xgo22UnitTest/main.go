package main

import (
	"fmt"
)

// 被测函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	fmt.Println()
	// 传统测试方法
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
	} else {
		fmt.Printf("addUpper 正确，返回值=%v 期望值=%v\n", res, 55)
	}
}
