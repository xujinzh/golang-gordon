package main

import (
	"fmt"
)

type Student struct{}
type USB interface{}

func TypeJudge(items ...interface{}) {

	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数 类型不确定，值是%v\n", index, x)
		}
	}
}

func main() {
	fmt.Println()
	var n1 bool = true
	var n2 float32 = 2.2
	var n3 float64 = 3.3
	var n4 int = 30
	var n5 string = "xjz"
	n6 := "北京"
	n7 := Student{}
	n77 := &Student{}
	var n8 USB
	TypeJudge(n1, n2, n3, n4, n5, n6, n7, n77, n8)
}
