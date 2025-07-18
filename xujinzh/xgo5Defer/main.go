package main

import (
	"fmt"
	"strings"
)

var Name = "xujinzh"

func sum(n1, n2 int) int {

	// 当执行到 defer 时，暂不执行 defer 后面的代码，而是将其压入栈中，等到函数执行完毕后再执行
	// 执行的顺序是先进后出
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func changeName() {
	Name = "jinzhongxu"
	fmt.Println("changeName Name=", Name)
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)

	// 变量作用域
	fmt.Println("Name=", Name)
	changeName()
	fmt.Println("Name=", Name)
	fmt.Println("main end")

	var str string = strings.Trim("!helloo!", "!o")
	fmt.Println("str after trim:", str)
}
