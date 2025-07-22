package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Cat struct {
	Name  string
	Age   int
	Color string
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp    Point
	rightDown Point
}

type Rect2 struct {
	leftUp    *Point
	rightDown *Point
}

type A struct {
	Num int
}

type B struct {
	Num int
}

type Dog struct {
	Name  string `json:"name"` // struct tag
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// 结构体 Dog 的方法
func (dog Dog) Eat() {
	fmt.Println(dog.Name, "啃骨头")
}

type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

type Person struct {
	Name string
	Age  int
}

// 给 Person 实现 String() 方法
func (person *Person) String() string {
	str := fmt.Sprintf("Name=[%v], Age=[%v]\n", person.Name, person.Age)
	return str
}

func main() {
	fmt.Println("hello")
	// 方式一
	var cat1 Cat
	fmt.Println(cat1)
	cat1.Name = "xiaobai"
	cat1.Age = 1
	cat1.Color = "white"
	fmt.Println(cat1)

	// 方式二
	cat2 := Cat{"xiaohua", 2, "color"}
	fmt.Println(cat2)

	// 方式三
	var cat3 *Cat = new(Cat)
	// golang 设计者为了程序员方便，底层进行了优化
	// 也可以用 cat3.Name = "xiaoqing"
	(*cat3).Name = "xiaoqing"
	fmt.Println(*cat3)

	// 方式四
	var cat4 *Cat = &Cat{}
	// var cat4 *Cat = &Cat{"xiaowang"}
	(*cat4).Name = "xiaowang"
	fmt.Println(*cat4)

	// struct 字段在内存中是连续的
	rect := Rect{Point{1, 2}, Point{3, 4}}
	// rect 的四个int，在内存中是连续分布的
	// 0xc000018200 0xc000018208 0xc000018210 0xc000018218
	fmt.Println(&rect.leftUp.x, &rect.leftUp.y, &rect.rightDown.x, &rect.rightDown.y)

	rect2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	// 0xc0000140a0 0xc0000140a8
	fmt.Println(&rect2.leftUp, &rect2.rightDown)

	// 结构体转换
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)

	// 将结构体序列号为json字符串
	dog1 := Dog{"xiaohei", 1, "gray"}
	jsonStr, err := json.Marshal(dog1)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(jsonStr)
		fmt.Printf("%q\n", string(jsonStr))
	}

	// 方法
	dog2 := Dog{"阿黄", 2, "yellow"}
	dog2.Eat()

	// 计算圆的面积
	circle := Circle{4.0}
	area := circle.Area()
	fmt.Printf("圆的面积：%.3f\n", area)

	// String() 方法改变 fmt.Println() 的行为
	person := Person{
		Name: "xujinzh",
		Age:  18,
	}
	fmt.Println(&person)
}
