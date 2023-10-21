package main

import (
	"fmt"
	_ "unsafe"
	_ "strconv"
)

func main() {
	var i = 10
	// ptr 是指针，指向 int 变量的内存，
	// 里面存的值是 &i，即 i 的内存地址
	var ptr *int = &i
	fmt.Println("变量 i 的内存地址：", &i)
	fmt.Println("ptr 也是变量，请把 ptr 的地址打印出来：", &ptr)
	fmt.Println("ptr 指针变量里面存的值（另一个变量的内存地址）：", ptr)
	fmt.Println("ptr 指向的变量在内存中的值：", *ptr)
}