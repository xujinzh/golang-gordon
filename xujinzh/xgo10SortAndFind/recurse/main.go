package main

import "fmt"

/*
递归需要遵守的重要原则
1. 执行一个函数时，就创建一个新的受保护的独立空间（新函数栈）
2. 函数的局部变量是独立的，不会相互影响；但是，如果希望各个函数栈使用同一个/组数据，可以使用全局变量或（指针）引用
3. 递归必须向退出递归的条件逼近，否则就是无限递归
4. 当一个函数执行完毕，或者遇到 return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或者返回时，该函数本身也会被系统销毁

常用递归解决的例子有八皇后问题、汉诺塔问题、阶乘问题、迷宫问题、球和篮子问题（GOOGLE编程大赛）等
*/

func recursePrint(n int) {
	if n > 2 {
		n--
		recursePrint(n)
	}
	fmt.Println("n =", n)
}

func recursePrintElse(n int) {
	if n > 2 {
		n--
		recursePrintElse(n)
	} else {
		fmt.Println("n =", n)
	}
}

func main() {
	n := 4
	recursePrint(n)
	fmt.Println()
	recursePrintElse(n)
}
