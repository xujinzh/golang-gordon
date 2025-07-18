package utils

import (
	"fmt"
)

var Age int
var Name string

func init() {
	fmt.Println("utils init function called")
	Age = 18
	Name = "xgo2"
}

func Test(n int) {
	if n > 2 {
		n--
		Test(n)
	} else {
		fmt.Println("n=", n)
	}
}

func Test2(n *int) {
	*n = *n + 1
	fmt.Println(*n)
}

func Test3(n, m int) (k int) {
	k = n + m
	return
}
