package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct { // B 和 A 组合
	a A // 有名结构体
}

func main() {
	fmt.Println("hello")
	b := B{}
	// b.Name = "jack" // error，在有名结构体中，必须指定结构体名 b.A.Name，因为它不会往里再找，认为只有a
	b.a.Name = "jack"
	b.a.age = 18
	fmt.Println(b)
}
