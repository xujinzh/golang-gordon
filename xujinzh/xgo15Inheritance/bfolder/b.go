package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	// A和B同级，不分先后
	A
	B
}

func main() {
	fmt.Println("hello")
	var c C
	// c.Name = "tom" // error，编译器不能确定到底是A还是B结构体里的字段
	fmt.Println(c)
}
