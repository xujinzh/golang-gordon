package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string // B 结构体中也有该字段，如果创建B的实例，想要赋值A结构体中该Name字段的值，需要显示指定 A.Name，否则就近原则默认访问B.Name
}

func (b *B) SayOk() {
	fmt.Println("B sayok", b.Name)
}

func main() {
	// var b B
	// b.A.Name = "tome"
	// b.A.age = 10
	// b.A.SayOk()
	// b.A.hello()

	// // 下面简化写
	// b.Name = "smith"
	// b.age = 20
	// b.SayOk()
	// b.hello()

	var b B
	b.Name = "jack" // 使用B的字段
	b.age = 100     // 使用A的字段
	b.SayOk()       // 使用B的方法
	b.hello()       // 使用A的方法，但是没有给 b.A.Name 赋值，所有a.Name为默认值空字符串。可以显示赋值解决
	b.A.Name = "lucy"
	b.hello()
}
