package main

import (
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int // 匿名字段，基本数据类型；只能有一个同名的匿名字段
	n   int
}

func main() {
	fmt.Println()
	e := E{}
	e.Monster.Name = "hulijing"
	e.Monster.Age = 500
	e.int = 10
	e.n = 10
	fmt.Println(e)
}
