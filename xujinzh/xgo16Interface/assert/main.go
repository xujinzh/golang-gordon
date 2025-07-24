package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	b = a.(Point) // 类型断言
	// b = a // error
	fmt.Println(b)
	fmt.Println(a)

	// 其他案例
	var x interface{}
	var t float32 = 1.1
	x = t
	y := x.(float32)
	fmt.Printf("%T, %v\n", y, y)

	// 断言检测
	var m interface{}
	var n float64 = 3.14
	m = n
	// k, ok := m.(float64) // 带检测 right
	// k, ok := m.(float32) // 带检测 right
	// if ok {
	// 	fmt.Printf("%T, %v\n", k, k)
	// } else {
	// 	fmt.Println("convert failed")
	// }

	if k, ok := m.(float32); ok {
		fmt.Printf("%T, %v\n", k, k)
	} else {
		fmt.Println("convert failed")
	}

}
