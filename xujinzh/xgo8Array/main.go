package main

import (
	"fmt"
)

func main() {
	// 定义一个数组
	var hens [6]float64
	// 初始化数组
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	// 计算数组的总和、平均值
	var sum float64 = 0.0
	for _, v := range hens {
		sum += v
	}
	average := sum / float64(len(hens))
	fmt.Printf("总和: %v, 平均值: %.2f\n", sum, average)

	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{4, 5, 6}
	var arr3 = [...]int{7, 8, 9} // 使用...自动推断长度
	var arr4 = [3]int{1: 10, 2: 20} // 指定索引初始化
	var arr5 = [...]int{1: 9, 2: 8, 3: 7} // 使用...自动推断长度，指定索引初始化
	arr6 := [...]string{1: "a", 2: "b", 3: "c"} // 使用...自动推断长度，指定索引初始化
	fmt.Println("arr:", arr)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println("arr4:", arr4)
	fmt.Println("arr5:", arr5)
	fmt.Println("arr6:", arr6)
}
