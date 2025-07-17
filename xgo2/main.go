package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("n=", n)
	}
}

func test2(n *int) {
	*n = *n + 1
	fmt.Println(*n)
}

func test3(n, m int) (k int) {
	k = n + m
	return
}

func main() {
	test(4)
	n := 1
	test2(&n)
	fmt.Println(n)
	fmt.Println(test3(2, 3))
}
